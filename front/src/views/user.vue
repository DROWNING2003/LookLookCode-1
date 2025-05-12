<template>
  <div
    class="min-h-screen bg-gradient-to-b from-[#121212] to-[#1a1a1a] overflow-auto flex flex-col pt-14 pb-20 mx-auto w-screen">
    <!-- header -->
    <div class="relative w-full h-[200px] overflow-hidden">
      <img src="/header.png" class="w-full h-full object-cover" />
      <div class="absolute bottom-0 left-0 w-full p-4">
        <div class="flex flex-row items-end">
          <img :src="userStore.avatar_url" class="w-16 h-16 rounded-full border-2 border-white" />
          <div class="ml-4 flex flex-col">
            <div class="flex flex-row items-center">
              <h1 class="text-xl font-semibold">{{ userStore.name }}</h1>
              <span class="ml-2 px-2 py-0.5 text-xs bg-blue-700 rounded-full">Pro
              </span>
            </div>
            <p class="text-sm text-gray-300 mt-1">
              {{ userStore.bio }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- 基本信息 -->
    <div class="px-4 py-6 grid grid-cols-4 gap-4 mb-6 ">
      <div class="text-center" v-for="(stat, index) in stats" :key="index">
        <div class="text-lg font-semibold text-gray-100">
          {{ stat.value }}
        </div>
        <div class="text-sm text-gray-500">{{ stat.label }}</div>
      </div>
    </div>

    <!-- 个人仓库 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 px-4">
      <RepoCard v-for="repo in repos" :key="repo.id" :repo="repo" />
    </div>

    <!-- 分页组件 -->
    <div class="flex w-full justify-center mt-12 mb-8">
      <Pagination v-if="total > 0" v-model:currentPage="currentPage" v-model:pageSize="pageSize" :total="total"
        class="backdrop-blur-lg bg-white/5 p-2 rounded-lg border border-white/10" />
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useUserStore } from "../stores/user";
import { ref, watch, onMounted } from "vue";
import Pagination from "../baseComponent/Pagination.vue";
import RepoCard from '../components/RepoCard.vue';
// import { useThemeStore } from "../stores/theme";
// import { useRouter } from "vue-router";
import {
  listUserRepositories,
  listUserStarredRepositories,
  type Repository,
} from "../api/repo";
const repos = ref<Repository[]>([]);
const currentPage = ref(1);
const selectedTab = ref(0);
const pageSize = ref(10);
const total = ref(0);
const loading = ref(false);
const error = ref("");

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
const userStore = useUserStore();
// const themeStore = useThemeStore();
// const router = useRouter();
const stats = ref([
  { value: '128', label: '仓库' },
  { value: '1.2k', label: 'Star' },
  { value: '328', label: '关注' },
  { value: '892', label: '粉丝' }
]);

// const formatDate = (dateString: string) => {
//   const date = new Date(dateString);
//   return (
//     date.getFullYear() +
//     "年" +
//     (date.getMonth() + 1) +
//     "月" +
//     date.getDate() +
//     "日"
//   );
// };

// const handleLogout = () => {
//   userStore.clearUser();
//   router.push("/login");
// };
</script>
<style scoped>
.github-profile {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
    Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
}
</style>
