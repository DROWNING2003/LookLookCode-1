<!-- 代码已包含 CSS：使用 TailwindCSS , 安装 TailwindCSS 后方可看到布局样式效果 -->
<template>
  <main class="h-screen overflow-auto flex flex-col pt-16 pb-20 px-4 max-w-7xl mx-auto w-full">
    <!-- 个人信息头部 -->
    <section class="card text-center mb-8">
      <div class="relative w-28 h-28 mx-auto mb-4">
        <img
          :src="userStore.avatar_url"
          alt="用户头像"
          class="rounded-full w-full h-full object-cover border-4 border-gray-700 border-dashed"
        />
        <div
          class="absolute -bottom-2 -right-2 bg-green-500 w-6 h-6 rounded-full flex items-center justify-center"
        >
          <i class="fas fa-check text-sm"></i>
        </div>
      </div>
      <h1 class="title-h1">
        {{ userStore.name }}
      </h1>
      <p class="title-h2 mb-3">@{{ userStore.login }}</p>
      <p class="title-h2 mb-4">"{{ userStore.bio }}"</p>
      <!-- 关注数据 -->
      <section class="mb-6 p-4">
        <div class="grid grid-cols-2 gap-4">
          <div class="text-center">
            <p class="title-h1 font-bold">{{ userStore.followers }}</p>
            <p class="title-h2">关注者</p>
          </div>
          <div class="text-center">
            <p class="title-h1 font-bold">{{ userStore.following }}</p>
            <p class="title-h2">正在关注</p>
          </div>
        </div>
      </section>
    </section>

    <!-- 基本信息卡片 -->
    <section class="card title-h2 p-4 mb-6 space-y-4">
      <div class="flex items-center">
        <i class="fas fa-building text-gray-400 w-6"></i>
        <span class="ml-3">{{ userStore.company }}</span>
      </div>
      <div class="flex items-center">
        <i class="fas fa-map-marker-alt text-gray-400 w-6"></i>
        <span class="ml-3">{{ userStore.location }}</span>
      </div>
      <div class="flex items-center">
        <i class="fas fa-link text-gray-400 w-6"></i>
        <a :href="userStore.blog ?? undefined" class="ml-3 text-blue-400">{{
          userStore.blog ?? ""
        }}</a>
      </div>
      <div class="flex items-center">
        <i class="fas fa-calendar text-gray-400 w-6"></i>
        <span class="ml-3">加入于 {{ formatDate(userStore.created_at) }}</span>
      </div>
    </section>
    <!-- 仓库统计 -->
    <section class="card mb-6">
      <div class="flex justify-between items-center">
        <div>
          <h3 class="text-lg font-semibold mb-1">公开仓库</h3>
          <p class="text-3xl font-bold">{{ userStore.public_repos }}</p>
        </div>
        <button
          class="!rounded-button bg-accent px-4 py-2 hover:bg-accent-hover"
        >
          查看全部
        </button>
      </div>
    </section>

    <!-- 退出登录按钮 -->
    <section class="card">
      <BButton text="退出登录" @click="handleLogout" />
    </section>
  </main>
</template>
<script lang="ts" setup>
import { useUserStore } from "../stores/user";
// import { useThemeStore } from "../stores/theme";
import { useRouter } from "vue-router";
import BButton from "../baseComponent/BButton.vue";
const userStore = useUserStore();
// const themeStore = useThemeStore();
const router = useRouter();

const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return (
    date.getFullYear() +
    "年" +
    (date.getMonth() + 1) +
    "月" +
    date.getDate() +
    "日"
  );
};

const handleLogout = () => {
  userStore.clearUser();
  router.push("/login");
};


</script>
<style scoped>
.github-profile {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
    Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
}
</style>
