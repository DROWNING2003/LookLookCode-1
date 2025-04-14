<template>
  <div class="codemirror">
    {{ props.option.options.language }}
    <Codemirror
      class="code"
      v-model="code"
      :style="{ height: '100%' }"
      :spellcheck="true"
      :autofocus="true"
      :indent-with-tab="true"
      :tabSize="4"
      :extensions="extensions"
      @ready="(res)=>{console.log('ready',res)}"
      @change="(res)=>{code=res}"
      @focus="(res)=>{console.log('focus',res)}"
      @blur="(res)=>{console.log('blur',res)}"
    />
  </div>
</template>
<script setup>
import { Codemirror } from "vue-codemirror";
import { javascript } from "@codemirror/lang-javascript";
import { css } from "@codemirror/lang-css";
import { python } from "@codemirror/lang-python";
import { yaml } from "@codemirror/lang-yaml";
import { java } from "@codemirror/lang-java";
import { cpp } from "@codemirror/lang-cpp";
import { vue } from "@codemirror/lang-vue";
import { rust } from "@codemirror/lang-rust";
import { html } from "@codemirror/lang-html";
import { oneDark } from "@codemirror/theme-one-dark";
// 引入vue模块
import { ref, reactive, onMounted } from "vue";
let mode = ref("vue");
let extensions = [];
// 定义从父组件接收的属性
const props = defineProps({
  option: Object,
});

const code = defineModel('codeModel')

onMounted(() => {
  mode.value = props.option.options.language;
//   code.value = props.option.options.code;
  switch (props.option.options.language) {
    case "javascript":
      extensions = [javascript(), oneDark];
      break;
    case "css":
      extensions = [css(), oneDark];
      break;
    case "python":
      extensions = [python(), oneDark];
      break;
    case "html":
      extensions = [html(), oneDark];
      break;
    case "yaml":
      extensions = [yaml(), oneDark];
      break;
    case "java":
      extensions = [java(), oneDark];
      break;
    case "cpp":
      extensions = [cpp(), oneDark];
      break;
    case "rust":
      extensions = [rust(), oneDark];
      break;
    case "vue":
      extensions = [vue(), oneDark];
      break;
    default:
      extensions = [javascript(), oneDark];
      break;
  }
  // if (props.option.options.language === "js") {
  //   extensions = [javascript(), oneDark];
  // } else if (props.option.options.language === "css") {
  //   extensions = [css(), oneDark];
  // } else {
  //   extensions = [html(), oneDark];
  // }
});
</script>
<style scoped>
.codemirror {
  width: 100%;
  height: 100%;
}
</style>
