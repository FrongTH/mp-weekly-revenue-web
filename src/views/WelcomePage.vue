<template>
  <div class="min-h-screen bg-gradient-to-br from-emerald-800 via-green-800 to-teal-700 relative overflow-hidden">
    <!-- Background Blobs -->
    <div class="absolute inset-0">
      <div class="absolute top-20 left-10 w-64 h-64 bg-emerald-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob"></div>
      <div class="absolute top-40 right-10 w-64 h-64 bg-green-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-2000"></div>
      <div class="absolute bottom-20 left-20 w-64 h-64 bg-teal-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-4000"></div>
    </div>

    <!-- Main Content -->
    <div class="relative z-10 min-h-screen flex flex-col">
      <!-- Header -->
      <div class="pt-12 pb-8 text-center">
        <h1 class="text-4xl font-black text-white mb-2">Welcome to</h1>
        <h2 class="text-2xl font-bold text-emerald-200">Food Delivery Revenue Tracking</h2>
        <p class="text-white/80 mt-2">Revenue Tracking Made Easy</p>
      </div>

      <!-- Slideshow Container -->
      <div class="flex-1 px-6 pb-8">
        <div class="max-w-sm mx-auto h-full flex flex-col">
          <!-- Slide Content -->
          <div 
            class="bg-white/95 backdrop-blur-xl rounded-3xl shadow-2xl flex-1 relative overflow-hidden"
            @touchstart="handleTouchStart"
            @touchend="handleTouchEnd"
            @touchmove="handleTouchMove"
          >
            <div 
              class="flex h-80 transition-transform duration-500 ease-in-out"
              :style="{ transform: `translateX(-${currentSlide * 100}%)` }"
            >
              <!-- Slide 1 -->
              <div class="w-full flex-shrink-0 p-8 flex flex-col items-center justify-center">
                <div class="w-48 h-48 bg-gradient-to-br from-orange-400 to-red-500 rounded-3xl flex items-center justify-center mb-6 shadow-lg">
                  <span class="text-6xl">🍔</span>
                </div>
                <h3 class="text-2xl font-bold text-gray-800 mb-4 text-center">Manage Restaurants</h3>
                <p class="text-gray-600 text-center leading-relaxed">
                  Keep track of all your restaurant partners in one place. Monitor performance, manage commissions, and grow your network.
                </p>
              </div>

              <!-- Slide 2 -->
              <div class="w-full flex-shrink-0 p-8 flex flex-col items-center justify-center">
                <div class="w-48 h-48 bg-gradient-to-br from-blue-400 to-indigo-500 rounded-3xl flex items-center justify-center mb-6 shadow-lg">
                  <span class="text-6xl">📱</span>
                </div>
                <h3 class="text-2xl font-bold text-gray-800 mb-4 text-center">Track Orders</h3>
                <p class="text-gray-600 text-center leading-relaxed">
                  Real-time order monitoring with detailed analytics. See what's working and optimize your delivery operations.
                </p>
              </div>

              <!-- Slide 3 -->
              <div class="w-full flex-shrink-0 p-8 flex flex-col items-center justify-center">
                <div class="w-48 h-48 bg-gradient-to-br from-green-400 to-emerald-500 rounded-3xl flex items-center justify-center mb-6 shadow-lg">
                  <span class="text-6xl">💰</span>
                </div>
                <h3 class="text-2xl font-bold text-gray-800 mb-4 text-center">Revenue Reports</h3>
                <p class="text-gray-600 text-center leading-relaxed">
                  Comprehensive revenue tracking with detailed reports and insights to help you make data-driven decisions.
                </p>
              </div>

              <!-- Slide 4 - Buy Me Coffee -->
              <div class="w-full flex-shrink-0 p-8 flex flex-col items-center justify-center">
                <!-- Buy Me Coffee Button (no background container) -->
                <div class="mb-6 flex items-center justify-center z-50 relative">
                  <a 
                    href="https://www.buymeacoffee.com/rattapon.san" 
                    target="_blank"
                    class="transition-transform hover:scale-105 block z-50 relative"
                    style="pointer-events: auto;"
                  >
                    <img 
                      src="https://cdn.buymeacoffee.com/buttons/v2/default-green.png" 
                      alt="Buy Me A Coffee" 
                      style="height: 60px !important;width: 217px !important; pointer-events: auto;" 
                      class="rounded-lg shadow-lg"
                    >
                  </a>
                </div>
                <h3 class="text-2xl font-bold text-gray-800 mb-4 text-center">Support the Developer</h3>
                <p class="text-gray-600 text-center leading-relaxed mb-6">
                  If you find this app helpful, consider buying me a coffee to support future development and improvements.
                </p>
              </div>
            </div>
          </div>

          <!-- Slide Indicators -->
          <div class="flex justify-center mt-6 mb-6">
            <div class="flex space-x-2">
              <button
                v-for="(slide, index) in totalSlides"
                :key="index"
                @click="goToSlide(index)"
                class="w-3 h-3 rounded-full transition-all duration-300"
                :class="currentSlide === index 
                  ? 'bg-white scale-125' 
                  : 'bg-white/50 hover:bg-white/75'"
              />
            </div>
          </div>

          <!-- Get Started Button -->
          <button
            @click="goToHome"
            :disabled="!canProceed"
            class="w-full py-4 rounded-2xl font-bold text-lg transition-all duration-300 transform"
            :class="canProceed
              ? 'bg-white text-emerald-600 hover:bg-emerald-50 hover:scale-105 shadow-lg'
              : 'bg-gray-400 text-gray-600 cursor-not-allowed opacity-50'"
          >
            {{ canProceed ? 'Get Started!' : 'Swipe to Continue' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Loading Screen -->
    <LoadingScreen 
      v-if="isLoading" 
      loading-text="Preparing your experience..."
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import LoadingScreen from '@/components/LoadingScreen.vue'

const router = useRouter()

// Loading state
const isLoading = ref(false)

// Slide state
const currentSlide = ref(0)
const totalSlides = 4

// Touch handling
const touchStartX = ref(0)
const touchEndX = ref(0)
const minSwipeDistance = 50

// Computed
const canProceed = computed(() => currentSlide.value === totalSlides - 1)

// Methods
const goToSlide = (index) => {
  if (index >= 0 && index < totalSlides) {
    currentSlide.value = index
  }
}

const nextSlide = () => {
  if (currentSlide.value < totalSlides - 1) {
    goToSlide(currentSlide.value + 1)
  }
}

const prevSlide = () => {
  if (currentSlide.value > 0) {
    goToSlide(currentSlide.value - 1)
  }
}

const goToHome = async () => {
  if (canProceed.value) {
    isLoading.value = true
    
    // Show loading for 2 seconds
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    // Navigate to SignIn page
    router.push({ name: 'SignIn' })
    
    isLoading.value = false
  }
}

// Touch handlers
const handleTouchStart = (e) => {
  // Check if touch started on a link or button
  const target = e.target.closest('a, button')
  if (target) {
    return // Don't handle swipe if touching a clickable element
  }
  touchStartX.value = e.changedTouches[0].screenX
}

const handleTouchMove = (e) => {
  // Check if touch is on a link or button
  const target = e.target.closest('a, button')
  if (target) {
    return // Don't prevent default for clickable elements
  }
  e.preventDefault() // Prevent scrolling
}

const handleTouchEnd = (e) => {
  // Check if touch ended on a link or button
  const target = e.target.closest('a, button')
  if (target) {
    return // Don't handle swipe if touching a clickable element
  }
  touchEndX.value = e.changedTouches[0].screenX
  handleSwipe()
}

const handleSwipe = () => {
  const swipeDistance = touchStartX.value - touchEndX.value
  
  if (Math.abs(swipeDistance) > minSwipeDistance) {
    if (swipeDistance > 0) {
      // Swiped left (next slide)
      nextSlide()
    } else {
      // Swiped right (previous slide)
      prevSlide()
    }
  }
}


// Auto-advance slides (optional)
const startAutoAdvance = () => {
  setInterval(() => {
    if (currentSlide.value < totalSlides - 1) {
      nextSlide()
    }
  }, 5000) // Auto advance every 5 seconds
}

onMounted(() => {
  // Uncomment if you want auto-advance
  // startAutoAdvance()
})
</script>