<template>
  <div class="manage-merchant-page">
    <!-- Header with Back Button -->
    <div class="page-header">
      <button @click="goBack" class="back-btn">
        <svg
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
        </svg>
        {{ t('back') }}
      </button>
      <h1>{{ t('manageMerchant') }}</h1>
    </div>

    <!-- Merchant Info -->
    <div class="merchant-info-card">
      <div class="merchant-icon">
        {{ merchantName.charAt(0).toUpperCase() }}
      </div>
      <div>
        <h2>{{ merchantName }}</h2>
        <p class="merchant-id">ID: {{ merchantId }}</p>
      </div>
    </div>

    <!-- Tabs -->
    <div class="tabs-container">
      <div class="tabs">
        <button
          @click="activeTab = 'main'"
          class="tab-btn"
          :class="{ active: activeTab === 'main' }"
        >
          {{ t('mainMenu') }}
        </button>
        <button
          @click="activeTab = 'extra'"
          class="tab-btn"
          :class="{ active: activeTab === 'extra' }"
        >
          {{ t('extraOnTop') }}
        </button>
      </div>
    </div>

    <!-- Tab Content -->
    <div class="tab-content">
      <!-- Main Menu Tab -->
      <div v-if="activeTab === 'main'" class="menu-section">
        <div class="section-header">
          <h3>{{ t('mainMenuItems') }}</h3>
          <button @click="openCreateMainModal" class="create-btn">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
              <path
                fill-rule="evenodd"
                d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
                clip-rule="evenodd"
              />
            </svg>
            {{ t('addMenuItem') }}
          </button>
        </div>

        <!-- Menu Items List -->
        <div v-if="mainMenuItems.length === 0" class="empty-state">
          <div class="empty-icon">
            <svg
              width="48"
              height="48"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="1.5"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 6.042A8.967 8.967 0 006 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 016 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 016-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0018 18a8.967 8.967 0 00-6 2.292m0-14.25v14.25"
              />
            </svg>
          </div>
          <h4>{{ t('noMenuItemsYet') }}</h4>
          <p>{{ t('startByAddingFirstMenuItem') }}</p>
        </div>

        <div v-else class="items-grid">
          <div v-for="item in mainMenuItems" :key="item.id" class="menu-item-card">
            <div class="item-header">
              <h4>{{ item.name }}</h4>
              <div class="action-buttons">
                <button @click="openEditModal(item, 'main')" class="edit-btn">
                  <svg width="18" height="18" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z"
                    />
                  </svg>
                </button>
                <button @click="deleteMenuItem(item.id, 'main')" class="delete-btn">
                  <svg width="18" height="18" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      fill-rule="evenodd"
                      d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                      clip-rule="evenodd"
                    />
                  </svg>
                </button>
              </div>
            </div>
            <div class="item-pricing">
              <div class="price-item">
                <span class="price-label">{{ t('general') }}:</span>
                <span class="price-value">{{ formatCurrency(item.general_price || 0) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Extra on Top Tab -->
      <div v-if="activeTab === 'extra'" class="menu-section">
        <div class="section-header">
          <h3>{{ t('extraOnTopItems') }}</h3>
          <button @click="openCreateExtraModal" class="create-btn">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
              <path
                fill-rule="evenodd"
                d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
                clip-rule="evenodd"
              />
            </svg>
            {{ t('addExtraItem') }}
          </button>
        </div>

        <!-- Extra Items List -->
        <div v-if="extraMenuItems.length === 0" class="empty-state">
          <div class="empty-icon">
            <svg
              width="48"
              height="48"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="1.5"
            >
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
            </svg>
          </div>
          <h4>{{ t('noExtraItemsYet') }}</h4>
          <p>{{ t('addExtraToppingsOrAddOns') }}</p>
        </div>

        <div v-else class="items-grid">
          <div v-for="item in extraMenuItems" :key="item.id" class="menu-item-card">
            <div class="item-header">
              <h4>{{ item.name }}</h4>
              <div class="action-buttons">
                <button @click="openEditModal(item, 'extra')" class="edit-btn">
                  <svg width="18" height="18" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z"
                    />
                  </svg>
                </button>
                <button @click="deleteMenuItem(item.id, 'extra')" class="delete-btn">
                  <svg width="18" height="18" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      fill-rule="evenodd"
                      d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                      clip-rule="evenodd"
                    />
                  </svg>
                </button>
              </div>
            </div>
            <div class="item-pricing">
              <div class="price-item">
                <span class="price-label">{{ t('general') }}:</span>
                <span class="price-value">{{ formatCurrency(item.general_price || 0) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Create Main Menu Modal -->
    <div v-if="showCreateMainModal" class="modal-overlay" @click.self="closeCreateMainModal">
      <div class="modal-content">
        <div class="modal-header">
          <h3>{{ t('addMainMenuItem') }}</h3>
          <button @click="closeCreateMainModal" class="modal-close">
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
            <label for="mainItemName">{{ t('menuItemName') }}</label>
            <input
              id="mainItemName"
              v-model="newMainItem.name"
              type="text"
              :placeholder="t('enterMenuItemName')"
              class="form-input"
              :disabled="isCreating"
            />
          </div>

          <div class="form-group">
            <label for="mainItemGeneralPrice">{{ t('generalSalePrice') }} (฿)</label>
            <input
              id="mainItemGeneralPrice"
              v-model="newMainItem.general_price"
              type="text"
              @input="validatePrice('main', 'general')"
              placeholder="0"
              class="form-input"
              :class="{ 'error-input': priceError }"
              :disabled="isCreating"
            />
            <p v-if="priceError" class="field-error">{{ priceError }}</p>
          </div>


          <p v-if="createError" class="error-message">{{ createError }}</p>
        </div>

        <div class="modal-footer">
          <button @click="closeCreateMainModal" class="btn-secondary" :disabled="isCreating">
            {{ t('cancel') }}
          </button>
          <button
            @click="createMainMenuItem"
            class="btn-primary"
            :disabled="!newMainItem.name || !isMainPriceValid || isCreating"
          >
            <span v-if="isCreating">{{ t('creating') }}...</span>
            <span v-else>{{ t('addItem') }}</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Create Extra Modal -->
    <div v-if="showCreateExtraModal" class="modal-overlay" @click.self="closeCreateExtraModal">
      <div class="modal-content">
        <div class="modal-header">
          <h3>{{ t('addExtraItem') }}</h3>
          <button @click="closeCreateExtraModal" class="modal-close">
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
            <label for="extraItemName">{{ t('extraItemName') }}</label>
            <input
              id="extraItemName"
              v-model="newExtraItem.name"
              type="text"
              :placeholder="t('enterExtraItemName')"
              class="form-input"
              :disabled="isCreating"
            />
          </div>

          <div class="form-group">
            <label for="extraItemGeneralPrice">{{ t('generalSalePrice') }} (฿)</label>
            <input
              id="extraItemGeneralPrice"
              v-model="newExtraItem.general_price"
              type="text"
              @input="validatePrice('extra', 'general')"
              placeholder="0"
              class="form-input"
              :class="{ 'error-input': priceError }"
              :disabled="isCreating"
            />
            <p v-if="priceError" class="field-error">{{ priceError }}</p>
          </div>


          <p v-if="createError" class="error-message">{{ createError }}</p>
        </div>

        <div class="modal-footer">
          <button @click="closeCreateExtraModal" class="btn-secondary" :disabled="isCreating">
            {{ t('cancel') }}
          </button>
          <button
            @click="createExtraMenuItem"
            class="btn-primary"
            :disabled="!newExtraItem.name || !isExtraPriceValid || isCreating"
          >
            <span v-if="isCreating">{{ t('creating') }}...</span>
            <span v-else>{{ t('addItem') }}</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Edit Item Modal -->
    <div v-if="showEditModal" class="modal-overlay" @click.self="closeEditModal">
      <div class="modal-content">
        <div class="modal-header">
          <h3>{{ t('editItem') }}</h3>
          <button @click="closeEditModal" class="modal-close">
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
            <label for="editItemName">{{ t('itemName') }}</label>
            <input
              id="editItemName"
              v-model="editForm.name"
              type="text"
              :placeholder="t('enterItemName')"
              class="form-input"
              :class="{ 'error-input': editFormErrors.name }"
              :disabled="isUpdating"
            />
            <p v-if="editFormErrors.name" class="field-error">{{ editFormErrors.name }}</p>
          </div>

          <div class="form-group">
            <label for="editItemGeneralPrice">{{ t('generalSalePrice') }} (฿)</label>
            <input
              id="editItemGeneralPrice"
              v-model="editForm.general_price"
              type="text"
              @input="validateEditPrice('general')"
              placeholder="0"
              class="form-input"
              :class="{ 'error-input': editFormErrors.general_price }"
              :disabled="isUpdating"
            />
            <p v-if="editFormErrors.general_price" class="field-error">
              {{ editFormErrors.general_price }}
            </p>
          </div>


          <p v-if="updateError" class="error-message">{{ updateError }}</p>
        </div>

        <div class="modal-footer">
          <button @click="closeEditModal" class="btn-secondary" :disabled="isUpdating">
            Cancel
          </button>
          <button
            @click="saveItemChanges"
            class="btn-primary"
            :disabled="!isEditFormValid || isUpdating"
          >
            <span v-if="isUpdating">{{ t('saving') }}</span>
            <span v-else>{{ t('ok') }}</span>
          </button>
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

// Reactive data
const merchantId = ref('')
const merchantName = ref('')
const activeTab = ref('main')
const mainMenuItems = ref([])
const extraMenuItems = ref([])
const showCreateMainModal = ref(false)
const showCreateExtraModal = ref(false)
const isCreating = ref(false)
const createError = ref('')
const priceError = ref('')
const showEditModal = ref(false)
const editingItem = ref(null)
const isUpdating = ref(false)
const updateError = ref('')
const editForm = ref({
  name: '',
  general_price: '',
})
const editFormErrors = ref({
  name: '',
  general_price: '',
})

// New item forms
const newMainItem = ref({
  name: '',
  general_price: '',
})

const newExtraItem = ref({
  name: '',
  general_price: '',
})

// Computed properties for button validation
const isMainPriceValid = computed(() => {
  const generalPrice = parseFloat(newMainItem.value.general_price)
  return (
    newMainItem.value.general_price !== '' &&
    !isNaN(generalPrice) &&
    generalPrice > 0
  )
})

const isExtraPriceValid = computed(() => {
  const generalPrice = parseFloat(newExtraItem.value.general_price)
  return (
    newExtraItem.value.general_price !== '' &&
    !isNaN(generalPrice) &&
    generalPrice > 0
  )
})

const isEditFormValid = computed(() => {
  const generalPrice = parseFloat(editForm.value.general_price)
  return (
    editForm.value.name !== '' &&
    editForm.value.general_price !== '' &&
    !isNaN(generalPrice) &&
    generalPrice > 0 &&
    !editFormErrors.value.name &&
    !editFormErrors.value.general_price
  )
})

// Format currency using language-aware formatter
const formatCurrency = (amount) => {
  return formatCurrencyLang(amount)
}

// Navigation
const goBack = () => {
  router.push('/home')
}

// Validate price input to allow only numbers
const validatePrice = (type, priceType) => {
  let value
  if (type === 'main') {
    value = newMainItem.value.general_price
  } else {
    value = newExtraItem.value.general_price
  }

  const originalValue = value.toString()

  // Check if user tried to enter non-numeric characters
  const hasInvalidChars = /[^0-9.]/.test(originalValue)

  // Remove any non-numeric characters except decimal point
  value = originalValue.replace(/[^0-9.]/g, '')

  // Ensure only one decimal point
  const parts = value.split('.')
  if (parts.length > 2) {
    value = parts[0] + '.' + parts.slice(1).join('')
  }

  // Limit to 2 decimal places
  if (parts[1] && parts[1].length > 2) {
    value = parts[0] + '.' + parts[1].substring(0, 2)
  }

  // Update the model
  if (type === 'main') {
    newMainItem.value.general_price = value
  } else {
    newExtraItem.value.general_price = value
  }

  // Set error message if invalid characters were entered
  if (hasInvalidChars) {
    priceError.value = 'Only numbers are allowed for price'
  } else if (value && parseFloat(value) <= 0) {
    priceError.value = 'Price must be greater than 0'
  } else {
    priceError.value = ''
  }
}

// Modal handlers for Main Menu
const openCreateMainModal = () => {
  showCreateMainModal.value = true
  createError.value = ''
  priceError.value = ''
}

const closeCreateMainModal = () => {
  showCreateMainModal.value = false
  newMainItem.value = { name: '', general_price: '' }
  createError.value = ''
  priceError.value = ''
}

// Modal handlers for Extra Menu
const openCreateExtraModal = () => {
  showCreateExtraModal.value = true
  createError.value = ''
  priceError.value = ''
}

const closeCreateExtraModal = () => {
  showCreateExtraModal.value = false
  newExtraItem.value = { name: '', general_price: '' }
  createError.value = ''
  priceError.value = ''
}

// Create Main Menu Item
const createMainMenuItem = async () => {
  if (
    !newMainItem.value.name ||
    !newMainItem.value.general_price
  ) {
    createError.value = 'Please fill in all required fields'
    return
  }

  const generalPrice = parseFloat(newMainItem.value.general_price)
  if (isNaN(generalPrice) || generalPrice <= 0) {
    createError.value = 'Please enter valid price'
    return
  }

  isCreating.value = true
  createError.value = ''

  try {
    // Make API call to create menu item
    const response = await fetch(`${API_BASE_URL}/menu-items`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        merchant_id: merchantId.value,
        item_name: newMainItem.value.name,
        general_price_sale: generalPrice,
        delivery_price_sale: generalPrice,
      }),
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const data = await response.json()

    // Add the new item to the local array
    mainMenuItems.value.push({
      id: data.id,
      name: data.item_name,
      general_price: data.general_price_sale,
      delivery_price: data.general_price_sale,
      type: 'main',
    })

    closeCreateMainModal()
    console.log('Created main menu item:', data)
  } catch (error) {
    console.error('Error creating menu item:', error)
    createError.value = 'Failed to create menu item. Please try again.'
  } finally {
    isCreating.value = false
  }
}

// Create Extra Menu Item
const createExtraMenuItem = async () => {
  if (
    !newExtraItem.value.name ||
    !newExtraItem.value.general_price
  ) {
    createError.value = 'Please fill in all required fields'
    return
  }

  const generalPrice = parseFloat(newExtraItem.value.general_price)
  if (isNaN(generalPrice) || generalPrice <= 0) {
    createError.value = 'Please enter valid price'
    return
  }

  isCreating.value = true
  createError.value = ''

  try {
    // Make API call to create extra item
    const response = await fetch(`${API_BASE_URL}/extra-items`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        merchant_id: merchantId.value,
        item_name: newExtraItem.value.name,
        general_price_sale: generalPrice,
        delivery_price_sale: generalPrice,
      }),
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const data = await response.json()

    // Add the new item to the local array
    extraMenuItems.value.push({
      id: data.id,
      name: data.item_name,
      general_price: data.general_price_sale,
      delivery_price: data.general_price_sale,
      type: 'extra',
    })

    closeCreateExtraModal()
    console.log('Created extra menu item:', data)
  } catch (error) {
    console.error('Error creating extra item:', error)
    createError.value = 'Failed to create extra item. Please try again.'
  } finally {
    isCreating.value = false
  }
}

// Delete Menu Item
const deleteMenuItem = async (itemId, type) => {
  if (!confirm('Are you sure you want to delete this item?')) {
    return
  }

  try {
    const endpoint = type === 'main' ? 'menu-items' : 'extra-items'
    const response = await fetch(`${API_BASE_URL}/${endpoint}/${itemId}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    // Remove from local array after successful deletion
    if (type === 'main') {
      mainMenuItems.value = mainMenuItems.value.filter((item) => item.id !== itemId)
    } else {
      extraMenuItems.value = extraMenuItems.value.filter((item) => item.id !== itemId)
    }

    console.log(`Deleted ${type} item:`, itemId)
  } catch (error) {
    console.error('Error deleting item:', error)
    alert('Failed to delete item. Please try again.')
  }
}

// Fetch menu items from API
const fetchMenuItems = async () => {
  if (!merchantId.value) return

  try {
    // Fetch main menu items
    const mainResponse = await fetch(`${API_BASE_URL}/menu-items?merchant_id=${merchantId.value}`)
    if (mainResponse.ok) {
      const mainData = await mainResponse.json()
      mainMenuItems.value = mainData.map((item) => ({
        id: item.id,
        name: item.item_name,
        general_price: item.general_price_sale || 0,
        delivery_price: item.delivery_price_sale || 0,
        type: 'main',
      }))
    }

    // Fetch extra items
    const extraResponse = await fetch(`${API_BASE_URL}/extra-items?merchant_id=${merchantId.value}`)
    if (extraResponse.ok) {
      const extraData = await extraResponse.json()
      extraMenuItems.value = extraData.map((item) => ({
        id: item.id,
        name: item.item_name,
        general_price: item.general_price_sale || 0,
        delivery_price: item.delivery_price_sale || 0,
        type: 'extra',
      }))
    }
  } catch (error) {
    console.error('Error fetching menu items:', error)
  }
}

// Edit modal functions
const openEditModal = (item, itemType) => {
  editingItem.value = { ...item, type: itemType }
  editForm.value = {
    name: item.name,
    general_price: item.general_price.toString(),
  }
  editFormErrors.value = {
    name: '',
    general_price: '',
  }
  updateError.value = ''
  showEditModal.value = true
}

const closeEditModal = () => {
  showEditModal.value = false
  editingItem.value = null
  editForm.value = {
    name: '',
    general_price: '',
  }
  editFormErrors.value = {
    name: '',
    general_price: '',
  }
  updateError.value = ''
}

const validateEditPrice = (priceType) => {
  let value = editForm.value.general_price
  const originalValue = value.toString()

  // Check if user tried to enter non-numeric characters
  const hasInvalidChars = /[^0-9.]/.test(originalValue)

  // Remove any non-numeric characters except decimal point
  value = originalValue.replace(/[^0-9.]/g, '')

  // Ensure only one decimal point
  const parts = value.split('.')
  if (parts.length > 2) {
    value = parts[0] + '.' + parts.slice(1).join('')
  }

  // Limit to 2 decimal places
  if (parts[1] && parts[1].length > 2) {
    value = parts[0] + '.' + parts[1].substring(0, 2)
  }

  // Update the model
  editForm.value.general_price = value

  // Set error message if invalid characters were entered
  if (hasInvalidChars) {
    editFormErrors.value.general_price = 'Only numbers are allowed for price'
  } else if (value && parseFloat(value) <= 0) {
    editFormErrors.value.general_price = 'Price must be greater than 0'
  } else {
    editFormErrors.value.general_price = ''
  }
}

const saveItemChanges = async () => {
  if (!editForm.value.name) {
    editFormErrors.value.name = 'Item name is required'
    return
  }

  const generalPrice = parseFloat(editForm.value.general_price)

  if (isNaN(generalPrice) || generalPrice <= 0) {
    editFormErrors.value.general_price = 'Please enter a valid price greater than 0'
    return
  }

  isUpdating.value = true
  updateError.value = ''

  try {
    const endpoint = editingItem.value.type === 'main' ? 'menu-items' : 'extra-items'
    const updateData = {
      item_name: editForm.value.name,
      general_price_sale: generalPrice,
      delivery_price_sale: generalPrice,
    }

    const response = await fetch(`${API_BASE_URL}/${endpoint}/${editingItem.value.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(updateData),
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const data = await response.json()

    // Update local data
    const targetArray = editingItem.value.type === 'main' ? mainMenuItems : extraMenuItems
    const itemIndex = targetArray.value.findIndex((item) => item.id === editingItem.value.id)
    if (itemIndex !== -1) {
      targetArray.value[itemIndex] = {
        ...targetArray.value[itemIndex],
        name: editForm.value.name,
        general_price: generalPrice,
        delivery_price: generalPrice,
      }
    }

    console.log(`Updated ${editingItem.value.type} item:`, data)
    closeEditModal()
  } catch (error) {
    console.error('Error updating item:', error)
    updateError.value = 'Failed to update item. Please try again.'
  } finally {
    isUpdating.value = false
  }
}

// Initialize
onMounted(() => {
  // Get merchant data from route params
  merchantId.value = route.params.merchantId || route.query.merchantId || ''
  merchantName.value = route.params.merchantName || route.query.merchantName || 'Merchant'

  // Fetch menu items from API
  console.log('Managing merchant:', merchantId.value, merchantName.value)
  fetchMenuItems()
})
</script>

<style scoped>
.manage-merchant-page {
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
  transform: translateX(-2px);
}

.page-header h1 {
  font-size: 28px;
  font-weight: 700;
  color: #111827;
  margin: 0;
}

/* Merchant Info Card */
.merchant-info-card {
  display: flex;
  align-items: center;
  gap: 16px;
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  margin-bottom: 24px;
}

.merchant-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 600;
}

.merchant-info-card h2 {
  font-size: 20px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 4px 0;
}

.merchant-id {
  font-size: 14px;
  color: #6b7280;
  margin: 0;
}

/* Tabs */
.tabs-container {
  margin-bottom: 24px;
}

.tabs {
  display: flex;
  background: white;
  border-radius: 12px;
  padding: 4px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.tab-btn {
  flex: 1;
  padding: 12px 20px;
  background: transparent;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  color: #6b7280;
  cursor: pointer;
  transition: all 0.2s;
}

.tab-btn:hover {
  color: #374151;
}

.tab-btn.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

/* Section */
.menu-section {
  background: white;
  border-radius: 16px;
  padding: 24px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.section-header h3 {
  font-size: 20px;
  font-weight: 600;
  color: #111827;
  margin: 0;
}

.create-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.2);
}

.create-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

/* Empty State */
.empty-state {
  text-align: center;
  padding: 48px 20px;
}

.empty-icon {
  width: 80px;
  height: 80px;
  margin: 0 auto 20px;
  background: #f3f4f6;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #9ca3af;
}

.empty-state h4 {
  font-size: 18px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 8px 0;
}

.empty-state p {
  font-size: 14px;
  color: #6b7280;
  margin: 0;
}

/* Items Grid */
.items-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 16px;
}

.menu-item-card {
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 16px;
  transition: all 0.2s;
}

.menu-item-card:hover {
  border-color: #d1d5db;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
}

.item-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 12px;
}

.item-header h4 {
  font-size: 16px;
  font-weight: 600;
  color: #111827;
  margin: 0;
  flex: 1;
}

.action-buttons {
  display: flex;
  gap: 4px;
}

.edit-btn {
  width: 28px;
  height: 28px;
  background: transparent;
  border: none;
  color: #3b82f6;
  cursor: pointer;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.edit-btn:hover {
  background: #dbeafe;
}

.delete-btn {
  width: 28px;
  height: 28px;
  background: transparent;
  border: none;
  color: #ef4444;
  cursor: pointer;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.delete-btn:hover {
  background: #fee2e2;
}

.item-price {
  font-size: 20px;
  font-weight: 700;
  color: #10b981;
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
}

.modal-content {
  background: white;
  border-radius: 16px;
  width: 100%;
  max-width: 500px;
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

.form-input.error-input {
  border-color: #ef4444;
  background: #fef2f2;
}

.form-input.error-input:focus {
  border-color: #ef4444;
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1);
}

.field-error {
  color: #ef4444;
  font-size: 12px;
  margin-top: 4px;
  margin-bottom: 0;
  font-weight: 500;
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

/* Item Pricing Display */
.item-pricing {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-top: 12px;
}

.price-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 6px 12px;
  background: #f9fafb;
  border-radius: 6px;
  border-left: 3px solid #10b981;
}

.price-item:last-child {
  border-left-color: #3b82f6;
}

.price-label {
  font-size: 12px;
  font-weight: 500;
  color: #6b7280;
  text-transform: uppercase;
}

.price-value {
  font-size: 14px;
  font-weight: 600;
  color: #10b981;
}

.price-value.delivery {
  color: #3b82f6;
}

/* Responsive */
@media (max-width: 768px) {
  .manage-merchant-page {
    padding: 16px;
  }

  .page-header h1 {
    font-size: 24px;
  }

  .items-grid {
    grid-template-columns: 1fr;
  }
}
</style>
