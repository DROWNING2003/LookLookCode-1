<!-- <!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>开源仓库</title>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Pacifico&display=swap" rel="stylesheet">
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css" rel="stylesheet">
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        tailwind.config = {
            theme: {
                extend: {
                    colors: {
                        primary: '#1a1a1a',
                        secondary: '#2d2d2d'
                    },
                    borderRadius: {
                        'none': '0px',
                        'sm': '8px',
                        DEFAULT: '12px',
                        'md': '16px',
                        'lg': '20px',
                        'xl': '24px',
                        '2xl': '28px',
                        '3xl': '32px',
                        'full': '9999px',
                        'button': '12px'
                    }
                }
            }
        }
    </script>
    <style>
        body {
            background-color: #121212;
            color: #ffffff;
        }
        .glass-effect {
            background: rgba(26, 26, 26, 0.8);
            backdrop-filter: blur(12px);
        }
        .card-glass {
            background: rgba(45, 45, 45, 0.7);
            backdrop-filter: blur(8px);
        }
    </style>
</head>
<body class="min-h-screen">
    <div class="relative w-[375px] mx-auto min-h-screen pb-[64px]">
        <div class="fixed top-0 w-full z-50">
            <div class="glass-effect px-4 py-3 flex items-center gap-3">
                <button class="text-white">
                    <i class="fas fa-bars text-xl"></i>
                </button>
                <div class="flex-1 relative">
                    <input type="text" placeholder="搜索仓库" class="w-full bg-[#2d2d2d] text-white px-4 py-2 !rounded-button text-sm border-none outline-none">
                    <i class="fas fa-search absolute right-3 top-1/2 -translate-y-1/2 text-gray-400"></i>
                </div>
                <button class="text-white">
                    <i class="fas fa-ellipsis-v text-xl"></i>
                </button>
            </div>
        </div>

        <div class="relative h-[200px] overflow-hidden">
            <img src="https://ai-public.mastergo.com/ai/img_res/1252829a77a67f4f395ee92e3dd72098.jpg" class="w-full h-full object-cover" alt="banner">
            <div class="absolute inset-0 bg-gradient-to-b from-transparent to-[#121212]"></div>
        </div>

        <div class="px-4 mt-4">
            <div class="flex items-center justify-between mb-6">
                <h2 class="text-xl font-semibold">热门仓库</h2>
                <button class="text-sm text-gray-400">
                    <i class="fas fa-sort-amount-down mr-1"></i>
                    排序
                </button>
            </div>

            <div class="space-y-4">
                <div class="card-glass p-4 !rounded-button">
                    <div class="flex justify-between items-start">
                        <h3 class="text-lg font-medium">Vue3-Admin-Template</h3>
                        <div class="flex items-center">
                            <i class="fas fa-star text-yellow-400"></i>
                            <span class="ml-1">328</span>
                        </div>
                    </div>
                    <p class="text-sm text-gray-400 mt-2">基于 Vue3 + TypeScript 的后台管理系统模板</p>
                    <div class="flex items-center mt-3 text-sm text-gray-500">
                        <i class="fas fa-code-branch mr-1"></i>
                        <span>42</span>
                        <span class="mx-2">•</span>
                        <span>2 天前更新</span>
                    </div>
                </div>

                <div class="card-glass p-4 !rounded-button">
                    <div class="flex justify-between items-start">
                        <h3 class="text-lg font-medium">React-Native-Shop</h3>
                        <div class="flex items-center">
                            <i class="fas fa-star text-yellow-400"></i>
                            <span class="ml-1">156</span>
                        </div>
                    </div>
                    <p class="text-sm text-gray-400 mt-2">跨平台电商 App 解决方案</p>
                    <div class="flex items-center mt-3 text-sm text-gray-500">
                        <i class="fas fa-code-branch mr-1"></i>
                        <span>28</span>
                        <span class="mx-2">•</span>
                        <span>5 天前更新</span>
                    </div>
                </div>

                <div class="card-glass p-4 !rounded-button">
                    <div class="flex justify-between items-start">
                        <h3 class="text-lg font-medium">Flutter-UI-Kit</h3>
                        <div class="flex items-center">
                            <i class="fas fa-star text-yellow-400"></i>
                            <span class="ml-1">89</span>
                        </div>
                    </div>
                    <p class="text-sm text-gray-400 mt-2">精美的 Flutter UI 组件库</p>
                    <div class="flex items-center mt-3 text-sm text-gray-500">
                        <i class="fas fa-code-branch mr-1"></i>
                        <span>15</span>
                        <span class="mx-2">•</span>
                        <span>1 周前更新</span>
                    </div>
                </div>
            </div>
        </div>

        <nav class="fixed bottom-0 w-full glass-effect">
            <div class="grid grid-cols-5 py-2">
                <a class="flex flex-col items-center text-white">
                    <i class="fas fa-home text-xl mb-1"></i>
                    <span class="text-xs">首页</span>
                </a>
                <a class="flex flex-col items-center text-gray-400">
                    <i class="fas fa-compass text-xl mb-1"></i>
                    <span class="text-xs">发现</span>
                </a>
                <a class="flex flex-col items-center text-gray-400">
                    <i class="fas fa-plus-circle text-2xl"></i>
                </a>
                <a class="flex flex-col items-center text-gray-400">
                    <i class="fas fa-bell text-xl mb-1"></i>
                    <span class="text-xs">通知</span>
                </a>
                <a class="flex flex-col items-center text-gray-400">
                    <i class="fas fa-user text-xl mb-1"></i>
                    <span class="text-xs">我的</span>
                </a>
            </div>
        </nav>
    </div>
</body>
</html> -->

<template>
  <div class="min-h-screen bg-gradient-to-b from-[#121212] to-[#1a1a1a] overflow-auto flex flex-col pt-16 pb-20 max-w-7xl mx-auto w-full">
  
      <!-- 标题和排序区域 -->
      <div class="relative w-full h-[200px] overflow-hidden rounded-b-[32px] mb-8">
        <img src="../../public/header.png" alt="" class="w-full h-full object-cover">
        <div class="absolute inset-0 bg-gradient-to-b from-transparent via-[#121212]/50 to-[#121212]">
          <div class="absolute bottom-8 left-8">
            <h1 class="text-3xl font-bold text-white mb-2">看看代码</h1>
            <p class="text-gray-300">学习优秀的开源项目</p>
          </div>
        </div>
      </div>
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
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 px-4">
        <RepoCard v-for="repo in repos" :key="repo.id" :repo="repo" />
      </div>

      <!-- 分页组件 -->
      <div class="flex justify-center mt-12 mb-8">
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
