<template>
  <div class="tracking-page">
    <!-- Header -->
    <div class="page-header">
      <button @click="goBack" class="back-btn">
        <svg
          width="20"
          height="20"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <path d="M19 12H5M12 19l-7-7 7-7" />
        </svg>
        {{ t('back') }}
      </button>
      <div class="header-content">
        <h1 class="page-title">{{ merchantName }} - {{ t('tracking') }}</h1>
      </div>
    </div>

    <!-- Action Section -->
    <div class="action-section">
      <button @click="openDateSelector" class="create-tracking-btn">
        <svg
          width="20"
          height="20"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <path d="M12 5v14M5 12h14" />
        </svg>
        {{ t('createDailyTracking') }}
      </button>
      <div class="view-toggles">
        <button
          @click="setViewMode('daily')"
          class="view-toggle-btn"
          :class="{ active: viewMode === 'daily' }"
        >
          {{ t('daily') }}
        </button>
        <button
          @click="setViewMode('weekly')"
          class="view-toggle-btn"
          :class="{ active: viewMode === 'weekly' }"
        >
          {{ t('weekly') }}
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="loading-state">
      <div class="loading-spinner"></div>
      <p>{{ t('loading') }}</p>
    </div>

    <!-- Daily/Weekly Cards Grid -->
    <div v-else-if="currentTrackingData.length > 0" class="tracking-list">
      <!-- Daily View -->
      <template v-if="viewMode === 'daily'">
        <div v-for="tracking in currentTrackingData" :key="tracking.id" class="daily-card-compact">
          <div class="card-main" @click="selectDailyTracking(tracking)">
            <div class="card-date">
              {{ formatTrackingDate(tracking.trackingDate) }}
            </div>
            <div class="card-stats-equal">
              <div class="stat-section">
                <span class="stat-value income">{{ formatCurrency(tracking.income) }}</span>
              </div>
              <div class="stat-section">
                <span class="stat-value outcome">{{ formatCurrency(tracking.outcome) }}</span>
              </div>
              <div class="stat-section">
                <span class="stat-value revenue" :class="getRevenueClass(tracking.revenue)">
                  {{ formatCurrency(tracking.revenue) }}
                </span>
              </div>
            </div>
          </div>
          <button
            @click.stop="deleteDailyTracking(tracking)"
            class="delete-btn-right"
            title="Delete tracking"
          >
            <svg
              width="18"
              height="18"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                d="M3 6h18M8 6V4a2 2 0 012-2h4a2 2 0 012 2v2M10 11v6M14 11v6M5 6l1 13a2 2 0 002 2h8a2 2 0 002-2l1-13"
              />
            </svg>
          </button>
        </div>
      </template>

      <!-- Weekly View -->
      <template v-if="viewMode === 'weekly'">
        <div
          v-for="week in currentTrackingData"
          :key="week.id"
          class="daily-card-compact weekly-card"
        >
          <div class="card-main" @click="selectWeeklyData(week)">
            <div class="card-date weekly-date">
              <div class="week-range">{{ formatWeekRange(week.startDate, week.endDate) }}</div>
              <div class="week-info">
                <span>{{ t('week') }} {{ week.weekNumber }} {{ t('of') }} {{ week.year }}</span>
                <span class="summary-badge">{{ t('summaryOnly') }}</span>
              </div>
            </div>
            <div class="card-stats-equal">
              <div class="stat-section">
                <span class="stat-value income">{{ formatCurrency(week.totalIncome) }}</span>
              </div>
              <div class="stat-section">
                <span class="stat-value outcome">{{ formatCurrency(week.totalOutcome) }}</span>
              </div>
              <div class="stat-section">
                <span class="stat-value revenue" :class="getRevenueClass(week.totalRevenue)">
                  {{ formatCurrency(week.totalRevenue) }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </template>
    </div>

    <!-- Empty State -->
    <div v-else-if="!isLoading" class="empty-state">
      <svg
        width="64"
        height="64"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="1.5"
      >
        <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
        <line x1="16" y1="2" x2="16" y2="6" />
        <line x1="8" y1="2" x2="8" y2="6" />
        <line x1="3" y1="10" x2="21" y2="10" />
      </svg>
      <p class="empty-title">{{ t('noDailyTracking') }}</p>
      <p class="empty-subtitle">{{ t('createDailyTrackingFirst') }}</p>
      <button @click="openDateSelector" class="empty-create-btn">
        {{ t('createDailyTracking') }}
      </button>
    </div>

    <!-- Daily Tracking Details Modal -->
    <div v-if="showTrackingDetails" class="modal-overlay" @click.self="closeTrackingDetails">
      <div class="modal-content tracking-details">
        <div class="modal-header">
          <h2>{{ formatTrackingDate(selectedTracking?.trackingDate) }}</h2>
          <button @click="closeTrackingDetails" class="close-btn">×</button>
        </div>

        <div class="modal-body">
          <!-- Tabs -->
          <div class="tabs">
            <button
              @click="activeTab = 'income'"
              class="tab-btn"
              :class="{ active: activeTab === 'income' }"
            >
              Income
            </button>
            <button
              @click="activeTab = 'outcome'"
              class="tab-btn"
              :class="{ active: activeTab === 'outcome' }"
            >
              Outcome
            </button>
          </div>

          <!-- Tab Content -->
          <div class="tab-content">
            <!-- Income Tab -->
            <div v-if="activeTab === 'income'" class="income-section">
              <div class="section-header">
                <h3>Income Details</h3>
                <button class="add-btn">
                  <svg
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                  >
                    <path d="M12 5v14M5 12h14" />
                  </svg>
                  Add Income
                </button>
              </div>

              <div v-if="selectedTracking?.incomeItems?.length > 0" class="items-list">
                <div v-for="item in selectedTracking.incomeItems" :key="item.id" class="item-row">
                  <span class="item-name">{{ item.name }}</span>
                  <span class="item-amount income">{{ formatCurrency(item.amount) }}</span>
                </div>
              </div>
              <div v-else class="empty-items">
                <p>No income recorded yet</p>
              </div>

              <div class="total-row">
                <span>Total Income</span>
                <span class="total-amount income">{{
                  formatCurrency(selectedTracking?.income || 0)
                }}</span>
              </div>
            </div>

            <!-- Outcome Tab -->
            <div v-if="activeTab === 'outcome'" class="outcome-section">
              <div class="section-header">
                <h3>Outcome Details</h3>
                <button class="add-btn">
                  <svg
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                  >
                    <path d="M12 5v14M5 12h14" />
                  </svg>
                  Add Expense
                </button>
              </div>

              <div v-if="selectedTracking?.outcomeItems?.length > 0" class="items-list">
                <div v-for="item in selectedTracking.outcomeItems" :key="item.id" class="item-row">
                  <span class="item-name">{{ item.name }}</span>
                  <span class="item-amount outcome">{{ formatCurrency(item.amount) }}</span>
                </div>
              </div>
              <div v-else class="empty-items">
                <p>No expenses recorded yet</p>
              </div>

              <div class="total-row">
                <span>Total Outcome</span>
                <span class="total-amount outcome">{{
                  formatCurrency(selectedTracking?.outcome || 0)
                }}</span>
              </div>
            </div>
          </div>

          <!-- Summary -->
          <div class="summary-section">
            <div class="summary-item">
              <span>Net Revenue</span>
              <span
                class="summary-value"
                :class="{
                  positive: selectedTracking?.revenue >= 0,
                  negative: selectedTracking?.revenue < 0,
                }"
              >
                {{ formatCurrency(selectedTracking?.revenue || 0) }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Date Selector Component -->
    <div v-if="showDateSelector" class="modal-overlay" @click.self="closeDateSelector">
      <div class="modal-content calendar-modal">
        <div class="modal-header">
          <h2>{{ t('selectDateForDailyTracking') }}</h2>
          <button @click="closeDateSelector" class="close-btn">×</button>
        </div>
        <div class="modal-body">
          <!-- Calendar Navigation -->
          <div class="calendar-header">
            <button @click="previousMonth" class="nav-btn">
              <svg
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path d="M15 18l-6-6 6-6" />
              </svg>
            </button>
            <h3 class="month-title">{{ formatMonthYear(currentMonth, currentYear) }}</h3>
            <button @click="nextMonth" class="nav-btn" :disabled="isCurrentMonth">
              <svg
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path d="M9 18l6-6-6-6" />
              </svg>
            </button>
          </div>

          <!-- Calendar Grid -->
          <div class="calendar-container">
            <!-- Day Headers -->
            <div class="calendar-grid">
              <div class="day-header" v-for="day in dayHeaders" :key="day">{{ day }}</div>

              <!-- Calendar Days -->
              <button
                v-for="date in calendarDays"
                :key="date.key"
                @click="selectDate(date)"
                class="calendar-day"
                :class="{
                  'other-month': !date.isCurrentMonth,
                  today: date.isToday,
                  selected: date.dateString === selectedTrackingDate,
                  disabled: date.isFuture,
                }"
                :disabled="date.isFuture || !date.isCurrentMonth"
              >
                {{ date.day }}
              </button>
            </div>
          </div>

          <div class="selected-date-display" v-if="selectedTrackingDate">
            <p>
              {{ t('selected') }}: <strong>{{ formatSelectedDate(selectedTrackingDate) }}</strong>
            </p>
          </div>
        </div>
        <div class="modal-footer">
          <button @click="closeDateSelector" class="btn-secondary">{{ t('cancel') }}</button>
          <button
            @click="createDailyTracking"
            class="btn-primary"
            :disabled="!selectedTrackingDate"
          >
            {{ t('createTracking') }}
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

