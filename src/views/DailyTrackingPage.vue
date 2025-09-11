<template>
  <div class="daily-tracking-page">
    <!-- Header -->
    <div class="page-header">
      <button @click="goBack" class="back-btn">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M19 12H5M12 19l-7-7 7-7"/>
        </svg>
        {{ t('back') }}
      </button>
      <div>
        <h1 class="page-title">{{ merchantName }}</h1>
        <p class="page-subtitle">{{ formatTrackingDate(trackingDate) }}</p>
      </div>
    </div>

    <!-- Summary Cards -->
    <div class="summary-cards">
      <div class="summary-card income-card">
        <div class="card-icon">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M12 2v20M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/>
          </svg>
        </div>
        <div class="card-content">
          <p class="card-label">{{ t('totalIncome') }}</p>
          <p class="card-value positive">{{ formatCurrency(totalIncome) }}</p>
        </div>
      </div>

      <div class="summary-card outcome-card">
        <div class="card-icon">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M3 3v18h18"/>
            <path d="M18 9l-5 5-3-3-5 5"/>
          </svg>
        </div>
        <div class="card-content">
          <p class="card-label">{{ t('totalOutcome') }}</p>
          <p class="card-value negative">{{ formatCurrency(totalOutcome) }}</p>
        </div>
      </div>

      <div class="summary-card revenue-card">
        <div class="card-icon">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M16 8l4 4-4 4M8 16l-4-4 4-4M12 2v20"/>
          </svg>
        </div>
        <div class="card-content">
          <p class="card-label">{{ t('netRevenue') }}</p>
          <p class="card-value" :class="{ positive: netRevenue >= 0, negative: netRevenue < 0 }">
            {{ formatCurrency(netRevenue) }}
          </p>
        </div>
      </div>
    </div>

    <!-- Tabs -->
    <div class="tabs-section">
      <div class="tabs">
        <button 
          @click="activeTab = 'income'" 
          class="tab-btn"
          :class="{ active: activeTab === 'income' }"
        >
          <h3>{{ t('Income') }}</h3>
        </button>
        <button 
          @click="activeTab = 'outcome'" 
          class="tab-btn"
          :class="{ active: activeTab === 'outcome', 'outcome-tab': activeTab === 'outcome' }"
        >
          <h3>{{ t('Outcome') }}</h3>
        </button>
      </div>

      <!-- Tab Content -->
      <div class="tab-content">
        <!-- Income Tab -->
        <div v-if="activeTab === 'income'" class="income-section">
          <div class="section-header">
            <h3>{{ t('incomeTransactions') }}</h3>
            <button @click="openAddIncomeModal" class="add-btn">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M12 5v14M5 12h14"/>
              </svg>
              {{ t('addIncome') }}
            </button>
          </div>
          
          <div v-if="incomeItems.length > 0" class="items-list">
            <div v-for="item in incomeItems" :key="item.transaction_id" class="item-card">
              <div class="item-info">
                <h4 class="item-name">{{ item.description }}</h4>
                <p class="item-date">{{ formatDate(item.transaction_date) }}</p>
                <span class="item-category">{{ item.category }}</span>
              </div>
              <div class="item-actions">
                <span class="item-amount income">{{ formatCurrency(item.amount) }}</span>
                <button @click="deleteItem(item.transaction_id, 'income')" class="delete-btn">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M3 6h18M8 6V4a2 2 0 012-2h4a2 2 0 012 2v2M10 11v6M14 11v6M5 6l1 13a2 2 0 002 2h8a2 2 0 002-2l1-13"/>
                  </svg>
                </button>
              </div>
            </div>
          </div>
          <div v-else class="empty-state">
            <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
              <path d="M12 2v20M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/>
            </svg>
            <p>{{ t('noIncomeRecorded') }}</p>
            <button @click="openAddIncomeModal" class="empty-add-btn">{{ t('addFirstIncome') }}</button>
          </div>
        </div>

        <!-- Outcome Tab -->
        <div v-if="activeTab === 'outcome'" class="outcome-section">
          <div class="section-header">
            <h3>{{ t('expenseTransactions') }}</h3>
            <button @click="openAddOutcomeModal" class="add-btn-expense">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M12 5v14M5 12h14"/>
              </svg>
              {{ t('addExpense') }}
            </button>
          </div>
          
          <div v-if="outcomeItems.length > 0" class="items-list">
            <div v-for="item in outcomeItems" :key="item.transaction_id" class="item-card">
              <div class="item-info">
                <h4 class="item-name">{{ item.description }}</h4>
                <p class="item-date">{{ formatDate(item.transaction_date) }}</p>
                <span class="item-category">{{ item.category }}</span>
              </div>
              <div class="item-actions">
                <span class="item-amount outcome">{{ formatCurrency(item.amount) }}</span>
                <button @click="deleteItem(item.transaction_id, 'outcome')" class="delete-btn">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M3 6h18M8 6V4a2 2 0 012-2h4a2 2 0 012 2v2M10 11v6M14 11v6M5 6l1 13a2 2 0 002 2h8a2 2 0 002-2l1-13"/>
                  </svg>
                </button>
              </div>
            </div>
          </div>
          <div v-else class="empty-state">
            <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
              <path d="M3 3v18h18"/>
              <path d="M18 9l-5 5-3-3-5 5"/>
            </svg>
            <p>{{ t('noExpensesRecorded') }}</p>
            <button @click="openAddOutcomeModal" class="empty-add-btn-expense">{{ t('addFirstExpense') }}</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Menu Selection Flow -->
    <!-- Step 1: Search Menu Items -->
    <div v-if="showAddIncomeModal" class="modal-overlay" @click.self="closeAddIncomeModal">
      <div class="modal-content large-modal">
        <div class="modal-header">
          <h2>{{ t('addSalesIncome') }}</h2>
          <button @click="closeAddIncomeModal" class="close-btn">×</button>
        </div>
        
        <div class="modal-body">
          <!-- Search Field -->
          <div class="form-group">
            <label for="menuSearch">{{ t('searchMenuItem') }}</label>
            <input 
              id="menuSearch"
              v-model="searchQuery" 
              type="text" 
              :placeholder="t('typeToSearchMenuItems')"
              class="form-input search-input"
              @input="searchMenuItems"
            />
          </div>
          
          <!-- Menu Items List -->
          <div v-if="filteredMenuItems.length > 0" class="menu-items-list">
            <h4>{{ t('mainMenuItems') }}</h4>
            <div class="items-grid">
              <div 
                v-for="item in filteredMenuItems" 
                :key="item.id"
                @click="selectMenuItem(item)"
                class="menu-item-card selectable"
              >
                <div class="item-info">
                  <h5>{{ item.item_name }}</h5>
                  <div class="price-info">
                    <span class="price-tag general">{{ t('general') }}: {{ formatCurrencyLang(item.general_price_sale) }}</span>
                    <span class="price-tag delivery">{{ t('delivery') }}: {{ formatCurrencyLang(item.delivery_price_sale) }}</span>
                  </div>
                </div>
                <div class="select-indicator">
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M9 18l6-6-6-6"/>
                  </svg>
                </div>
              </div>
            </div>
          </div>
          
          <div v-else-if="searchQuery && !isLoadingMenu" class="empty-search">
            <p>No menu items found for "{{ searchQuery }}"</p>
          </div>
          
          <div v-if="isLoadingMenu" class="loading-state">
            <div class="loading-spinner"></div>
            <p>Loading menu items...</p>
          </div>
        </div>
        
        <div class="modal-footer">
          <button @click="closeAddIncomeModal" class="btn-secondary">{{ t('cancel') }}</button>
        </div>
      </div>
    </div>

    <!-- Step 2: Price Type Selection -->
    <div v-if="showPriceSelectionModal" class="modal-overlay" @click.self="closePriceSelectionModal">
      <div class="modal-content">
        <div class="modal-header">
          <h2>{{ t('selectPriceType') }}</h2>
          <button @click="closePriceSelectionModal" class="close-btn">×</button>
        </div>
        
        <div class="modal-body">
          <div class="selected-item-info">
            <h4>{{ selectedMenuItem?.item_name }}</h4>
            <p>{{ t('choosePriceType') }}</p>
          </div>
          
          <div class="price-selection-grid">
            <button 
              @click="selectPriceType('general')"
              class="price-option-card"
            >
              <div class="price-type-icon general">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
                </svg>
              </div>
              <h5>{{ t('generalSale') }}</h5>
              <p class="price">{{ formatCurrency(selectedMenuItem?.general_price_sale) }}</p>
              <p class="description">{{ t('inStoreOrPickup') }}</p>
            </button>
            
            <button 
              @click="selectPriceType('delivery')"
              class="price-option-card"
            >
              <div class="price-type-icon delivery">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M16 3h3v5l-3-1V3zM2 17h20l-4 4H6l-4-4z"/>
                  <path d="M8 12V3h8v9"/>
                </svg>
              </div>
              <h5>{{ t('deliverySale') }}</h5>
              <p class="price">{{ formatCurrency(selectedMenuItem?.delivery_price_sale) }}</p>
              <p class="description">{{ t('deliveryOrders') }}</p>
            </button>
          </div>
        </div>
        
        <div class="modal-footer">
          <button @click="goBackToMenuSelection" class="btn-secondary">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M19 12H5M12 19l-7-7 7-7"/>
            </svg>
            {{t('back')}}
          </button>
        </div>
      </div>
    </div>

    <!-- Step 3: Extra Items Selection -->
    <div v-if="showExtraSelectionModal" class="modal-overlay" @click.self="closeExtraSelectionModal">
      <div class="modal-content large-modal">
        <div class="modal-header">
          <h2>{{ t('addExtraItems') }}</h2>
          <button @click="closeExtraSelectionModal" class="close-btn">×</button>
        </div>
        
        <div class="modal-body">
          <div class="selected-item-summary">
            <h4>Selected: {{ selectedMenuItem?.item_name }}</h4>
            <p class="selected-price">{{ selectedPriceType === 'general' ? 'General' : 'Delivery' }} Price: {{ formatCurrency(getSelectedPrice()) }}</p>
          </div>
          
          <!-- Extra Items Search -->
          <div class="form-group">
            <label for="extraSearch">{{ t('searchExtraItemsOptional') }}</label>
            <input 
              id="extraSearch"
              v-model="extraSearchQuery" 
              type="text" 
              :placeholder="t('typeToSearchExtraItems')"
              class="form-input search-input"
              @input="searchExtraItems"
            />
          </div>
          
          <!-- Extra Items List -->
          <div v-if="filteredExtraItems.length > 0" class="extra-items-list">
            <h4>{{ t('extraItems') }}</h4>
            <div class="items-grid">
              <div 
                v-for="item in filteredExtraItems" 
                :key="item.id"
                @click="toggleExtraItem(item)"
                class="extra-item-card"
                :class="{ selected: isExtraSelected(item.id) }"
              >
                <div class="item-info">
                  <h5>{{ item.item_name }}</h5>
                  <p class="extra-price">{{ formatCurrency(getExtraPrice(item)) }}</p>
                </div>
                <div class="selection-indicator">
                  <div v-if="getExtraQuantity(item.id) > 0" class="quantity-badge">
                    {{ getExtraQuantity(item.id) }}
                  </div>
                  <div class="checkbox" :class="{ checked: isExtraSelected(item.id) }">
                    <svg v-if="isExtraSelected(item.id)" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M20 6L9 17l-5-5"/>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>
          
          <div v-else-if="extraSearchQuery && !isLoadingExtras" class="empty-search">
            <p>No extra items found for "{{ extraSearchQuery }}"</p>
          </div>
          
          <div v-if="isLoadingExtras" class="loading-state">
            <div class="loading-spinner"></div>
            <p>Loading extra items...</p>
          </div>
        </div>
        
        <div class="modal-footer">
          <button @click="goBackToPriceSelection" class="btn-secondary">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M19 12H5M12 19l-7-7 7-7"/>
            </svg>
            Back to Price
          </button>
          <button @click="proceedToSummary" class="btn-primary">
            Continue
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M5 12h14M12 5l7 7-7 7"/>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Step 4: Order Summary -->
    <div v-if="showOrderSummaryModal" class="modal-overlay" @click.self="closeOrderSummaryModal">
      <div class="modal-content">
        <div class="modal-header">
          <h2>Order Summary</h2>
          <button @click="closeOrderSummaryModal" class="close-btn">×</button>
        </div>
        
        <div class="modal-body">
          <div class="order-summary">
            <!-- Main Item -->
            <div class="summary-section">
              <h4>Main Item</h4>
              <div class="summary-item main-item">
                <div class="item-details">
                  <h5>{{ selectedMenuItem?.item_name }}</h5>
                  <p class="item-type">{{ selectedPriceType === 'general' ? 'General Sale' : 'Delivery Sale' }}</p>
                </div>
                <div class="item-price">{{ formatCurrency(getSelectedPrice()) }}</div>
              </div>
            </div>
            
            <!-- Extra Items -->
            <div v-if="selectedExtraItems.length > 0" class="summary-section">
              <h4>{{ t('extraItems') }}</h4>
              <div v-for="extra in selectedExtraItems" :key="extra.id" class="summary-item extra-item">
                <div class="item-details">
                  <h5>{{ extra.item_name }}</h5>
                  <p class="item-quantity">Quantity: {{ extra.quantity }}</p>
                </div>
                <div class="item-actions">
                  <div class="quantity-controls">
                    <button @click="decreaseExtraQuantity(extra.id)" class="qty-btn">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M5 12h14"/>
                      </svg>
                    </button>
                    <span class="qty-display">{{ extra.quantity }}</span>
                    <button @click="increaseExtraQuantity(extra.id)" class="qty-btn">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M12 5v14M5 12h14"/>
                      </svg>
                    </button>
                  </div>
                  <div class="item-price">{{ formatCurrency(getExtraPrice(extra) * extra.quantity) }}</div>
                  <button @click="removeExtraItem(extra.id)" class="remove-btn">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M3 6h18M8 6V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2M10 11v6M14 11v6"/>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
            
            <!-- Total -->
            <div class="summary-total">
              <div class="total-row">
                <span>Total Amount:</span>
                <span class="total-amount">{{ formatCurrency(calculateTotalAmount()) }}</span>
              </div>
            </div>
          </div>
        </div>
        
        <div class="modal-footer">
          <button @click="goBackToExtraSelection" class="btn-secondary">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M19 12H5M12 19l-7-7 7-7"/>
            </svg>
            Back to Extras
          </button>
          <button @click="saveOrderToIncome" class="btn-primary" :disabled="isSavingOrder">
            <span v-if="isSavingOrder">Saving...</span>
            <span v-else>Add to Income</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Add Outcome Modal -->
    <div v-if="showAddOutcomeModal" class="modal-overlay" @click.self="closeAddOutcomeModal">
      <div class="modal-content">
        <div class="modal-header">
          <h2>{{ t('addExpense') }}</h2>
          <button @click="closeAddOutcomeModal" class="close-btn">×</button>
        </div>
        
        <div class="modal-body">
          <!-- Suggestions Section -->
          <div v-if="recentOutcomes.length > 0" class="suggestions-section">
            <h4>{{ t('quickSelect') }}</h4>
            <p class="suggestions-subtitle">{{ t('chooseFromPrevious') }}</p>
            <div class="suggestions-tags">
              <button 
                v-for="outcome in recentOutcomes" 
                :key="outcome.name"
                @click="selectRecentOutcome(outcome)"
                class="suggestion-tag"
              >
                {{ outcome.name }}
              </button>
            </div>
          </div>
          
          <div v-if="isLoadingRecentOutcomes" class="loading-suggestions">
            <div class="loading-spinner-small"></div>
            <p>Loading recent expenses...</p>
          </div>
          
          <!-- Form Section -->
          <div class="form-section" :class="{ 'with-margin': recentOutcomes.length > 0 }">
            <div class="form-group">
              <label for="outcomeDescription">{{ t('listing') }}</label>
              <input 
                id="outcomeDescription"
                v-model="newOutcome.description" 
                type="text" 
                :placeholder="t('enterExpenseListing')"
                class="form-input"
              />
            </div>
            
            <div class="form-group">
              <label for="outcomeAmount">{{ t('amount') }} (฿)</label>
              <input 
                id="outcomeAmount"
                v-model="newOutcome.amount" 
                type="text"
                @input="validateAmount('outcome')"
                placeholder="0"
                class="form-input"
                :class="{ 'error-input': amountError }"
              />
              <p v-if="amountError" class="field-error">{{ amountError }}</p>
            </div>
          </div>
        </div>
        
        <div class="modal-footer">
          <button @click="closeAddOutcomeModal" class="btn-secondary">{{ t('cancel') }}</button>
          <button @click="saveOutcome" class="btn-expense-primary" :disabled="!isOutcomeValid">{{ t('addExpense') }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import API_BASE_URL from '@/config/api'
import { useLanguage } from '@/composables/useLanguage'

const router = useRouter()
const route = useRoute()
const { t, formatCurrency: formatCurrencyLang } = useLanguage()

// Route params
const trackingId = ref('')
const merchantId = ref('')
const merchantName = ref('')
const trackingDate = ref('')

// Data
const activeTab = ref('income')
const incomeItems = ref([])
const outcomeItems = ref([])
const showAddIncomeModal = ref(false)
const showAddOutcomeModal = ref(false)
const amountError = ref('')

// New transaction forms
const newIncome = ref({
  description: '',
  amount: '',
  category: 'sales'
})

const newOutcome = ref({
  description: '',
  amount: '',
  category: 'other' // Default category since it's still required by backend
})

// Data for totals (from API)
const totalIncome = ref(0)
const totalOutcome = ref(0)

// Menu selection system
const showPriceSelectionModal = ref(false)
const showExtraSelectionModal = ref(false)
const showOrderSummaryModal = ref(false)
const searchQuery = ref('')
const extraSearchQuery = ref('')
const menuItems = ref([])
const extraItems = ref([])
const filteredMenuItems = ref([])
const filteredExtraItems = ref([])
const selectedMenuItem = ref(null)
const selectedPriceType = ref('') // 'general' or 'delivery'
const selectedExtraItems = ref([])
const isLoadingMenu = ref(false)
const isLoadingExtras = ref(false)
const isSavingOrder = ref(false)

// Recent outcomes for suggestions
const recentOutcomes = ref([])
const isLoadingRecentOutcomes = ref(false)

// Computed
const netRevenue = computed(() => {
  return totalIncome.value - totalOutcome.value
})

const isIncomeValid = computed(() => {
  return newIncome.value.description && 
    newIncome.value.amount && 
    !amountError.value
})

const isOutcomeValid = computed(() => {
  return newOutcome.value.description && 
    newOutcome.value.amount && 
    !amountError.value
})

// Format functions
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('th-TH', {
    style: 'currency',
    currency: 'THB',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(amount || 0)
}

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('th-TH', {
    day: 'numeric',
    month: 'short',
    year: 'numeric'
  })
}

const formatTrackingDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('th-TH', {
    weekday: 'long',
    day: 'numeric',
    month: 'long',
    year: 'numeric'
  })
}

