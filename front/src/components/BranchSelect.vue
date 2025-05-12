<template>
  <div class="relative">
    <button
      @click="isOpen = !isOpen"
      class="flex items-center gap-2 px-3 py-1.5 text-sm text-gray-300 bg-gray-800 hover:bg-gray-700 rounded-md transition-colors"
    >
      <i class="fas fa-code-branch"></i>
      <span>{{ modelValue }}</span>
      <i
        :class="['fas fa-chevron-down transition-transform', { 'transform rotate-180': isOpen }]"
      ></i>
    </button>

    <div
      v-if="isOpen"
      class="absolute z-10 w-48 mt-2 py-1 bg-gray-800 rounded-md shadow-lg"
      @click.stop
    >
      <div class="max-h-48 overflow-y-auto">
        <button
          v-for="branch in branches"
          :key="branch.name"
          @click="selectBranch(branch.name)"
          class="w-full px-4 py-2 text-sm text-left text-gray-300 hover:bg-gray-700 transition-colors"
          :class="{ 'bg-gray-700': branch.name === modelValue }"
        >
          {{ branch.name }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref,onUnmounted } from 'vue';
import type { Branch } from '../api/repo';

defineProps<{
  modelValue: string;
  branches: Branch[];
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const isOpen = ref(false);

const selectBranch = (branchName: string) => {
  emit('update:modelValue', branchName);
  isOpen.value = false;
};

// 点击外部关闭下拉框
const closeDropdown = (e: MouseEvent) => {
  if (isOpen.value) {
    isOpen.value = false;
  }
};

document.addEventListener('click', closeDropdown);

onUnmounted(() => {
  document.removeEventListener('click', closeDropdown);
});
</script>