import githubInstance from "./user";
import {GetInstance} from "./user";
import { Base64 } from 'js-base64';
// 仓库信息接口定义
export interface Repository {
  id: number;
  owner: {
    login: string;
    id: number;
    avatar_url: string;
    html_url: string;
    type: string;
  };
  name: string;
  full_name: string;
  private: boolean;
  description: string | null;
  fork: boolean;
  created_at: string;
  updated_at: string;
  pushed_at: string;
  homepage: string | null;
  size: number;
  stargazers_count: number;
  watchers_count: number;
  language: string | null;
  forks_count: number;
  open_issues_count: number;
  default_branch: string;
  topics: string[];
  visibility: string;
  html_url: string;
}

// 仓库列表查询参数接口
export interface ListRepositoriesParams {
  type?: 'all' | 'owner' | 'public' | 'private' | 'member';
  sort?: 'created' | 'updated' | 'pushed' | 'full_name';
  direction?: 'asc' | 'desc';
  per_page?: number;
  page?: number;
}

// 仓库搜索参数接口
export interface SearchRepositoriesParams {
  q: string;
  sort?: 'stars' | 'forks' | 'help-wanted-issues' | 'updated';
  order?: 'desc' | 'asc';
  per_page?: number;
  page?: number;
}

/**
 * 获取当前用户的仓库列表
 * @param params 可选的查询参数
 * @example
 * // 获取所有仓库，按创建时间降序排序
 * const repos = await listUserRepositories({
 *   type: 'all',
 *   sort: 'created',
 *   direction: 'desc'
 * });
 */
export const listUserRepositories = (params?: ListRepositoriesParams) => {
  return githubInstance.get<Repository[]>('/user/repos', { params });
};

/**
 * 获取指定用户的仓库列表
 * @param username GitHub 用户名
 * @param params 可选的查询参数
 * @example
 * // 获取用户 octocat 的公开仓库
 * const repos = await listUserRepositoriesByUsername('octocat', {
 *   type: 'public'
 * });
 */
export const listUserRepositoriesByUsername = (username: string, params?: ListRepositoriesParams) => {
  return githubInstance.get<Repository[]>(`/users/${username}/repos`, { params });
};

/**
 * 获取仓库详细信息
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @example
 * // 获取仓库详情
 * const repoInfo = await getRepository('octocat', 'Hello-World');
 */
export const getRepository = (owner: string, repo: string) => {
  return githubInstance.get<Repository>(`/repos/${owner}/${repo}`);
};

/**
 * 搜索仓库
 * @param params 搜索参数
 * @example
 * // 搜索带有 "javascript" 标签的仓库，按星标数降序排序
 * const searchResult = await searchRepositories({
 *   q: 'language:javascript',
 *   sort: 'stars',
 *   order: 'desc'
 * });
 */
export const searchRepositories = (params: SearchRepositoriesParams) => {
  return githubInstance.get('/search/repositories', { params });
};

/**
 * 为仓库添加星标
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @example
 * // 为仓库添加星标
 * await starRepository('octocat', 'Hello-World');
 */
export const starRepository = (owner: string, repo: string) => {
  return githubInstance.put(`/user/starred/${owner}/${repo}`);
};

/**
 * 取消仓库星标
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @example
 * // 取消仓库星标
 * await unstarRepository('octocat', 'Hello-World');
 */
export const unstarRepository = (owner: string, repo: string) => {
  return githubInstance.delete(`/user/starred/${owner}/${repo}`);
};

/**
 * 检查当前用户是否为指定仓库添加了星标
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @example
 * // 检查是否已添加星标
 * const isStarred = await checkRepositoryIsStarred('octocat', 'Hello-World');
 */
export const checkRepositoryIsStarred = (owner: string, repo: string) => {
  return githubInstance.get(`/user/starred/${owner}/${repo}`)
    .then(() => true)
    .catch(() => false);
};

/**
 * 获取当前用户的star仓库列表
 * @param params 可选的查询参数
 * @example
 * // 获取所有star仓库，按创建时间降序排序
 * const repos = await listUserStarredRepositories({
 *   sort: 'created',
 *   direction: 'desc'
 * });
 */
export const listUserStarredRepositories = (params?: ListRepositoriesParams) => {
  return githubInstance.get<Repository[]>('/user/starred', { params });
};

// 提交信息接口定义
export interface Commit {
  sha: string;
  commit: {
    author: {
      name: string;
      email: string;
      date: string;
    };
    message: string;
  };
  html_url: string;
  author: {
    login: string;
    avatar_url: string;
  };
}

// 文件内容接口定义
export interface FileContent {
  name: string;
  path: string;
  sha: string;
  size: number;
  url: string;
  html_url: string;
  git_url: string;
  download_url: string | null;
  type: 'file' | 'dir';
  content?: string;
  encoding?: string;
  icon?: string;
  message?: string;
  time?: string;
}

// 分支信息接口定义
export interface Branch {
  name: string;
  commit: {
    sha: string;
    url: string;
  };
  protected: boolean;
}

// 贡献者信息接口定义
export interface Contributor {
  login: string;
  id: number;
  avatar_url: string;
  html_url: string;
  contributions: number;
}

