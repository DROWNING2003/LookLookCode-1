<template>
  <div
    class="w-64 h-screen flex flex-col p-4 gap-x-2.5 gap-y-0 flex-wrap content-start"
  >
    <div class="">
      <i class="fas fa-home text-xl mb-1 text-gray-100"></i>
      <span
        class="font-roboto text-xl font-bold leading-7 tracking-normal text-gray-100"
      >
        looklookcode</span
      >
    </div>
    <nav class=" mt-16 font-roboto text-base font-normal leading-6 text-gray-100">
      <router-link
        to="/"
        class="w-[224px] rounded-[1.25rem] h-[48px] flex items-center p-3 gap-5 flex-wrap content-center"
        :class="
          currentRoute === '/'
            ? 'text-[var(--text-primary)] bg-gray-800'
            : 'text-[var(--text-secondary)]'
        "
      >
        <i class="fas fa-home text-xl mb-1"></i>
        <span>首页</span>
      </router-link>
      <router-link
        to="/discover"
        class="w-[224px] rounded-[1.25rem] h-[48px] flex items-center p-3 gap-5 flex-wrap content-center"
        :class="
          currentRoute === '/discover'
            ? 'text-[var(--text-primary)] bg-gray-800'
            : 'text-[var(--text-secondary)]'
        "
      >
        <i class="fas fa-compass text-xl mb-1"></i>
        <span>发现</span>
      </router-link>
      <router-link
        to="/notification"
        class="w-[224px] rounded-[1.25rem] h-[48px] flex items-center p-3 gap-5 flex-wrap content-center"
        :class="
          currentRoute === '/notification'
            ? 'text-[var(--text-primary)] bg-gray-800'
            : 'text-[var(--text-secondary)]'
        "
      >
        <i class="fas fa-bell text-xl mb-1"></i>
        <span>通知</span>
      </router-link>
      <router-link
        to="/user"
        class="w-[224px] rounded-[1.25rem] h-[48px] flex items-center p-3 gap-5 flex-wrap content-center"
        :class="
          currentRoute === '/user'
            ? 'text-[var(--text-primary)] bg-gray-800'
            : 'text-[var(--text-secondary)]'
        "
      >
        <i class="fas fa-user text-xl mb-1"></i>
        <span>我的</span>
      </router-link>
    </nav>
    <div class="fixed bottom-10 rounded-3xl bg-gray-800 w-56 h-[4.75rem] flex flex-row items-center p-4 gap-5 content-center">
      <img :src="userStore.avatar_url" class="rounded-[624.94rem] w-10 h-10" />
      <div class="flex flex-col truncate font-roboto text-base font-medium leading-6 ">
        <h1 class="text-gray-100">   {{ userStore.name }} </h1>
        <h1 class="font-roboto text-sm font-normal leading-5 text-gray-400"> {{ userStore.bio }}</h1>
      </div>
    </div>
  </div>
  <!-- 平板 -->
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { useRoute } from "vue-router";
import { useUserStore } from "../stores/user";
const userStore = useUserStore();
const route = useRoute();
const currentRoute = computed(() => route.path);
</script>
