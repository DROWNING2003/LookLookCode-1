<template>
  <div class="min-h-screen bg-gradient-to-b from-[#121212] to-[#1a1a1a] overflow-auto flex flex-col pt-14 pb-20 mx-auto w-full">
  
      <!-- header&banner -->
      <Banner :banners="[
        {
          image: '/header.png',
          title: '看看代码',
          description: '学习优秀的开源项目',
          url: 'https://github.com'
        },
        {
          image: '/default.jpg',
          title: '看看代码2',
          description: '学习优秀的开源项目',
          url: 'http://localhost:8081/user'
        },
        {
          image: '/vite.svg',
          title: '广告位1',
          description: '',
          url: 'https://vitejs.dev'
        },
      ]" />

      <!-- 最近仓库 -->
      <div class="flex items-center justify-between mb-8 px-4">
        <h2 class="text-2xl font-bold bg-gradient-to-r from-white to-gray-300 bg-clip-text text-transparent">最近仓库</h2>
        <div class="flex items-center gap-4">
          <button class="flex items-center gap-2 px-4 py-2 rounded-lg bg-white/5 backdrop-blur-lg text-white hover:bg-white/10 transition-all duration-300 border border-white/10">
            <i class="fas fa-sort-amount-down"></i>
            <span>最近更新</span>
          </button>
        </div>
      </div>  

      <!-- 仓库列表 -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 px-4">
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
import Pagination from "../baseComponent/Pagination.vue";
import RepoCard from '../components/RepoCard.vue';
import Banner from '../components/Banner.vue';
import {
  listUserRepositories,
  listUserStarredRepositories,
  type Repository,
} from "../api/repo";

const loading = ref(false);
const error = ref("");
const selectedTab = ref(0);
const repos = ref<Repository[]>([]);
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);

async function loadData(index: number = 0) {
  loading.value = true;
  error.value = "";
  try {
    const response = await (index === 0
      ? listUserRepositories({
        sort: "updated",
        direction: "desc",
        per_page: pageSize.value,
        page: currentPage.value,
      })
      : listUserStarredRepositories({
        sort: "updated",
        direction: "desc",
        per_page: pageSize.value,
        page: currentPage.value,
      }));
    repos.value = response.data;
    // 从响应头获取总数
    const linkHeader = response.headers.link;
    if (linkHeader) {
      const lastLink = linkHeader
        .split(",")
        .find((link: string | string[]) => link.includes('rel="last"'));
      if (lastLink) {
        const match = lastLink.match(/[&?]page=(\d+)/i);
        if (match) {
          total.value = Number(match[1]) * pageSize.value;
        }
      }
    } else {
      total.value = repos.value.length;
    }
  } catch (e) {
    error.value = "加载仓库失败，请稍后重试";
    console.error("Failed to load repositories:", e);
  } finally {
    loading.value = false;
  }
}

// 监听分页变化
watch([currentPage, pageSize], () => {
  loadData(selectedTab.value);
});

// 切换标签时重置分页
watch(selectedTab, () => {
  currentPage.value = 1;
  loadData(selectedTab.value);
});

onMounted(() => {
  loadData(selectedTab.value);
});
</script>

<style scoped>
.repo-card {
  transition: all 0.3s ease;
}

.repo-card:hover {
  transform: translateY(-2px);
}

input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>
