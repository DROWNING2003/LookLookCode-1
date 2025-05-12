
<template>
   <div class="min-h-screen bg-gradient-to-b from-[#121212] to-[#1a1a1a] overflow-auto flex flex-col pt-14 pb-20 mx-auto w-screen">
    <div class="flex items-center justify-between mt-1 mb-8 px-4">
        <div class="flex items-center gap-4">
          <h2 class="text-2xl font-bold bg-gradient-to-r from-white to-gray-300 bg-clip-text text-transparent">
            {{ searchQuery ? '搜索结果' : '热门仓库' }}
          </h2>
          <span v-if="searchQuery" class="text-white/60 text-sm">{{ searchQuery }}</span>
        </div>
        <div class="flex items-center gap-4">
          <button class="flex items-center gap-2 px-4 py-2 rounded-lg bg-white/5 backdrop-blur-lg text-white hover:bg-white/10 transition-all duration-300 border border-white/10">
            <i class="fas fa-sort-amount-down"></i>
            <span>最近更新</span>
          </button>
        </div>
      </div>  
        <!-- 仓库列表 -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 px-4">
        <RepoCard v-for="repo in repos" :key="repo.id" :repo="repo" />
      </div>

      <!-- 分页组件 -->
      <div class="flex  justify-center mt-12 mb-8">
        <Pagination 
          v-if="total > 0" 
          v-model:currentPage="currentPage" 
          v-model:pageSize="pageSize" 
          :total="total"
          class="backdrop-blur-lg bg-white/5 p-2 rounded-lg border border-white/10" 
        />
      </div>
   </div>
</template>
<script lang="ts" setup>
import { ref, onMounted, watch } from "vue";
import { useRoute } from "vue-router";
import Pagination from "../baseComponent/Pagination.vue";
import RepoCard from '../components/RepoCard.vue';
import {
  searchRepositories,
  type Repository,
} from "../api/repo";

const route = useRoute();
const loading = ref(false);
const error = ref("");
const repos = ref<Repository[]>([]);
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);

const searchQuery = ref("");

watch(() => route.params.query, (newQuery) => {
  searchQuery.value = typeof newQuery === 'string' ? decodeURIComponent(newQuery) : '';
  currentPage.value = 1; // 重置页码
  loadData();
}, { immediate: true });

async function loadData() {
  loading.value = true;
  error.value = "";
  try {
    const oneWeekAgo = new Date();
    oneWeekAgo.setDate(oneWeekAgo.getDate() - 7);
    const dateFilter = oneWeekAgo.toISOString().split('T')[0];
    
    const query = searchQuery.value
      ? searchQuery.value
      : `pushed:>${dateFilter} stars:>100`; // 没有搜索词时显示最近一周内有100+ stars的热门仓库
    
    const response = await searchRepositories({
      q: query,
      sort: "stars",
      order: "desc",
      per_page: pageSize.value,
      page: currentPage.value,
    });
    repos.value = response.data.items;
    total.value = Math.min(response.data.total_count, 1000); // GitHub API限制最多返回1000条结果
  } catch (e) {
    error.value = "加载仓库失败，请稍后重试";
    console.error("Failed to load repositories:", e);
  } finally {
    loading.value = false;
  }
}

// 监听分页变化
watch([currentPage, pageSize], () => {
  loadData();
});

onMounted(() => {
  loadData();
});
</script>
<style scoped>
</style>
