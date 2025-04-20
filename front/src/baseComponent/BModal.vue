<template>
  <div v-if="props.isOpen" class="fixed inset-0 z-50 flex items-center justify-center">
    <div class="absolute inset-0 bg-black bg-opacity-50" @click="close"></div>
    <div class="relative bg-[var(--background-color)] rounded-lg shadow-xl w-96 p-6">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-semibold">{{ title }}</h3>
        <button @click="close" class="text-gray-400 hover:text-gray-600">
          <i class="fas fa-times"></i>
        </button>
      </div>
      <div class="mb-4">
        <slot></slot>
      </div>
      <div class="flex justify-end space-x-2">
        <button
          @click="close"
          class="px-4 py-2 text-sm rounded-md bg-gray-200 hover:bg-gray-300 text-gray-700"
        >
          取消
        </button>
        <button
          @click="confirm"
          class="px-4 py-2 text-sm rounded-md bg-[var(--primary-color)] hover:bg-[var(--primary-color)]-700 text-white"
        >
          确认
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title: string;
  isOpen: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'confirm'): void;
}>();

const close = () => {
  emit('close');
};

const confirm = () => {
  emit('confirm');
};
</script>