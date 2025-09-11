<template>
  <div class="modal-overlay" @click.self="close">
    <div class="modal-content">
      <div class="modal-header">
        <h2>Select Weekly Period</h2>
        <button @click="close" class="close-btn">×</button>
      </div>
      
      <div class="modal-body">
        <!-- Month/Year Navigation -->
        <div class="calendar-nav">
          <button @click="previousMonth" class="nav-btn">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M15 18l-6-6 6-6"/>
            </svg>
          </button>
          <div class="current-month">
            {{ monthNames[currentMonth] }} {{ currentYear }}
          </div>
          <button @click="nextMonth" class="nav-btn">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M9 18l6-6-6-6"/>
            </svg>
          </button>
        </div>

        <!-- Calendar Table -->
        <div class="calendar-wrapper">
          <table class="calendar-table">
            <thead>
              <tr>
                <th v-for="day in weekDays" :key="day" class="weekday">{{ day }}</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(week, weekIndex) in calendarWeeks" :key="weekIndex">
                <td 
                  v-for="(day, dayIndex) in week" 
                  :key="`${weekIndex}-${dayIndex}`"
                  @click="selectDate(day)"
                  class="calendar-day"
                  :class="{
                    'other-month': day.otherMonth,
                    'selected': isInSelectedRange(day),
                    'selected-start': isRangeStart(day),
                    'selected-end': isRangeEnd(day),
                    'today': isToday(day),
                    'disabled': day.disabled
                  }"
                >
                  <span v-if="day.date">{{ day.date }}</span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Selected Period Display -->
        <div v-if="selectedStartDate && selectedEndDate" class="selected-period">
          <div class="period-icon">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
              <line x1="16" y1="2" x2="16" y2="6"/>
              <line x1="8" y1="2" x2="8" y2="6"/>
              <line x1="3" y1="10" x2="21" y2="10"/>
            </svg>
          </div>
          <div class="period-text">
            <div class="period-label">Selected Period</div>
            <div class="period-dates">{{ formatDisplayDate(selectedStartDate) }} - {{ formatDisplayDate(selectedEndDate) }}</div>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="action-buttons">
          <button @click="close" class="btn-cancel">Cancel</button>
          <button 
            @click="confirmSelection" 
            :disabled="!selectedStartDate || !selectedEndDate"
            class="btn-confirm"
          >
            Confirm Selection
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