// Navigation
const goBack = () => {
  router.push({
    name: 'TrackingPage',
    params: { merchantId: merchantId.value },
    query: { merchantName: merchantName.value }
  })
}

// Body Scroll Lock Helpers
const lockBodyScroll = () => {
  document.body.classList.add('modal-open')
}

const unlockBodyScroll = () => {
  document.body.classList.remove('modal-open')
}

const isAnyModalOpen = () => {
  return showAddIncomeModal.value || 
         showPriceSelectionModal.value || 
         showExtraSelectionModal.value || 
         showOrderSummaryModal.value || 
         showAddOutcomeModal.value
}

// Menu Selection Methods
const fetchMenuItems = async () => {
  if (!merchantId.value) return
  
  isLoadingMenu.value = true
  try {
    const response = await fetch(`${API_BASE_URL}/menu-items?merchant_id=${merchantId.value}`)
    if (response.ok) {
      const data = await response.json()
      menuItems.value = data.map(item => ({
        id: item.id,
        item_name: item.item_name,
        general_price_sale: item.general_price_sale || 0,
        delivery_price_sale: item.delivery_price_sale || 0
      }))
      filteredMenuItems.value = menuItems.value
    }
  } catch (error) {
    console.error('Error fetching menu items:', error)
  } finally {
    isLoadingMenu.value = false
  }
}

