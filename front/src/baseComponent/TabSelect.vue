<template>
  <div class="mb-4">
    <div
      class="flex space-x-2 card-layer p-1 rounded-xl"
    >
      <template v-for="(tab, index) in tabs" :key="index">
        <button
          class="flex-1 py-2 px-4 rounded-lg transition-colors duration-300 ease-in-out"
          :class="[
            modelValue === index
              ? 'bg-[var(--primary-color)] text-[var(--text-primary)]'
              : 'text-[var(--text-secondary)] hover:bg-[var(--primary-color)]/10',
          ]"
          @click="handleSelect(index)"
        >
          <slot :name="`tab-${index}`">
            {{ tab }}
          </slot>
        </button>
      </template>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { defineProps, defineEmits } from 'vue';

interface Props {
  modelValue: number;
  tabs: string[];
}

const props = defineProps<Props>();
const emit = defineEmits<{
  (e: 'update:modelValue', value: number): void;
}>();

const handleSelect = (index: number) => {
  if (index !== props.modelValue) {
    emit('update:modelValue', index);
  }
};
</script>