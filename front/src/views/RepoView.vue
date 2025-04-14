<!--
 * @description 仓库详情页面组件
 * @features
 * - 显示仓库基本信息（名称、描述、统计数据等）
 * - 支持分支切换和文件浏览
 * - 集成文件图标系统
 * - 显示文件的最后提交信息
 * @dependencies
 * - Vue 3 Composition API
 * - Vue Router
 * - TailwindCSS
 * - Font Awesome 图标
 -->
<template>
  <div class="min-h-screen bg-primary text-primary">
    <TopNavBar
      :title="repository?.full_name || '仓库详情'"
      @search="handleSearch"
    />
    <!-- 主要内容区域 -->
    <main class="pt-16 pb-20 px-4 max-w-7xl mx-auto w-full">
      <!-- 加载状态 -->
      <div v-if="loading" class="flex items-center justify-center py-12">
        <div
          class="animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-blue-500"
        ></div>
      </div>
      <!-- 错误提示 -->
      <div v-else-if="error" class="text-red-500 text-center py-8">
        {{ error }}
      </div>
      <!-- 仓库内容 -->
      <template v-else>
        <!-- 仓库信息部分保持不变 -->
        <div class="mb-8" v-if="repository">
          <div class="flex items-center gap-2 mb-2">
            <i class="far fa-book text-gray-400"></i>
            <h1 class="text-xl font-semibold text-blue-400">
              {{ repository.full_name }}
            </h1>
          </div>
          <p class="text-gray-400 mb-4">{{ repository.description }}</p>
          <!-- 统计信息 -->
          <div class="flex gap-6 mb-4">
            <button class="flex items-center gap-1 hover:text-blue-400">
              <i class="far fa-star"></i>
              <span>{{ repository.stargazers_count }} Star</span>
            </button>
            <button class="flex items-center gap-1 hover:text-blue-400">
              <i class="fas fa-code-branch"></i>
              <span>{{ repository.forks_count }} Fork</span>
            </button>
            <button class="flex items-center gap-1 hover:text-blue-400">
              <i class="far fa-eye"></i>
              <span>{{ repository.watchers_count }} Watch</span>
            </button>
          </div>
          <!-- 技术标签 -->
          <div class="flex flex-wrap gap-2">
            <span
              v-if="repository.language"
              class="px-2 py-1 text-sm bg-gray-700 rounded-full"
              >{{ repository.language }}</span
            >
            <span
              v-for="topic in repository.topics"
              :key="topic"
              class="px-2 py-1 text-sm bg-gray-700 rounded-full"
              >{{ topic }}</span
            >
          </div>
        </div>
      </template>
      <!-- 分支选择和文件列表 -->
      <div class="mb-8">
        <div class="flex items-center gap-4 mb-4 relative">
          <button
            @click="showBranchDropdown = !showBranchDropdown"
            class="flex items-center gap-2 bg-card px-3 py-2 rounded-lg hover:bg-card-hover transition-colors duration-200"
          >
            <i class="fas fa-code-branch"></i>
            <span>{{ currentBranch }}</span>
            <i class="fas fa-chevron-down text-sm"></i>
          </button>
          <div
            v-if="showBranchDropdown"
            class="absolute top-full left-0 mt-2 w-48 bg-card rounded-lg shadow-lg overflow-hidden z-50"
          >
            <button
              v-for="branch in branches"
              :key="branch.name"
              @click="switchBranch(branch.name)"
              class="w-full px-4 py-2 text-left hover:bg-card-hover transition-colors duration-200"
            >
              {{ branch.name }}
            </button>
          </div>
          <a
            v-if="repository"
            :href="repository.html_url"
            target="_blank"
            class="flex items-center gap-2 bg-accent hover:bg-accent-hover text-white px-6 py-2 rounded-lg transition-colors duration-200"
          >
            <i class="fas fa-book-open"></i>
            <span>在 GitHub 上查看</span>
          </a>
        </div>
        <!-- 文件列表 -->
        <div class="bg-card rounded-lg overflow-hidden">
          <div v-if="currentPath" class="border-b border-gray-700">
            <button
              @click="goToParentDirectory"
              class="flex items-center gap-3 px-4 py-3 hover:bg-card-hover w-full text-left"
            >
              <i class="fas fa-arrow-left text-gray-400"></i>
              <span>返回上级目录</span>
            </button>
          </div>
          <div
            v-for="file in files"
            :key="file.path"
            class="border-b border-gray-700 last:border-b-0"
          >
            <div
              class="flex items-center justify-between px-4 py-3 hover:bg-card-hover cursor-pointer"
              @click="navigateToFile(file)"
            >
              <div class="flex items-center gap-3 flex-1">
                <i
                  :class="getFileIcon(file.type, file.name)"
                  class="text-gray-400"
                ></i>
                <div class="min-w-0">
                  <p class="text-sm font-medium truncate">{{ file.name }}</p>
                  <p v-if="file.message" class="text-xs text-gray-400 truncate">
                    {{ file.message }}
                  </p>
                </div>
              </div>
              <span
                v-if="file.time"
                class="text-xs text-gray-400 flex-shrink-0"
                >{{ formatDate(file.time) }}</span
              >
            </div>
          </div>
        </div>
        <!-- readme 预览 -->
        <MdPreview id="preview-only" :modelValue="readme" />
      </div>
    </main>
  </div>
