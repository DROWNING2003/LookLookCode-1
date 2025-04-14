<!-- 代码已包含 CSS：使用 TailwindCSS , 安装 TailwindCSS 后方可看到布局样式效果 -->
<template>
  <div
    class="h-screen overflow-auto flex flex-col pt-16 pb-20 px-4 max-w-7xl mx-auto w-full"
  >
    <DiscoverSkeletonLoader v-if="isLoading" />
    <!-- 热门推荐 -->
    <section class="flex flex-col items-center mb-8">
      <h2 class="text-white text-lg font-medium mb-4 w-full">热门推荐</h2>
      <div class="grid grid-cols-1 gap-4">
        <div
          v-for="repo in hotRepos"
          :key="repo.id"
          class="flex flex-col p-4 rounded-xl transition-colors duration-200 card hover:card-hover"
          @click="
            () => {
              router.push(`/repo/${repo.full_name}`);
            }
          "
        >
          <div class="flex items-start gap-3">
            <img
              :src="`https://github.com/${repo.full_name.split('/')[0]}.png`"
              alt="avatar"
              class="w-10 h-10 rounded-full object-cover flex-shrink-0"
            />
            <div class="flex-1 min-w-0">
              <h3 class="text-white font-medium mb-1 truncate">
                {{ repo.name }}
              </h3>
              <p class="text-gray-400 text-sm mb-3 line-clamp-2">
                {{ repo.description }}
              </p>
              <div class="flex items-center gap-4 flex-wrap">
                <span class="text-gray-400 text-sm flex items-center gap-1">
                  <i class="fas fa-star"></i>
                  {{ repo.stargazers_count }}
                </span>
                <span class="flex items-center gap-1">
                  <span
                    :class="[
                      'w-2 h-2 rounded-full',
                      getLanguageColor(repo.language),
                    ]"
                  ></span>
                  <span class="text-gray-400 text-sm">{{ repo.language }}</span>
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    <!-- 趋势榜单 -->
    <section class="flex-1 mb-8">
      <h2 class="text-white text-lg font-medium mb-4">趋势榜单</h2>
      <div class="grid gap-4 grid-cols-1 md:grid-cols-2 xl:grid-cols-3">
        <div
          v-for="repo in trendingRepos"
          :key="repo.id"
          class="flex flex-col p-4 rounded-xl transition-colors duration-200 card hover:card-hover"
        >
          <div class="flex items-center justify-between mb-2">
            <h3 class="text-white font-medium truncate flex-1 mr-4">
              {{ repo.name }}
            </h3>
            <span
              class="text-accent text-sm flex items-center gap-1 flex-shrink-0"
            >
              <i class="fas fa-arrow-up"></i>
              {{ repo.growth }}%
            </span>
          </div>
          <p class="text-gray-400 text-sm line-clamp-2">
            {{ repo.description }}
          </p>
        </div>
      </div>
    </section>
    <!-- 个性化推荐 -->
    <section class="flex-1">
      <h2 class="text-white text-lg font-medium mb-4">为你推荐</h2>
      <div class="flex flex-wrap gap-2 mb-4">
        <button
          v-for="tag in tags"
          :key="tag"
          class="px-3 py-1 rounded-full text-sm transition-colors duration-200"
          :class="[
            selectedTag === tag
              ? 'bg-[var(--primary-color)] text-[var(--text-primary)]'
              : 'bg-[var(--card-layer)] text-[var(--text-secondary)] hover:bg-[var(--primary-color)]/10'
          ]"
          @click="handleTagSelect(tag)"
        >
          {{ tag }}
        </button>
      </div>
      <div class="flex overflow-x-auto -mx-4 px-4 scrollbar-none">
        <div class="inline-flex gap-4 pb-4 flex-nowrap">
          <div
            v-for="repo in recommendedRepos"
            :key="repo.id"
            class="flex flex-col w-[280px] sm:w-[320px] md:w-[280px] lg:w-[320px] flex-shrink-0 card rounded-lg overflow-hidden hover:card-hover transition-colors duration-200"
          >
            <img
              :src="`https://opengraph.githubassets.com/1/${repo.full_name}`"
              class="w-full h-40 object-cover"
              :alt="repo.name"
            />
            <div class="p-4">
              <h3 class="text-white font-medium mb-1 truncate">
                {{ repo.name }}
              </h3>
              <p class="text-gray-400 text-sm mb-3 line-clamp-2">
                {{ repo.description }}
              </p>
              <div class="flex items-center justify-between">
                <span class="text-gray-400 text-sm flex items-center gap-1">
                  <i class="fas fa-star"></i>
                  {{ repo.stargazers_count }}
                </span>
                <span
                  class="text-accent text-sm hover:text-accent-hover cursor-pointer"
                  >查看详情</span
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script lang="ts" setup>
import { ref, onMounted } from "vue";
// import { useThemeStore } from "../stores/theme";
import DiscoverSkeletonLoader from "../components/DiscoverSkeletonLoader.vue";

import { searchRepositories } from "../api/repo";
import type { Repository } from "../api/repo";
import { subDays, format } from "date-fns";
import { useRouter } from "vue-router";
const router = useRouter();
const tags = ["推荐", "Swift", "Python", "JavaScript", "Java", "Go", "Rust"];
const selectedTag = ref("推荐");
const hotRepos = ref<Repository[]>([]);
const trendingRepos = ref<(Repository & { growth: number })[]>([]);
const recommendedRepos = ref<Repository[]>([]);
const isLoading = ref(true);

// 获取热门仓库
const fetchHotRepos = async () => {
  try {
    const lastWeek = format(subDays(new Date(), 31), "yyyy-MM-dd");

    const response = await searchRepositories({
      q: `stars:>1000 created:>${lastWeek}`,
      sort: "stars",
      order: "desc",
      per_page: 10,
    });
    hotRepos.value = response.data.items;
  } catch (error) {
    console.error("获取热门仓库失败:", error);
  }
};

// 获取趋势榜单
const fetchTrendingRepos = async () => {
  try {
    const response = await searchRepositories({
      q: "created:>2024-01-01",
      sort: "stars",
      order: "desc",
      per_page: 2,
    });
    trendingRepos.value = response.data.items.map((repo: any) => ({
      ...repo,
      growth: Math.floor(Math.random() * 200) + 50, // 模拟增长率
    }));
  } catch (error) {
    console.error("获取趋势榜单失败:", error);
  }
};

// 获取推荐仓库
const fetchRecommendedRepos = async () => {
  try {
    const response = await searchRepositories({
      q:
        selectedTag.value === "推荐"
          ? "good-first-issues:>0"
          : `language:${selectedTag.value.toLowerCase()}`,
      sort: "updated",
      order: "desc",
      per_page: 3,
    });
    recommendedRepos.value = response.data.items;
  } catch (error) {
    console.error("获取推荐仓库失败:", error);
  }
};

// 处理标签选择
const handleTagSelect = async (tag: string) => {
  selectedTag.value = tag;
  await fetchRecommendedRepos();
};

// 获取编程语言对应的颜色
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

onMounted(async () => {
  // themeStore.initTheme();
  try {
    await Promise.all([
      fetchHotRepos(),
      fetchTrendingRepos(),
      fetchRecommendedRepos(),
    ]);
  } catch (error) {
    console.error("数据加载失败:", error);
  } finally {
    isLoading.value = false;
  }
});
</script>
<style scoped>
.scrollbar-none {
  scrollbar-width: none;
  -ms-overflow-style: none;
}
.scrollbar-none::-webkit-scrollbar {
  display: none;
}
</style>
