<template>
  <div class="home-page">
    <!-- Header Card -->
    <div class="header-card">
      <div class="welcome-section">
        <div class="welcome-text">
          <h1>{{ t('welcomeBack') }}</h1>
          <p class="date">{{ trackingPeriod }}</p>
          <LanguageToggle />
        </div>
        <div class="header-actions">
          <div class="app-icon">🚀</div>
        </div>
      </div>
      <div class="total-revenue">
        <p class="revenue-label">{{ t('totalRevenue') }}</p>
        <h2 class="revenue-amount">{{ formatCurrency(totalRevenue) }}</h2>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="loading-container">
      <div class="loading-spinner"></div>
      <p>{{ t('loading') }}</p>
    </div>

    <!-- Content when not loading -->
    <template v-else>
      <!-- Merchant Selection Section -->
      <div class="merchant-selection" v-if="merchants.length > 0">
        <div class="merchant-nav-container">
          <button
            class="nav-arrow left"
            @click="scrollMerchants('left')"
            :disabled="!canScrollLeft"
          >
            <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
              <path
                d="M12.707 14.707a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 1.414L9.414 10l3.293 3.293a1 1 0 010 1.414z"
              />
            </svg>
          </button>
          <div class="merchant-buttons-wrapper" ref="merchantScroll">
            <div class="merchant-buttons">
              <button
                v-for="merchant in merchants"
                :key="merchant.id"
                class="merchant-button"
                :class="{ active: selectedMerchantId === merchant.id }"
                @click="selectMerchant(merchant.id)"
              >
                <div class="merchant-button-icon">
                  {{ merchant.name.charAt(0).toUpperCase() }}
                </div>
                <span>{{ merchant.name }}</span>
              </button>
            </div>
          </div>
          <button
            class="nav-arrow right"
            @click="scrollMerchants('right')"
            :disabled="!canScrollRight"
          >
            <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
              <path
                d="M7.293 14.707a1 1 0 001.414 0l4-4a1 1 0 000-1.414l-4-4a1 1 0 00-1.414 1.414L10.586 10l-3.293 3.293a1 1 0 000 1.414z"
              />
            </svg>
          </button>
        </div>
      </div>
      <div class="starter-guide">
        <h3>{{ t('gettingStartedGuide') }}</h3>
        <!-- Sliding Cards Container -->
        <div class="guide-carousel-container">
          <div class="guide-steps-wrapper">
            <div
              class="guide-steps"
              :style="`transform: translateX(-${currentGuideStep * 100}%)`"
              @touchstart="handleTouchStart"
              @touchmove="handleTouchMove"
              @touchend="handleTouchEnd"
            >
              <div class="guide-step">
                <span class="step-number">1</span>
                <div class="step-content">
                  <h4>{{ t('step1Title') }}</h4>
                  <p>{{ t('step1Desc') }}</p>
                </div>
              </div>
              <div class="guide-step">
                <span class="step-number">2</span>
                <div class="step-content">
                  <h4>{{ t('step2Title') }}</h4>
                  <p>{{ t('step2Desc') }}</p>
                </div>
              </div>
              <div class="guide-step">
                <span class="step-number">3</span>
                <div class="step-content">
                  <h4>{{ t('step3Title') }}</h4>
                  <p>{{ t('step3Desc') }}</p>
                </div>
              </div>
              <div class="guide-step">
                <span class="step-number">4</span>
                <div class="step-content">
                  <h4>{{ t('step4Title') }}</h4>
                  <p>{{ t('step4Desc') }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
        <!-- Bullet Navigation -->
        <div class="guide-bullets">
          <button
            v-for="(step, index) in 4"
            :key="index"
            @click="currentGuideStep = index"
            class="bullet"
            :class="{ active: currentGuideStep === index }"
          ></button>
        </div>
      </div>

      <!-- Merchant Section for New Owners -->
      <div v-if="merchants.length === 0" class="merchant-create-section">
        <div class="section-header">
          <h2>{{ t('merchantDetails') }}</h2>
        </div>
        <div class="empty-merchant-card">
          <div class="empty-merchant-icon">
            <svg
              width="40"
              height="40"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="1.5"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M13.5 21v-7.5a.75.75 0 01.75-.75h3a.75.75 0 01.75.75V21m-4.5 0H2.36m11.14 0H18m0 0h3.64m-1.39 0V9.349m-16.5 11.65V9.35m0 0a3.001 3.001 0 003.75-.615A2.993 2.993 0 009.75 9.75c.896 0 1.7-.393 2.25-1.016a2.993 2.993 0 002.25 1.016c.896 0 1.7-.393 2.25-1.016a3.001 3.001 0 003.75.614m-16.5 0a3.004 3.004 0 01-.621-4.72L4.318 3.44A1.5 1.5 0 015.378 3h13.243a1.5 1.5 0 011.06.44l1.19 1.189a3 3 0 01-.621 4.72m-13.5 8.65h3.75a.75.75 0 00.75-.75V13.5a.75.75 0 00-.75-.75H6.75a.75.75 0 00-.75.75v3.75c0 .415.336.75.75.75z"
              />
            </svg>
          </div>
          <h3>{{ t('noMerchantsYet') }}</h3>
          <p>{{ t('startByCreating') }}</p>
          <button @click="showCreateMerchantModal = true" class="create-merchant-btn">
            <svg width="18" height="18" viewBox="0 0 20 20" fill="currentColor">
              <path
                fill-rule="evenodd"
                d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
                clip-rule="evenodd"
              />
            </svg>
            {{ t('createMerchant') }}
          </button>
        </div>
      </div>

      <!-- Merchant Details Section -->
      <div v-if="merchants.length > 0" class="merchant-details-section">
        <div class="section-header">
          <h2>{{ t('merchantDetails') }}</h2>
          <div class="action-buttons">
            <button class="create-btn" @click="showCreateMerchantModal = true">
              <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
                <path
                  fill-rule="evenodd"
                  d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
                  clip-rule="evenodd"
                />
              </svg>
              {{ t('create') }}
            </button>
            <!--ignore this time... -->
            <!-- <button class="filter-btn" @click="openFilter">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
              <path d="M5 4a1 1 0 00-2 0v7.268a2 2 0 000 3.464V16a1 1 0 102 0v-1.268a2 2 0 000-3.464V4zM11 4a1 1 0 10-2 0v1.268a2 2 0 000 3.464V16a1 1 0 102 0V8.732a2 2 0 000-3.464V4zM16 3a1 1 0 011 1v7.268a2 2 0 010 3.464V16a1 1 0 11-2 0v-1.268a2 2 0 010-3.464V4a1 1 0 011-1z"/>
            </svg>
            {{ t('filter') }}
          </button>
          <button class="sort-btn" @click="toggleSort">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
              <path d="M3 3a1 1 0 000 2h11a1 1 0 100-2H3zM3 7a1 1 0 000 2h7a1 1 0 100-2H3zM3 11a1 1 0 100 2h4a1 1 0 100-2H3z"/>
            </svg>
            {{ t('sort') }}
          </button> -->
          </div>
        </div>

        <!-- Merchant Cards -->
        <div class="merchant-cards" :id="`merchant-${selectedMerchantId}`">
          <div
            v-for="merchant in filteredAndSortedMerchants"
            :key="merchant.id"
            class="merchant-card"
            :class="{ highlighted: merchant.id === selectedMerchantId }"
          >
            <div class="merchant-card-header">
              <div class="merchant-icon">
                {{ merchant.name.charAt(0).toUpperCase() }}
              </div>
              <div class="merchant-info">
                <h3>{{ merchant.name }}</h3>
                <p class="merchant-category">{{ t('merchant') }}</p>
              </div>
              <div class="more-menu-wrapper">
                <button class="more-btn" @click="toggleMerchantMenu(merchant.id)">
                  <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      d="M10 6a2 2 0 110-4 2 2 0 010 4zM10 12a2 2 0 110-4 2 2 0 010 4zM10 18a2 2 0 110-4 2 2 0 010 4z"
                    />
                  </svg>
                </button>
                <div v-if="activeMerchantMenu === merchant.id" class="dropdown-menu">
                  <button @click="deleteMerchant(merchant)" class="dropdown-item delete-item">
                    <svg width="16" height="16" viewBox="0 0 20 20" fill="currentColor">
                      <path
                        fill-rule="evenodd"
                        d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                        clip-rule="evenodd"
                      />
                    </svg>
                    {{ t('deleteMerchant') }}
                  </button>
                </div>
              </div>
            </div>

            <div class="merchant-metrics">
              <div class="stat-section">
                <div class="stat-label">{{ t('income') }}</div>
                <span class="stat-value income">{{ formatCurrency(merchant.income) }}</span>
              </div>

              <div class="stat-section">
                <div class="stat-label">{{ t('outcome') }}</div>
                <span class="stat-value outcome">{{ formatCurrency(merchant.outcome) }}</span>
              </div>

              <div class="stat-section">
                <div class="stat-label">{{ t('revenue') }}</div>
                <span class="stat-value revenue">{{ formatCurrency(merchant.revenue) }}</span>
              </div>
            </div>

            <div class="merchant-footer">
              <div class="merchant-actions">
                <button @click="goToTracking(merchant)" class="tracking-btn">
                  <svg width="16" height="16" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      fill-rule="evenodd"
                      d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11.707 4.293a1 1 0 00-1.414-1.414L10 9.172 8.707 7.879a1 1 0 00-1.414 0l-2 2a1 1 0 101.414 1.414L8 10l1.293 1.293a1 1 0 001.414 0l4-4z"
                      clip-rule="evenodd"
                    />
                  </svg>
                  {{ t('tracking') }}
                </button>
                <button @click="goToManage(merchant)" class="manage-btn">
                  <svg width="16" height="16" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      fill-rule="evenodd"
                      d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z"
                      clip-rule="evenodd"
                    />
                  </svg>
                  {{ t('manage') }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>

  <!-- Delete Merchant Confirmation Modal -->
  <div v-if="showDeleteModal" class="modal-overlay" @click.self="closeDeleteModal">
    <div class="modal-content delete-modal">
      <div class="modal-header">
        <h3>{{ t('deleteMerchant') }}</h3>
        <button @click="closeDeleteModal" class="modal-close">
          <svg
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path d="M6 6l12 12M6 18L18 6" />
          </svg>
        </button>
      </div>

      <div class="modal-body">
        <div class="delete-warning">
          <div class="warning-icon">
            <svg
              width="48"
              height="48"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
              />
            </svg>
          </div>
          <h4>{{ t('areYouSureDelete') }}</h4>
          <p>{{ t('actionCannotBeUndone') }}</p>
        </div>
      </div>

      <div class="modal-footer">
        <button @click="closeDeleteModal" class="btn-secondary" :disabled="isDeletingMerchant">
          {{ t('cancel') }}
        </button>
        <button @click="confirmDeleteMerchant" class="btn-danger" :disabled="isDeletingMerchant">
          <span v-if="isDeletingMerchant">{{ t('deleting') }}</span>
          <span v-else>{{ t('deleteMerchant') }}</span>
        </button>
      </div>
    </div>
  </div>

  <!-- Create Merchant Modal -->
  <div v-if="showCreateMerchantModal" class="modal-overlay" @click.self="closeCreateMerchantModal">
    <div class="modal-content">
      <div class="modal-header">
        <h3>{{ t('createMerchant') }}</h3>
        <button @click="closeCreateMerchantModal" class="modal-close">
          <svg
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path d="M6 6l12 12M6 18L18 6" />
          </svg>
        </button>
      </div>

      <div class="modal-body">
        <div class="form-group">
          <label for="merchantName">{{ t('merchantName') }}</label>
          <input
            id="merchantName"
            v-model="newMerchantName"
            type="text"
            :placeholder="t('enterMerchantName')"
            class="form-input"
            @keyup.enter="createMerchant"
            :disabled="isCreatingMerchant"
          />
          <p v-if="merchantError" class="error-message">{{ merchantError }}</p>
        </div>
      </div>

      <div class="modal-footer">
        <button
          @click="closeCreateMerchantModal"
          class="btn-secondary"
          :disabled="isCreatingMerchant"
        >
          {{ t('cancel') }}
        </button>
        <button
          @click="createMerchant"
          class="btn-primary"
          :disabled="!newMerchantName || isCreatingMerchant"
        >
          <span v-if="isCreatingMerchant">{{ t('loading') }}</span>
          <span v-else>{{ t('createMerchant') }}</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import API_BASE_URL from '@/config/api'
import { useLanguage } from '@/composables/useLanguage'
import LanguageToggle from '@/components/LanguageToggle.vue'

// Router
const router = useRouter()

// Language
const { t, formatCurrency } = useLanguage()

// Reactive data
const merchants = ref([])
const selectedMerchantId = ref(null)
const canScrollLeft = ref(false)
const canScrollRight = ref(false)
const merchantScroll = ref(null)
const sortBy = ref('name')
const sortOrder = ref('asc')
const filterOptions = ref({
  minRevenue: 0,
  category: 'all',
})
const isLoading = ref(true)
const dashboardData = ref(null)
const currentGuideStep = ref(0)

// Touch handling for mobile swipe
const touchStartX = ref(0)
const touchEndX = ref(0)
const isSwiping = ref(false)

// Create Merchant Modal
const showCreateMerchantModal = ref(false)
const newMerchantName = ref('')
const isCreatingMerchant = ref(false)
const merchantError = ref('')

// Delete Merchant Modal
const showDeleteModal = ref(false)
const merchantToDelete = ref(null)
const isDeletingMerchant = ref(false)
const activeMerchantMenu = ref(null)

// Computed properties
const trackingPeriod = computed(() => {
  // Check if we have dashboard data with first order date
  if (dashboardData.value?.first_order_date) {
    const startDate = new Date(dashboardData.value.first_order_date)
    const endDate = new Date()

    const formatDate = (date) => {
      const day = date.getDate()
      const month = date.toLocaleDateString('en-US', { month: 'short' })
      const year = date.getFullYear()
      return `${day} ${month} ${year}`
    }

    return `${formatDate(startDate)} - ${formatDate(endDate)}`
  } else {
    // No orders yet, just show today
    const currentDate = new Date()
    const day = currentDate.getDate()
    const month = currentDate.toLocaleDateString('en-US', { month: 'short' })
    const year = currentDate.getFullYear()
    return `Today, ${day} ${month} ${year}`
  }
})

const totalRevenue = computed(() => {
  return dashboardData.value?.total_revenue || 0
})

const filteredAndSortedMerchants = computed(() => {
  let result = [...merchants.value]

  // Apply filters
  if (filterOptions.value.minRevenue > 0) {
    result = result.filter((m) => m.revenue >= filterOptions.value.minRevenue)
  }
  if (filterOptions.value.category !== 'all') {
    result = result.filter((m) => m.category === filterOptions.value.category)
  }

  // Apply sorting
  result.sort((a, b) => {
    let comparison = 0
    switch (sortBy.value) {
      case 'name':
        comparison = a.name.localeCompare(b.name)
        break
      case 'revenue':
        comparison = a.revenue - b.revenue
        break
      case 'income':
        comparison = a.income - b.income
        break
      case 'outcome':
        comparison = a.outcome - b.outcome
        break
    }
    return sortOrder.value === 'asc' ? comparison : -comparison
  })

  return result
})

// Methods
const scrollMerchants = (direction) => {
  if (!merchantScroll.value) return

  const scrollAmount = 200
  const currentScroll = merchantScroll.value.scrollLeft

  if (direction === 'left') {
    merchantScroll.value.scrollTo({
      left: currentScroll - scrollAmount,
      behavior: 'smooth',
    })
  } else {
    merchantScroll.value.scrollTo({
      left: currentScroll + scrollAmount,
      behavior: 'smooth',
    })
  }

  setTimeout(updateScrollButtons, 300)
}

const updateScrollButtons = () => {
  if (!merchantScroll.value) return

  canScrollLeft.value = merchantScroll.value.scrollLeft > 0
  canScrollRight.value =
    merchantScroll.value.scrollLeft <
    merchantScroll.value.scrollWidth - merchantScroll.value.clientWidth
}

const selectMerchant = async (merchantId) => {
  selectedMerchantId.value = merchantId
  await nextTick()

  const element = document.getElementById(`merchant-${merchantId}`)
  if (element) {
    element.scrollIntoView({ behavior: 'smooth', block: 'center' })
  }
}

const openFilter = () => {
  console.log('Filter clicked')
  // Implement filter modal/dropdown
}

const toggleSort = () => {
  if (sortOrder.value === 'asc') {
    sortOrder.value = 'desc'
  } else if (sortBy.value === 'name') {
    sortBy.value = 'revenue'
    sortOrder.value = 'desc'
  } else if (sortBy.value === 'revenue') {
    sortBy.value = 'income'
    sortOrder.value = 'desc'
  } else if (sortBy.value === 'income') {
    sortBy.value = 'outcome'
    sortOrder.value = 'desc'
  } else {
    sortBy.value = 'name'
    sortOrder.value = 'asc'
  }
}

// Touch/swipe handlers for mobile
const handleTouchStart = (e) => {
  touchStartX.value = e.changedTouches[0].screenX
  isSwiping.value = true
}

const handleTouchMove = (e) => {
  if (!isSwiping.value) return
  touchEndX.value = e.changedTouches[0].screenX
}

const handleTouchEnd = () => {
  if (!isSwiping.value) return

  const swipeThreshold = 50 // Minimum distance for swipe
  const diff = touchStartX.value - touchEndX.value

  if (Math.abs(diff) > swipeThreshold) {
    if (diff > 0) {
      // Swiped left - go to next step
      if (currentGuideStep.value < 3) {
        currentGuideStep.value++
      }
    } else {
      // Swiped right - go to previous step
      if (currentGuideStep.value > 0) {
        currentGuideStep.value--
      }
    }
  }

  isSwiping.value = false
  touchStartX.value = 0
  touchEndX.value = 0
}

// Merchant Menu methods
const toggleMerchantMenu = (merchantId) => {
  if (activeMerchantMenu.value === merchantId) {
    activeMerchantMenu.value = null
  } else {
    activeMerchantMenu.value = merchantId
  }
}

// Delete Merchant methods
const deleteMerchant = (merchant) => {
  merchantToDelete.value = merchant
  showDeleteModal.value = true
  activeMerchantMenu.value = null // Close dropdown
}

const closeDeleteModal = () => {
  showDeleteModal.value = false
  merchantToDelete.value = null
  isDeletingMerchant.value = false
}

const confirmDeleteMerchant = async () => {
  if (!merchantToDelete.value) return

  isDeletingMerchant.value = true

  try {
    const response = await fetch(`${API_BASE_URL}/merchants/${merchantToDelete.value.id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!response.ok) {
      let errorMessage = 'Failed to delete merchant'
      try {
        const errorData = await response.json()
        errorMessage = errorData.message || errorMessage
      } catch {
        errorMessage = `Server error: ${response.status}`
      }
      throw new Error(errorMessage)
    }

    console.log('Merchant deleted successfully')

    // Close modal and refresh dashboard
    closeDeleteModal()
    await fetchDashboardData()
  } catch (error) {
    console.error('Error deleting merchant:', error)
    alert(error.message || 'Failed to delete merchant. Please try again.')
  } finally {
    isDeletingMerchant.value = false
  }
}

// Create Merchant methods
const closeCreateMerchantModal = () => {
  showCreateMerchantModal.value = false
  newMerchantName.value = ''
  merchantError.value = ''
}

const createMerchant = async () => {
  if (!newMerchantName.value || newMerchantName.value.trim() === '') {
    merchantError.value = 'Please enter a merchant name'
    return
  }

  isCreatingMerchant.value = true
  merchantError.value = ''

  try {
    // Get owner info from localStorage
    const userStr = localStorage.getItem('user')
    if (!userStr) {
      merchantError.value = 'User not found. Please login again.'
      isCreatingMerchant.value = false
      return
    }

    const user = JSON.parse(userStr)
    const ownerID = user.owner_id

    // Create merchant via API
    console.log('Creating merchant with:', {
      owner_id: ownerID,
      merchant_name: newMerchantName.value.trim(),
    })

    const response = await fetch(`${API_BASE_URL}/merchants`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        owner_id: ownerID,
        merchant_name: newMerchantName.value.trim(),
      }),
    })

    if (!response.ok) {
      let errorMessage = 'Failed to create merchant'
      try {
        const errorData = await response.json()
        errorMessage = errorData.message || errorMessage
      } catch {
        errorMessage = `Server error: ${response.status}`
      }
      throw new Error(errorMessage)
    }

    const data = await response.json()
    console.log('Merchant created:', data)

    // Close modal and refresh dashboard
    closeCreateMerchantModal()
    await fetchDashboardData()
  } catch (error) {
    console.error('Error creating merchant:', error)
    if (error.message === 'Failed to fetch') {
      merchantError.value = 'Network error. Please check your connection and try again.'
    } else {
      merchantError.value = error.message || 'Failed to create merchant. Please try again.'
    }
  } finally {
    isCreatingMerchant.value = false
  }
}

// Fetch dashboard data from API
const fetchDashboardData = async () => {
  try {
    // Get owner info from localStorage
    const userStr = localStorage.getItem('user')
    if (!userStr) {
      console.error('No user data found')
      isLoading.value = false
      return
    }

    const user = JSON.parse(userStr)
    const ownerID = user.owner_id

    // Fetch merchants first
    const merchantResponse = await fetch(`${API_BASE_URL}/merchants?owner_id=${ownerID}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!merchantResponse.ok) {
      throw new Error('Failed to fetch merchants')
    }

    const merchantData = await merchantResponse.json()
    console.log('Fetched merchants:', merchantData)

    // Calculate totals from daily tracking data for each merchant
    const merchantsWithTotals = []
    let totalIncome = 0
    let totalOutcome = 0
    let totalRevenue = 0
    let earliestDate = null

    for (const merchant of merchantData) {
      try {
        // Fetch daily tracking data for this merchant (last 90 days)
        const endDate = new Date().toISOString().split('T')[0]
        const startDate = new Date(Date.now() - 90 * 24 * 60 * 60 * 1000)
          .toISOString()
          .split('T')[0]

        const trackingResponse = await fetch(
          `${API_BASE_URL}/daily-tracking?merchant_id=${merchant.merchant_id}&start_date=${startDate}&end_date=${endDate}`,
        )

        let merchantIncome = 0
        let merchantOutcome = 0
        let merchantRevenue = 0

        if (trackingResponse.ok) {
          const trackingData = await trackingResponse.json()
          console.log(`Tracking data for merchant ${merchant.merchant_name}:`, trackingData)

          if (trackingData && Array.isArray(trackingData)) {
            // Calculate totals for this merchant
            trackingData.forEach((tracking) => {
              const income = tracking.total_income || 0
              const outcome = tracking.total_outcome || 0
              const revenue = income - outcome // Calculate revenue as income - outcome

              merchantIncome += income
              merchantOutcome += outcome
              merchantRevenue += revenue

              // Track earliest date
              if (tracking.tracking_date) {
                const trackingDate = new Date(tracking.tracking_date)
                if (!earliestDate || trackingDate < earliestDate) {
                  earliestDate = trackingDate
                }
              }
            })
          }
        }

        // Add to totals
        totalIncome += merchantIncome
        totalOutcome += merchantOutcome
        totalRevenue += merchantRevenue

        // Add merchant with calculated totals
        merchantsWithTotals.push({
          id: merchant.merchant_id,
          name: merchant.merchant_name,
          category: merchant.category || 'Restaurant',
          income: merchantIncome,
          outcome: merchantOutcome,
          revenue: merchantRevenue,
          incomeChange: 0, // We can calculate this later if needed
          outcomeChange: 0,
          revenueChange: 0,
        })
      } catch (merchantError) {
        console.error(`Error fetching data for merchant ${merchant.merchant_name}:`, merchantError)
        // Add merchant with zero values if there's an error
        merchantsWithTotals.push({
          id: merchant.merchant_id,
          name: merchant.merchant_name,
          category: merchant.category || 'Restaurant',
          income: 0,
          outcome: 0,
          revenue: 0,
          incomeChange: 0,
          outcomeChange: 0,
          revenueChange: 0,
        })
      }
    }

    // Update dashboard data with calculated totals
    dashboardData.value = {
      total_revenue: totalRevenue,
      total_income: totalIncome,
      total_outcome: totalOutcome,
      total_merchants: merchantsWithTotals.length,
      first_order_date: earliestDate ? earliestDate.toISOString().split('T')[0] : null,
      merchants: merchantsWithTotals,
    }

    merchants.value = merchantsWithTotals

    if (merchants.value.length > 0) {
      selectedMerchantId.value = merchants.value[0].id
    }

    console.log('Final dashboard data:', dashboardData.value)
  } catch (error) {
    console.error('Error fetching dashboard data:', error)
    // Keep merchants empty for new users
    merchants.value = []
    dashboardData.value = {
      total_revenue: 0,
      total_income: 0,
      total_outcome: 0,
      total_merchants: 0,
      first_order_date: null,
      merchants: [],
    }
  } finally {
    isLoading.value = false
  }
}

// Navigation to manage merchant page
const goToManage = (merchant) => {
  router.push({
    name: 'ManageMerchant',
    params: {
      merchantId: merchant.id,
    },
    query: {
      merchantName: merchant.name,
    },
  })
}

const goToTracking = (merchant) => {
  console.log('🔥 TRACKING BUTTON CLICKED!')
  console.log('Merchant data:', merchant)

  try {
    console.log('📍 Navigating to TrackingPage...')
    router.push({
      name: 'TrackingPage',
      params: {
        merchantId: merchant.id,
      },
      query: {
        merchantName: merchant.name,
      },
    })
    console.log('✅ Router.push called successfully')
  } catch (error) {
    console.error('❌ Error in goToTracking:', error)
  }
}

// Click outside handler to close dropdown
const handleClickOutside = (event) => {
  const dropdownMenu = event.target.closest('.more-menu-wrapper')
  if (!dropdownMenu) {
    activeMerchantMenu.value = null
  }
}

// Initialize on mount
onMounted(async () => {
  await fetchDashboardData()

  // Setup scroll button visibility check
  if (merchantScroll.value) {
    merchantScroll.value.addEventListener('scroll', updateScrollButtons)
    updateScrollButtons()
  }

  // Add click outside listener
  document.addEventListener('click', handleClickOutside)
})

// Cleanup on unmount
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.home-page {
  padding: 24px;
  max-width: 1400px;
  margin: 0 auto;
  background: #f9fafb;
  min-height: 100vh;
}

/* Header Card */
.header-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 20px;
  padding: 32px;
  color: white;
  margin-bottom: 32px;
  box-shadow: 0 10px 30px rgba(102, 126, 234, 0.3);
}

.welcome-section {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 24px;
}

.welcome-text h1 {
  font-size: 32px;
  font-weight: 700;
  margin: 0 0 8px 0;
}

.welcome-text .date {
  font-size: 14px;
  opacity: 0.9;
  margin-bottom: 20px;
}

.app-icon {
  background: rgba(255, 255, 255, 0.2);
  padding: 12px;
  border-radius: 16px;
  backdrop-filter: blur(10px);
}

.total-revenue {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 20px;
  backdrop-filter: blur(10px);
}

.revenue-label {
  font-size: 14px;
  margin: 0 0 8px 0;
  opacity: 0.9;
}

.revenue-amount {
  font-size: 36px;
  font-weight: 700;
  margin: 0 0 4px 0;
}

.revenue-period {
  font-size: 12px;
  opacity: 0.8;
  margin: 0;
}

/* Merchant Selection */
.merchant-selection {
  margin-bottom: 32px;
}

.merchant-nav-container {
  display: flex;
  align-items: center;
  gap: 12px;
}

.nav-arrow {
  flex-shrink: 0;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  border: 1px solid #e5e7eb;
  background: white;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
  color: #6b7280;
}

.nav-arrow:hover:not(:disabled) {
  background: #f3f4f6;
  color: #111827;
}

.nav-arrow:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}

