<template>
  <div class="fixed inset-0 bg-gradient-to-br from-emerald-800 via-green-800 to-teal-700 z-50 flex items-center justify-center">
    <!-- Background Blobs -->
    <div class="absolute inset-0">
      <div class="absolute top-20 left-10 w-64 h-64 bg-emerald-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob"></div>
      <div class="absolute top-40 right-10 w-64 h-64 bg-green-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-2000"></div>
      <div class="absolute bottom-20 left-20 w-64 h-64 bg-teal-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-4000"></div>
    </div>

    <!-- Loading Content -->
    <div class="relative z-10 text-center">
      <!-- App Logo -->
      <div class="w-24 h-24 mx-auto bg-white/20 backdrop-blur-sm rounded-3xl flex items-center justify-center mb-8 shadow-lg">
        <span class="text-4xl">🍔</span>
      </div>

      <!-- App Title -->
      <h1 class="text-3xl font-black text-white mb-2">Food Delivery</h1>
      <h2 class="text-xl font-bold text-emerald-200 mb-8">Revenue Tracking</h2>

      <!-- Loading Animation -->
      <div class="flex items-center justify-center space-x-2 mb-4">
        <div class="w-3 h-3 bg-white rounded-full animate-bounce" style="animation-delay: 0ms;"></div>
        <div class="w-3 h-3 bg-white rounded-full animate-bounce" style="animation-delay: 150ms;"></div>
        <div class="w-3 h-3 bg-white rounded-full animate-bounce" style="animation-delay: 300ms;"></div>
      </div>

      <!-- Loading Text -->
      <p class="text-white/80 text-lg font-medium">{{ loadingText }}</p>

      <!-- Progress Bar -->
      <div class="w-64 h-2 bg-white/20 rounded-full mx-auto mt-6 overflow-hidden">
        <div 
          class="h-full bg-white rounded-full transition-all duration-300 ease-out"
          :style="{ width: `${progress}%` }"
        ></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const props = defineProps({
  loadingText: {
    type: String,
    default: 'Loading...'
  }
})

const progress = ref(0)

onMounted(() => {
  // Animate progress bar
  const interval = setInterval(() => {
    if (progress.value < 100) {
      progress.value += 2
    } else {
      clearInterval(interval)
    }
  }, 30)
})
</script>