const props = defineProps({
  existingWeeks: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['close', 'select'])

// Calendar data
const currentDate = new Date()
const currentMonth = ref(currentDate.getMonth())
const currentYear = ref(currentDate.getFullYear())
const selectedStartDate = ref(null)
const selectedEndDate = ref(null)

const weekDays = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
const monthNames = [
  'January', 'February', 'March', 'April', 'May', 'June',
  'July', 'August', 'September', 'October', 'November', 'December'
]

// Check if current month is displayed
const isCurrentMonth = computed(() => {
  const now = new Date()
  return currentMonth.value === now.getMonth() && currentYear.value === now.getFullYear()
})

// Generate calendar weeks
const calendarWeeks = computed(() => {
  const year = currentYear.value
  const month = currentMonth.value
  const firstDay = new Date(year, month, 1)
  const lastDay = new Date(year, month + 1, 0)
  const startCalendar = new Date(firstDay)
  const endCalendar = new Date(lastDay)
  
  // Adjust to start from Sunday
  startCalendar.setDate(startCalendar.getDate() - firstDay.getDay())
  
  // Adjust to end on Saturday
  const daysToAdd = 6 - lastDay.getDay()
  endCalendar.setDate(endCalendar.getDate() + daysToAdd)
  
  const weeks = []
  let currentWeek = []
  const today = new Date()
  today.setHours(0, 0, 0, 0)
  
  for (let d = new Date(startCalendar); d <= endCalendar; d.setDate(d.getDate() + 1)) {
    const dayObj = {
      date: d.getDate(),
      fullDate: new Date(d),
      otherMonth: d.getMonth() !== month,
      disabled: false
    }
    
    currentWeek.push(dayObj)
    
    if (currentWeek.length === 7) {
      weeks.push(currentWeek)
      currentWeek = []
    }
  }
  
  return weeks
})

// Navigation functions
const previousMonth = () => {
  if (currentMonth.value === 0) {
    currentMonth.value = 11
    currentYear.value--
  } else {
    currentMonth.value--
  }
}

const nextMonth = () => {
  if (currentMonth.value === 11) {
    currentMonth.value = 0
    currentYear.value++
  } else {
    currentMonth.value++
  }
}

// Date selection
const selectDate = (day) => {
  if (day.disabled || day.otherMonth) return
  
  // Calculate the 7-day period starting from selected date
  const startDate = new Date(day.fullDate)
  const endDate = new Date(day.fullDate)
  endDate.setDate(endDate.getDate() + 6)
  
  selectedStartDate.value = startDate
  selectedEndDate.value = endDate
}

// Check if date is in selected range
const isInSelectedRange = (day) => {
  if (!selectedStartDate.value || !selectedEndDate.value || day.otherMonth) return false
  
  const date = new Date(day.fullDate)
  date.setHours(0, 0, 0, 0)
  const start = new Date(selectedStartDate.value)
  start.setHours(0, 0, 0, 0)
  const end = new Date(selectedEndDate.value)
  end.setHours(0, 0, 0, 0)
  
  return date >= start && date <= end
}

// Check if date is range start
const isRangeStart = (day) => {
  if (!selectedStartDate.value || day.otherMonth) return false
  
  const date = new Date(day.fullDate)
  date.setHours(0, 0, 0, 0)
  const start = new Date(selectedStartDate.value)
  start.setHours(0, 0, 0, 0)
  
  return date.getTime() === start.getTime()
}

// Check if date is range end
const isRangeEnd = (day) => {
  if (!selectedEndDate.value || day.otherMonth) return false
  
  const date = new Date(day.fullDate)
  date.setHours(0, 0, 0, 0)
  const end = new Date(selectedEndDate.value)
  end.setHours(0, 0, 0, 0)
  
  return date.getTime() === end.getTime()
}

// Check if date is today
const isToday = (day) => {
  const today = new Date()
  const date = new Date(day.fullDate)
  
  return date.getDate() === today.getDate() &&
    date.getMonth() === today.getMonth() &&
    date.getFullYear() === today.getFullYear()
}

// Format date for display
const formatDisplayDate = (date) => {
  if (!date) return ''
  const d = new Date(date)
  const day = d.getDate()
  const month = monthNames[d.getMonth()]
  const year = d.getFullYear()
  return `${day} ${month} ${year}`
}

// Format date range for API
const formatDateRange = (start, end) => {
  const s = new Date(start)
  const e = new Date(end)
  const monthNames = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
  
  let dateStr = ''
  // If same month, show: "8 - 14 Sep 2025"
  if (s.getMonth() === e.getMonth() && s.getFullYear() === e.getFullYear()) {
    dateStr = `${s.getDate()} - ${e.getDate()} ${monthNames[s.getMonth()]} ${s.getFullYear()}`
  }
  // If different months same year, show: "28 Aug - 3 Sep 2025"
  else if (s.getFullYear() === e.getFullYear()) {
    dateStr = `${s.getDate()} ${monthNames[s.getMonth()]} - ${e.getDate()} ${monthNames[e.getMonth()]} ${s.getFullYear()}`
  }
  // If different years, show: "28 Dec 2024 - 3 Jan 2025"
  else {
    dateStr = `${s.getDate()} ${monthNames[s.getMonth()]} ${s.getFullYear()} - ${e.getDate()} ${monthNames[e.getMonth()]} ${e.getFullYear()}`
  }
  
  return dateStr
}

// Get week number
const getWeekNumber = (date) => {
  const d = new Date(Date.UTC(date.getFullYear(), date.getMonth(), date.getDate()))
  const dayNum = d.getUTCDay() || 7
  d.setUTCDate(d.getUTCDate() + 4 - dayNum)
  const yearStart = new Date(Date.UTC(d.getUTCFullYear(), 0, 1))
  return Math.ceil((((d - yearStart) / 86400000) + 1) / 7)
}

// Confirm selection
const confirmSelection = () => {
  if (!selectedStartDate.value || !selectedEndDate.value) return
  
  const weekNum = getWeekNumber(selectedStartDate.value)
  const year = selectedStartDate.value.getFullYear()
  const dateRange = formatDateRange(selectedStartDate.value, selectedEndDate.value)
  
  emit('select', {
    weekId: `${year}-W${weekNum}`,
    weekLabel: `Week ${weekNum}`,
    dateRange: dateRange,
    startDate: selectedStartDate.value.toISOString().split('T')[0],
    endDate: selectedEndDate.value.toISOString().split('T')[0]
  })
}

// Close modal
const close = () => {
  emit('close')
}
</script>

<style scoped>
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
  max-height: 90vh;
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

/* Calendar Navigation */
.calendar-nav {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.nav-btn {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  color: #374151;
  cursor: pointer;
  transition: all 0.2s;
}

.nav-btn:hover:not(:disabled) {
  background: #f3f4f6;
  border-color: #d1d5db;
}

.nav-btn:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}

.current-month {
  font-size: 16px;
  font-weight: 600;
  color: #111827;
}

/* Calendar Table */
.calendar-wrapper {
  margin-bottom: 20px;
}

.calendar-table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0;
}