// Language
const { t, formatCurrency: formatCurrencyLang } = useLanguage()

// Reactive data
const merchantId = ref('')
const merchantName = ref('')
const dailyTrackings = ref([])
const weeklyData = ref([])
const viewMode = ref('daily') // 'daily' or 'weekly'
const showDateSelector = ref(false)
const showTrackingDetails = ref(false)
const selectedTracking = ref(null)
const selectedTrackingDate = ref('')
const activeTab = ref('income')
const isLoading = ref(false)

// Calendar data
const currentMonth = ref(new Date().getMonth())
const currentYear = ref(new Date().getFullYear())
const dayHeaders = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']

// Computed
const currentTrackingData = computed(() => {
  return viewMode.value === 'daily' ? dailyTrackings.value : weeklyData.value
})

// Use language-aware currency formatting
const formatCurrency = formatCurrencyLang

// Get revenue color class based on value
const getRevenueClass = (revenue) => {
  if (revenue > 0) return 'positive'
  if (revenue === 0) return 'neutral'
  return 'negative'
}

// Format tracking date
const formatTrackingDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('th-TH', {
    weekday: 'short',
    day: 'numeric',
    month: 'short',
    year: 'numeric',
  })
}

// Calendar format functions
const formatMonthYear = (month, year) => {
  const date = new Date(year, month)
  return date.toLocaleDateString('th-TH', {
    month: 'long',
    year: 'numeric',
  })
}