const fetchExtraItems = async () => {
  if (!merchantId.value) return
  
  isLoadingExtras.value = true
  try {
    const response = await fetch(`${API_BASE_URL}/extra-items?merchant_id=${merchantId.value}`)
    if (response.ok) {
      const data = await response.json()
      extraItems.value = data.map(item => ({
        id: item.id,
        item_name: item.item_name,
        general_price_sale: item.general_price_sale || 0,
        delivery_price_sale: item.delivery_price_sale || 0
      }))
      filteredExtraItems.value = extraItems.value
    }
  } catch (error) {
    console.error('Error fetching extra items:', error)
  } finally {
    isLoadingExtras.value = false
  }
}

const searchMenuItems = () => {
  if (!searchQuery.value.trim()) {
    filteredMenuItems.value = menuItems.value
  } else {
    filteredMenuItems.value = menuItems.value.filter(item => 
      item.item_name.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }
}

const searchExtraItems = () => {
  if (!extraSearchQuery.value.trim()) {
    filteredExtraItems.value = extraItems.value
  } else {
    filteredExtraItems.value = extraItems.value.filter(item => 
      item.item_name.toLowerCase().includes(extraSearchQuery.value.toLowerCase())
    )
  }
}

const selectMenuItem = (item) => {
  selectedMenuItem.value = item
  showAddIncomeModal.value = false
  showPriceSelectionModal.value = true
}

const selectPriceType = (type) => {
  selectedPriceType.value = type
  showPriceSelectionModal.value = false
  showExtraSelectionModal.value = true
  // Body scroll stays locked since we're transitioning between modals
  
  // Load extra items if not already loaded
  if (extraItems.value.length === 0) {
    fetchExtraItems()
  }
}

// Extra Items Management
const toggleExtraItem = (item) => {
  const existingIndex = selectedExtraItems.value.findIndex(extra => extra.id === item.id)
  
  if (existingIndex >= 0) {
    // If already selected, increase quantity
    selectedExtraItems.value[existingIndex].quantity += 1
  } else {
    // Add new extra item with quantity 1
    selectedExtraItems.value.push({
      ...item,
      quantity: 1
    })
  }
}

const isExtraSelected = (itemId) => {
  return selectedExtraItems.value.some(extra => extra.id === itemId)
}

const getExtraQuantity = (itemId) => {
  const extra = selectedExtraItems.value.find(extra => extra.id === itemId)
  return extra ? extra.quantity : 0
}

const increaseExtraQuantity = (itemId) => {
  const extra = selectedExtraItems.value.find(extra => extra.id === itemId)
  if (extra) {
    extra.quantity += 1
  }
}

const decreaseExtraQuantity = (itemId) => {
  const extra = selectedExtraItems.value.find(extra => extra.id === itemId)
  if (extra) {
    if (extra.quantity > 1) {
      extra.quantity -= 1
    } else {
      removeExtraItem(itemId)
    }
  }
}

const removeExtraItem = (itemId) => {
  selectedExtraItems.value = selectedExtraItems.value.filter(extra => extra.id !== itemId)
}

// Modal Navigation Methods
const goBackToMenuSelection = () => {
  showPriceSelectionModal.value = false
  showAddIncomeModal.value = true
}

const goBackToPriceSelection = () => {
  showExtraSelectionModal.value = false
  showPriceSelectionModal.value = true
}

const goBackToExtraSelection = () => {
  showOrderSummaryModal.value = false
  showExtraSelectionModal.value = true
}

const proceedToSummary = () => {
  showExtraSelectionModal.value = false
  showOrderSummaryModal.value = true
  // Body scroll stays locked since we're transitioning between modals
}

const closePriceSelectionModal = () => {
  showPriceSelectionModal.value = false
  resetMenuSelectionState()
  
  // Only unlock body scroll if no other modals are open
  if (!isAnyModalOpen()) {
    unlockBodyScroll()
  }
}

const closeExtraSelectionModal = () => {
  showExtraSelectionModal.value = false
  resetMenuSelectionState()
  
  // Only unlock body scroll if no other modals are open
  if (!isAnyModalOpen()) {
    unlockBodyScroll()
  }
}

const closeOrderSummaryModal = () => {
  showOrderSummaryModal.value = false
  resetMenuSelectionState()
  
  // Only unlock body scroll if no other modals are open
  if (!isAnyModalOpen()) {
    unlockBodyScroll()
  }
}

const resetMenuSelectionState = () => {
  selectedMenuItem.value = null
  selectedPriceType.value = ''
  selectedExtraItems.value = []
  searchQuery.value = ''
  extraSearchQuery.value = ''
  filteredMenuItems.value = menuItems.value
  filteredExtraItems.value = extraItems.value
}

// Business Logic Methods
const getSelectedPrice = () => {
  if (!selectedMenuItem.value || !selectedPriceType.value) return 0
  
  return selectedPriceType.value === 'general' 
    ? selectedMenuItem.value.general_price_sale 
    : selectedMenuItem.value.delivery_price_sale
}

const getExtraPrice = (item) => {
  if (!selectedPriceType.value) return 0
  
  return selectedPriceType.value === 'general' 
    ? item.general_price_sale 
    : item.delivery_price_sale
}

const calculateTotalAmount = () => {
  let total = getSelectedPrice()
  
  selectedExtraItems.value.forEach(extra => {
    total += getExtraPrice(extra) * extra.quantity
  })
  
  return total
}

const saveOrderToIncome = async () => {
  if (!selectedMenuItem.value || isSavingOrder.value) return
  
  isSavingOrder.value = true
  
  try {
    // Create description from selected items
    const mainItemDesc = `${selectedMenuItem.value.item_name} (${selectedPriceType.value === 'general' ? t('general') : t('delivery')})`
    const extraItemsDesc = selectedExtraItems.value.length > 0
      ? ` + ${selectedExtraItems.value.map(extra => `${extra.item_name} x${extra.quantity}`).join(', ')}`
      : ''
    
    const description = mainItemDesc + extraItemsDesc
    const totalAmount = calculateTotalAmount()
    
    // Save as income transaction
    const response = await fetch(`${API_BASE_URL}/daily-tracking/${trackingId.value}/income`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        description: description,
        amount: totalAmount,
        category: 'sales',
        transaction_date: trackingDate.value
      })
    })
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    
    const data = await response.json()
    
    // Add to local income items list
    incomeItems.value.push({
      transaction_id: data.transaction_id,
      description: description,
      amount: totalAmount,
      category: 'sales',
      transaction_date: trackingDate.value
    })
    
    // Update total income
    totalIncome.value += totalAmount
    
    // Close modals and reset state
    closeOrderSummaryModal()
    
    console.log('Order saved to income:', data)
    
  } catch (error) {
    console.error('Error saving order to income:', error)
    alert('Failed to save order. Please try again.')
  } finally {
    isSavingOrder.value = false
  }
}

