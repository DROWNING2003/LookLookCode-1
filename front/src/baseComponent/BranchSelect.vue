<template>
  <div class="relative">
   
    <button
      @click="isOpen = !isOpen"
      class="flex items-center gap-2 px-3 py-2 rounded-lg bg-amber-900 hover:bg-[var(--card-layer-hover)] transition-colors duration-200"
    >
      <i class="fas fa-code-branch"></i>
      <span class="text-[var(--text-primary)]">{{ props.modelValue.name }}</span>
      <i
        class="fas fa-chevron-down text-xs transition-transform duration-200"
        :class="{ 'transform rotate-180': isOpen }"
      ></i>
    </button>

    <div
      v-if="isOpen"
      class="absolute top-full left-0 mt-2 w-48 bg-gray-700 rounded-lg shadow-lg overflow-hidden z-[100]"
    >
      <div class="py-1">
        <button
          v-for="branch in props.branches"
          :key="branch"
          @click="selectBranch(branch)"
          class="w-full px-4 py-2 text-left hover:bg-[var(--card-layer-hover)] transition-colors duration-200 flex items-center gap-2"
          :class="{
            'text-[var(--primary-color)]': branch === modelValue,
            'text-[var(--text-primary)]': branch !== modelValue
          }"
        >
          <i
            class="fas fa-check text-xs"
            :class="{ invisible: branch !== modelValue }"
          ></i>
          {{ branch.name }}
        </button>
      </div>
    </div>

    <!-- 点击外部关闭下拉框 -->
    <div
      v-if="isOpen"
      class="fixed inset-0 z-40"
      @click="isOpen = false"
    ></div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
// API 和类型定义导入
import {
  type Branch,
} from "../api/repo";
interface Props {
  modelValue: Branch;
  branches: Branch[];
}

const props = defineProps<Props>();
const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const isOpen = ref(false);

const selectBranch = (branch: string) => {
  emit('update:modelValue', branch);
  isOpen.value = false;
};
</script>