.calendar-table th {
  padding: 8px 0;
  text-align: center;
  font-size: 12px;
  font-weight: 600;
  color: #6b7280;
  text-transform: uppercase;
}

.calendar-day {
  position: relative;
  padding: 0;
  text-align: center;
  cursor: pointer;
  transition: all 0.2s;
  border: 1px solid transparent;
}

.calendar-day span {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 36px;
  font-size: 14px;
  color: #374151;
}

.calendar-day.other-month span {
  color: #d1d5db;
}

.calendar-day.today span {
  font-weight: 700;
  color: #10b981;
}

.calendar-day.disabled {
  cursor: not-allowed;
}

.calendar-day.disabled span {
  color: #d1d5db;
}

.calendar-day:hover:not(.disabled):not(.other-month) {
  background: #3b82f6;
}

/* Selected Range Styles */
.calendar-day.selected {
  background: #dbeafe;
}

.calendar-day.selected span {
  color: #1d4ed8;
}

.calendar-day.selected-start {
  border-top-left-radius: 8px;
  border-bottom-left-radius: 8px;
  background: #3b82f6;
}

.calendar-day.selected-start span {
  color: white;
  font-weight: 600;
}

.calendar-day.selected-end {
  border-top-right-radius: 8px;
  border-bottom-right-radius: 8px;
  background: #3b82f6;
}

.calendar-day.selected-end span {
  color: white;
  font-weight: 600;
}

/* Selected Period Display */
.selected-period {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px;
  background: #f0f9ff;
  border: 1px solid #bae6fd;
  border-radius: 12px;
  margin-bottom: 20px;
}

.period-icon {
  color: #0284c7;
}

.period-text {
  flex: 1;
}

.period-label {
  font-size: 12px;
  font-weight: 500;
  color: #6b7280;
  margin-bottom: 4px;
}

.period-dates {
  font-size: 14px;
  font-weight: 600;
  color: #0c4a6e;
}

/* Action Buttons */
.action-buttons {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.btn-cancel,
.btn-confirm {
  padding: 10px 20px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  border: none;
}

.btn-cancel {
  background: #f3f4f6;
  color: #374151;
}

.btn-cancel:hover {
  background: #e5e7eb;
}

.btn-confirm {
  background: #3b82f6;
  color: white;
}

.btn-confirm:hover:not(:disabled) {
  background: #2563eb;
}

.btn-confirm:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>