// Amount validation
const validateAmount = (type) => {
  let value = type === 'income' ? newIncome.value.amount : newOutcome.value.amount
  const originalValue = value.toString()
  const hasInvalidChars = /[^0-9.]/.test(originalValue)
  
  value = originalValue.replace(/[^0-9.]/g, '')
  const parts = value.split('.')
  if (parts.length > 2) {
    value = parts[0] + '.' + parts.slice(1).join('')
  }
  
  if (type === 'income') {
    newIncome.value.amount = value
  } else {
    newOutcome.value.amount = value
  }
  
  if (hasInvalidChars) {
    amountError.value = 'Only numbers are allowed'
  } else if (parseFloat(value) <= 0 && value !== '') {
    amountError.value = 'Amount must be greater than 0'
  } else {
    amountError.value = ''
  }
}

// Income Modal - Updated for Menu Selection
const openAddIncomeModal = () => {
  showAddIncomeModal.value = true
  amountError.value = ''
  lockBodyScroll()
  
  // Load menu items if not already loaded
  if (menuItems.value.length === 0) {
    fetchMenuItems()
  }
}

const closeAddIncomeModal = () => {
  showAddIncomeModal.value = false
  newIncome.value = {
    description: '',
    amount: '',
    category: 'sales'
  }
  amountError.value = ''
  resetMenuSelectionState()
  
  // Only unlock body scroll if no other modals are open
  if (!isAnyModalOpen()) {
    unlockBodyScroll()
  }
}

