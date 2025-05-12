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
        <div class="flex items-center gap-2 mb-4"></div>
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
              @click="navigateToFile(file, owner, repo, branch, file.path)"
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
<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
// import {navigateToFile} from "../utils/repo";
import { getFileIcon } from "../utils/icon";
import { onMounted, ref, watch } from "vue";
import { MdPreview } from "md-editor-v3";
import "md-editor-v3/lib/preview.css";
import {
  getRepositoryReadmeDir,
  getRepository,
  type Repository,
  getDirectoryContentsWithCommits,
} from "../api/repo";
const router = useRouter();
const repository = ref<Repository>(); // 仓库基本信息
const files = ref(); // 文件列表
const readme = ref(""); // 仓库 README 内容

const route = useRoute();
const path = Array.isArray(route.params.path)
  ? route.params.path.join("/")
  : route.params.path || "";
const branch = route.params.branch as string;
const repo = route.params.repo as string;
const owner = route.params.owner as string;
// console.log(path,branch,repo,owner);

onMounted(() => {
  getRepository(owner, repo).then((res) => {
    repository.value = res.data;
  });
  getDirectoryContentsWithCommits(owner, repo, path, branch).then((res) => {
    files.value = res.sort((a, b) => {
      if (a.type === "file" && b.type === "dir") {
        return 1; // a 排在 b 前面
      } else if (a.type === "dir" && b.type === "file") {
        return -1; // b 排在 a 前面
      } else {
        return 0; // 保持原顺序
      }
    });
  });
  getRepositoryReadmeDir(owner, repo, path).then((res) => {
    readme.value = res;
  });
});

watch(
  () => route.params.path,
  (newPath) => {
    const path = Array.isArray(newPath) ? newPath.join("/") : newPath || "";
    getDirectoryContentsWithCommits(owner, repo, path, branch).then((res) => {
      files.value = res;
    });
  }
);

const navigateToFile = (
  file: any,
  owner: any,
  repo: any,
  branch: any,
  path: string
) => {
  // console.log("文件详情:", file);
  // 如果是文件，获取文件内容
  if (file.type === "file") {
    router
      .push({
        name: "RepoBlob",
        params: { owner, repo, branch, path },
        force: true,
      })
      .catch((err) => {
        console.error("Error fetching directory contents:", err);
        // error.value = "获取目录内容失败，请稍后重试";
      });
    // getRepositoryContents(owner, repo, file.path as string, branch.value)
    //   .then((response) => {
    //     const content = response.data;
    //     console.log("文件内容:", content);
    //   })
    //   .catch((err) => {
    //     console.error("获取文件内容失败:", err);
    //     error.value = "获取文件内容失败，请稍后重试";
    //   })
  }

  if (file.type === "dir") {
    // 如果是目录，更新当前路径并重新获取内容
    router
      .push({
        name: "RepoTree",
        params: { owner, repo, branch, path },
        force: true,
      })
      .catch((err) => {
        console.error("Error fetching directory contents:", err);
        // error.value = "获取目录内容失败，请稍后重试";
      });
    //   router.replace({
    //     name: "RepoTree",
    //     params: { owner, repo,branch, path},
    //   })
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
</script>
