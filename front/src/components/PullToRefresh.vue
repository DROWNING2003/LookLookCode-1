<template>
  <van-pull-refresh v-model="isRefreshing" @refresh="handleRefresh">
    <slot></slot>
  </van-pull-refresh>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps<{
  onRefresh?: () => Promise<void>;
}>();

const isRefreshing = ref(false);

const handleRefresh = async () => {
  if (props.onRefresh) {
    try {
      await props.onRefresh();
    } finally {
      isRefreshing.value = false;
    }
  }
};
</script>