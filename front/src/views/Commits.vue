<template>
  <div
    class="min-h-screen bg-gradient-to-b bg-background overflow-auto flex flex-col pt-14 pb-20 mx-auto w-screen"
  >
    <main class="pt-8 pb-12 px-4 max-w-5xl mx-auto w-full">
      <div class="mb-8" v-if="repository">
        <div class="flex items-center gap-2 mb-2">
          <button
            @click="router.back()"
            class="flex items-center gap-2 text-gray-300 hover:text-white transition-colors"
          >
            <i class="text-gray-300 hover:text-white fas fa-arrow-left"></i>
          </button>
          <h1 class="text-xl font-semibold text-blue-400">
            {{ repository.full_name }}{{ path ? "/" + path : "" }}
          </h1>
        </div>
      </div>
    </main>
  </div>
</template>
<script setup lang="ts">
// API 和类型定义导入
import { onMounted, ref, computed } from "vue";
import {
  getRepository,
  //   getRepositoryReadme,
  //   getRepositoryBranches,
  //   getRepositoryContributors,
  getRepositoryContents,
  // getRepositoryCommits,
  type Repository,
  //   type Branch,
  type FileContent,
  //   type Commit,
} from "../api/repo";
const repository = ref<Repository>(); // 仓库基本信息
import { useRoute, useRouter } from "vue-router";
const router = useRouter();
const route = useRoute();
const path = Array.isArray(route.params.path)
  ? route.params.path.join("/")
  : route.params.path || "";
</script>
<style></style>
