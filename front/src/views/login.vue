<template>
  <div class="continer">
    <!-- 背景代码雨效果 -->
    <div class="absolute inset-0 z-0 matrix-bg">
      <div
        class="absolute inset-0 bg-gradient-to-b from-transparent via-green-900/10 to-green-900/30 backdrop-blur-sm"
      ></div>
    </div>

    <!-- 主要内容区域 -->
    <div
      class="relative z-10 w-full h-screen flex flex-col items-center justify-center px-6"
    >
      <!-- Logo区域 -->
      <div class="text-center mb-16">
        <h1 class="text-4xl font-mono mb-4 font-bold text-gradient">
          代码阅读
        </h1>
        <p class="text-secondary font-mono text-lg tracking-wider">
          阅读好的代码是种享受
        </p>
      </div>

      <!-- Github登录按钮 -->
      <div class="glass-button-container mb-8">
        <button
          class="glass-button flex items-center justify-center space-x-3 w-64"
          @click="handleGithubLogin"
        >
          <i class="fab fa-github text-2xl"></i>
          <span>使用 Github 登录</span>
        </button>
      </div>
      <!-- 装饰线条 -->
      <div class="cyber-corners"></div>
    </div>

    <!-- 底部版权信息 -->
    <div class="fixed bottom-4 w-full text-center z-10">
      <p class="text-secondary text-sm font-mono" @click="showEnvEditor = true">
        &copy; 2024 代码阅读. 保留所有权利
      </p>
    </div>
    <EnvEditor :isOpen="showEnvEditor" @close="showEnvEditor = false"/>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useUserStore } from "../stores/user";
import { useThemeStore } from "../stores/theme";
import { getUserInfo } from "../api/user";
import EnvEditor from "../components/EnvEditor.vue";

const userStore = useUserStore();
const themeStore = useThemeStore();
const showEnvEditor = ref(false);


const handleGithubLogin = () => {
  window.location.href = themeStore.VITE_BACK_URL+"/api/auth/github";
};

const handleCallback = () => {
  const urlParams = new URLSearchParams(window.location.search);
  const code = urlParams.get("code");
  const state = urlParams.get("state");

  if (code) {
    fetch(
      `${themeStore.VITE_BACK_URL}/api/auth/github/callback?code=${code}&state=${state}`
    )
      .then((response) => response.json())
      .then((data) => {
        const accessToken = data.access_token;
        userStore.setAccessToken(accessToken);
        getUserInfo().then((res) => {
          userStore.setUser(res.data);
          window.location.href = "/";
        });
      })
      .catch((error) => {
        console.error("获取访问令牌时出错:", error);
      });
  }
};

onMounted(() => {
  handleCallback();
});
</script>

<style scoped>
.text-gradient {
  background: linear-gradient(135deg, #00ff00 0%, #4ade80 50%, #22d3ee 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  animation: gradient 8s linear infinite;
  background-size: 200% 200%;
  font-size: 2.5rem;
  letter-spacing: 0.1em;
}
.glass-button-container {
  position: relative;
}
.glass-button {
  background: rgba(98, 0, 238, 0.2);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(98, 0, 238, 0.5);
  color: white;
  padding: 1rem;
  border-radius: 8px;
  font-family: monospace;
  transition: all 0.3s ease;
  text-shadow: 0 0 5px rgba(255, 255, 255, 0.5);
}
.glass-button:hover {
  background: rgba(98, 0, 238, 0.3);
  box-shadow: 0 0 20px rgba(98, 0, 238, 0.5);
  transform: translateY(-2px);
}
</style>
