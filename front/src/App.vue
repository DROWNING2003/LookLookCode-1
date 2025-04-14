<script setup lang="ts">
import {RouterView } from 'vue-router'
//主题自定义
import {useThemeStore} from "./stores/theme";
const ThemeStore = useThemeStore()
//判断本地主题是否存在
if (ThemeStore.theme[ThemeStore.mode]!=null){
  ThemeStore.applyTheme()
  console.log("主题切换",ThemeStore.theme[ThemeStore.mode]);
}
</script>

<template>
  <RouterView />
</template>

<style scoped>
</style>
