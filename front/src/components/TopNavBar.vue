<template>
  <SideBar v-model="showSideBar" />
  <nav
    class="w-full text-[var(--text-primary)] min-w-[320px] bg-background-700 backdrop-blur-md z-50 px-4 h-14 flex items-center justify-between fixed top-0"
  >
    <button class="p-2" @click="showSideBar = true">
      <i class="fas fa-bars text-lg"></i>
    </button>
    <div class="flex-1 mx-4 relative">
      <i class="fas fa-search absolute left-3 top-1/2 transform -translate-y-1/2 text-[var(--text-secondary)]"></i>
      <input
        type="text"
        class="w-full h-10 pl-10 shadow-2xl pr-4 bg-background-900/90 rounded-2xl text-[var(--text-primary)] border-none focus:outline-none placeholder-[var(--text-secondary)]"
        placeholder="搜索仓库..."
        v-model="searchQuery"
        @input="$emit('search', searchQuery)"
      />
    </div>
    <button class="p-2" @click="$emit('more')">
      <i class="fas fa-ellipsis-v text-lg"></i>
    </button>
  </nav>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import SideBar from './SideBar.vue';

const searchQuery = ref('');
const showSideBar = ref(false);

defineEmits<{
  (e: 'search', query: string): void;
  (e: 'more'): void;
}>();
</script>