<template>
  <div
    class="bg-gray-800 w-screen min-h-[100vh] flex flex-col items-center justify-start gap-6 p-8"
  >
    <h1 class="text-white text-2xl font-bold">CodeEdit 组件测试</h1>

    <div class="w-full max-w-4xl space-y-4">
      <!-- 控制面板 -->
      <div class="flex items-center justify-between bg-gray-700 p-4 rounded-lg">
        <div class="flex items-center gap-4">
          <select
            v-model="currentLanguage"
            class="bg-gray-600 text-white px-3 py-2 rounded-md"
            @change="updateEditorOptions"
          >
            <option value="js">JavaScript</option>
            <option value="go">go</option>
            <option value="ts">TypeScript</option>
            <option value="java">Java</option>
            <option value="cpp">C++</option>
            <option value="rust">Rust</option>
            <option value="vue">Vue</option>
            <option value="py">python</option>
            <option value="css">CSS</option>
            <option value="html">HTML</option>
          </select>
        </div>
      </div>

      <!-- 编辑器容器 -->
      <div class="h-[500px] w-full bg-gray-900 rounded-lg overflow-hidden">
        <CodeEdit
          v-model:codeModel="currentCode"
          :option="editorOptions"
          @update:modelValue="handleCodeChange"
        />
      </div>

      <!-- 代码预览 -->
      <div class="bg-gray-700 p-4 rounded-lg">
        <h3 class="text-white font-medium mb-2">ai 回复：</h3>
        <div class="text-gray-300">
          <div v-for="m in messages" :key="m.id">
            <div v-if="m.role != 'user'">
              <MdPreview id="preview-only" :modelValue="m.content" />
            </div>
          </div>
        </div>
      </div>
      <button @click="aichat">ai 分析</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import CodeEdit from "../components/CodeEdit.vue";
import { MdPreview } from "md-editor-v3";
import "md-editor-v3/lib/preview.css";
import { useChat } from "@ai-sdk/vue";

const { messages, input, handleSubmit } = useChat({
  api: "http://localhost:8080/api/chat",
});
function aichat() {
  input.value = currentCode.value;
  handleSubmit();

  // qwenAPI(currentCode.value).then((res) => {
  //   console.log(res);
  // });
}
// 当前选择的语言
const currentLanguage = ref("vue");

// 当前代码内容
const currentCode = ref("// 在这里输入代码");

// 编辑器配置
const editorOptions = reactive({
  options: {
    mode: "javascript",
    language: "vue",
    code: currentCode.value,
  },
});

// 更新编辑器配置
const updateEditorOptions = () => {
  editorOptions.options.language = currentLanguage.value;
  editorOptions.options.mode =
    currentLanguage.value === "js" ? "javascript" : currentLanguage.value;
};

// 处理代码变更
const handleCodeChange = (newCode: string) => {
  currentCode.value = newCode;
};
</script>
