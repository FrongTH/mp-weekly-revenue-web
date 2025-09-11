<template>
  <div 
    v-if="isVisible" 
    class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 backdrop-blur-sm"
    @click="handleBackdropClick"
  >
    <div class="bg-white rounded-2xl shadow-2xl max-w-sm w-full mx-4 transform transition-all duration-300 scale-100">
      <!-- Header -->
      <div class="bg-gradient-to-r from-emerald-500 to-teal-600 text-white p-6 rounded-t-2xl">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-white/20 rounded-full flex items-center justify-center mr-3">
              <span class="text-xl">📱</span>
            </div>
            <div>
              <h3 class="text-lg font-bold">SMS Simulation</h3>
              <p class="text-emerald-100 text-sm">Testing OTP Code</p>
            </div>
          </div>
          <button
            @click="closeModal"
            class="w-8 h-8 bg-white/20 hover:bg-white/30 rounded-full flex items-center justify-center transition-colors"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
      </div>

      <!-- Content -->
      <div class="p-6">
        <!-- Phone Number -->
        <div class="text-center mb-4">
          <p class="text-gray-600 text-sm mb-2">SMS sent to:</p>
          <p class="font-semibold text-gray-800">{{ phoneNumber }}</p>
        </div>

        <!-- OTP Code Display -->
        <div class="bg-gray-50 border-2 border-emerald-200 rounded-xl p-6 mb-6">
          <p class="text-center text-gray-600 text-sm mb-3">Your verification code is:</p>
          <div class="text-center">
            <p class="text-3xl font-black text-emerald-600 tracking-wider font-mono">
              {{ formattedOTP }}
            </p>
          </div>
          <p class="text-center text-xs text-gray-500 mt-3">
            This code expires in 5 minutes
          </p>
        </div>

        <!-- Message -->
        <div class="bg-blue-50 border border-blue-200 rounded-lg p-4 mb-6">
          <div class="flex items-start">
            <svg class="w-5 h-5 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <div>
              <p class="text-blue-800 text-sm font-medium">For Testing Only</p>
              <p class="text-blue-700 text-xs mt-1">
                In production, this code would be sent via SMS to your phone. Copy this code and enter it in the verification form.
              </p>
            </div>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="flex gap-3">
          <button
            @click="copyOTP"
            class="flex-1 py-3 px-4 bg-emerald-600 hover:bg-emerald-700 text-white font-semibold rounded-lg transition-colors duration-200 flex items-center justify-center"
          >
            <svg v-if="!copied" class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
            </svg>
            <svg v-else class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
            {{ copied ? 'Copied!' : 'Copy Code' }}
          </button>
          <button
            @click="closeModal"
            class="flex-1 py-3 px-4 bg-gray-200 hover:bg-gray-300 text-gray-700 font-semibold rounded-lg transition-colors duration-200"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

const props = defineProps({
  isVisible: {
    type: Boolean,
    default: false
  },
  otpCode: {
    type: String,
    default: ''
  },
  phoneNumber: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['close', 'otp-copied'])

const copied = ref(false)

// Format OTP with spaces for better readability
const formattedOTP = computed(() => {
  if (!props.otpCode) return ''
  return props.otpCode.split('').join(' ')
})

// Methods
const closeModal = () => {
  emit('close')
}

const handleBackdropClick = (event) => {
  // Only close if clicking the backdrop, not the modal content
  if (event.target === event.currentTarget) {
    closeModal()
  }
}

const copyOTP = async () => {
  try {
    await navigator.clipboard.writeText(props.otpCode)
    copied.value = true
    emit('otp-copied', props.otpCode)
    
    // Auto-close modal after successful copy with countdown
    setTimeout(() => {
      closeModal()
    }, 1500) // Close after 1.5 seconds to show "Copied!" feedback
    
  } catch (err) {
    console.error('Failed to copy OTP:', err)
    // Fallback for older browsers
    const textArea = document.createElement('textarea')
    textArea.value = props.otpCode
    document.body.appendChild(textArea)
    textArea.select()
    try {
      document.execCommand('copy')
      copied.value = true
      emit('otp-copied', props.otpCode)
      
      // Auto-close modal after successful copy with countdown
      setTimeout(() => {
        closeModal()
      }, 1500) // Close after 1.5 seconds to show "Copied!" feedback
      
    } catch (fallbackErr) {
      console.error('Fallback copy failed:', fallbackErr)
    }
    document.body.removeChild(textArea)
  }
}

// Reset copied state when modal is opened/closed
watch(() => props.isVisible, (newValue) => {
  if (newValue) {
    copied.value = false
  }
})
</script>