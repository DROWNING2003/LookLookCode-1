<template>
  <!-- {{path}} -->
  <div
    class="min-h-screen bg-gradient-to-b bg-background overflow-auto flex flex-col pt-14 pb-20 mx-auto w-screen"
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
            {{ repository.full_name }}{{ path ? "/" + path : "" }}
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
      </div>

      <div class="mt-8 bg-card rounded-lg">
        <BButton @click="ai" text="ai 分析" />
        <MdPreview
          class="p-4 bg-[#121212]"
          id="preview-only"
          :modelValue="code"
        />
        <div v-for="m in messages" :key="m.id">
          <div v-if="m.role != 'user'">
            <MdPreview
              class="p-4 bg-[#121212]"
              id="preview-only"
              :modelValue="m.content"
            />
          </div>
        </div>
      </div>
    </main>
  </div>
</template>
<script setup lang="ts">
import BButton from "@/baseComponent/BButton.vue";
import { useRoute, useRouter } from "vue-router";
import { onMounted, ref, computed } from "vue";
const router = useRouter();
import { MdPreview } from "md-editor-v3";
import "md-editor-v3/lib/preview.css";
// API 和类型定义导入
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
import { Base64 } from "js-base64";
//ai
import { useChat } from "@ai-sdk/vue";
import { useThemeStore } from "../stores/theme";
const themeStore = useThemeStore();
const { messages, input, handleSubmit } = useChat({
  api: `http://${themeStore.VITE_BACK_URL}/api/chatOllama`,
});

function ai() {
  input.value = <string>file.value?.content;
  handleSubmit();
}

const code = computed(() => {
  if (!file.value?.content) return "";
  return "```" + file.value.name + "\n" + file.value.content + "\n```";
});
const route = useRoute();
const repository = ref<Repository>(); // 仓库基本信息
const file = ref<FileContent>(); // 文件列表
// const readme = ref(""); // 仓库 README 内容

const path = Array.isArray(route.params.path)
  ? route.params.path.join("/")
  : route.params.path || "";
const branch = route.params.branch as string;
const repo = route.params.repo as string;
const owner = route.params.owner as string;
onMounted(() => {
  getRepository(owner, repo).then((res) => {
    repository.value = res.data;
  });
  getRepositoryContents(owner, repo, path, branch)
    .then((response) => {
      const content = response.data;
      console.log("文件内容:", content);
      file.value = <FileContent>content;
      file.value.content = Base64.decode(<string>file.value.content);
    })
    .catch((err) => {
      console.error("获取文件内容失败:", err);
    });
});
</script>
<style>
.md-editor {
  --md-bk-color: var(--background);
  --md-color: #fcfeff;
}
</style>
