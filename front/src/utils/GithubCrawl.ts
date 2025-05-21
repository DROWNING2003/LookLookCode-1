// API 和类型定义导入
import {
  getRepository,
  getRepositoryCommits,
  getRepositoryReadme,
  getRepositoryBranches,
  getRepositoryContributors,
  getRepositoryContents,
  // getRepositoryCommits,
  type Repository,
  type Branch,
  type FileContent,
  // type Commit,
} from "../api/repo";

import { Base64 } from "js-base64";
// 类型定义
// interface GithubFileResult {
//   [filePath: string]: string;
// }

// interface CrawlStats {
//   downloadedCount: number;
//   skippedCount: number;
//   skippedFiles: Array<[string, number]>;
//   basePath?: string;
//   includePatterns?: Set<string>;
//   excludePatterns?: Set<string>;
// }

interface CrawlOptions {
  token?: string;
  maxFileSize?: number;
  useRelativePaths?: boolean;
  includePatterns?: string | string[];
  excludePatterns?: string | string[];
}

interface GithubContentItem {
  type: "file" | "dir";
  path: string;
  name: string;
  size?: number;
  download_url?: string | null;
  url: string;
  content?: string;
  encoding?: string;
}

// 主函数
export async function crawlGithubFiles(
  owner: string,
  repo: string,
  branch: string,
  targetPath: string,
  options: CrawlOptions = {
    includePatterns: ["*"],
    excludePatterns: [
      ".gitignore",
      "LICENSE",
      "node_modules",
      ".git",
      ".github",
      ".vscode",
      "dist",
      "build",
      "*.test.*",
      "*.spec.*",
      ".DS_Store",
      ".env",
      ".env.*",
      "*.log",
    ],
  }
) {
  const {
    maxFileSize = 1 * 1024 * 1024,
    useRelativePaths = false,
    includePatterns,
    excludePatterns,
  } = options;
  const files: Record<string, string> = {};
  const skippedFiles: Array<[string, number]> = [];

  console.log(owner, repo, branch, targetPath);

  // 并发控制参数
  const MAX_CONCURRENT_REQUESTS = 5;
  let currentRequests = 0;
  const requestQueue: (() => Promise<void>)[] = [];

  // 并发控制函数
  const executeWithConcurrencyControl = async (fn: () => Promise<void>) => {
    if (currentRequests >= MAX_CONCURRENT_REQUESTS) {
      return new Promise<void>((resolve) => {
        requestQueue.push(async () => {
          await fn();
          resolve();
        });
      });
    }

    currentRequests++;
    try {
      await fn();
    } finally {
      currentRequests--;
      if (requestQueue.length > 0) {
        const nextRequest = requestQueue.shift();
        if (nextRequest) {
          executeWithConcurrencyControl(nextRequest);
        }
      }
    }
  };

  // 处理单个文件内容
  const fetchFileContent = async (item: GithubContentItem, relativePath: string) => {
    try {
      const response = await getRepositoryContents(owner, repo, item.path, "");
      const content = <FileContent>response.data;
      if (content !== null) {
        files[relativePath] = Base64.decode(<string>content.content);
      }
    } catch (error) {
      console.error(`Error fetching file ${item.path}:`, error);
    }
  };

  // 递归获取内容
  const fetchContents = async (currentPath: string): Promise<void> => {
    try {
      const response = await getRepositoryContents(owner, repo, currentPath, branch);
      const filePromises: Promise<void>[] = [];

      for (const item of response.data) {
        const relativePath = useRelativePaths && targetPath
          ? item.path.replace(`${targetPath}/`, "")
          : item.path;

        // 文件过滤
        if (!matchPatterns(relativePath, includePatterns, excludePatterns)) {
          console.log("过滤的文件", relativePath);
          continue;
        }

        // 大小检查
        const fileSize = item.size || 0;
        if (fileSize > maxFileSize) {
          skippedFiles.push([relativePath, fileSize]);
          continue;
        }

        if (item.type === "dir") {
          console.log("文件夹", item.path);
          filePromises.push(fetchContents(item.path));
        } else {
          console.log("文件", item.path);
          filePromises.push(
            executeWithConcurrencyControl(() => fetchFileContent(item, relativePath))
          );
        }
      }

      await Promise.all(filePromises);
    } catch (error) {
      console.error(`Error processing directory ${currentPath}:`, error);
    }
  };

  // 开始爬取
  await fetchContents(targetPath);
  // if (content !== null) {
  //   files[relativePath] = content;
  // }
  return {
    files,
    // stats: {
    //   downloadedCount: Object.keys(files).length,
    //   skippedCount: skippedFiles.length,
    //   skippedFiles,
    //   basePath: useRelativePaths ? targetPath : undefined,
    //   includePatterns: includeSet,
    //   excludePatterns: excludeSet,
    // },
  };
}