.merchant-buttons-wrapper {
  flex: 1;
  overflow-x: auto;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.merchant-buttons-wrapper::-webkit-scrollbar {
  display: none;
}

.merchant-buttons {
  display: flex;
  gap: 12px;
  padding: 4px;
}

.merchant-button {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  background: white;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.2s;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
}

.merchant-button:hover {
  border-color: #4f46e5;
  background: #f9f9ff;
}

.merchant-button.active {
  background: #4f46e5;
  border-color: #4f46e5;
  color: white;
}

.merchant-button-icon {
  width: 32px;
  height: 32px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: 600;
}

.merchant-button.active .merchant-button-icon {
  background: white;
  color: #4f46e5;
}

/* Empty State */
.empty-state {
  text-align: center;
  padding: 64px 32px;
  background: white;
  border-radius: 20px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.empty-icon {
  margin: 0 auto 24px;
  width: fit-content;
}

.empty-state h3 {
  font-size: 20px;
  font-weight: 600;
  margin: 0 0 8px 0;
  color: #111827;
}

.empty-state p {
  font-size: 14px;
  color: #6b7280;
  margin: 0 0 24px 0;
}

.create-merchant-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  background: #4f46e5;
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.create-merchant-btn:hover {
  background: #4338ca;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(79, 70, 229, 0.3);
}

/* Merchant Details Section */
.merchant-details-section {
  margin-top: 32px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.section-header h2 {
  font-size: 24px;
  font-weight: 700;
  color: #111827;
  margin: 0;
}

.action-buttons {
  display: flex;
  gap: 12px;
}

.create-btn,
.filter-btn,
.sort-btn {
  display: flex;
  align-items: left;
  gap: 6px;
  padding: 10px 16px;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
  cursor: pointer;
  transition: all 0.2s;
}

.create-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.2);
}

.create-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.filter-btn:hover,
.sort-btn:hover {
  background: #f3f4f6;
  border-color: #d1d5db;
  transform: translateY(-1px);
}

/* Merchant Cards */
.merchant-cards {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(380px, 1fr));
  gap: 24px;
}

