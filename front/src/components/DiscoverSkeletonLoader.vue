<template>
  <div class="animate-pulse">
    <div class="h-screen overflow-auto flex flex-col pt-16 pb-20 px-4 max-w-7xl mx-auto w-full">
      <!-- 热门推荐骨架 -->
      <section class="flex flex-col items-center mb-8">
        <div class="h-7 w-32 bg-gray-700 rounded mb-4"></div>
        <div class="grid grid-cols-1 gap-4 w-full">
          <div v-for="n in 3" :key="n" class="flex flex-col p-4 rounded-xl card">
            <div class="flex items-start gap-3">
              <div class="w-10 h-10 rounded-full bg-gray-700 flex-shrink-0"></div>
              <div class="flex-1 min-w-0">
                <div class="h-5 w-48 bg-gray-700 rounded mb-2"></div>
                <div class="h-4 w-full bg-gray-700 rounded mb-3"></div>
                <div class="flex items-center gap-4">
                  <div class="h-4 w-16 bg-gray-700 rounded"></div>
                  <div class="h-4 w-24 bg-gray-700 rounded"></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- 趋势榜单骨架 -->
      <section class="flex-1 mb-8">
        <div class="h-7 w-32 bg-gray-700 rounded mb-4"></div>
        <div class="grid gap-4 grid-cols-1 md:grid-cols-2 xl:grid-cols-3">
          <div v-for="n in 6" :key="n" class="flex flex-col p-4 rounded-xl card">
            <div class="flex items-center justify-between mb-2">
              <div class="h-5 w-32 bg-gray-700 rounded"></div>
              <div class="h-5 w-16 bg-gray-700 rounded"></div>
            </div>
            <div class="h-4 w-full bg-gray-700 rounded"></div>
          </div>
        </div>
      </section>

      <!-- 个性化推荐骨架 -->
      <section class="flex-1">
        <div class="h-7 w-32 bg-gray-700 rounded mb-4"></div>
        <div class="flex overflow-x-auto -mx-4 px-4 scrollbar-none">
          <div class="inline-flex gap-4 pb-4 flex-nowrap">
            <div v-for="n in 3" :key="n" class="flex flex-col w-[280px] sm:w-[320px] md:w-[280px] lg:w-[320px] flex-shrink-0 card rounded-lg overflow-hidden">
              <div class="w-full h-40 bg-gray-700"></div>
              <div class="p-4">
                <div class="h-5 w-48 bg-gray-700 rounded mb-2"></div>
                <div class="h-4 w-full bg-gray-700 rounded mb-3"></div>
                <div class="flex items-center justify-between">
                  <div class="h-4 w-16 bg-gray-700 rounded"></div>
                  <div class="h-4 w-20 bg-gray-700 rounded"></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script lang="ts" setup>
// 无需任何 props 或逻辑
</script>

<style scoped>
.scrollbar-none {
  scrollbar-width: none;
  -ms-overflow-style: none;
}
.scrollbar-none::-webkit-scrollbar {
  display: none;
}
</style>