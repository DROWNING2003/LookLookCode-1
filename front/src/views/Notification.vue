<template>
  <div class="h-screen overflow-auto flex flex-col pt-16 pb-20 px-4 max-w-7xl mx-auto w-full">
    <h2 class="text-white text-lg font-medium mb-4">通知</h2>
    <div class="flex flex-col gap-4">
      <!-- 通知列表 -->
      <div class="flex flex-col">
        <div v-for="i in 5" :key="i" class="p-4 rounded-xl transition-colors duration-200 card hover:card-hover mb-4">
          <div class="flex items-start gap-3">
            <img src="/default.jpg" alt="avatar" class="w-10 h-10 rounded-full object-cover flex-shrink-0" />
            <div class="flex-1 min-w-0">
              <p class="text-gray-400 text-sm mb-2">用户xxx 关注了你</p>
              <span class="text-gray-500 text-xs">2小时前</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
// 组件逻辑待实现
</script>