const saveIncome = async () => {
  if (!isIncomeValid.value) return

  try {
    const response = await fetch(`${API_BASE_URL}/daily-tracking/${trackingId.value}/income`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        description: newIncome.value.description,
        amount: parseFloat(newIncome.value.amount),
        category: newIncome.value.category,
        transaction_date: trackingDate.value
      })
    })

    if (!response.ok) {
      throw new Error('Failed to add income')
    }

    const savedIncome = await response.json()
    
    // Refresh the tracking details
    await fetchTrackingDetails()
    
    closeAddIncomeModal()
  } catch (error) {
    console.error('Error saving income:', error)
    alert('Failed to save income. Please try again.')
  }
}

// Fetch recent outcomes for suggestions - simplified version using localStorage or existing data
const fetchRecentOutcomes = async () => {
  isLoadingRecentOutcomes.value = true
  
  try {
    // For now, use current outcome items and localStorage to build suggestions
    const currentOutcomes = outcomeItems.value || []
    
    // Try to get additional outcomes from localStorage if available
    const storageKey = `recent_outcomes_${merchantId.value}`
    const storedOutcomes = JSON.parse(localStorage.getItem(storageKey) || '[]')
    
    // Combine current and stored outcomes
    const allOutcomes = [...currentOutcomes, ...storedOutcomes]
    
    if (allOutcomes.length > 0) {
      // Get unique expense names (descriptions) sorted by most recent usage
      const uniqueNames = new Map()
      allOutcomes.forEach(outcome => {
        if (outcome.description && outcome.description.trim()) {
          const name = outcome.description
          const date = outcome.transaction_date || new Date().toISOString()
          if (!uniqueNames.has(name) || new Date(date) > new Date(uniqueNames.get(name).date)) {
            uniqueNames.set(name, { name, date })
          }
        }
      })
      
      // Convert to array, sort by date (most recent first), and take top 8
      recentOutcomes.value = Array.from(uniqueNames.values())
        .sort((a, b) => new Date(b.date) - new Date(a.date))
        .slice(0, 8)
        .map(item => ({ name: item.name }))
    } else {
      recentOutcomes.value = []
    }
  } catch (error) {
    console.error('Error processing recent outcomes:', error)
    recentOutcomes.value = []
  } finally {
    isLoadingRecentOutcomes.value = false
  }
}

