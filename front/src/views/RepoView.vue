<template>
  <div
    class="min-h-screen bg-gradient-to-b from-[#121212] to-[#1a1a1a] overflow-auto flex flex-col pt-14 pb-20 mx-auto w-screen"
  >
    <!-- 主要内容区域 -->
    <main class="pt-8 pb-12 px-4 max-w-5xl mx-auto w-full">
      <!-- 仓库信息部分 -->
      <div class="mb-8" v-if="repository">
        <div class="flex items-center gap-2 mb-2">
          <button
            @click="router.back()"
            class="flex items-center gap-2 text-gray-300 hover:text-white transition-colors"
          >
            <i class="text-gray-300 hover:text-white fas fa-arrow-left"></i>
          </button>
          <h1 class="text-xl font-semibold text-blue-400">
            {{ repository.full_name }}
          </h1>
        </div>
        <p class="text-gray-400 mb-4">{{ repository.description }}</p>

        <!-- 修改后的统计信息 -->
        <div class="flex gap-6 mb-4">
          <div class="flex items-center gap-1 text-gray-300">
            <i class="far fa-star"></i>
            <span>{{ repository.stargazers_count / 1000 }}k</span>
          </div>
          <div class="flex items-center gap-1 text-gray-300">
            <i class="fas fa-code-branch"></i>
            <span>{{ repository.forks_count }}</span>
          </div>
          <div class="flex items-center gap-1 text-gray-300">
            <i class="far fa-eye"></i>
            <span>{{ repository.watchers_count }}</span>
          </div>
        </div>
        <!-- 分支切换 -->
        <div class="flex items-center gap-2 mb-4">
          <BranchSelect
            @update:modelValue="(value:Branch)=>{
              currentBranch = value;
              fetchRepoData();
            }"
            v-model="currentBranch"
            :branches="branches"
          />
        </div>
        <!-- 文件列表部分 -->
        <div class="bg-card rounded-lg overflow-hidden">
          <div class="px-4 py-2 border-b border-gray-700 bg-gray-800">
            <span class="text-sm text-gray-400">代码</span>
          </div>
          <div
            v-for="file in files"
            :key="file.path"
            class="border-b border-gray-700 last:border-b-0"
          >
            <div
              class="flex items-center justify-between px-4 py-3 hover:bg-card-hover"
              @click="navigateToFile(file)"
            >
              <div class="flex items-center gap-3 flex-1">
                <i
                  :class="getFileIcon(file.type, file.name)"
                  class="text-gray-400"
                ></i>
                <div class="min-w-0">
                  <p class="text-sm font-medium text-gray-100">
                    {{ file.name }}
                  </p>
                  <p class="text-xs text-gray-400">{{ file.time }}</p>
                </div>
              </div>
              <span class="text-xs text-gray-400">
                <!-- {{ formatCompactDate(file.time) }} -->
              </span>
            </div>
          </div>
        </div>

        <!-- README 部分 -->
        <div class="mt-8 bg-card rounded-lg overflow-hidden">
          <div class="px-4 py-2 border-b border-gray-700 bg-gray-800">
            <span class="text-sm text-gray-400">README.md</span>
          </div>
          <MdPreview
            class="p-4"
            theme="dark"
            id="preview-only"
            :modelValue="readme"
          />
        </div>
      </div>
    </main>
  </div>
</template>

