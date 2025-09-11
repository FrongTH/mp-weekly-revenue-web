<template>
  <div class="min-h-screen bg-gradient-to-br from-emerald-800 via-green-800 to-teal-700 relative overflow-hidden flex flex-col items-center justify-center px-6 py-8">
    <!-- Background Blobs -->
    <div class="absolute inset-0">
      <div class="absolute top-20 left-10 w-64 h-64 bg-emerald-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob"></div>
      <div class="absolute top-40 right-10 w-64 h-64 bg-green-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-2000"></div>
      <div class="absolute bottom-20 left-20 w-64 h-64 bg-teal-200 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-4000"></div>
    </div>
    
    <!-- Main Content -->
    <div class="relative z-10 flex flex-col items-center">
      <!-- Logo and Title -->
      <!-- <div class="flex items-center mb-6 text-2xl font-semibold text-white">
        <div class="w-8 h-8 mr-2 bg-white/20 backdrop-blur-sm rounded-lg flex items-center justify-center">
          <span class="text-white text-sm">🍔</span>
        </div>
        Food Delivery Revenue Tracking
      </div> -->
      
      <h1 class="text-xl mb-6 font-bold leading-tight tracking-tight text-white md:text-2xl">
        <span class="text-white text-sm">🍔</span>
          {{ isSignUp ? 'Register your account' : 'Sign in to your account' }}
        </h1> 

      <!-- Sign In Form -->
      <div class="w-full bg-white/95 backdrop-blur-xl rounded-lg shadow-2xl max-w-md border border-white/20">
      <div class="p-6 space-y-4 sm:p-8">
        <!-- <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl">
          {{ isSignUp ? 'Register your account' : 'Sign in to your account' }}
        </h1> -->

        <form class="space-y-4 md:space-y-6" @submit.prevent="handleSubmit">
          <!-- Phone Number Field -->
          <div>
            <label for="phone" class="block mb-2 text-sm font-medium text-gray-900">
              Your phone number
            </label>
            <input 
              type="tel" 
              name="phone" 
              id="phone" 
              v-model="phoneNumber"
              class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-emerald-600 focus:border-emerald-600 block w-full p-2.5"
              :class="{ 'border-red-500 focus:border-red-500 focus:ring-red-600': phoneError }"
              placeholder="09x-xxx-xxxx"
              required
              @input="formatPhoneNumber"
              @keydown="handlePhoneKeydown"
              @blur="validatePhone"
            >
            <p v-if="phoneError" class="mt-1 text-xs text-red-600">{{ phoneError }}</p>
          </div>

          <!-- Password Field -->
          <div>
            <label for="password" class="block mb-2 text-sm font-medium text-gray-900">
              Password
            </label>
            <div class="relative">
              <input 
                :type="showPassword ? 'text' : 'password'" 
                name="password" 
                id="password" 
                v-model="password"
                class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-emerald-600 focus:border-emerald-600 block w-full p-2.5 pr-10"
                :class="{ 'border-red-500 focus:border-red-500 focus:ring-red-600': passwordError }"
                required
                @input="validatePassword"
              >
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-400 hover:text-gray-600"
              >
                <svg v-if="!showPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                </svg>
                <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"></path>
                </svg>
              </button>
            </div>
            
            <!-- Password Requirements Checklist -->
            <div v-if="password" class="mt-2 space-y-1">
              <div class="flex items-center text-xs">
                <div 
                  class="w-3 h-3 rounded-full mr-2 flex items-center justify-center"
                  :class="passwordChecks.hasUppercase ? 'bg-green-500' : 'bg-red-500'"
                >
                  <svg v-if="passwordChecks.hasUppercase" class="w-2 h-2 text-white" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
                  </svg>
                  <svg v-else class="w-2 h-2 text-white" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                  </svg>
                </div>
                <span :class="passwordChecks.hasUppercase ? 'text-green-600' : 'text-red-600'">
                  At least one uppercase letter
                </span>
              </div>
              
              <div class="flex items-center text-xs">
                <div 
                  class="w-3 h-3 rounded-full mr-2 flex items-center justify-center"
                  :class="passwordChecks.hasLowercase ? 'bg-green-500' : 'bg-red-500'"
                >
                  <svg v-if="passwordChecks.hasLowercase" class="w-2 h-2 text-white" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
                  </svg>
                  <svg v-else class="w-2 h-2 text-white" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                  </svg>
                </div>
                <span :class="passwordChecks.hasLowercase ? 'text-green-600' : 'text-red-600'">
                  At least one lowercase letter
                </span>
              </div>
              
              <div class="flex items-center text-xs">
                <div 
                  class="w-3 h-3 rounded-full mr-2 flex items-center justify-center"
                  :class="passwordChecks.hasNumber ? 'bg-green-500' : 'bg-red-500'"
                >
                  <svg v-if="passwordChecks.hasNumber" class="w-2 h-2 text-white" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
                  </svg>
                  <svg v-else class="w-2 h-2 text-white" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                  </svg>
                </div>
                <span :class="passwordChecks.hasNumber ? 'text-green-600' : 'text-red-600'">
                  At least one number
                </span>
              </div>
              
              <div class="flex items-center text-xs">
                <div 
                  class="w-3 h-3 rounded-full mr-2 flex items-center justify-center"
                  :class="passwordChecks.noSpecialChars ? 'bg-green-500' : 'bg-red-500'"
                >
                  <svg v-if="passwordChecks.noSpecialChars" class="w-2 h-2 text-white" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
                  </svg>
                  <svg v-else class="w-2 h-2 text-white" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                  </svg>
                </div>
                <span :class="passwordChecks.noSpecialChars ? 'text-green-600' : 'text-red-600'">
                  No special characters allowed
                </span>
              </div>
            </div>

            <p v-if="passwordError" class="mt-1 text-xs text-red-600">{{ passwordError }}</p>
          </div>

          <!-- Remember Me / Forgot Password -->
          <div class="flex items-center justify-between" v-if="!isSignUp">
            <div class="flex items-start">
              <div class="flex items-center h-5">
                <input 
                  id="remember" 
                  type="checkbox" 
                  v-model="rememberMe"
                  class="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-emerald-300"
                >
              </div>
              <div class="ml-3 text-sm">
                <label for="remember" class="text-gray-500">Remember me</label>
              </div>
            </div>
            <a href="#" class="text-sm font-medium text-emerald-600 hover:underline">
              Forgot password?
            </a>
          </div>

          <!-- Submit Button -->
          <button 
            type="submit" 
            :disabled="!isFormValid || isSubmitting"
            class="w-full text-white font-medium rounded-lg text-sm px-5 py-2.5 text-center transition-all duration-300"
            :class="isFormValid && !isSubmitting 
              ? 'bg-emerald-600 hover:bg-emerald-700 focus:ring-4 focus:outline-none focus:ring-emerald-300' 
              : 'bg-gray-400 cursor-not-allowed'"
          >
            <div v-if="isSubmitting" class="flex items-center justify-center">
              <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
              {{ isSignUp ? 'Registering...' : 'Signing In...' }}
            </div>
            <span v-else>
              {{ isSignUp ? 'Register' : 'Sign in' }}
            </span>
          </button>

          <!-- Toggle Sign Up/Sign In -->
          <p class="text-sm font-light text-gray-500">
            {{ isSignUp ? 'Already have an account?' : "Don't have an account yet?" }}
            <a 
              href="#" 
              @click.prevent="toggleSignUp" 
              class="font-medium text-emerald-600 hover:underline ml-1"
            >
              {{ isSignUp ? 'Sign in' : 'Register' }}
            </a>
          </p>
        </form>
      </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import API_BASE_URL from '@/config/api'

