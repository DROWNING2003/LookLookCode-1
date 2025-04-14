<template>
  <div class="animate-pulse">
    <!-- 搜索框骨架 -->
    <div class="mb-4">
      <div class="relative">
        <div class="w-full h-12 bg-gray-700 rounded-xl"></div>
      </div>
    </div>

    <!-- 分类标签骨架 -->
    <div class="card mb-4 flex gap-4">
      <div class="h-8 w-20 bg-gray-700 rounded"></div>
      <div class="h-8 w-20 bg-gray-700 rounded"></div>
    </div>

    <!-- 仓库卡片骨架 -->
    <div class="grid grid-cols-1 gap-4">
      <div v-for="n in count" :key="n" class="p-4 bg-gray-800/80 rounded-xl">
        <div class="flex items-center gap-3 mb-3">
          <div class="h-5 w-5 bg-gray-700 rounded"></div>
          <div class="h-5 w-48 bg-gray-700 rounded"></div>
        </div>
        <div class="h-4 w-3/4 bg-gray-700 rounded mb-2"></div>
        <div class="h-4 w-1/2 bg-gray-700 rounded"></div>
        <div class="flex gap-4 mt-4">
          <div class="h-4 w-16 bg-gray-700 rounded"></div>
          <div class="h-4 w-16 bg-gray-700 rounded"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
defineProps<{
  count?: number
}>();
</script>