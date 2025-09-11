<template>
  <div class="min-h-screen relative overflow-hidden">
    <!-- Animated Background -->
    <div class="absolute inset-0 bg-gradient-to-br from-emerald-800 via-teal-800 to-green-700">
      <div class="absolute top-10 left-10 w-72 h-72 bg-green-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob"></div>
      <div class="absolute top-0 right-4 w-72 h-72 bg-emerald-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-2000"></div>
      <div class="absolute -bottom-8 left-20 w-72 h-72 bg-teal-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-4000"></div>
    </div>
    
    <!-- Content -->
    <div class="relative z-10 flex items-center justify-center min-h-screen p-6">
      <div class="bg-white/90 backdrop-blur-xl rounded-3xl shadow-2xl p-8 max-w-lg w-full border border-white/20">
        <!-- Floating Mobile Icon -->
        <div class="relative mb-8">
          <div class="w-24 h-24 mx-auto bg-gradient-to-br from-red-400 to-pink-600 rounded-2xl flex items-center justify-center shadow-lg transform rotate-3 hover:rotate-0 transition-transform duration-300">
            <svg class="w-12 h-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
            </svg>
          </div>
          <div class="absolute -top-2 -right-2 w-6 h-6 bg-red-500 rounded-full flex items-center justify-center">
            <span class="text-white text-xs font-bold">!</span>
          </div>
        </div>
        
        <!-- Title with Gradient Text -->
        <h1 class="text-4xl font-black text-transparent bg-clip-text bg-gradient-to-r from-emerald-600 to-teal-600 mb-6 text-center">
          Mobile Only Access
        </h1>
        
        <!-- Stylish Message -->
        <div class="bg-gradient-to-r from-emerald-50 to-teal-50 rounded-2xl p-6 mb-8 border border-emerald-100">
          <p class="text-xl text-gray-700 text-center font-medium leading-relaxed">
            This food delivery platform is designed exclusively for 
            <span class="text-emerald-600 font-bold">mobile devices</span> 
            to ensure the best user experience! 📱
          </p>
        </div>
        
        <!-- Device Info Card -->
        <div class="bg-gradient-to-br from-gray-50 to-gray-100 rounded-2xl p-6 mb-8 shadow-inner">
          <div class="flex items-center justify-between mb-3">
            <span class="text-gray-600 font-medium">Current Device:</span>
            <span class="bg-red-100 text-red-700 px-3 py-1 rounded-full text-sm font-semibold">{{ deviceType }}</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600 font-medium">Screen Size:</span>
            <span class="text-gray-800 font-mono text-sm">{{ screenSize }}</span>
          </div>
        </div>
        
        <!-- Beautiful Instructions -->
        <div class="mb-8">
          <h3 class="text-xl font-bold text-gray-800 mb-6 text-center">
            🚀 How to Access
          </h3>
          <div class="space-y-4">
            <div class="flex items-center p-4 bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl border border-blue-100">
              <div class="w-10 h-10 bg-blue-500 rounded-full flex items-center justify-center mr-4 flex-shrink-0">
                <span class="text-white font-bold">1</span>
              </div>
              <span class="text-gray-700 font-medium">Open this link on your smartphone</span>
            </div>
            <div class="flex items-center p-4 bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl border border-green-100">
              <div class="w-10 h-10 bg-green-500 rounded-full flex items-center justify-center mr-4 flex-shrink-0">
                <span class="text-white font-bold">2</span>
              </div>
              <span class="text-gray-700 font-medium">Use your tablet in portrait mode</span>
            </div>
            <div class="flex items-center p-4 bg-gradient-to-r from-teal-50 to-emerald-50 rounded-xl border border-teal-100">
              <div class="w-10 h-10 bg-teal-500 rounded-full flex items-center justify-center mr-4 flex-shrink-0">
                <span class="text-white font-bold">3</span>
              </div>
              <span class="text-gray-700 font-medium">Scan QR code with mobile device</span>
            </div>
          </div>
        </div>
        
        <!-- Beautiful Refresh Button -->
        <button 
          @click="refreshCheck" 
          class="w-full bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-700 hover:to-teal-700 text-white font-bold py-4 px-8 rounded-2xl transition-all duration-300 transform hover:scale-105 hover:shadow-xl focus:outline-none focus:ring-4 focus:ring-emerald-300"
        >
          <span class="flex items-center justify-center">
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            Check Device Again
          </span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getDeviceType } from '@/utils/deviceDetection'

const deviceType = ref('')
const screenSize = ref('')

const updateDeviceInfo = () => {
  deviceType.value = getDeviceType()
  screenSize.value = `${window.innerWidth} x ${window.innerHeight}`
}

const refreshCheck = () => {
  updateDeviceInfo()
  window.location.reload()
}

onMounted(() => {
  updateDeviceInfo()
})
</script>