.merchant-card {
  background: white;
  border-radius: 16px;
  padding: 24px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  transition: all 0.3s;
}

.merchant-card:hover {
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
}

.merchant-card.highlighted {
  border: 2px solid #4f46e5;
  box-shadow: 0 0 0 4px rgba(79, 70, 229, 0.1);
}

.merchant-card-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 24px;
}

.merchant-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 20px;
  font-weight: 600;
}

.merchant-info {
  flex: 1;
}

.merchant-info h3 {
  font-size: 18px;
  font-weight: 600;
  margin: 0 0 4px 0;
  color: #111827;
}

.merchant-category {
  font-size: 12px;
  color: #6b7280;
  margin: 0;
}

.more-btn {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  border: none;
  background: transparent;
  color: #9ca3af;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.more-btn:hover {
  background: #f3f4f6;
  color: #374151;
}

.merchant-metrics {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  gap: 12px;
  margin-bottom: 20px;
}

.merchant-metrics .stat-section {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 6px;
  padding: 12px 8px;
}

.stat-label {
  font-size: 12px;
  color: #6b7280;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.stat-value {
  font-size: 15px;
  font-weight: 700;
  padding: 8px 10px;
  border-radius: 8px;
  background: #f9fafb;
  white-space: nowrap;
  min-height: 36px;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.stat-value.income {
  color: #10b981;
  background: #f0fdf4;
}

.stat-value.outcome {
  color: #ef4444;
  background: #fef2f2;
}

.stat-value.revenue {
  color: #4f46e5;
  background: #eef2ff;
}

.metric-change {
  font-size: 11px;
  font-weight: 500;
  margin: 0;
}

.metric-change.positive {
  color: #10b981;
}

.metric-change.negative {
  color: #ef4444;
}

.merchant-footer {
  padding-top: 16px;
  border-top: 1px solid #e5e7eb;
}

.merchant-actions {
  display: flex;
  gap: 12px;
  width: 100%;
}

.tracking-btn,
.manage-btn {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  padding: 10px;
  border: none;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  border-radius: 8px;
}

.tracking-btn {
  background: #10b981;
  color: white;
}

.tracking-btn:hover {
  background: #059669;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.manage-btn {
  background: #3b82f6;
  color: white;
}

.manage-btn:hover {
  background: #2563eb;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

/* Loading State */
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px;
  background: white;
  border-radius: 20px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  margin-top: 32px;
}

.loading-spinner {
  width: 48px;
  height: 48px;
  border: 4px solid #e5e7eb;
  border-top-color: #4f46e5;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 16px;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.loading-container p {
  color: #6b7280;
  font-size: 14px;
  margin: 0;
}

/* Empty State Styles */
.empty-state-main {
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%);
  border-radius: 20px;
  padding: 48px;
  text-align: center;
  margin-bottom: 32px;
}