const router = useRouter()

// Form state
const phoneNumber = ref('')
const password = ref('')
const showPassword = ref(false)
const rememberMe = ref(false)
const isSignUp = ref(false)
const isSubmitting = ref(false)

// Error states
const phoneError = ref('')
const passwordError = ref('')

// Password validation
const passwordChecks = computed(() => ({
  hasUppercase: /[A-Z]/.test(password.value),
  hasLowercase: /[a-z]/.test(password.value),
  hasNumber: /\d/.test(password.value),
  noSpecialChars: !/[^A-Za-z0-9]/.test(password.value)
}))

const isPasswordValid = computed(() => {
  return passwordChecks.value.hasUppercase && 
         passwordChecks.value.hasLowercase && 
         passwordChecks.value.hasNumber && 
         passwordChecks.value.noSpecialChars &&
         password.value.length >= 6
})

const isFormValid = computed(() => {
  return phoneNumber.value && 
         isPasswordValid.value && 
         !phoneError.value && 
         !passwordError.value
})

// Handle special keys for phone input
const handlePhoneKeydown = (event) => {
  const input = event.target
  const cursorPosition = input.selectionStart
  const value = input.value
  
  // Handle backspace
  if (event.key === 'Backspace' && !event.ctrlKey && !event.metaKey) {
    event.preventDefault()
    
    if (cursorPosition === 0) return
    
    // If cursor is right after a dash, move cursor to before the dash
    if (cursorPosition > 0 && value[cursorPosition - 1] === '-') {
      input.setSelectionRange(cursorPosition - 1, cursorPosition - 1)
      return
    }
    
    // Remove the character before cursor and reformat
    const beforeCursor = value.substring(0, cursorPosition - 1)
    const afterCursor = value.substring(cursorPosition)
    const newValue = beforeCursor + afterCursor
    
    // Remove all formatting and reformat
    const digitsOnly = newValue.replace(/\D/g, '')
    const formatted = formatPhoneDigits(digitsOnly)
    
    phoneNumber.value = formatted
    input.value = formatted
    
    // Calculate new cursor position - should be after the last remaining digit
    const digitsBeforeCursor = beforeCursor.replace(/\D/g, '').length
    let newCursorPos = formatted.length // Default to end
    
    if (digitsBeforeCursor === 0) {
      newCursorPos = 0
    } else {
      let digitCount = 0
      for (let i = 0; i < formatted.length; i++) {
        if (formatted[i] !== '-') {
          digitCount++
        }
        if (digitCount === digitsBeforeCursor) {
          newCursorPos = i + 1
          break
        }
      }
    }
    
    setTimeout(() => {
      input.setSelectionRange(newCursorPos, newCursorPos)
    }, 0)
  }
  
  // Handle delete key
  if (event.key === 'Delete' && !event.ctrlKey && !event.metaKey) {
    event.preventDefault()
    
    if (cursorPosition >= value.length) return
    
    // If cursor is right before a dash, move cursor to after the dash
    if (cursorPosition < value.length && value[cursorPosition] === '-') {
      input.setSelectionRange(cursorPosition + 1, cursorPosition + 1)
      return
    }
    
    // Remove the character after cursor and reformat
    const beforeCursor = value.substring(0, cursorPosition)
    const afterCursor = value.substring(cursorPosition + 1)
    const newValue = beforeCursor + afterCursor
    
    // Remove all formatting and reformat
    const digitsOnly = newValue.replace(/\D/g, '')
    const formatted = formatPhoneDigits(digitsOnly)
    
    phoneNumber.value = formatted
    input.value = formatted
    
    // Calculate new cursor position - keep cursor at same digit position
    const digitsBeforeCursor = beforeCursor.replace(/\D/g, '').length
    let newCursorPos = formatted.length // Default to end
    
    if (digitsBeforeCursor === 0) {
      newCursorPos = 0
    } else {
      let digitCount = 0
      for (let i = 0; i < formatted.length; i++) {
        if (formatted[i] !== '-') {
          digitCount++
        }
        if (digitCount === digitsBeforeCursor) {
          newCursorPos = i + 1
          break
        }
      }
    }
    
    setTimeout(() => {
      input.setSelectionRange(newCursorPos, newCursorPos)
    }, 0)
  }
}

