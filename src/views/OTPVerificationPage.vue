<template>
  <div class="min-h-screen bg-gradient-to-br from-emerald-800 via-green-800 to-teal-700 relative overflow-hidden flex flex-col items-center justify-center px-6 py-8">
    <!-- Background Blobs -->
    <div class="absolute inset-0">
      <div class="absolute top-20 left-10 w-64 h-64 bg-emerald-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob"></div>
      <div class="absolute top-40 right-10 w-64 h-64 bg-green-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-2000"></div>
      <div class="absolute bottom-20 left-20 w-64 h-64 bg-teal-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-4000"></div>
    </div>
    
    <!-- Main Content -->
    <div class="relative z-10 flex flex-col items-center max-w-md w-full">
      <!-- Header -->
      <div class="text-center mb-8">
        <div class="w-16 h-16 mx-auto bg-white/20 backdrop-blur-sm rounded-2xl flex items-center justify-center mb-4">
          <span class="text-2xl">📱</span>
        </div>
        <h1 class="text-2xl font-bold text-white mb-2">Verify Your Phone</h1>
        <p class="text-white/80 text-sm">
          We've sent a 6-digit code to<br>
          <span class="font-semibold">{{ phoneNumber }}</span>
        </p>
      </div>

      <!-- OTP Form -->
      <div class="w-full bg-white/95 backdrop-blur-xl rounded-lg shadow-2xl border border-white/20 p-6">
        <form @submit.prevent="handleSubmit">
          <!-- OTP Input Fields -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-3 text-center">
              Enter verification code
            </label>
            <div class="flex justify-center space-x-3">
              <input
                v-for="(digit, index) in otpDigits"
                :key="index"
                :ref="el => otpInputs[index] = el"
                v-model="otpDigits[index]"
                type="text"
                maxlength="1"
                class="w-12 h-12 text-center text-xl font-bold border-2 rounded-lg focus:border-emerald-500 focus:outline-none transition-colors"
                :class="[
                  otpError ? 'border-red-500' : 'border-gray-300',
                  otpDigits[index] ? 'bg-emerald-50' : 'bg-gray-50'
                ]"
                @input="handleOTPInput(index, $event)"
                @keydown="handleKeydown(index, $event)"
                @paste="handlePaste"
              >
            </div>
            <p v-if="otpError" class="mt-2 text-xs text-red-600 text-center">{{ otpError }}</p>
          </div>

          <!-- Timer -->
          <div class="text-center mb-6">
            <p v-if="timeRemaining > 0" class="text-sm text-gray-600">
              Code expires in <span class="font-semibold text-emerald-600">{{ formatTime(timeRemaining) }}</span>
            </p>
            <p v-else class="text-sm text-red-600">
              Your code has expired
            </p>
          </div>

          <!-- Submit Button -->
          <button
            type="submit"
            :disabled="!isOTPComplete || isSubmitting"
            class="w-full py-3 rounded-lg font-semibold text-sm transition-all duration-300 transform"
            :class="isOTPComplete && !isSubmitting
              ? 'bg-emerald-600 hover:bg-emerald-700 text-white hover:scale-105 shadow-lg'
              : 'bg-gray-400 text-gray-600 cursor-not-allowed'"
          >
            <div v-if="isSubmitting" class="flex items-center justify-center">
              <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
              Verifying...
            </div>
            <span v-else>
              Confirm Code
            </span>
          </button>

          <!-- Resend Code -->
          <div class="text-center mt-4">
            <button
              type="button"
              @click="resendOTP"
              :disabled="timeRemaining > 0 || isResending"
              class="text-sm transition-colors duration-200"
              :class="timeRemaining > 0 || isResending 
                ? 'text-gray-400 cursor-not-allowed' 
                : 'text-emerald-600 hover:text-emerald-700 font-medium'"
            >
              <div v-if="isResending" class="flex items-center justify-center">
                <div class="animate-spin rounded-full h-3 w-3 border-b-2 border-emerald-600 mr-1"></div>
                Sending...
              </div>
              <span v-else>
                Didn't receive code? Resend
              </span>
            </button>
          </div>

          <!-- Back to Registration -->
          <div class="text-center mt-4">
            <button
              type="button"
              @click="goBack"
              class="text-sm text-gray-500 hover:text-gray-700 transition-colors duration-200"
            >
              ← Back to registration
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- OTP Modal -->
    <OTPModal
      :is-visible="showOTPModal"
      :otp-code="receivedOTP"
      :phone-number="phoneNumber"
      @close="showOTPModal = false"
      @otp-copied="handleOTPCopied"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, nextTick } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import OTPModal from '@/components/OTPModal.vue'
import API_BASE_URL from '@/config/api'

const router = useRouter()
const route = useRoute()

// Props from route
const phoneNumber = ref(route.query.phone || '')
const registrationData = ref({
  phone: route.query.phone || '',
  password: route.query.password || ''
})

// OTP state
const otpDigits = ref(['', '', '', '', '', ''])
const otpInputs = ref([])
const otpError = ref('')
const isSubmitting = ref(false)
const isResending = ref(false)