.empty-icon {
  width: 120px;
  height: 120px;
  margin: 0 auto 24px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 60px;
}

.empty-state-main h2 {
  font-size: 28px;
  font-weight: 700;
  color: #111827;
  margin: 0 0 12px 0;
}

.empty-state-main p {
  font-size: 16px;
  color: #6b7280;
  margin: 0 0 32px 0;
  max-width: 500px;
  margin-left: auto;
  margin-right: auto;
}

.get-started-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 16px 32px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.get-started-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
}

/* Starter Guide */
.starter-guide {
  background: white;
  border-radius: 20px;
  padding: 32px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  margin-bottom: 32px;
}

.starter-guide h3 {
  font-size: 20px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 24px 0;
}

/* Guide Carousel */
.guide-carousel-container {
  position: relative;
  overflow: hidden;
  border-radius: 12px;
  margin-bottom: 20px;
}

.guide-steps-wrapper {
  overflow: hidden;
  width: 100%;
}

.guide-steps {
  display: flex;
  transition: transform 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  touch-action: pan-y;
  user-select: none;
  -webkit-user-select: none;
}

.guide-step {
  flex: 0 0 100%;
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 24px;
  background: #f9fafb;
  min-height: 100px;
}

.guide-step .step-content {
  flex: 1;
}