// Helper function to format phone digits
const formatPhoneDigits = (digits) => {
  if (digits.length === 0) return ''
  
  // Limit to 10 digits
  if (digits.length > 10) {
    digits = digits.substring(0, 10)
  }
  
  // Apply 09x-xxx-xxxx format
  let formatted = digits
  if (digits.length >= 3) {
    formatted = digits.substring(0, 3) + '-' + digits.substring(3)
  }
  if (digits.length >= 6) {
    formatted = digits.substring(0, 3) + '-' + digits.substring(3, 6) + '-' + digits.substring(6)
  }
  
  return formatted
}

// Phone formatting method for regular input
const formatPhoneNumber = (event) => {
  const input = event.target.value.replace(/\D/g, '') // Remove all non-digits
  const formatted = formatPhoneDigits(input)
  
  phoneNumber.value = formatted
  event.target.value = formatted
}

// Validation methods
const validatePhone = async () => {
  phoneError.value = ''
  const phoneRegex = /^0\d{2}-\d{3}-\d{4}$/
  
  if (!phoneNumber.value) {
    phoneError.value = 'Phone number is required'
    return
  } 
  
  if (!phoneRegex.test(phoneNumber.value)) {
    phoneError.value = 'Please enter a valid Thai phone number (09x-xxx-xxxx)'
    return
  }

  // Check phone number existence for sign in/register validation
  try {
    const response = await fetch(`${API_BASE_URL}/auth/check-phone?phone=${encodeURIComponent(phoneNumber.value)}`)
    const data = await response.json()
    
    if (isSignUp.value && data.exists) {
      phoneError.value = 'This phone number is member please sign in or other number register'
    } else if (!isSignUp.value && !data.exists) {
      phoneError.value = 'This mobile phone is not member please register before.'
    }
  } catch (error) {
    console.error('Phone validation error:', error)
    // Don't show error for network issues during validation
  }
}

