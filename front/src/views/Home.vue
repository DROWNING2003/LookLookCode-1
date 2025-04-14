<template>
  <background
    class="absolute blur-(--blur) z-10 top-0 left-0 w-full min-h-screen bg-center bg-contain bg-(image:--background-image) bg-fixed"
  >
  </background>
  <div class="continer">
    <TopNavBar class="fixed top-0" :title="pageTitle" />
    <Suspense>
      <template #default>
        <RouterView class="overflow-auto" />
      </template>
      <template #fallback>
        <HomeSkeletonLoader />
      </template>
    </Suspense>
    <BottomNavBar class="fixed bottom-0" />
  </div>
</template>
<script setup lang="ts">
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import TopNavBar from "../components/TopNavBar.vue";
import BottomNavBar from "../components/BottomNavBar.vue";
import HomeSkeletonLoader from "../components/HomeSkeletonLoader.vue";

const route = useRoute();
const pageTitle = computed(() => {
  switch (route.path) {
    case '/':
      return '主页';
    case '/discover':
      return '发现';
    case '/user':
      return '用户';
    default:
      return '主页';
  }
});

</script>