<script lang="ts" setup>
// Vue 核心功能导入
import BranchSelect from "@/baseComponent/BranchSelect.vue";
import { ref, watch, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { MdPreview } from "md-editor-v3";
import "md-editor-v3/lib/preview.css";
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

// 路由参数获取
const route = useRoute();
const router = useRouter();
const owner = route.params.owner as string;
const repo = route.params.repo as string;

// 状态管理
const currentPath = ref(""); // 当前浏览的文件路径

/**
 * 处理搜索事件
 * @todo 实现仓库内容搜索功能
 */
// const handleSearch = () => {
//   console.log("Search triggered");
// };

// 仓库数据相关的响应式状态
const repository = ref<Repository | null>(null); // 仓库基本信息
const readme = ref(""); // 仓库 README 内容
const branches = ref<Branch[]>([]); // 仓库分支列表
const contributors = ref<number>(0); // 贡献者数量
const currentBranch = ref<Branch>(); // 当前选中的分支
const files = ref<FileContent[]>([]); // 当前目录下的文件列表
// const commits = ref<Commit[]>([]); // 提交历史

// UI 状态控制
// const showBranchDropdown = ref(false); // 分支下拉菜单显示状态
const loading = ref(true); // 加载状态
const error = ref<string | null>(null); // 错误信息

/**
 * 获取仓库所有相关数据
 * 包括：基本信息、README、分支列表、贡献者、目录内容、提交历史
 */
const fetchRepoData = async () => {
  loading.value = true;
  error.value = null;
  //先获取分支
  getRepositoryBranches(owner, repo)
  .then((response) => {
    branches.value = response.data;
    currentBranch.value = branches.value.find(branch => branch.name === "main"||branch.name === "master") || branches.value[0];
  })
  try {
    const [
      repoData,
      readmeContent,
      contributorsData,
      contentsData,
      // commitsData,
    ] = await Promise.all([
      getRepository(owner, repo),
      getRepositoryReadme(owner, repo),
    
      getRepositoryContributors(owner, repo),
      getRepositoryContents(
        owner,
        repo,
        currentPath.value,
        currentBranch.value?.name
      ),
      // getRepositoryCommits(owner, repo, {
      //   per_page: 10,
      //   sha: currentBranch.value,
      // }),
    ]);

    repository.value = repoData.data;
    readme.value = readmeContent;
    contributors.value = contributorsData.data.length;

    // 为每个文件添加最后一次提交信息
    const filesWithCommitInfo = Array.isArray(contentsData.data)
      ? await Promise.all(
          contentsData.data.map(async (file) => {
            try {
              const fileCommits = await getRepositoryCommits(owner, repo, {
                per_page: 1,
                path: file.path,
                sha: currentBranch.value?.name,
              });
              if (fileCommits.data.length > 0) {
                const lastCommit = fileCommits.data[0];
                return {
                  ...file,
                  checked: file.name === "build.gradle",
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
    // commits.value = commitsData.data;
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
  console.log("文件详情:", file);
  // 如果是文件，获取文件内容
  if (file.type === "file") {
    // getRepositoryContents(owner, repo, file.path as string, currentBranch.value)
    //   .then((response) => {
    //     const content = response.data;
    //     console.log("文件内容:", content);
    //   })
    //   .catch((err) => {
    //     console.error("获取文件内容失败:", err);
    //     error.value = "获取文件内容失败，请稍后重试";
    //   });
  }
  if (file.type === "dir") {
    // 如果是目录，更新当前路径并重新获取内容
    router.push({
      name: "RepoTree",
      params: { owner, repo, branch: currentBranch.value?.name, path: file.path },
      force: true,
    });
    // currentPath.value = file.path as string;
    // console.log(currentPath.value);

    // getRepositoryContents(owner, repo, currentPath.value, currentBranch.value)
    //   .then((response) => {
    //     files.value = Array.isArray(response.data)
    //       ? response.data
    //       : [response.data];
    //   })
    //   .catch((err) => {
    //     console.error("Error fetching directory contents:", err);
    //     error.value = "获取目录内容失败，请稍后重试";
    //   });
  } else {
    // 获取文件内容
  }
};

/**
 * 返回上级目录
 * 通过分割当前路径并移除最后一个部分来实现
 */
// const goToParentDirectory = () => {
//   if (currentPath.value) {
//     const parentPath = currentPath.value.split("/").slice(0, -1).join("/");
//     currentPath.value = parentPath;
//     getRepositoryContents(owner, repo, parentPath, currentBranch.value)
//       .then((response) => {
//         files.value = Array.isArray(response.data)
//           ? response.data
//           : [response.data];
//       })
//       .catch((err) => {
//         console.error("Error fetching parent directory contents:", err);
//         error.value = "获取目录内容失败，请稍后重试";
//       });
//   }
// };
/**
 * 切换分支
 * @param branchName 目标分支名称
 */
// const switchBranch = async (branchName: string) => {
//   currentBranch.value = branchName;
//   showBranchDropdown.value = false;
//   await fetchRepoData();
// };

/**
 * 紧凑时间格式化（匹配图片显示）
 */
// const formatCompactDate = (dateString: string) => {
//   const date = new Date(dateString);
//   const now = new Date();
//   const diffDays = Math.floor((now - date) / (1000 * 60 * 60 * 24));

//   if (diffDays === 0) return "今天";
//   if (diffDays === 1) return "昨天";
//   if (diffDays < 7) return `${diffDays}天前`;
//   if (diffDays < 30) return `${Math.floor(diffDays / 7)}周前`;
//   return `${Math.floor(diffDays / 30)}月前`;
// };

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

// watch(() => currentBranch.value, fetchRepoData);

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

/* 添加复选框样式 */
.form-checkbox {
  border-color: #4b5563;
  background-color: #1f2937;
}
.form-checkbox:checked {
  border-color: #3b82f6;
  background-color: #3b82f6;
}
</style>