// Store outcome to localStorage for future suggestions
const storeOutcomeForSuggestions = (outcome) => {
  try {
    const storageKey = `recent_outcomes_${merchantId.value}`
    const storedOutcomes = JSON.parse(localStorage.getItem(storageKey) || '[]')
    
    // Add new outcome to stored outcomes (keep last 20 for performance)
    const updatedOutcomes = [outcome, ...storedOutcomes.slice(0, 19)]
    localStorage.setItem(storageKey, JSON.stringify(updatedOutcomes))
  } catch (error) {
    console.error('Error storing outcome for suggestions:', error)
  }
}

const selectRecentOutcome = (outcome) => {
  newOutcome.value.description = outcome.name
  // Don't auto-fill amount, let user enter it manually
}

// Outcome Modal
const openAddOutcomeModal = () => {
  showAddOutcomeModal.value = true
  amountError.value = ''
  lockBodyScroll()
  
  // Fetch recent outcomes for suggestions
  fetchRecentOutcomes()
}

const closeAddOutcomeModal = () => {
  showAddOutcomeModal.value = false
  newOutcome.value = {
    description: '',
    amount: '',
    category: 'other'
  }
  amountError.value = ''
  
  // Only unlock body scroll if no other modals are open
  if (!isAnyModalOpen()) {
    unlockBodyScroll()
  }
}

const saveOutcome = async () => {
  if (!isOutcomeValid.value) return

  try {
    const response = await fetch(`${API_BASE_URL}/daily-tracking/${trackingId.value}/outcome`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        description: newOutcome.value.description,
        amount: parseFloat(newOutcome.value.amount),
        category: newOutcome.value.category,
        transaction_date: trackingDate.value
      })
    })

    if (!response.ok) {
      throw new Error('Failed to add outcome')
    }

    const savedOutcome = await response.json()
    
    // Store outcome for future suggestions
    storeOutcomeForSuggestions({
      description: newOutcome.value.description,
      amount: parseFloat(newOutcome.value.amount),
      category: newOutcome.value.category,
      transaction_date: trackingDate.value
    })
    
    // Refresh the tracking details
    await fetchTrackingDetails()
    
    closeAddOutcomeModal()
  } catch (error) {
    console.error('Error saving outcome:', error)
    alert('Failed to save outcome. Please try again.')
  }
}

