<template>
  <nav
    class="bg-background-700 z-50 h-full flex flex-col items-center justify-around"
  >
    <router-link
      to="/"
      class="flex flex-col items-center flex-1"
      :class="currentRoute === '/' ? 'text-[var(--text-primary)]' : 'text-[var(--text-secondary)]'"
    >
      <i class="fas fa-home text-xl mb-1"></i>
      <span class="text-xs">首页</span>
    </router-link>
    <router-link
      to="/discover"
      class="flex flex-col items-center flex-1"
      :class="currentRoute === '/discover' ? 'text-[var(--text-primary)]' : 'text-[var(--text-secondary)]'"
    >
      <i class="fas fa-compass text-xl mb-1"></i>
      <span class="text-xs">发现</span>
    </router-link>
    <router-link
      to="/notification"
      class="flex flex-col items-center flex-1"
      :class="currentRoute === '/notification' ? 'text-[var(--text-primary)]' : 'text-[var(--text-secondary)]'"
    >
      <i class="fas fa-bell text-xl mb-1"></i>
      <span class="text-xs">通知</span>
    </router-link>
    <router-link
      to="/user"
      class="flex flex-col items-center flex-1"
      :class="currentRoute === '/user' ? 'text-[var(--text-primary)]' : 'text-[var(--text-secondary)]'"
    >
      <i class="fas fa-user text-xl mb-1"></i>
      <span class="text-xs">我的</span>
    </router-link>
  </nav>
  <!-- 平板 -->
</template>

<script lang="ts" setup>
import { computed } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const currentRoute = computed(() => route.path);
</script>