</template>

<script lang="ts" setup>
// Vue 核心功能导入
import { ref, watch, onMounted } from "vue";
import { useRoute } from "vue-router";
import { MdPreview} from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
// 组件导入
import TopNavBar from "../components/TopNavBar.vue";

// API 和类型定义导入
import {
  getRepository,
  getRepositoryReadme,
  getRepositoryBranches,
  getRepositoryContributors,
  getRepositoryContents,
  getRepositoryCommits,
  type Repository,
  type Branch,
  type FileContent,
  type Commit,
} from "../api/repo";

// 路由参数获取
const route = useRoute();
const owner = route.params.owner as string;
const repo = route.params.repo as string;

// 状态管理
const currentPath = ref(""); // 当前浏览的文件路径

/**
 * 处理搜索事件
 * @todo 实现仓库内容搜索功能
 */
const handleSearch = () => {
  console.log("Search triggered");
};

// 仓库数据相关的响应式状态
const repository = ref<Repository | null>(null); // 仓库基本信息
const readme = ref(""); // 仓库 README 内容
const branches = ref<Branch[]>([]); // 仓库分支列表
const contributors = ref<number>(0); // 贡献者数量
const currentBranch = ref("main"); // 当前选中的分支
const files = ref<FileContent[]>([]); // 当前目录下的文件列表
const commits = ref<Commit[]>([]); // 提交历史

// UI 状态控制
const showBranchDropdown = ref(false); // 分支下拉菜单显示状态
const loading = ref(true); // 加载状态
const error = ref<string | null>(null); // 错误信息

/**
 * 获取仓库所有相关数据
 * 包括：基本信息、README、分支列表、贡献者、目录内容、提交历史
 */
const fetchRepoData = async () => {
  loading.value = true;
  error.value = null;
  try {
    const [
      repoData,
      readmeContent,
      branchesData,
      contributorsData,
      contentsData,
      commitsData,
    ] = await Promise.all([
      getRepository(owner, repo),
      getRepositoryReadme(owner, repo),
      getRepositoryBranches(owner, repo),
      getRepositoryContributors(owner, repo),
      getRepositoryContents(
        owner,
        repo,
        currentPath.value,
        currentBranch.value
      ),
      getRepositoryCommits(owner, repo, {
        per_page: 10,
        sha: currentBranch.value,
      }),
    ]);

    repository.value = repoData.data;
    readme.value = readmeContent;
    branches.value = branchesData.data;
    contributors.value = contributorsData.data.length;

    // 为每个文件添加最后一次提交信息
    const filesWithCommitInfo = Array.isArray(contentsData.data)
      ? await Promise.all(
          contentsData.data.map(async (file) => {
            try {
              const fileCommits = await getRepositoryCommits(owner, repo, {
                per_page: 1,
                path: file.path,
                sha: currentBranch.value,
              });
              if (fileCommits.data.length > 0) {
                const lastCommit = fileCommits.data[0];
                return {
                  ...file,
                  message: lastCommit.commit.message,
                  time: lastCommit.commit.author.date,
                };
              }
              return file;
            } catch (error) {
              console.error(`Error fetching commits for ${file.path}:`, error);
              return file;
            }
          })
        )
      : [];

    files.value = filesWithCommitInfo;
    commits.value = commitsData.data;
  } catch (error) {
    console.error("Error fetching repository data:", error);
    // error.value = "获取仓库数据失败，请稍后重试" as string;
  } finally {
    loading.value = false;
  }
};

/**
 * 处理文件导航
 * @param file 选中的文件或目录
 */
