<!-- 代码已包含 CSS：使用 TailwindCSS , 安装 TailwindCSS 后方可看到布局样式效果 -->
<template>
  <background
    class="absolute blur-(--blur) z-10 top-0 left-0 w-full min-h-screen bg-center bg-contain bg-(image:--background-image) bg-fixed"
  >
  </background>
  <div class="continer">
    <div
      v-if="showAddDialog"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
    >
      <div class="bg-[#1a1f2c] rounded-lg p-6 w-[90%] max-w-md">
        <h3 class="text-lg font-semibold mb-4">添加新属性</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm opacity-80 mb-2">属性名称</label>
            <input
              v-model="newProperty.key"
              type="text"
              placeholder="请输入CSS变量名称"
              class="w-full bg-[#242937] border border-[#2a2a2a] rounded px-3 py-2 text-sm"
            />
          </div>
          <div>
            <label class="block text-sm opacity-80 mb-2">属性值</label>
            <input
              v-model="newProperty.value"
              type="text"
              placeholder="请输入属性值"
              class="w-full bg-[#242937] border border-[#2a2a2a] rounded px-3 py-2 text-sm"
            />
          </div>
        </div>
        <div class="flex justify-end space-x-3 mt-6">
          <button
            class="!rounded-button bg-[#242937] px-4 py-2 text-sm"
            @click="closeAddDialog"
          >
            取消
          </button>
          <button
            class="!rounded-button bg-[#413edc] px-4 py-2 text-sm"
            @click="confirmAddProperty"
          >
            确定
          </button>
        </div>
      </div>
    </div>
    <nav
      class="fixed top-0 left-0 w-full bg-[#111722] border-b border-[#2a2a2a] p-4 z-50"
    >
      <h1 class="text-xl font-semibold">主题设置{{ ThemeStore.mode }}</h1>
    </nav>
    <div
      class="z-50 fixed bottom-0 left-0 w-full bg-[#111722] border-t border-[#2a2a2a] p-4 flex justify-between items-center"
    >
      <button
        class="!rounded-button bg-[#242937] px-6 py-2"
        @click="resetTheme"
      >
        恢复默认
      </button>
      <button class="!rounded-button bg-[#413edc] px-6 py-2" @click="saveTheme">
        保存主题
      </button>
    </div>
    <div class="h-screen w-full flex flex-row items-center justify-center p-4">
      <div
        class="h-full gap-4 w-[50%] flex flex-col items-center justify-center p-4"
        name="组件测试"
      >
        <BButton text="确认" />
        <BottomNavBar />
        <TopNavBar title="主页" />
        
      </div>
      <div name="编辑区域" class="w-[50%] h-full overflow-auto">
        <div class="mt-16 mb-24 space-y-6">
          <button
            class="!rounded-button w-full bg-[#1a1f2c] p-4 mb-4 flex items-center justify-center space-x-2"
            @click="addNewProperty"
          >
            <i class="fas fa-plus text-sm"></i>
            <span>添加新属性</span>
          </button>
          <div
            v-for="(value, key) in themeVariables"
            :key="key"
            class="card p-4 transition-all duration-300 hover:shadow-lg"
          >
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-4">
                <span class="text-sm opacity-80">{{ key }}</span>
                <div
                  v-if="isColor(value)"
                  class="w-6 h-6 rounded-full border border-[#2a2a2a]"
                  :style="{ backgroundColor: value }"
                ></div>
              </div>
              <button
                class="!rounded-button bg-[#413edc] px-4 py-2 text-sm"
                @click="toggleEdit(key)"
              >
                编辑
              </button>
            </div>
            <div v-if="editingKey === key" class="mt-4">
              <div v-if="isColor(value)" class="space-y-4">
                <input
                  type="color"
                  :value="value"
                  @input="(e: Event) => {
                    const target = e.target as HTMLInputElement;
                    updateValue(key, target.value);
                  }"
                  class="w-full h-10 rounded bg-transparent border border-[#2a2a2a]"
                />
                <input
                  type="text"
                  :value="value"
                  @input="(e: Event) => {
                    const target = e.target as HTMLInputElement;
                    updateValue(key, target.value);
                  }"
                  class="w-full bg-[#242937] border border-[#2a2a2a] rounded px-3 py-2 text-sm"
                />
              </div>
              <div v-else-if="isSize(value)" class="space-y-4">
                <input
                  type="range"
                  :value="String(parseFloat(value))"
                  :min="0"
                  :max="50"
                  @input="(e: Event) => {
                    const target = e.target as HTMLInputElement;
                    updateValue(key, `${target.value}${getSizeUnit(value)}`)
                  }"
                  class="w-full"
                />
                <input
                  type="text"
                  :value="value"
                  @input="(e: Event) => {
                    const target = e.target as HTMLInputElement;
                    updateValue(key, target.value);
                  }"
                  class="w-full bg-[#242937] border border-[#2a2a2a] rounded px-3 py-2 text-sm"
                />
              </div>
              <div v-else class="space-y-4">
                <input
                  type="text"
                  :value="value"
                  @input="(e: Event) => { const target = e.target as HTMLInputElement; updateValue(key, target.value); }"
                  class="w-full bg-[#242937] border border-[#2a2a2a] rounded px-3 py-2 text-sm"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import BButton from "@/baseComponent/BButton.vue";
