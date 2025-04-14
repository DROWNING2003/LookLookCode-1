<template>
  <div
    class="flex flex-col p-4 rounded-xl transition-all duration-200 bg-gray-800/80 hover:bg-gray-700/80 hover:transform hover:-translate-y-1 cursor-pointer"
    @click="handleClick"
  >
    <div class="flex items-start gap-3">
      <img
        :src="
          repo.owner?.avatar_url ||
          `https://github.com/${repo.full_name?.split('/')[0]}.png`
        "
        :alt="repo.name"
        class="w-10 h-10 rounded-full object-cover flex-shrink-0"
      />
      <div class="flex-1 min-w-0">
        <h3 class="text-white font-medium mb-1 truncate hover:text-blue-400">
          {{ repo.name }}
        </h3>
        <p class="text-gray-400 text-sm mb-3 line-clamp-2 hover:text-gray-300">
          {{ repo.description || "暂无描述" }}
        </p>
        <div class="flex items-center justify-between w-full">
          <div class="flex items-center gap-4">
            <span class="text-gray-400 text-sm flex items-center gap-1">
              <i class="fas fa-star"></i>
              {{ repo.stargazers_count || 0 }}
            </span>
            <span class="text-gray-400 text-sm flex items-center gap-1">
              <i class="fas fa-code-fork"></i>
              {{ repo.forks_count || 0 }}
            </span>
            <span class="text-gray-400 text-sm flex items-center gap-1">
              <i class="fas fa-history"></i>
              {{ formatDate(repo.updated_at) }} 更新
            </span>
          </div>
          <span class="flex items-center gap-1" v-if="repo.language">
            <span
              :class="['w-2 h-2 rounded-full', getLanguageColor(repo.language)]"
            ></span>
            <span class="text-gray-400 text-sm">{{ repo.language }}</span>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
import type { Repository } from "../api/repo";

const router = useRouter();

const props = defineProps<{
  repo: Repository;
}>();

const getLanguageColor = (language: string | null) => {
  const colors: Record<string, string> = {
    JavaScript: "bg-yellow-400",
    Python: "bg-blue-400",
    Java: "bg-red-400",
    Go: "bg-cyan-400",
    Rust: "bg-orange-400",
    Swift: "bg-pink-400",
  };
  return colors[language || ""] || "bg-gray-400";
};

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

const handleClick = () => {
  router.push(`/repo/${props.repo.full_name}`);
};
</script>
