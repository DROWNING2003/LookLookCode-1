<template>
  <div>
    <!-- 遮罩层 -->
    <div
      v-show="modelValue"
      class="fixed inset-0 bg-black/50 z-51"
      @click="$emit('update:modelValue', false)"
    ></div>
    <!-- 侧边栏 -->
    <div
      class="fixed top-0 left-0 h-full w-80 bg-background-700 z-52 transform transition-transform duration-300 ease-in-out"
      :class="modelValue ? 'translate-x-0' : '-translate-x-full'"
    >
      <!-- 侧边栏头部 -->
      <div class="h-14 flex items-center justify-between px-4 border-b border-background-800">
        <h2 class="text-[var(--text-primary)] text-lg font-medium">设置</h2>
        <button
          class="p-2 text-[var(--text-secondary)] hover:text-[var(--text-primary)]"
          @click="$emit('update:modelValue', false)"
        >
          <i class="fas fa-times text-lg"></i>
        </button>
      </div>
      <!-- 侧边栏内容 -->
      <div class="p-4">
        <router-link
          to="/theme"
          class="flex items-center p-3 rounded-lg hover:bg-background-800 text-[var(--text-primary)]"
          @click="$emit('update:modelValue', false)"
        >
          <i class="fas fa-palette mr-3"></i>
          <span>主题设置</span>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
defineProps<{
  modelValue: boolean;
}>();

defineEmits<{
  (e: 'update:modelValue', value: boolean): void;
}>();
</script>