import TopNavBar from "../components/TopNavBar.vue";
import BottomNavBar from "@/components/BottomNavBar.vue";
import { ref, reactive, watch } from "vue";
//导入当前主题
import { useThemeStore } from "../stores/theme";
const ThemeStore = useThemeStore();
let themeVariables: Theme = ThemeStore.theme["customer"]; //当前主题

//导入Theme类型
import type { Theme } from "../stores/theme";
const editingKey = ref("");

const isColor = (value: string) => {
  return value.startsWith("#") || value.startsWith("rgb");
};

const isSize = (value: string) => {
  return value.includes("px") || value.includes("rem") || value.includes("em");
};

const getSizeUnit = (value: string) => {
  if (value.includes("px")) return "px";
  if (value.includes("rem")) return "rem";
  if (value.includes("em")) return "em";
  return "";
};

const toggleEdit = (key: keyof Theme) => {
  editingKey.value = editingKey.value === String(key) ? "" : String(key);
};

const updateValue = (key: keyof Theme, value: string) => {
  themeVariables[key] = value;
};
const resetTheme = () => {
  Object.assign(themeVariables, ThemeStore.theme["default"]);
  console.log("恢复默认主题");

  ThemeStore.setCustomTheme(themeVariables);
};
const saveTheme = () => {
  ThemeStore.setCustomTheme(themeVariables);
};
const showAddDialog = ref(false);
const newProperty = reactive({
  key: "",
  value: "",
});

const addNewProperty = () => {
  showAddDialog.value = true;
  newProperty.key = "";
  newProperty.value = "";
};

const closeAddDialog = () => {
  showAddDialog.value = false;
};

const confirmAddProperty = () => {
  if (!newProperty.key || !newProperty.value) return;

  const formattedKey = newProperty.key.startsWith("--")
    ? newProperty.key
    : `--${newProperty.key}`;
  themeVariables[formattedKey] = newProperty.value;
  showAddDialog.value = false;
  editingKey.value = formattedKey;
};

//监听themeVariables变化
watch(
  () => themeVariables,
  (newValue) => {
    if (newValue === ThemeStore.theme["default"]) {
      //恢复默认主题
      console.log(newValue, ThemeStore.theme["default"]);
      return;
    }
    ThemeStore.setCustomTheme(newValue);
  },
  { deep: true }
);
</script>
<style scoped>
input[type="range"] {
  appearance: none;
  background: #242937;
  height: 4px;
  border-radius: 2px;
  outline: none;
}
input[type="range"]::-webkit-slider-thumb {
  appearance: none;
  width: 16px;
  height: 16px;
  background: #413edc;
  border-radius: 50%;
  cursor: pointer;
}
input[type="color"] {
  appearance: none;
  padding: 0;
}
input[type="color"]::-webkit-color-swatch-wrapper {
  padding: 0;
}
input[type="color"]::-webkit-color-swatch {
  border: none;
  border-radius: 4px;
}
</style>
