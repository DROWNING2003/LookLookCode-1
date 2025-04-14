<!--
  @component BButton
  @description 一个基础的按钮组件，提供了标准的按钮样式和交互效果。
  
  @props {
    text: string - 按钮显示的文本内容
    disabled?: boolean - 是否禁用按钮，默认为 false
  }
  
  @events {
    click: (event: MouseEvent) - 点击按钮时触发，当按钮未被禁用时生效
  }
  
  @example 基础用法
  ```vue
  <BButton text="确认" @click="handleClick" />
  ```
  
  @example 禁用状态
  ```vue
  <BButton text="禁用按钮" :disabled="true" />
  ```
  
  @example 在表单中使用
  ```vue
  <form @submit.prevent="onSubmit">
    <BButton text="提交表单" @click="submitForm" />
  </form>
  ```
-->

<template>
  <button
    :class="[
      'rounded-[var(--border-radius)] hover:bg-[var(--primary-color)]-700 bg-[var(--primary-color)] text-[var(--text-primary)] text-(length:--front-size)',
      'w-full h-12 flex items-center justify-center',
      'shadow-lg transition-all duration-200',
      'active:opacity-80',
      { 'btn-disabled': disabled },
    ]"
    :disabled="disabled"
    @click="handleClick"
  >
    <span class="tracking-wide">{{ text }}</span>
  </button>
</template>

<script setup lang="ts">
interface Props {
  text: string;
  disabled?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  disabled: false,
});

const emit = defineEmits<{
  (e: "click", event: MouseEvent): void;
}>();

const handleClick = (event: MouseEvent) => {
  if (!props.disabled) {
    emit("click", event);
  }
};
</script>
<style scoped>
button:disabled {
  pointer-events: none;
}

button:active {
  transform: scale(0.98);
}
</style>