const formatSelectedDate = (dateString) => {
  if (!dateString) return ''

  // Parse the date string as local date to avoid timezone issues
  // dateString format is "YYYY-MM-DD"
  const [year, month, day] = dateString.split('-').map(Number)
  const date = new Date(year, month - 1, day) // month is 0-indexed

  return date.toLocaleDateString('th-TH', {
    weekday: 'long',
    day: 'numeric',
    month: 'long',
    year: 'numeric',
  })
}

// Weekly format functions
const formatWeekRange = (startDate, endDate) => {
  const start = new Date(startDate)
  const end = new Date(endDate)

  const startFormatted = start.toLocaleDateString('th-TH', {
    day: 'numeric',
    month: 'short',
  })

  const endFormatted = end.toLocaleDateString('th-TH', {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
  })

  return `${startFormatted} - ${endFormatted}`
}

// View mode management
const setViewMode = (mode) => {
  viewMode.value = mode
  if (mode === 'weekly') {
    generateWeeklyData()
  }
}

// Navigation
const goBack = () => {
  router.push('/home')
}

// Calendar computed properties
const isCurrentMonth = computed(() => {
  const today = new Date()
  return currentMonth.value === today.getMonth() && currentYear.value === today.getFullYear()
})

const calendarDays = computed(() => {
  const firstDay = new Date(currentYear.value, currentMonth.value, 1)
  // const lastDay = new Date(currentYear.value, currentMonth.value + 1, 0) // Not used
  const startDate = new Date(firstDay)
  startDate.setDate(startDate.getDate() - firstDay.getDay())

  const days = []
  const today = new Date()
  today.setHours(0, 0, 0, 0)

  for (let i = 0; i < 42; i++) {
    const date = new Date(startDate)
    date.setDate(startDate.getDate() + i)

    const isCurrentMonth = date.getMonth() === currentMonth.value
    // Create date string in local timezone to avoid UTC conversion issues
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    const day = String(date.getDate()).padStart(2, '0')
    const dateString = `${year}-${month}-${day}`
    const isToday = date.getTime() === today.getTime()
    const isFuture = date > today

    days.push({
      day: date.getDate(),
      date: date,
      dateString: dateString,
      isCurrentMonth: isCurrentMonth,
      isToday: isToday,
      isFuture: isFuture,
      key: `${date.getFullYear()}-${date.getMonth()}-${date.getDate()}`,
    })
  }

  return days
})

