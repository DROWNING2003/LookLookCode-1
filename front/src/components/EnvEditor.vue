<template>
  <van-dialog
    v-model:show="props.isOpen"
    title="编辑后端地址"
    show-cancel-button
    @confirm="handleConfirm"
  >
    <van-field
      v-model="backendUrl"
      label="后端地址"
      placeholder="请输入后端地址"
    />
  </van-dialog>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
import { useThemeStore } from '../stores/theme'
import { Dialog as VanDialog, Field as VanField } from 'vant'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const themeStore = useThemeStore()
const backendUrl = ref(themeStore.VITE_BACK_URL)

watch(() => props.isOpen, (val) => {
  if (val) {
    backendUrl.value = themeStore.VITE_BACK_URL
  }
})

const handleConfirm = () => {
  themeStore.setBackendUrl(backendUrl.value)
  emit('close')
}
</script>