// Modal state
const showOTPModal = ref(false)
const receivedOTP = ref('')

// Timer state
const timeRemaining = ref(300) // 5 minutes
let timerInterval = null

// Computed
const isOTPComplete = computed(() => otpDigits.value.every(digit => digit !== ''))
const otpCode = computed(() => otpDigits.value.join(''))

// Methods
const handleOTPInput = (index, event) => {
  const value = event.target.value.replace(/[^0-9]/g, '')
  otpDigits.value[index] = value
  
  // Clear error when user starts typing
  if (otpError.value) {
    otpError.value = ''
  }
  
  // Auto-focus next input
  if (value && index < 5) {
    nextTick(() => {
      otpInputs.value[index + 1]?.focus()
    })
  }
}

const handleKeydown = (index, event) => {
  // Handle backspace
  if (event.key === 'Backspace' && !otpDigits.value[index] && index > 0) {
    otpInputs.value[index - 1]?.focus()
  }
  
  // Handle arrow keys
  if (event.key === 'ArrowLeft' && index > 0) {
    otpInputs.value[index - 1]?.focus()
  }
  if (event.key === 'ArrowRight' && index < 5) {
    otpInputs.value[index + 1]?.focus()
  }
}

const handlePaste = (event) => {
  event.preventDefault()
  const pasteData = event.clipboardData.getData('text').replace(/[^0-9]/g, '')
  
  if (pasteData.length === 6) {
    otpDigits.value = pasteData.split('')
    otpError.value = ''
    // Focus last input
    nextTick(() => {
      otpInputs.value[5]?.focus()
    })
  }
}

const handleSubmit = async () => {
  if (!isOTPComplete.value) return
  
  isSubmitting.value = true
  otpError.value = ''
  
  try {
    const response = await fetch(`${API_BASE_URL}/auth/verify-otp`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        phone: registrationData.value.phone,
        otp: otpCode.value
      })
    })
    
    const data = await response.json()
    
    if (data.success) {
      // Store owner data in localStorage
      localStorage.setItem('user', JSON.stringify(data.owner))
      
      // Navigate to home page
      router.push({ name: 'Home' })
    } else {
      otpError.value = data.message
      // Clear OTP inputs on error
      otpDigits.value = ['', '', '', '', '', '']
      nextTick(() => {
        otpInputs.value[0]?.focus()
      })
    }
    
  } catch (error) {
    console.error('OTP verification error:', error)
    otpError.value = 'Network error. Please check your connection.'
  } finally {
    isSubmitting.value = false
  }
}

const resendOTP = async () => {
  isResending.value = true
  
  try {
    const response = await fetch(`${API_BASE_URL}/auth/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(registrationData.value)
    })
    
    const data = await response.json()
    
    if (data.success) {
      // Store new OTP and show modal
      receivedOTP.value = data.otp_code || ''
      showOTPModal.value = true
      
      // Reset timer
      timeRemaining.value = 300
      startTimer()
      
      // Clear current OTP
      otpDigits.value = ['', '', '', '', '', '']
      otpError.value = ''
      
      // Focus first input
      nextTick(() => {
        otpInputs.value[0]?.focus()
      })
    } else {
      otpError.value = data.message
    }
    
  } catch (error) {
    console.error('Resend OTP error:', error)
    otpError.value = 'Failed to resend code. Please try again.'
  } finally {
    isResending.value = false
  }
}

const goBack = () => {
  router.push({ name: 'SignIn' })
}

const formatTime = (seconds) => {
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = seconds % 60
  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`
}

const handleOTPCopied = (otpCode) => {
  // Auto-fill the OTP when copied
  if (otpCode && otpCode.length === 6) {
    otpDigits.value = otpCode.split('')
    otpError.value = ''
    nextTick(() => {
      otpInputs.value[5]?.focus()
    })
  }
}

// Show OTP modal on page load with initial OTP
const showInitialOTP = () => {
  // Get OTP from route query or generate a display OTP
  const initialOTP = route.query.otp || receivedOTP.value
  if (initialOTP) {
    receivedOTP.value = initialOTP
    showOTPModal.value = true
  }
}

const startTimer = () => {
  if (timerInterval) {
    clearInterval(timerInterval)
  }
  
  timerInterval = setInterval(() => {
    if (timeRemaining.value > 0) {
      timeRemaining.value--
    } else {
      clearInterval(timerInterval)
    }
  }, 1000)
}

// Lifecycle
onMounted(() => {
  // Check if we have required data
  if (!registrationData.value.phone || !registrationData.value.password) {
    router.push({ name: 'SignIn' })
    return
  }
  
  // Show initial OTP modal
  showInitialOTP()
  
  // Start timer
  startTimer()
  
  // Focus first input
  nextTick(() => {
    otpInputs.value[0]?.focus()
  })
})

onUnmounted(() => {
  if (timerInterval) {
    clearInterval(timerInterval)
  }
})
</script>