// Calendar navigation
const previousMonth = () => {
  if (currentMonth.value === 0) {
    currentMonth.value = 11
    currentYear.value--
  } else {
    currentMonth.value--
  }
}

const nextMonth = () => {
  if (isCurrentMonth.value) return // Don't allow future months

  if (currentMonth.value === 11) {
    currentMonth.value = 0
    currentYear.value++
  } else {
    currentMonth.value++
  }
}

// Date Selector handlers
const openDateSelector = () => {
  showDateSelector.value = true
  selectedTrackingDate.value = ''
  // Reset to current month
  const today = new Date()
  currentMonth.value = today.getMonth()
  currentYear.value = today.getFullYear()
}

const closeDateSelector = () => {
  showDateSelector.value = false
  selectedTrackingDate.value = ''
}

const selectDate = (dateObj) => {
  if (dateObj.isFuture || !dateObj.isCurrentMonth) return
  selectedTrackingDate.value = dateObj.dateString
}

const createDailyTracking = async () => {
  if (!selectedTrackingDate.value) return

  try {
    // Call backend API to create daily tracking
    const response = await fetch(`${API_BASE_URL}/daily-tracking`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        merchant_id: merchantId.value,
        tracking_date: selectedTrackingDate.value,
      }),
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const data = await response.json()

    // If tracking already exists, navigate to it
    if (data.exists) {
      console.log('Daily tracking already exists:', data.tracking_id)
      router.push({
        name: 'DailyTrackingPage',
        params: { trackingId: data.tracking_id },
        query: {
          merchantId: merchantId.value,
          merchantName: merchantName.value,
        },
      })
    } else {
      // Add new tracking to the list
      const income = data.total_income || 0
      const outcome = data.total_outcome || 0
      const revenue = income - outcome // Calculate revenue as income - outcome

      const newTracking = {
        id: data.tracking_id,
        merchantId: data.merchant_id,
        trackingDate: data.tracking_date,
        income: income,
        outcome: outcome,
        revenue: revenue,
      }

      dailyTrackings.value.push(newTracking)
      console.log('Created daily tracking:', data.tracking_id)
    }

    closeDateSelector()
  } catch (error) {
    console.error('Error creating daily tracking:', error)
    alert('Failed to create daily tracking. Please try again.')
  }
}

