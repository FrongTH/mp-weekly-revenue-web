<template>
  <div class="top-bar flex items-center gap-4">
    <!-- Language Toggle -->
    <div class="language-toggle">
      <button @click="toggleDropdown" class="language-button">
        <span class="current-lang">{{ currentLanguage === 'th' ? 'TH' : 'EN' }}</span>
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="chevron">
          <path d="M6 9l6 6 6-6"/>
        </svg>
      </button>
      <!-- Dropdown positioned within language-toggle container -->
      <div v-if="showDropdown" class="language-dropdown">
        <button 
          @click="selectLanguage('en')" 
          class="language-option"
          :class="{ active: currentLanguage === 'en' }"
        >
          <span>{{ t('english') }}</span>
          <span class="lang-code">EN</span>
        </button>
        <button 
          @click="selectLanguage('th')" 
          class="language-option"
          :class="{ active: currentLanguage === 'th' }"
        >
          <span>{{ t('thai') }}</span>
          <span class="lang-code">TH</span>
        </button>
      </div>
    </div>
    
    <!-- Buy Me A Coffee -->
    <a 
      href="https://www.buymeacoffee.com/rattapon.san" 
      target="_blank"
      class="coffee-button transition-transform hover:scale-105 block"
    >
      <img 
        src="https://cdn.buymeacoffee.com/buttons/v2/default-green.png" 
        alt="Buy Me A Coffee" 
        class="rounded-lg shadow-lg coffee-image"
      >
    </a>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useLanguage } from '@/composables/useLanguage'

const { currentLanguage, setLanguage, t } = useLanguage()
const showDropdown = ref(false)

const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value
}

const selectLanguage = (lang) => {
  setLanguage(lang)
  showDropdown.value = false
}

// Close dropdown when clicking outside
const handleClickOutside = (event) => {
  const languageToggle = event.target.closest('.language-toggle')
  const coffeeButton = event.target.closest('.coffee-button')
  
  // Close dropdown if clicking outside language toggle but not on coffee button
  if (!languageToggle && !coffeeButton) {
    showDropdown.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.top-bar {
  display: flex;
  align-items: center;
  gap: 12px; /* spacing between language toggle and button */
}

.language-toggle {
  position: relative;
  display: inline-block;
}

.language-button {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: white;
  border-radius: 8px;
  color: #374151;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.language-button:hover {
  background: #f9fafb;
  border-color: #d1d5db;
}

.current-lang {
  font-weight: 600;
  color: #4f46e5;
}

.chevron {
  transition: transform 0.2s;
}

.language-button:hover .chevron {
  transform: rotate(180deg);
}

.language-dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 4px;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  z-index: 9999;
  min-width: 160px;
  padding: 4px;
  pointer-events: auto;
}

.language-option {
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

.language-option:hover {
  background: #f3f4f6;
}

.language-option.active {
  background: #eef2ff;
  color: #4f46e5;
  font-weight: 600;
}

.flag {
  font-size: 16px;
}

.lang-code {
  margin-left: auto;
  font-size: 12px;
  color: #6b7280;
  font-weight: 500;
}

.language-option.active .lang-code {
  color: #4f46e5;
}

.language-dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 4px;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  z-index: 9999;
  min-width: 160px;
  padding: 4px;
  pointer-events: auto;
}

.coffee-button {
  pointer-events: auto;
  z-index: 1;
  display: block;
  text-decoration: none;
}

.coffee-image {
  height: 36px;
  width: 140px;
  pointer-events: auto;
  display: block;
}

/* Mobile responsive */
@media (max-width: 768px) {
  .language-button {
    padding: 6px 10px;
    font-size: 13px;
  }
  
  .language-dropdown {
    left: 0;
    right: auto;
    min-width: 140px;
  }
  
  .coffee-image {
    height: 32px;
    width: 120px;
  }
}
</style>