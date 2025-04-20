<template>
  <div class="relative w-full h-[200px] overflow-hidden rounded-b-[32px] mb-8">
    <div
      ref="sliderContainer"
      class="h-full flex transition-transform duration-500 ease-out"
      :style="{ transform: `translateX(-${currentIndex * 100}%)` }"
      @touchstart="handleTouchStart"
      @touchmove="handleTouchMove"
      @touchend="handleTouchEnd"
    >
      <div
        v-for="(banner, index) in banners"
        :key="index"
        class="relative w-full h-full flex-shrink-0 cursor-pointer"
        @click="banner.url && handleBannerClick(banner.url)"
      >
        <img :src="banner.image" :alt="banner.title" class="w-full h-full object-cover">
        <div class="absolute inset-0 bg-gradient-to-b from-transparent via-[#121212]/50 to-[#121212]">
          <div class="absolute bottom-8 left-8">
            <h1 class="text-3xl font-bold text-white mb-2">{{ banner.title }}</h1>
            <p class="text-gray-300">{{ banner.description }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- 分页指示器 -->
    <div class="absolute bottom-4 left-1/2 transform -translate-x-1/2 flex space-x-2">
      <button
        v-for="(_, index) in banners"
        :key="index"
        class="w-2 h-2 rounded-full transition-all duration-300"
        :class="index === currentIndex ? 'bg-white' : 'bg-white/50'"
        @click="goToSlide(index)"
      ></button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';

interface Banner {
  image: string;
  title: string;
  description: string;
  url?: string;
}

const props = defineProps<{
  banners: Banner[];
}>();

const currentIndex = ref(0);
const autoPlayInterval = ref<number>();
const touchStartX = ref(0);
const touchDeltaX = ref(0);
const sliderContainer = ref<HTMLElement | null>(null);
const router = useRouter();

// 处理banner点击
const handleBannerClick = (url: string) => {
  if (url.startsWith('http')) {
    window.open(url, '_blank');
  } else {
    router.push(url);
  }
};

// 自动播放
const startAutoPlay = () => {
  autoPlayInterval.value = window.setInterval(() => {
    nextSlide();
  }, 3000);
};

const stopAutoPlay = () => {
  if (autoPlayInterval.value) {
    clearInterval(autoPlayInterval.value);
  }
};

// 切换到指定幻灯片
const goToSlide = (index: number) => {
  currentIndex.value = index;
};

// 下一张幻灯片
const nextSlide = () => {
  currentIndex.value = (currentIndex.value + 1) % props.banners.length;
};

// 上一张幻灯片
const prevSlide = () => {
  currentIndex.value = (currentIndex.value - 1 + props.banners.length) % props.banners.length;
};

// 触摸事件处理
const handleTouchStart = (e: TouchEvent) => {
  touchStartX.value = e.touches[0].clientX;
  stopAutoPlay();
};

const handleTouchMove = (e: TouchEvent) => {
  if (!sliderContainer.value) return;
  touchDeltaX.value = e.touches[0].clientX - touchStartX.value;
  const translateX = -(currentIndex.value * 100) + (touchDeltaX.value / sliderContainer.value.offsetWidth * 100);
  sliderContainer.value.style.transform = `translateX(${translateX}%)`;
};

const handleTouchEnd = () => {
  if (Math.abs(touchDeltaX.value) > 50) {
    if (touchDeltaX.value > 0) {
      prevSlide();
    } else {
      nextSlide();
    }
  }
  touchDeltaX.value = 0;
  startAutoPlay();
};

onMounted(() => {
  startAutoPlay();
});

onUnmounted(() => {
  stopAutoPlay();
});
</script>

<style scoped>
.transition-transform {
  transition: transform 0.3s ease-in-out;
}
</style>