// Weekly data generation
const generateWeeklyData = () => {
  if (dailyTrackings.value.length === 0) {
    weeklyData.value = []
    return
  }

  // Group daily trackings by week
  const weekGroups = new Map()

  dailyTrackings.value.forEach((tracking) => {
    const date = new Date(tracking.trackingDate)
    const weekStart = getWeekStart(date)
    const weekKey = weekStart.toISOString().split('T')[0]

    if (!weekGroups.has(weekKey)) {
      const weekEnd = new Date(weekStart)
      weekEnd.setDate(weekStart.getDate() + 6)

      weekGroups.set(weekKey, {
        id: `week-${weekKey}`,
        startDate: weekKey,
        endDate: weekEnd.toISOString().split('T')[0],
        weekNumber: getWeekNumber(weekStart),
        year: weekStart.getFullYear(),
        totalIncome: 0,
        totalOutcome: 0,
        totalRevenue: 0,
        dailyTrackings: [],
      })
    }

    const week = weekGroups.get(weekKey)
    week.totalIncome += tracking.income
    week.totalOutcome += tracking.outcome
    week.totalRevenue += tracking.revenue
    week.dailyTrackings.push(tracking)
  })

  // Convert to array and sort by date (most recent first)
  weeklyData.value = Array.from(weekGroups.values()).sort(
    (a, b) => new Date(b.startDate) - new Date(a.startDate),
  )
}

// Helper functions for weekly data
const getWeekStart = (date) => {
  const start = new Date(date)
  const day = start.getDay()
  // Monday as start of week (1), adjust Sunday (0) to be 7
  const adjustedDay = day === 0 ? 7 : day
  const diff = start.getDate() - adjustedDay + 1
  start.setDate(diff)
  start.setHours(0, 0, 0, 0)
  return start
}

const getWeekNumber = (date) => {
  // ISO week calculation (Monday-based)
  const target = new Date(date.valueOf())
  const dayNr = (date.getDay() + 6) % 7 // Monday = 0, Sunday = 6
  target.setDate(target.getDate() - dayNr + 3) // Nearest Thursday
  const jan4 = new Date(target.getFullYear(), 0, 4)
  const dayDiff = (target - jan4) / 86400000
  return 1 + Math.ceil(dayDiff / 7)
}

// Select daily tracking for details
const selectDailyTracking = (tracking) => {
  console.log('Selecting daily tracking:', tracking)

  router.push({
    name: 'DailyTrackingPage',
    params: { trackingId: tracking.id },
    query: {
      merchantId: merchantId.value,
      merchantName: merchantName.value,
    },
  })
}

// Select weekly data (show error warning)
const selectWeeklyData = (week) => {
  console.log('Selecting weekly data:', week)

  // Show warning message to switch to daily view
  const shouldSwitchToDaily = confirm(t('weeklyViewWarning'))

  if (shouldSwitchToDaily) {
    setViewMode('daily')
  }
}

const closeTrackingDetails = () => {
  showTrackingDetails.value = false
  selectedTracking.value = null
}