.guide-step .step-content h4 {
  font-size: 16px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 4px 0;
}

.guide-step .step-content p {
  font-size: 14px;
  color: #6b7280;
  margin: 0;
}

/* Bullet Navigation */
.guide-bullets {
  display: flex;
  justify-content: center;
  gap: 8px;
  padding: 12px 0;
}

.guide-bullets .bullet {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #d1d5db;
  border: none;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  padding: 0;
}

.guide-bullets .bullet:hover {
  background: #9ca3af;
}

.guide-bullets .bullet.active {
  width: 24px;
  border-radius: 4px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.step-number {
  flex-shrink: 0;
  width: 32px;
  height: 32px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 14px;
}

.step-content h4 {
  font-size: 16px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 4px 0;
}

.step-content p {
  font-size: 14px;
  color: #6b7280;
  margin: 0;
  line-height: 1.5;
}

/* Empty State for Merchant Selection */
.empty-merchant-button {
  padding: 16px 24px;
  background: #f3f4f6;
  border: 2px dashed #d1d5db;
  border-radius: 12px;
  color: #6b7280;
  font-size: 14px;
  cursor: not-allowed;
  opacity: 0.7;
}

/* Disabled Filter/Sort Buttons */
.action-button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  background: #f3f4f6;
  color: #9ca3af;
}

