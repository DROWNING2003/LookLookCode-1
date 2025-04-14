<template>
  <div class="animate-pulse">
    <!-- 顶部导航栏骨架 -->
    <div class="fixed top-0 left-0 right-0 h-16 bg-gray-800/80 z-50 flex items-center px-4">
      <div class="h-6 w-24 bg-gray-700 rounded"></div>
    </div>

    <!-- 主要内容区域骨架 -->
    <div class="h-screen overflow-auto flex flex-col pt-16 pb-20 px-4 max-w-7xl mx-auto w-full">
      <!-- 内容区域占位 -->
      <div class="flex-1">
        <div class="grid gap-4">
          <div v-for="n in 3" :key="n" class="card p-4">
            <div class="flex items-center gap-3 mb-3">
              <div class="h-10 w-10 bg-gray-700 rounded-full"></div>
              <div class="flex-1">
                <div class="h-5 w-48 bg-gray-700 rounded mb-2"></div>
                <div class="h-4 w-32 bg-gray-700 rounded"></div>
              </div>
            </div>
            <div class="h-4 w-3/4 bg-gray-700 rounded mb-2"></div>
            <div class="h-4 w-1/2 bg-gray-700 rounded"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- 底部导航栏骨架 -->
    <div class="fixed bottom-0 left-0 right-0 h-20 bg-gray-800/80 z-50 flex items-center justify-around px-4">
      <div v-for="n in 3" :key="n" class="flex flex-col items-center gap-1">
        <div class="h-6 w-6 bg-gray-700 rounded"></div>
        <div class="h-4 w-16 bg-gray-700 rounded"></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
// 无需任何 props 或逻辑
</script>