const validatePassword = () => {
  passwordError.value = ''
  
  if (!password.value) {
    passwordError.value = 'Password is required'
  } else if (password.value.length < 6) {
    passwordError.value = 'Password must be at least 6 characters'
  } else if (!isPasswordValid.value) {
    passwordError.value = 'Password does not meet requirements'
  }
}

// API configuration is imported from @/config/api

// Form actions
const toggleSignUp = () => {
  isSignUp.value = !isSignUp.value
  phoneError.value = ''
  passwordError.value = ''
}

const handleSubmit = async () => {
  validatePhone()
  validatePassword()
  
  if (!isFormValid.value) return
  
  isSubmitting.value = true
  
  try {
    const endpoint = isSignUp.value ? '/auth/register' : '/auth/signin'
    const url = `${API_BASE_URL}${endpoint}`
    
    console.log('Making request to:', url)
    console.log('Request body:', {
      phone: phoneNumber.value,
      password: password.value
    })
    
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        phone: phoneNumber.value,
        password: password.value
      })
    })
    
    console.log('Response status:', response.status)
    console.log('Response headers:', response.headers)
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    
    const data = await response.json()
    console.log('Response data:', data)
    
    if (data.success) {
      if (isSignUp.value) {
        // Registration successful - navigate to OTP verification with OTP code
        router.push({ 
          name: 'OTPVerification',
          query: {
            phone: phoneNumber.value,
            password: password.value,
            otp: data.otp_code // Pass the received OTP for testing
          }
        })
      } else {
        // Sign in successful - store owner data and navigate to home
        localStorage.setItem('user', JSON.stringify(data.owner))
        router.push({ name: 'Home' })
      }
    } else {
      // Show error message
      phoneError.value = data.message
    }
    
  } catch (error) {
    console.error('Authentication error:', error)
    phoneError.value = 'Network error. Please check your connection.'
  } finally {
    isSubmitting.value = false
  }
}
</script>