const navigateToFile = (file: FileContent) => {
  console.log('文件详情:', file);
  // 如果是文件，获取文件内容
  if (file.type === 'file') {
    getRepositoryContents(owner, repo, file.path as string, currentBranch.value)
      .then((response) => {
        const content = response.data;
        console.log('文件内容:', content);
      })
      .catch((err) => {
        console.error('获取文件内容失败:', err);
        error.value = '获取文件内容失败，请稍后重试';
      });
  }
  if (file.type === "dir") {
    // 如果是目录，更新当前路径并重新获取内容
    currentPath.value = file.path as string;
    getRepositoryContents(owner, repo, currentPath.value, currentBranch.value)
      .then((response) => {
        files.value = Array.isArray(response.data) ? response.data : [response.data];
      })
      .catch((err) => {
        console.error("Error fetching directory contents:", err);
        error.value = "获取目录内容失败，请稍后重试";
      });
  } else {
    // 获取文件内容
   
  }
};

/**
 * 返回上级目录
 * 通过分割当前路径并移除最后一个部分来实现
 */
const goToParentDirectory = () => {
  if (currentPath.value) {
    const parentPath = currentPath.value.split("/").slice(0, -1).join("/");
    currentPath.value = parentPath;
    getRepositoryContents(owner, repo, parentPath, currentBranch.value)
      .then((response) => {
        files.value = Array.isArray(response.data) ? response.data : [response.data];
      })
      .catch((err) => {
        console.error("Error fetching parent directory contents:", err);
        error.value = "获取目录内容失败，请稍后重试";
      });
  }
};
/**
 * 切换分支
 * @param branchName 目标分支名称
 */
const switchBranch = async (branchName: string) => {
  currentBranch.value = branchName;
  showBranchDropdown.value = false;
  await fetchRepoData();
};

/**
 * 格式化日期为相对时间
 * @param dateString ISO 格式的日期字符串
 * @returns 格式化后的相对时间描述
 */
const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  const now = new Date();
  const diff = now.getTime() - date.getTime();
  const days = Math.floor(diff / (1000 * 60 * 60 * 24));

  if (days === 0) return "今天";
  if (days === 1) return "昨天";
  if (days < 7) return `${days}天前`;
  if (days < 30) return `${Math.floor(days / 7)}周前`;
  if (days < 365) return `${Math.floor(days / 30)}月前`;
  return `${Math.floor(days / 365)}年前`;
};

/**
 * 根据文件类型和名称获取对应的图标类名
 * @param type 文件类型（'dir' 或 'file'）
 * @param name 文件名
 * @returns Font Awesome 图标类名
 */
const getFileIcon = (type: string, name: string) => {
  if (type === "dir") return "fas fa-folder text-blue-300";

  // 根据文件扩展名返回不同图标
  const ext = name.split(".").pop()?.toLowerCase();
  switch (ext) {
    case "md":
      return "fab fa-markdown text-blue-400";
    case "json":
      return "fas fa-file-code text-yellow-400";
    case "js":
    case "ts":
    case "jsx":
    case "tsx":
      return "fab fa-js-square text-yellow-300";
    case "html":
    case "htm":
      return "fab fa-html5 text-orange-500";
    case "css":
    case "scss":
    case "less":
      return "fab fa-css3-alt text-blue-500";
    case "py":
      return "fab fa-python text-blue-400";
    case "java":
      return "fab fa-java text-red-500";
    case "php":
      return "fab fa-php text-purple-500";
    case "go":
      return "fab fa-golang text-blue-500";
    case "rb":
      return "fas fa-gem text-red-500";
    case "sh":
      return "fas fa-terminal text-green-500";
    case "png":
    case "jpg":
    case "jpeg":
    case "gif":
    case "svg":
      return "fas fa-image text-purple-400";
    case "zip":
    case "tar":
    case "gz":
    case "rar":
      return "fas fa-file-archive text-yellow-500";
    case "pdf":
      return "fas fa-file-pdf text-red-500";
    case "doc":
    case "docx":
      return "fas fa-file-word text-blue-500";
    case "xls":
    case "xlsx":
      return "fas fa-file-excel text-green-500";
    case "ppt":
    case "pptx":
      return "fas fa-file-powerpoint text-orange-500";
    case "mp3":
    case "wav":
    case "ogg":
      return "fas fa-file-audio text-purple-400";
    case "mp4":
    case "mov":
    case "avi":
      return "fas fa-file-video text-red-400";
    default:
      return "fas fa-file text-gray-400";
  }
};

watch(() => currentBranch.value, fetchRepoData);

onMounted(() => {
  fetchRepoData();
});
</script>
<style scoped>
.prose {
  color: inherit;
}
input:focus {
  outline: none;
}
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
::-webkit-scrollbar-track {
  background: #1f2937;
}
::-webkit-scrollbar-thumb {
  background: #4b5563;
  border-radius: 4px;
}
::-webkit-scrollbar-thumb:hover {
  background: #6b7280;
}
</style>