.action-button:disabled:hover {
  background: #f3f4f6;
  transform: none;
}

/* Merchant Create Section */
.merchant-create-section {
  margin-bottom: 32px;
}

.empty-merchant-card {
  background: white;
  border-radius: 20px;
  padding: 40px;
  text-align: center;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.07);
  border: 1px solid #e5e7eb;
}

.empty-merchant-icon {
  width: 72px;
  height: 72px;
  margin: 0 auto 20px;
  background: linear-gradient(135deg, #eef2ff 0%, #e0e7ff 100%);
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6366f1;
}

.empty-merchant-card h3 {
  font-size: 18px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 8px 0;
}

.empty-merchant-card p {
  font-size: 14px;
  color: #6b7280;
  margin: 0 0 24px 0;
  line-height: 1.5;
}

.create-merchant-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  box-shadow: 0 4px 14px rgba(102, 126, 234, 0.25);
}

.create-merchant-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.35);
}

.create-merchant-btn:active {
  transform: translateY(0);
}

/* Dropdown Menu */
.more-menu-wrapper {
  position: relative;
}

.dropdown-menu {
  position: absolute;
  right: 0;
  top: 36px;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  z-index: 100;
  min-width: 180px;
  padding: 8px;
}

.dropdown-item {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 12px;
  border: none;
  background: none;
  font-size: 14px;
  color: #374151;
  cursor: pointer;
  transition: all 0.2s;
  border-radius: 6px;
  text-align: left;
}

