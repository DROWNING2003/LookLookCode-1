<template>
  <div class="flex items-center justify-center space-x-2">
    <button
      class="px-3 py-1 rounded-lg text-sm transition-all duration-200 active:scale-95"
      :class="[
        currentPage === 1
          ? 'text-gray-500 bg-gray-800/30 cursor-not-allowed'
          : 'text-gray-200 bg-gray-800/80 hover:bg-gray-700 active:bg-gray-600',
      ]"
      :disabled="currentPage === 1"
      @click="handlePageChange(currentPage - 1)"
    >
      上一页
    </button>

    <div class="flex items-center space-x-1">
      <button
        v-for="page in displayPages"
        :key="page"
        class="w-8 h-8 rounded-lg text-sm transition-all duration-200 active:scale-95 flex items-center justify-center"
        :class="[
          page === currentPage
            ? 'bg-blue-500 text-white'
            : 'text-gray-200 bg-gray-800/80 hover:bg-gray-700 active:bg-gray-600',
        ]"
        @click="handlePageChange(page)"
      >
        {{ page }}
      </button>
    </div>

    <button
      class="px-3 py-1 rounded-lg text-sm transition-all duration-200 active:scale-95"
      :class="[
        currentPage === totalPages
          ? 'text-gray-500 bg-gray-800/30 cursor-not-allowed'
          : 'text-gray-200 bg-gray-800/80 hover:bg-gray-700 active:bg-gray-600',
      ]"
      :disabled="currentPage === totalPages"
      @click="handlePageChange(currentPage + 1)"
    >
      下一页
    </button>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue';

interface Props {
  currentPage: number;
  pageSize: number;
  total: number;
}

const props = defineProps<Props>();

const emit = defineEmits<{
  (e: 'update:currentPage', page: number): void;
}>();

const totalPages = computed(() =>
  Math.max(1, Math.ceil(props.total / props.pageSize))
);

const displayPages = computed(() => {
  const pages: number[] = [];
  const maxDisplayPages = 5;
  let start = Math.max(
    1,
    props.currentPage - Math.floor(maxDisplayPages / 2)
  );
  let end = Math.min(totalPages.value, start + maxDisplayPages - 1);

  if (end - start + 1 < maxDisplayPages) {
    start = Math.max(1, end - maxDisplayPages + 1);
  }

  for (let i = start; i <= end; i++) {
    pages.push(i);
  }

  return pages;
});

const handlePageChange = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    emit('update:currentPage', page);
  }
};


</script>