// Delete item
const deleteItem = async (itemId, type) => {
  if (!confirm('Are you sure you want to delete this item?')) return

  try {
    console.log(`Deleting ${type} item:`, itemId)
    
    // Call backend API to delete the transaction
    const endpoint = type === 'income' ? 'income' : 'outcome'
    const response = await fetch(`${API_BASE_URL}/daily-tracking/${trackingId.value}/${endpoint}/${itemId}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      }
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    console.log(`Successfully deleted ${type} item:`, itemId)
    
    // Remove from local state after successful API call
    if (type === 'income') {
      const deletedItem = incomeItems.value.find(item => item.transaction_id === itemId)
      if (deletedItem) {
        totalIncome.value -= deletedItem.amount
        incomeItems.value = incomeItems.value.filter(item => item.transaction_id !== itemId)
      }
    } else {
      const deletedItem = outcomeItems.value.find(item => item.transaction_id === itemId)
      if (deletedItem) {
        totalOutcome.value -= deletedItem.amount
        outcomeItems.value = outcomeItems.value.filter(item => item.transaction_id !== itemId)
      }
    }
    
    console.log(`Updated ${type} items and totals`)
    
  } catch (error) {
    console.error(`Error deleting ${type} item:`, error)
    alert(`Failed to delete ${type} item. Please try again.`)
  }
}

// Fetch tracking details from API
const fetchTrackingDetails = async () => {
  if (!trackingId.value) {
    console.warn('No tracking ID, cannot fetch details')
    return
  }
  
  try {
    console.log('Fetching daily tracking details for ID:', trackingId.value)
    const response = await fetch(`${API_BASE_URL}/daily-tracking/${trackingId.value}`)
    if (!response.ok) {
      throw new Error('Failed to fetch daily tracking details')
    }
    
    const data = await response.json()
    console.log('Daily tracking details from API:', data)
    
    if (data.tracking) {
      trackingDate.value = data.tracking.tracking_date
      
      // Update merchant info if not set
      if (!merchantName.value && data.tracking.merchant_name) {
        merchantName.value = data.tracking.merchant_name
      }
      
      totalIncome.value = data.tracking.total_income || 0
      totalOutcome.value = data.tracking.total_outcome || 0
    }
    
    // Update income items
    if (data.income_items && Array.isArray(data.income_items)) {
      incomeItems.value = data.income_items.map(item => ({
        transaction_id: item.transaction_id,
        description: item.description,
        amount: item.amount,
        category: item.category,
        transaction_date: item.transaction_date
      }))
    }
    
    // Update outcome items
    if (data.outcome_items && Array.isArray(data.outcome_items)) {
      outcomeItems.value = data.outcome_items.map(item => ({
        transaction_id: item.transaction_id,
        description: item.description,
        amount: item.amount,
        category: item.category,
        transaction_date: item.transaction_date
      }))
    }
  } catch (error) {
    console.error('Error fetching daily tracking details:', error)
  }
}

// Initialize
onMounted(async () => {
  console.log('=== DailyTrackingPage Mounted ===')
  console.log('Route params:', route.params)
  console.log('Route query:', route.query)
  
  // Get tracking ID from route params (most important)
  trackingId.value = route.params.trackingId || ''
  
  // Get initial values from query (might be empty)
  merchantId.value = route.query.merchantId || ''
  merchantName.value = route.query.merchantName || 'Merchant'
  
  console.log('Tracking ID:', trackingId.value)
  
  if (!trackingId.value) {
    console.error('No tracking ID provided!')
    alert('No tracking ID provided. Please go back and select a daily tracking.')
    router.push('/home')
    return
  }
  
  // Fetch tracking details from API
  await fetchTrackingDetails()
  
  console.log('Final initialization:', {
    trackingId: trackingId.value,
    trackingDate: trackingDate.value,
    merchantName: merchantName.value
  })
})
</script>

<style scoped>
.daily-tracking-page {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
  background: #f9fafb;
  min-height: 100vh;
}

/* Header */
.page-header {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 24px;
}

.back-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 16px;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 10px;
  color: #374151;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.back-btn:hover {
  background: #f3f4f6;
}

.page-title {
  font-size: 24px;
  font-weight: 700;
  color: #111827;
  margin: 0 0 4px 0;
}

.page-subtitle {
  font-size: 14px;
  color: #6b7280;
  margin: 0;
}

/* Summary Cards */
.summary-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}

.summary-card {
  background: white;
  border-radius: 12px;
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.card-icon {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
}

.income-card .card-icon {
  background: #d1fae5;
  color: #065f46;
}

.outcome-card .card-icon {
  background: #fee2e2;
  color: #991b1b;
}

.revenue-card .card-icon {
  background: #dbeafe;
  color: #1e40af;
}

.card-content {
  flex: 1;
}

.card-label {
  font-size: 13px;
  color: #6b7280;
  margin: 0 0 4px 0;
}

.card-value {
  font-size: 24px;
  font-weight: 700;
  color: #111827;
  margin: 0;
}

.card-value.positive {
  color: #10b981;
}

.card-value.negative {
  color: #ef4444;
}

/* Tabs */
.tabs-section {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.tabs {
  display: flex;
  border-bottom: 1px solid #e5e7eb;
}

.tab-btn {
  flex: 1;
  padding: 16px;
  background: transparent;
  border: none;
  color: #6b7280;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  position: relative;
  transition: all 0.2s;
}

.tab-btn.active {
  color: #10b981;
  background: #f0fdf4;
}

.tab-btn.active.outcome-tab {
  color: #dc2626;
  background: #fef2f2;
}

.tab-btn.active::after {
  content: '';
  position: absolute;
  bottom: -1px;
  left: 0;
  right: 0;
  height: 2px;
  background: #10b981;
}

.tab-btn.active.outcome-tab::after {
  background: #dc2626;
}

.tab-content {
  padding: 24px;
  min-height: 400px;
}

/* Section Header */
.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.section-header h3 {
  font-size: 18px;
  font-weight: 600;
  color: #111827;
  margin: 0;
}

.add-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 16px;
  background: #10b981;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.add-btn:hover {
  background: #059669;
}

.add-btn-expense {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 16px;
  background: #ef4444;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.add-btn-expense:hover {
  background: #ef4444;
}


/* Items List */
.items-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.item-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background: #f9fafb;
  border-radius: 10px;
  transition: all 0.2s;
}

.item-card:hover {
  background: #f3f4f6;
}

.item-info {
  flex: 1;
}

.item-name {
  font-size: 15px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 4px 0;
}

.item-date {
  font-size: 13px;
  color: #6b7280;
  margin: 0 0 8px 0;
}

.item-category {
  display: inline-block;
  padding: 4px 10px;
  background: #e5e7eb;
  color: #374151;
  font-size: 12px;
  font-weight: 500;
  border-radius: 12px;
}

.item-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.item-amount {
  font-size: 18px;
  font-weight: 700;
}

.item-amount.income {
  color: #10b981;
}

.item-amount.outcome {
  color: #ef4444;
}

.delete-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fee2e2;
  border: none;
  border-radius: 6px;
  color: #ef4444;
  cursor: pointer;
  transition: all 0.2s;
}

.delete-btn:hover {
  background: #fecaca;
}

/* Empty State */
.empty-state {
  text-align: center;
  padding: 60px 20px;
}

.empty-state svg {
  color: #d1d5db;
  margin: 0 auto 20px;
}

.empty-state p {
  font-size: 16px;
  color: #6b7280;
  margin: 0 0 20px 0;
}

.empty-add-btn {
  padding: 10px 20px;
  background: #10b981;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.empty-add-btn:hover {
  background: #059669;
}

.empty-add-btn-expense {
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

.empty-add-btn-expense:hover {
  background: #ef4444;
}


/* Modal */
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
  overflow: hidden; /* Prevent background scroll */
}

.modal-content {
  background: white;
  border-radius: 16px;
  width: 100%;
  max-width: 500px;
  max-height: 90vh;
  display: flex;
  flex-direction: column;
  overflow: hidden; /* Remove scroll from modal content */
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24px;
  border-bottom: 1px solid #e5e7eb;
}

.modal-header h2 {
  font-size: 20px;
  font-weight: 600;
  color: #111827;
  margin: 0;
}

.close-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f3f4f6;
  border: none;
  border-radius: 8px;
  font-size: 24px;
  color: #6b7280;
  cursor: pointer;
  transition: all 0.2s;
}

.close-btn:hover {
  background: #e5e7eb;
  color: #374151;
}

.modal-body {
  padding: 24px;
  overflow-y: auto; /* Only modal body should scroll */
  flex: 1; /* Take available space */
  max-height: calc(90vh - 140px); /* Account for header and footer */
}

.form-group {
  margin-bottom: 20px;
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
  color: #111827;
  transition: all 0.2s;
}

.form-input:focus {
  outline: none;
  border-color: #10b981;
  box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.1);
}

.form-input.error-input {
  border-color: #ef4444;
}

.field-error {
  font-size: 12px;
  color: #ef4444;
  margin: 4px 0 0 0;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 16px 24px;
  border-top: 1px solid #e5e7eb;
}

.btn-secondary,
.btn-expense-primary,
.btn-primary {
  padding: 10px 20px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  border: none;
}

.btn-secondary {
  background: #f3f4f6;
  color: #374151;
}

.btn-secondary:hover {
  background: #e5e7eb;
}

.btn-primary {
  background: #10b981;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background: #059669;
}

.btn-primary:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-expense-primary {
  background: #ef4444;
  color: white;
}

.btn-expense-primary:hover:not(:disabled) {
  background: #ef4444;
}

.btn-expense-primary:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Modal Body Scroll Prevention */
body.modal-open {
  overflow: hidden;
}

/* Menu Selection Modal Styles */
.large-modal {
  max-width: 700px;
}

.search-input {
  position: relative;
  background: #f9fafb;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  padding: 12px 16px;
  font-size: 14px;
  transition: all 0.2s;
}

.search-input:focus {
  border-color: #10b981;
  box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.1);
  background: white;
}

.menu-items-list, .extra-items-list {
  margin-top: 20px;
}

.menu-items-list h4, .extra-items-list h4 {
  font-size: 16px;
  font-weight: 600;
  color: #374151;
  margin: 0 0 12px 0;
}

.items-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 12px;
  max-height: 400px;
  overflow-y: auto;
}

.menu-item-card, .extra-item-card {
  background: white;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.menu-item-card:hover, .extra-item-card:hover {
  border-color: #10b981;
  transform: translateY(-1px);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.menu-item-card.selectable:hover {
  background: #f0fdf4;
}

.extra-item-card.selected {
  border-color: #10b981;
  background: #f0fdf4;
}

.item-info h5 {
  font-size: 14px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 8px 0;
}

.price-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.price-tag {
  font-size: 12px;
  padding: 2px 6px;
  border-radius: 4px;
  font-weight: 500;
}

.price-tag.general {
  background: #d1fae5;
  color: #047857;
}

.price-tag.delivery {
  background: #dbeafe;
  color: #1d4ed8;
}

.extra-price {
  font-size: 13px;
  color: #6b7280;
  margin: 0;
}

.select-indicator {
  color: #6b7280;
}

.selection-indicator {
  display: flex;
  align-items: center;
  gap: 8px;
}

.quantity-badge {
  background: #10b981;
  color: white;
  border-radius: 50%;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 600;
}

.checkbox {
  width: 20px;
  height: 20px;
  border: 2px solid #d1d5db;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.checkbox.checked {
  background: #10b981;
  border-color: #10b981;
  color: white;
}

/* Price Selection Modal */
.selected-item-info {
  text-align: center;
  margin-bottom: 24px;
  padding: 16px;
  background: #f9fafb;
  border-radius: 12px;
}

.selected-item-info h4 {
  font-size: 18px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 8px 0;
}

.selected-item-info p {
  color: #6b7280;
  margin: 0;
}

.price-selection-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
}

.price-option-card {
  background: white;
  border: 2px solid #e5e7eb;
  border-radius: 16px;
  padding: 24px;
  text-align: center;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.price-option-card:hover {
  border-color: #10b981;
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
}

.price-type-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 8px;
}

.price-type-icon.general {
  background: #d1fae5;
  color: #047857;
}

.price-type-icon.delivery {
  background: #dbeafe;
  color: #1d4ed8;
}

.price-option-card h5 {
  font-size: 16px;
  font-weight: 600;
  color: #111827;
  margin: 0;
}

.price-option-card .price {
  font-size: 20px;
  font-weight: 700;
  color: #10b981;
  margin: 0;
}

.price-option-card .description {
  font-size: 12px;
  color: #6b7280;
  margin: 0;
}

/* Extra Selection Modal */
.selected-item-summary {
  background: #f0fdf4;
  border: 1px solid #bbf7d0;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 20px;
}

.selected-item-summary h4 {
  font-size: 16px;
  font-weight: 600;
  color: #047857;
  margin: 0 0 4px 0;
}

.selected-price {
  font-size: 14px;
  color: #047857;
  margin: 0;
  font-weight: 500;
}

/* Order Summary Modal */
.order-summary {
  max-height: 500px;
  overflow-y: auto;
}

.summary-section {
  margin-bottom: 24px;
}

.summary-section h4 {
  font-size: 16px;
  font-weight: 600;
  color: #374151;
  margin: 0 0 12px 0;
  padding-bottom: 8px;
  border-bottom: 1px solid #e5e7eb;
}

.summary-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px;
  background: #f9fafb;
  border-radius: 8px;
  margin-bottom: 8px;
}

.summary-item.main-item {
  background: #f0fdf4;
  border: 1px solid #bbf7d0;
}

.item-details h5 {
  font-size: 14px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 4px 0;
}

.item-type, .item-quantity {
  font-size: 12px;
  color: #6b7280;
  margin: 0;
}

.item-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.quantity-controls {
  display: flex;
  align-items: center;
  gap: 8px;
}

.qty-btn {
  width: 24px;
  height: 24px;
  border: 1px solid #d1d5db;
  background: white;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
}

.qty-btn:hover {
  border-color: #10b981;
  background: #f0fdf4;
}

.qty-display {
  font-size: 14px;
  font-weight: 600;
  color: #374151;
  min-width: 20px;
  text-align: center;
}

.item-price {
  font-size: 14px;
  font-weight: 600;
  color: #10b981;
}

.remove-btn {
  width: 24px;
  height: 24px;
  border: 1px solid #fca5a5;
  background: #fef2f2;
  color: #dc2626;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
}

.remove-btn:hover {
  background: #fecaca;
  border-color: #dc2626;
}

.summary-total {
  border-top: 2px solid #e5e7eb;
  padding-top: 16px;
  margin-top: 16px;
}

.total-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 18px;
  font-weight: 700;
}

.total-amount {
  color: #10b981;
}

/* Loading and Empty States */
.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px;
  color: #6b7280;
}

.loading-spinner {
  width: 32px;
  height: 32px;
  border: 3px solid #f3f4f6;
  border-top: 3px solid #10b981;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 12px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.empty-search {
  text-align: center;
  padding: 40px;
  color: #9ca3af;
}

.empty-search p {
  margin: 0;
  font-size: 14px;
}

/* Responsive Design */
@media (max-width: 768px) {
  .large-modal {
    max-width: 95vw;
    margin: 10px;
  }
  
  .items-grid {
    grid-template-columns: 1fr;
    max-height: 300px;
  }
  
  .price-selection-grid {
    grid-template-columns: 1fr;
  }
  
  .summary-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
  
  .item-actions {
    align-self: flex-end;
  }
}

/* Outcome Modal Suggestions */
.suggestions-section {
  margin-bottom: 24px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e5e7eb;
}

.suggestions-section h4 {
  font-size: 16px;
  font-weight: 600;
  color: #374151;
  margin: 0 0 4px 0;
}

.suggestions-subtitle {
  font-size: 13px;
  color: #6b7280;
  margin: 0 0 16px 0;
}

.suggestions-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  max-height: 120px;
  overflow-y: auto;
}

.suggestion-tag {
  background: #ffe2e2;
  border-radius: 20px;
  padding: 6px 12px;
  font-size: 13px;
  color: #ef4444;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
  font-weight: 500;
}

.suggestion-tag:hover {
  background: #dcfce7;
  border-color: #10b981;
  transform: translateY(-1px);
}

.loading-suggestions {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 20px;
  color: #6b7280;
  font-size: 13px;
}

.loading-spinner-small {
  width: 16px;
  height: 16px;
  border: 2px solid #f3f4f6;
  border-top: 2px solid #dc2626;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

.form-section.with-margin {
  margin-top: 20px;
}

@media (max-width: 768px) {
  .suggestions-tags {
    max-height: 100px;
  }
}
</style>