// Delete daily tracking
const deleteDailyTracking = async (tracking) => {
  if (
    !confirm(
      `Are you sure you want to delete tracking for "${formatTrackingDate(tracking.trackingDate)}"?`,
    )
  ) {
    return
  }

  try {
    const response = await fetch(`${API_BASE_URL}/daily-tracking/${tracking.id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    // Remove from local list
    dailyTrackings.value = dailyTrackings.value.filter((t) => t.id !== tracking.id)

    console.log('Deleted daily tracking:', tracking.id)
  } catch (error) {
    console.error('Error deleting daily tracking:', error)
    alert('Failed to delete daily tracking. Please try again.')
  }
}

// Fetch daily trackings from backend
const fetchDailyTrackings = async () => {
  if (!merchantId.value) {
    console.warn('No merchant ID provided, cannot fetch daily trackings')
    return
  }

  isLoading.value = true
  try {
    // Get recent daily trackings (last 30 days)
    const endDate = new Date().toISOString().split('T')[0]
    const startDate = new Date(Date.now() - 30 * 24 * 60 * 60 * 1000).toISOString().split('T')[0]

    const url = `${API_BASE_URL}/daily-tracking?merchant_id=${merchantId.value}&start_date=${startDate}&end_date=${endDate}`
    console.log('Fetching daily trackings from:', url)

    const response = await fetch(url)

    console.log('API Response status:', response.status)

    if (!response.ok) {
      const errorText = await response.text()
      console.error('API Error response:', errorText)
      throw new Error(`HTTP error! status: ${response.status}, message: ${errorText}`)
    }

    const data = await response.json()
    console.log('Raw API data:', data)

    // Handle case where API returns empty array or null
    if (!data || !Array.isArray(data)) {
      console.log('No daily trackings data received or invalid format')
      dailyTrackings.value = []
      return
    }

    // Transform data for frontend
    dailyTrackings.value = data
      .map((tracking) => {
        const income = tracking.total_income || 0
        const outcome = tracking.total_outcome || 0
        const revenue = income - outcome // Calculate revenue as income - outcome

        return {
          id: tracking.tracking_id,
          merchantId: tracking.merchant_id,
          trackingDate: tracking.tracking_date,
          income: income,
          outcome: outcome,
          revenue: revenue,
        }
      })
      .sort((a, b) => new Date(b.trackingDate) - new Date(a.trackingDate)) // Sort by date descending

    console.log('Processed daily trackings:', dailyTrackings.value.length, dailyTrackings.value)

    // Generate weekly data after fetching daily data
    generateWeeklyData()
  } catch (error) {
    console.error('Error fetching daily trackings:', error)
    // Set empty array so UI shows empty state instead of nothing
    dailyTrackings.value = []
    weeklyData.value = []
  } finally {
    isLoading.value = false
  }
}

// Initialize
onMounted(async () => {
  console.log('🚀 TRACKING PAGE LOADED!')
  console.log('=== TrackingPage Mounted ===')
  console.log('Route params:', route.params)
  console.log('Route query:', route.query)

  // Get merchant data from route params
  merchantId.value = route.params.merchantId || route.query.merchantId || ''
  merchantName.value = route.query.merchantName || 'Merchant'

  console.log('Merchant ID:', merchantId.value)
  console.log('Merchant Name:', merchantName.value)
  console.log('API Base URL:', API_BASE_URL)

  if (!merchantId.value) {
    console.error('No merchant ID found in route params or query!')
    return
  }

  // Fetch existing daily trackings from API
  console.log('Starting to fetch daily trackings...')
  await fetchDailyTrackings()
  console.log('Finished fetching daily trackings')
})
</script>

<style scoped>
.tracking-page {
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

.header-content {
  flex: 1;
  display: flex;
  justify-content: space-between;
  align-items: center;
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
  margin: 0;
}

/* Action Section */
.action-section {
  margin-bottom: 24px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.create-tracking-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  background: #10b981;
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.create-tracking-btn:hover {
  background: #059669;
}

/* View Toggles */
.view-toggles {
  display: flex;
  background: #f3f4f6;
  border-radius: 10px;
  padding: 4px;
  gap: 2px;
}

.view-toggle-btn {
  padding: 8px 16px;
  background: transparent;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  color: #6b7280;
  cursor: pointer;
  transition: all 0.2s;
}

.view-toggle-btn.active {
  background: white;
  color: #374151;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.view-toggle-btn:hover:not(.active) {
  color: #374151;
}

/* Tracking List */
.tracking-list {
  display: flex;
  flex-direction: column;
  gap: 14px;
  margin-bottom: 24px;
}

.daily-card-compact {
  background: white;
  border-radius: 12px;
  display: flex;
  align-items: center;
  transition: all 0.2s;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  border: 1px solid #e5e7eb;
  overflow: hidden;
  min-height: 120px;
}

.daily-card-compact:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border-color: #d1d5db;
}

.card-main {
  display: flex;
  align-items: center;
  padding: 20px 24px;
  flex: 1;
  cursor: pointer;
  transition: background-color 0.2s;
  min-height: 120px;
}

.card-main:hover {
  background-color: #f9fafb;
}

.card-date {
  font-size: 16px;
  font-weight: 600;
  color: #374151;
  min-width: 160px;
  margin-right: 24px;
}

.card-stats-equal {
  display: flex;
  flex: 1;
}

.stat-section {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

.stat-section:first-child {
  flex: 0 0 33%;
}

.stat-section:nth-child(2) {
  flex: 0 0 34%;
}

.stat-section:last-child {
  flex: 0 0 33%;
}

.stat-value {
  font-size: 15px !important;
  font-weight: 700;
  padding: 8px 12px;
  border-radius: 10px;
  background: #f9fafb;
  white-space: nowrap;
  min-height: 40px;
  min-width: 90px;
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

.stat-value.revenue.positive {
  color: #10b981;
  background: #f0fdf4;
}

.stat-value.revenue.negative {
  color: #ef4444;
  background: #fef2f2;
}

.stat-value.revenue.neutral {
  color: #9ca3af;
  background: #f3f4f6;
}

.delete-btn-right {
  width: 44px;
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fee2e2;
  border: none;
  border-radius: 10px;
  color: #dc2626;
  cursor: pointer;
  transition: all 0.2s;
  margin-right: 12px;
}

.delete-btn-right:hover {
  background: #fee2e2;
  color: #dc2626;
  transform: scale(1.1);
}

/* Weekly Card Styles */
.weekly-card {
  border-left: 4px solid #3b82f6;
}

.weekly-date {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.week-range {
  font-size: 15px;
  font-weight: 600;
  color: #374151;
}

.week-info {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
  color: #6b7280;
  font-weight: 500;
}

.summary-badge {
  background: #fef3c7;
  color: #92400e;
  padding: 2px 6px;
  border-radius: 4px;
  font-size: 10px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

/* Loading State */
.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 400px;
  text-align: center;
}

.loading-spinner {
  width: 48px;
  height: 48px;
  border: 4px solid #f3f4f6;
  border-top: 4px solid #10b981;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 20px;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.loading-state p {
  font-size: 16px;
  color: #6b7280;
  margin: 0;
}

/* Empty State */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 400px;
  text-align: center;
}

.empty-state svg {
  color: #d1d5db;
  margin-bottom: 20px;
}

.empty-title {
  font-size: 18px;
  font-weight: 600;
  color: #374151;
  margin: 0 0 8px 0;
}

.empty-subtitle {
  font-size: 14px;
  color: #6b7280;
  margin: 0 0 24px 0;
  max-width: 400px;
}

.empty-create-btn {
  padding: 12px 24px;
  background: #10b981;
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.empty-create-btn:hover {
  background: #059669;
}

/* Modal Overlay */
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
  max-width: 600px;
  max-height: 80vh;
  overflow-y: auto;
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
}

/* Tabs */
.tabs {
  display: flex;
  gap: 8px;
  margin-bottom: 24px;
  border-bottom: 1px solid #e5e7eb;
}

.tab-btn {
  padding: 12px 24px;
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

.tab-content {
  min-height: 200px;
}

/* Section Header */
.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.section-header h3 {
  font-size: 16px;
  font-weight: 600;
  color: #374151;
  margin: 0;
}

.add-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  background: #f3f4f6;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  color: #374151;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.add-btn:hover {
  background: #e5e7eb;
}

/* Items List */
.items-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 20px;
}

.item-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px;
  background: #f9fafb;
  border-radius: 8px;
}

.item-name {
  font-size: 14px;
  color: #374151;
}

.item-amount {
  font-size: 14px;
  font-weight: 600;
}

.item-amount.income {
  color: #10b981;
}

.item-amount.outcome {
  color: #ef4444;
}

.empty-items {
  text-align: center;
  padding: 40px 20px;
  color: #9ca3af;
  font-size: 14px;
}

/* Total Row */
.total-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background: #f3f4f6;
  border-radius: 8px;
  margin-top: 20px;
}

.total-row span:first-child {
  font-size: 14px;
  font-weight: 500;
  color: #374151;
}

.total-amount {
  font-size: 18px;
  font-weight: 700;
}

.total-amount.income {
  color: #10b981;
}

.total-amount.outcome {
  color: #ef4444;
}

/* Summary Section */
.summary-section {
  margin-top: 24px;
  padding: 20px;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%);
  border-radius: 12px;
}

.summary-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.summary-item span:first-child {
  font-size: 16px;
  font-weight: 500;
  color: #374151;
}

.summary-value {
  font-size: 24px;
  font-weight: 700;
  color: #111827;
}

.summary-value.positive {
  color: #10b981;
}

.summary-value.negative {
  color: #ef4444;
}

/* Form Styles for Date Selector */
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

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 16px 24px;
  border-top: 1px solid #e5e7eb;
}

.btn-secondary,
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

/* Calendar Styles */
.calendar-modal {
  max-width: 400px;
}

.calendar-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
  padding: 0 10px;
}

.nav-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  background: #f3f4f6;
  border: none;
  border-radius: 8px;
  color: #374151;
  cursor: pointer;
  transition: all 0.2s;
}

.nav-btn:hover:not(:disabled) {
  background: #e5e7eb;
}

.nav-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.month-title {
  font-size: 18px;
  font-weight: 600;
  color: #111827;
  margin: 0;
}

.calendar-container {
  background: #f9fafb;
  border-radius: 12px;
  padding: 16px;
}

.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 2px;
}

.day-header {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 40px;
  font-size: 12px;
  font-weight: 600;
  color: #6b7280;
  text-transform: uppercase;
}

.calendar-day {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 40px;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
  cursor: pointer;
  transition: all 0.2s;
}

.calendar-day:hover:not(:disabled) {
  background: #f3f4f6;
  border-color: #10b981;
}

.calendar-day.other-month {
  color: #d1d5db;
  background: #f9fafb;
  cursor: not-allowed;
}

.calendar-day.today {
  background: #dbeafe;
  color: #1e40af;
  border-color: #3b82f6;
  font-weight: 700;
}

.calendar-day.selected {
  background: #10b981;
  color: white;
  border-color: #10b981;
  font-weight: 700;
}

.calendar-day.selected:hover {
  background: #059669;
  border-color: #059669;
}

.calendar-day.disabled {
  opacity: 0.3;
  cursor: not-allowed;
}

.selected-date-display {
  margin-top: 20px;
  padding: 12px 16px;
  background: #f0fdf4;
  border: 1px solid #bbf7d0;
  border-radius: 8px;
  text-align: center;
}

.selected-date-display p {
  margin: 0;
  font-size: 14px;
  color: #166534;
}

.selected-date-display strong {
  color: #15803d;
}

/* Responsive Design */
@media (max-width: 768px) {
  .daily-card-compact {
    position: relative;
  }

  .card-main {
    flex-direction: column;
    align-items: flex-start;
    padding: 12px 16px;
    padding-right: 50px;
  }

  .card-date {
    min-width: auto;
    margin-right: 0;
    margin-bottom: 10px;
    font-size: 13px;
  }

  .card-stats-equal {
    width: 100%;
  }

  .stat-section {
    justify-content: flex-start;
  }

  .stat-section:first-child,
  .stat-section:nth-child(2),
  .stat-section:last-child {
    flex: 0 0 33.33%;
  }

  .stat-value {
    font-size: 14px;
    padding: 4px 8px;
  }

  .delete-btn-right {
    position: absolute;
    top: 12px;
    right: 8px;
    width: 32px;
    height: 32px;
    margin-right: 0;
  }

  .delete-btn-right svg {
    width: 16px;
    height: 16px;
  }
}
</style>