.dropdown-item:hover {
  background: #f3f4f6;
}

.dropdown-item.delete-item {
  color: #ef4444;
}

.dropdown-item.delete-item:hover {
  background: #fee2e2;
}

/* Delete Modal Styles */
.delete-modal .modal-body {
  padding: 32px 24px;
}

.delete-warning {
  text-align: center;
}

.warning-icon {
  width: 72px;
  height: 72px;
  margin: 0 auto 20px;
  background: #fef2f2;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ef4444;
}

.delete-warning h4 {
  font-size: 18px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 12px 0;
}

.delete-warning p {
  font-size: 14px;
  color: #6b7280;
  margin: 0;
  line-height: 1.5;
}

.delete-warning strong {
  color: #111827;
  font-weight: 600;
}

.btn-danger {
  padding: 10px 20px;
  background: #ef4444;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-danger:hover:not(:disabled) {
  background: #dc2626;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3);
}

.btn-danger:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 20px;
}

.modal-content {
  background: white;
  border-radius: 16px;
  width: 100%;
  max-width: 450px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24px;
  border-bottom: 1px solid #e5e7eb;
}

.modal-header h3 {
  font-size: 20px;
  font-weight: 600;
  color: #111827;
  margin: 0;
}

.modal-close {
  width: 32px;
  height: 32px;
  border: none;
  background: transparent;
  color: #6b7280;
  cursor: pointer;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.modal-close:hover {
  background: #f3f4f6;
  color: #111827;
}

.modal-body {
  padding: 24px;
}

.form-group {
  margin-bottom: 0;
}

.form-group label {
  display: block;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
  margin-bottom: 8px;
}

.form-input {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 14px;
  transition: all 0.2s;
}

.form-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.form-input:disabled {
  background: #f3f4f6;
  cursor: not-allowed;
}

.error-message {
  color: #ef4444;
  font-size: 12px;
  margin-top: 8px;
  margin-bottom: 0;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 24px;
  border-top: 1px solid #e5e7eb;
}

.btn-secondary {
  padding: 10px 20px;
  border: 1px solid #d1d5db;
  background: white;
  color: #374151;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-secondary:hover:not(:disabled) {
  background: #f3f4f6;
}

.btn-primary {
  padding: 10px 20px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.btn-primary:disabled,
.btn-secondary:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Responsive Design */
@media (max-width: 768px) {
  .home-page {
    padding: 16px;
  }

  .header-card {
    padding: 24px;
  }

  .welcome-text h1 {
    font-size: 24px;
  }

  .revenue-amount {
    font-size: 28px;
  }

  .stat-value {
    font-size: 16px;
    min-width: auto;
    padding: 6px 10px;
  }

  .tracking-btn,
  .manage-btn {
    font-size: 13px;
    padding: 8px;
  }

  .section-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .action-buttons {
    width: 100%;
    justify-content: flex-end;
  }

  .empty-state-main {
    padding: 32px 20px;
  }

  .empty-icon {
    width: 100px;
    height: 100px;
    font-size: 48px;
  }

  .empty-state-main h2 {
    font-size: 24px;
  }

  .starter-guide {
    padding: 24px 20px;
  }

  .guide-step {
    padding: 20px 16px;
    min-height: 90px;
  }

  .guide-bullets {
    padding: 8px 0;
  }

  .empty-merchant-card {
    padding: 32px 20px;
  }

  .empty-merchant-icon {
    width: 60px;
    height: 60px;
  }

  .empty-merchant-card h3 {
    font-size: 16px;
  }

  .create-merchant-btn {
    padding: 10px 20px;
    font-size: 13px;
  }
}
</style>