//   // 开始爬取
//   await fetchContents(targetPath);

//   return {
//     files,
//     stats: {
//       downloadedCount: Object.keys(files).length,
//       skippedCount: skippedFiles.length,
//       skippedFiles,
//       basePath: useRelativePaths ? targetPath : undefined,
//       includePatterns: includeSet,
//       excludePatterns: excludeSet,
//     },
//   };
// }

// 辅助函数 ---------------------------------------------------

/** 处理glob模式 */
function processPatterns(patterns?: string | string[]): Set<string> {
  if (!patterns) return new Set();
  if (typeof patterns === "string") return new Set([patterns]);
  return new Set(patterns);
}

// 原生模式匹配
function matchPatterns(
  filePath: string,
  includes?: string | string[],
  excludes?: string | string[]
): boolean {
  const fileName = filePath.split("/").pop() || "";

  // 包含模式
  if (includes) {
    const includePatterns = Array.isArray(includes) ? includes : [includes];
    if (!includePatterns.some((p) => globMatch(fileName, p))) return false;
  }

  // 排除模式
  if (excludes) {
    const excludePatterns = Array.isArray(excludes) ? excludes : [excludes];
    if (excludePatterns.some((p) => globMatch(filePath, p))) return false;
  }

  return true;
}

// 简单通配符匹配
function globMatch(str: string, pattern: string): boolean {
  const regex = new RegExp(
    `^${pattern
      .replace(/\./g, "\\.")
      .replace(/\*/g, ".*")
      .replace(/\?/g, ".")}$`
  );
  return regex.test(str);
}

// /** 获取文件内容 */
// async function getFileContent(
//   item: GithubContentItem,
//   headers: Record<string, string>,
//   maxSize: number
// ): Promise<string | null> {
//   try {
//     // 优先使用download_url
//     if (item.download_url) {
//       const response = await axios.get<string>(item.download_url, { headers });
//       return response.data;
//     }

//     // 备用方案：通过API获取内容
//     const response = await axios.get<GithubContentItem>(item.url, { headers });
//     const data = response.data;

//     if (data.encoding === 'base64' && data.content) {
//       const content = Buffer.from(data.content, 'base64').toString('utf-8');
//       if (content.length > maxSize) return null;
//       return content;
//     }

//     return null;
//   } catch (error) {
//     console.error(`Failed to fetch ${item.path}: ${error.message}`);
//     return null;
//   }
// }

// /** 处理API错误 */
// async handleApiError(
//   error: AxiosError,
//   headers: Record<string, string>,
//   currentPath: string
// ): Promise<void> {
//   const response = error.response;

//   if (!response) {
//     console.error(`Network error: ${error.message}`);
//     return;
//   }

//   // 处理速率限制
//   if (response.status === 403 && response.headers['x-ratelimit-remaining'] === '0') {
//     const resetTime = parseInt(response.headers['x-ratelimit-reset'] || '0', 10);
//     const waitTime = Math.max(resetTime * 1000 - Date.now(), 0) + 1000;
//     console.log(`Rate limit exceeded. Waiting ${Math.round(waitTime / 1000)}s...`);
//     await new Promise(resolve => setTimeout(resolve, waitTime));
//     return;
//   }

//   // 处理404错误
//   if (response.status === 404) {
//     const isPrivateRepo = !headers.Authorization;
//     const errorMsg = isPrivateRepo
//       ? 'Private repository requires authentication'
//       : `Path '${currentPath}' not found`;
//     console.error(`404 Error: ${errorMsg}`);
//     return;
//   }

//   console.error(`API Error ${response.status}: ${response.data?.message || ''}`);
// }

// // 使用示例 ---------------------------------------------------
// // (async () => {
// //   try {
// //     const result = await crawlGithubFiles(
// //       'https://github.com/microsoft/TypeScript/tree/main/src',
// //       {
// //         token: process.env.GITHUB_TOKEN,
// //         maxFileSize: 2 * 1024 * 1024, // 2MB
// //         useRelativePaths: true,
// //         includePatterns: ['*.ts', '*.md'],
// //         excludePatterns: ['*test*']
// //       }
// //     );

// //     console.log(`Downloaded ${result.stats.downloadedCount} files`);
// //     console.log('Sample files:', Object.keys(result.files).slice(0, 5));
// //   } catch (error) {
// //     console.error('Crawl failed:', error.message);
// //   }
// // })();
