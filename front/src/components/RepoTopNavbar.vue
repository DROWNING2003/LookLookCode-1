<template>
  <nav
    class="w-full font-roboto text-lg font-medium leading-7 h-[56px] flex items-center px-4 text-gray-300 gap-10 flex-wrap content-center fixed top-0"
  >
    <button
      @click="router.back()"
      class="flex items-center gap-2 text-gray-300 hover:text-white transition-colors"
    >
      <i class="text-gray-300 hover:text-white fas fa-arrow-left"></i>
    </button>
    <slot> </slot>

    <!-- <button class="p-2" @click="$emit('more')">
      <i class="fas fa-ellipsis-v text-lg"></i>
    </button> -->
  </nav>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";
const router = useRouter();
</script>