/**
 * 获取仓库的 README 内容
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @example
 * // 获取仓库的 README 内容
 * const readme = await getRepositoryReadme('octocat', 'Hello-World');
 */
export const getRepositoryReadme = (owner: string, repo: string) => {
  return githubInstance.get<{ content: string }>(`/repos/${owner}/${repo}/readme`)
    .then(response => {
      // GitHub API 返回的是 Base64 编码的内容
      // 使用 atob 进行 base64 解码，这是浏览器原生支持的方法
      return Base64.decode(response.data.content);
    });
};

/**
 * 获取仓库的 README 内容
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @example
 * // 获取仓库的 README 内容
 * const readme = await getRepositoryReadme('octocat', 'Hello-World');
 */
export const getRepositoryReadmeDir = (owner: string, repo: string,dir:string) => {
  return githubInstance.get<{ content: string }>(`/repos/${owner}/${repo}/readme/${dir}`)
    .then(response => {
      // GitHub API 返回的是 Base64 编码的内容
      // 使用 atob 进行 base64 解码，这是浏览器原生支持的方法
      return Base64.decode(response.data.content);
    });
};

/**
 * 获取仓库的提交历史
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @param params 可选的查询参数
 * @example
 * // 获取最近的 10 条提交记录
 * const commits = await getRepositoryCommits('octocat', 'Hello-World', {
 *   per_page: 10,
 *   page: 1
 * });
 */
export const getRepositoryCommits = (owner: string, repo: string, params?: { per_page?: number; page?: number; sha?: string; path?: string }) => {
  return githubInstance.get<Commit[]>(`/repos/${owner}/${repo}/commits`, { params });
};

/**
 * wner, repo, path, branch
 * 获取仓库的文件列表
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @param path 文件路径，默认为根目录
 * @param ref 分支或提交的引用
 * @example
 * // 获取主分支根目录的文件列表
 * const files = await getRepositoryContents('octocat', 'Hello-World', '', 'main');
 */
export const getRepositoryContents = (owner: string, repo: string, path: string = '', ref?: string) => {
  return githubInstance.get<FileContent | FileContent[]>(`/repos/${owner}/${repo}/contents/${path}`, {
    params: { ref }
  });
};

/**
 * 获取仓库的分支列表
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @example
 * // 获取仓库的所有分支
 * const branches = await getRepositoryBranches('octocat', 'Hello-World');
 */
export const getRepositoryBranches = (owner: string, repo: string) => {
  return githubInstance.get<Branch[]>(`/repos/${owner}/${repo}/branches`);
};

/**
 * 获取仓库的贡献者列表
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @example
 * // 获取仓库的所有贡献者
 * const contributors = await getRepositoryContributors('octocat', 'Hello-World');
 */
export const getRepositoryContributors = (owner: string, repo: string) => {
  return githubInstance.get<Contributor[]>(`/repos/${owner}/${repo}/contributors`);
};

/**
 * 获取仓库文件内容
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @param path 文件路径
 * @param ref 分支或提交的引用
 * @example
 * // 获取主分支下指定文件的内容
 * const fileContent = await getRepositoryFileContent('octocat', 'Hello-World', 'README.md', 'main');
 */
export const getRepositoryFileContent = (owner: string, repo: string, path: string, ref?: string) => {
  return githubInstance.get<FileContent>(`/repos/${owner}/${repo}/contents/${path}`, {
    params: { ref }
  }).then(response => {
    if (response.data.content && response.data.encoding === 'base64') {
      return {
        ...response.data,
        content: atob(response.data.content)
      };
    }
    return response.data;
  });
};

/**
 * 获取仓库文件内容
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @param branch 分支
 * @param ref 分支或提交的引用
 * @example
 * // 获取指定文件的内容
 * const fileContent = await getRepositoryFileContent('octocat', 'Hello-World', 'README.md', 'main');
 */
export const getRepositoryFile= (owner: string, repo: string, branch: string, ref?: string) => {
  return GetInstance.get(`/repos/${owner}/${repo}/blob/${branch}/${ref}`).then(response => {
    response.data = response.data.payload
    return response.data;
  });
};


/**
 * 获取仓库文件提交信息
 * @param owner 仓库所有者
 * @param repo 仓库名称
 * @param branch 分支
 * @param ref 分支或提交的引用
 * @example
 * // 获取指定文件的内容
 * const fileContent = await getRepositoryFileContent('octocat', 'Hello-World', 'README.md', 'main');
 */
export const getDirectoryContentsWithCommits = async (
  owner: string,
  repo: string,
  path: string,
  branch: string = 'develop'
) => {
  const response = await githubInstance.get(`/repos/${owner}/${repo}/contents/${path}`, {
    params: { ref: branch }
  });

  // 获取每个文件的提交信息
  const filesWithCommits = await Promise.all(
    response.data.map(async (file: any) => {
      const commits = await getFileCommits(owner, repo, file.path, branch);
      return {
        ...file,
        lastCommit: commits.data[0] // 取最近一次提交
      };
    })
  );

  return filesWithCommits;
};

export const getFileCommits = (
  owner: string,
  repo: string,
  path: string,
  branch: string
) => {
  return githubInstance.get(`/repos/${owner}/${repo}/commits`, {
    params: {
      path: path,
      sha: branch,
      per_page: 1 // 只获取最近一次提交
    }
  });
};