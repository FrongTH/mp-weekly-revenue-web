import { ref, computed } from 'vue'

// Language state
const currentLanguage = ref(localStorage.getItem('language') || 'en')

// Language translations
const translations = {
  en: {
    // Navigation
    home: 'Home',
    tracking: 'Tracking',
    manage: 'Manage',

    // Common
    create: 'Create',
    edit: 'Edit',
    delete: 'Delete',
    cancel: 'Cancel',
    save: 'Save',
    loading: 'Loading...',
    search: 'Search',
    filter: 'Filter',
    sort: 'Sort',

    // Dashboard/HomePage
    welcomeBack: 'Welcome back!',
    totalRevenue: 'Total Revenue',
    todayDate: 'Today',
    merchantDetails: 'Merchant Details',
    createMerchant: 'Create Merchant',
    noMerchantsYet: 'No merchants yet',
    startByCreating: 'Start by creating your first merchant to track revenue',
    gettingStartedGuide: 'Getting Started Guide',
    merchant: 'Merchant',

    // Guide Steps
    step1Title: 'Create a Merchant',
    step1Desc: 'Add your restaurant or food business',
    step2Title: 'Add Menu Items',
    step2Desc: 'Set up your products with costs and prices',
    step3Title: 'Add on Extra Items',
    step3Desc: 'Set up your extra items with costs and prices',
    step4Title: 'Track Orders',
    step4Desc: 'Monitor sales and calculate revenue automatically',

    // Financial Terms
    income: 'Income',
    outcome: 'Outcome',
    revenue: 'Revenue',
    totalIncome: 'Total Income',
    totalOutcome: 'Total Outcome',
    netRevenue: 'Net Revenue',

    // Daily Tracking
    createDailyTracking: 'Create Daily Tracking',
    addIncome: 'Add Income',
    addExpense: 'Add Expense',
    incomeTransactions: 'Income Transactions',
    expenseTransactions: 'Expense Transactions',
    noIncomeRecorded: 'No income recorded yet',
    noExpensesRecorded: 'No expenses recorded yet',
    addFirstIncome: 'Add your first income',
    addFirstExpense: 'Add your first expense',

    // Modal/Form
    description: 'Description',
    listing: 'Listing',
    amount: 'Amount',
    category: 'Category',
    merchantName: 'Merchant Name',
    enterMerchantName: 'Enter merchant name',
    enterExpenseListing: 'Enter expense listing',

    // Weekly View
    daily: 'Daily',
    weekly: 'Weekly',
    summaryOnly: 'Summary Only',
    weeklyViewWarning:
      'Weekly view shows summary data only.\n\nTo view and edit individual tracking details, please switch to Daily view.\n\nWould you like to switch to Daily view now?',

    // Categories
    ingredients: 'Ingredients',
    supplies: 'Supplies',
    deliveryCost: 'Delivery Cost',
    utilities: 'Utilities',
    rent: 'Rent',
    salary: 'Salary',
    other: 'Other',

    // Manage Merchant
    menuItems: 'Menu Items',
    extraItems: 'Extra Items',
    generalPrice: 'General Price',
    deliveryPrice: 'Delivery Price',

    // Quick Select
    quickSelect: 'Quick Select',
    chooseFromPrevious: 'Choose from your previous expense names',
    recentExpenses: 'Recent Expenses',
    quicklyAddFromRecent: 'Quickly add from your recent expenses',

    // Error Messages
    failedToDelete: 'Failed to delete',
    failedToSave: 'Failed to save',
    pleaseEnterMerchantName: 'Please enter a merchant name',
    networkError: 'Network error. Please check your connection and try again.',

    // Delete Confirmation
    deleteMerchant: 'Delete Merchant',
    areYouSureDelete: 'Are you sure you want to delete this merchant?',
    actionCannotBeUndone: 'This action cannot be undone.',
    deleting: 'Deleting...',
    inStoreOrPickup: 'In-store or pickup',

    // Language Toggle
    language: 'Language',
    thai: 'ไทย',
    english: 'English',

    // ManageMerchantPage
    back: 'Back',
    manageMerchant: 'Manage Merchant',
    mainMenu: 'Main Menu',
    extraOnTop: 'Extra on Top',
    mainMenuItems: 'Main Menu Items',
    addMenuItem: 'Add Menu Item',
    noMenuItemsYet: 'No menu items yet',
    startByAddingFirstMenuItem: 'Start by adding your first menu item',
    extraOnTopItems: 'Extra on Top Items',
    addExtraItem: 'Add Extra Item',
    noExtraItemsYet: 'No extra items yet',
    addExtraToppingsOrAddOns: 'Add extra toppings or add-ons for your menu',
    general: 'General',
    delivery: 'Delivery',
    addMainMenuItem: 'Add Main Menu Item',
    menuItemName: 'Menu Item Name',
    enterMenuItemName: 'Enter menu item name',
    generalSalePrice: 'General Sale Price',
    deliverySalePrice: 'Delivery Sale Price',
    creating: 'Creating...',
    addItem: 'Add Item',
    extraItemName: 'Extra Item Name',
    enterExtraItemName: 'Enter extra item name',
    editItem: 'Edit Item',
    itemName: 'Item Name',
    enterItemName: 'Enter item name',
    saving: 'Saving...',
    ok: 'OK',
    restaurant: 'Restaurant',
    addQuickIncome: 'Add Quick Income',
    enterIncomeDescription: 'Enter income description',
    continue: 'Continue',
    quantity: 'Quantity',

    // DailyTrackingPage
    generalSale: 'General Sale',
    deliverySale: 'Delivery Sale',
    deliveryOrders: 'Delivery orders',
    typeToSearchMenuItems: 'Type to search menu items...',
    addExtraItems: 'Add Extra Items',
    searchExtraItemsOptional: 'Search Extra Items (Optional)',
    typeToSearchExtraItems: 'Type to search extra items...',
    searchMenuItem: 'Search Menu Item',
    addSalesIncome: 'Add Sales Income',
    selectPriceType: 'Select Price Type', // IGNORE
    choosePriceType: 'Please choose the price type for this sale: General or Delivery.',
    general: 'General',
    delivery: 'Delivery',
    addExpense: 'Add Expense', // IGNORE
    quickSelect: 'Quick Select',
    chooseFromPrevious: 'Choose from your previous expense names',
    week: 'Week',
    of: 'of',
    selected: 'Selected',
    selectDateForDailyTracking: 'Select Date for Daily Tracking',
    createTracking: 'Create Tracking',
    noDailyTracking: 'No daily tracking yet',
    createDailyTrackingFirst:
      'Create your first daily tracking to start monitoring income and expenses',
    createDailyTracking: 'Create Daily Tracking', // IGNORE
    incomeDescription: 'Income Description',
    orderSummary: 'Order Summary',
    totalAmount: 'Total Amount',
  },
  th: {
    orderSummary: 'สรุปคำสั่งซื้อ',
    noDailyTracking: 'ยังไม่มีการติดตามรายวัน',
    createDailyTrackingFirst: 'สร้างการติดตามรายวันครั้งแรกของคุณเพื่อเริ่มติดตามรายรับและรายจ่าย',
    merchant: 'ร้านค้า',
    selectDateForDailyTracking: 'เลือกวันที่สำหรับการติดตามรายวัน',
    selected: 'เลือก',
    createTracking: 'สร้าง',
    of: 'ของปี',
    choosePriceType: 'กรุณาเลือกประเภทราคาสำหรับการขายนี้: ทั่วไป หรือ เดลิเวอรี่.', // IGNORE
    // Navigation
    home: 'หน้าหลัก',
    tracking: 'การติดตาม',
    manage: 'จัดการ',
    restaurant: 'ร้านอาหาร',
    incomeDescription: 'ชื่อรายการ',
    enterIncomeDescription: 'กรอกชื่อรายการ',
    continue: 'เพิ่มรายการ',
    price: 'ราคา',

    // Common
    create: 'สร้าง',
    edit: 'แก้ไข',
    delete: 'ลบ',
    cancel: 'ยกเลิก',
    save: 'บันทึก',
    loading: 'กำลังโหลด...',
    search: 'ค้นหา',
    filter: 'กรอง',
    sort: 'เรียง',
    inStoreOrPickup: 'ขายด้วยตนเอง',
    addQuickIncome: 'เพิ่มรายรับอื่นๆ',

    // Dashboard/HomePage
    welcomeBack: 'ยินดีต้อนรับกลับมา!',
    totalRevenue: 'รายได้รวม',
    todayDate: 'วันนี้',
    merchantDetails: 'รายละเอียดร้านค้า',
    createMerchant: 'สร้างร้านค้า',
    noMerchantsYet: 'ยังไม่มีร้านค้า',
    startByCreating: 'เริ่มต้นด้วยการสร้างร้านค้าแรกเพื่อติดตามรายได้',
    gettingStartedGuide: 'คู่มือการเริ่มต้น',

    // Guide Steps
    step1Title: 'สร้างร้านค้า',
    step1Desc: 'เพิ่มร้านอาหารหรือธุรกิจอาหารของคุณ',
    step2Title: 'เพิ่มรายการเมนู',
    step2Desc: 'ตั้งค่าสินค้าพร้อมต้นทุนและราคา',
    step3Title: 'เพิ่มรายการเสริม',
    step3Desc: 'ตั้งค่ารายการเสริมพร้อมต้นทุนและราคา',
    step4Title: 'ติดตามคำสั่งซื้อ',
    step4Desc: 'ตรวจสอบยอดขายและคำนวณรายได้อัตโนมัติ',

    // Financial Terms
    income: 'รายรับ',
    outcome: 'รายจ่าย',
    revenue: 'กำไร',
    totalIncome: 'รายรับรวม',
    totalOutcome: 'รายจ่ายรวม',
    netRevenue: 'กำไรสุทธิ',

    // Daily Tracking
    createDailyTracking: 'สร้าง',
    addIncome: 'เพิ่มรายรับ',
    addExpense: 'เพิ่มรายจ่าย',
    Income: 'รายรับ',
    Outcome: 'รายจ่าย',
    incomeTransactions: 'รายการรายรับ',
    expenseTransactions: 'รายการรายจ่าย',
    noIncomeRecorded: 'ยังไม่มีการบันทึกรายรับ',
    noExpensesRecorded: 'ยังไม่มีการบันทึกรายจ่าย',
    addFirstIncome: 'เพิ่มรายรับแรกของคุณ',
    addFirstExpense: 'เพิ่มรายจ่ายแรกของคุณ',
    totalAmount: 'จำนวนเงินรวม',

    // Modal/Form
    description: 'คำอธิบาย',
    listing: 'รายการ',
    amount: 'จำนวนเงิน',
    category: 'หมวดหมู่',
    merchantName: 'ชื่อร้านค้า',
    enterMerchantName: 'กรอกชื่อร้านค้า',
    enterExpenseListing: 'กรอกรายการรายจ่าย',

    // Weekly View
    daily: 'รายวัน',
    week: 'สัปดาห์ที่',
    weekly: 'รายสัปดาห์',
    summaryOnly: 'สรุปเท่านั้น',
    weeklyViewWarning:
      'มุมมองรายสัปดาห์แสดงข้อมูลสรุปเท่านั้น\n\nเพื่อดูและแก้ไขรายละเอียดการติดตาม กรุณาเปลี่ยนเป็นมุมมองรายวัน\n\nคุณต้องการเปลี่ยนเป็นมุมมองรายวันหรือไม่?',

    // Categories
    ingredients: 'วัตถุดิบ',
    supplies: 'อุปกรณ์',
    deliveryCost: 'ค่าขนส่ง',
    utilities: 'สาธารณูปโภค',
    rent: 'ค่าเช่า',
    salary: 'เงินเดือน',
    other: 'อื่นๆ',

    // Manage Merchant
    menuItems: 'รายการเมนู',
    extraItems: 'รายการเสริม',
    generalPrice: 'ราคาทั่วไป',
    deliveryPrice: 'ราคาส่ง',

    // Quick Select
    quickSelect: 'รายการจ่ายเมื่อเร็วๆนี้',
    chooseFromPrevious: 'เลือกจากชื่อรายจ่ายก่อนหน้าของคุณ',
    recentExpenses: 'รายจ่ายล่าสุด',
    quicklyAddFromRecent: 'เพิ่มอย่างรวดเร็วจากรายจ่ายล่าสุดของคุณ',

    // Error Messages
    failedToDelete: 'ลบไม่สำเร็จ',
    failedToSave: 'บันทึกไม่สำเร็จ',
    pleaseEnterMerchantName: 'กรุณากรอกชื่อร้านค้า',
    networkError: 'เกิดข้อผิดพลาดเครือข่าย กรุณาตรวจสอบการเชื่อมต่อและลองอีกครั้ง',

    // Delete Confirmation
    deleteMerchant: 'ลบร้านค้า',
    areYouSureDelete: 'คุณแน่ใจหรือไม่ว่าต้องการลบร้านค้านี้?',
    actionCannotBeUndone: 'การกระทำนี้ไม่สามารถยกเลิกได้',
    deleting: 'กำลังลบ...',

    // Language Toggle
    language: 'ภาษา',
    thai: 'ไทย',
    english: 'English',

    // ManageMerchantPage
    back: 'กลับ',
    manageMerchant: 'จัดการร้านค้า',
    mainMenu: 'เมนูหลัก',
    extraOnTop: 'รายการเสริม',
    mainMenuItems: 'รายการเมนูหลัก',
    addMenuItem: 'เพิ่มรายการเมนู',
    noMenuItemsYet: 'ยังไม่มีรายการเมนู',
    startByAddingFirstMenuItem: 'เริ่มต้นด้วยการเพิ่มรายการเมนูแรกของคุณ',
    extraOnTopItems: 'รายการเสริม',
    addExtraItem: 'เพิ่มรายการเสริม',
    noExtraItemsYet: 'ยังไม่มีรายการเสริม',
    addExtraToppingsOrAddOns: 'เพิ่มท็อปปิ้งหรือรายการเสริมสำหรับเมนูของคุณ',
    general: 'ทั่วไป',
    delivery: 'เดลิเวอรี่',
    addMainMenuItem: 'เพิ่มรายการเมนูหลัก',
    menuItemName: 'ชื่อรายการเมนู',
    enterMenuItemName: 'กรอกชื่อรายการเมนู',
    generalSalePrice: 'ราคาขายทั่วไป',
    deliverySalePrice: 'ราคาขายเดลิเวอรี่',
    creating: 'กำลังสร้าง...',
    addItem: 'เพิ่มรายการ',
    extraItemName: 'ชื่อรายการเสริม',
    enterExtraItemName: 'กรอกชื่อรายการเสริม',
    editItem: 'แก้ไขรายการ',
    itemName: 'ชื่อรายการ',
    enterItemName: 'กรอกชื่อรายการ',
    saving: 'กำลังบันทึก...',
    ok: 'ตกลง',
    quantity: 'จำนวน',

    // DailyTrackingPage
    generalSale: 'ขายทั่วไป',
    deliverySale: 'ขายเดลิเวอรี่',
    deliveryOrders: 'คำสั่งซื้อเดลิเวอรี่',
    typeToSearchMenuItems: 'พิมพ์เพื่อค้นหารายการเมนู...',
    addExtraItems: 'เพิ่มรายการเสริม',
    searchExtraItemsOptional: 'ค้นหารายการเสริม',
    typeToSearchExtraItems: 'พิมพ์เพื่อค้นหารายการเสริม...',
    searchMenuItem: 'ค้นหารายการเมนู',
    addSalesIncome: 'เพิ่มรายรับจากการขาย',
    selectPriceType: 'เลือกประเภทราคา', // IGNORE
  },
}

// Composable
export function useLanguage() {
  // Set language
  const setLanguage = (lang) => {
    currentLanguage.value = lang
    localStorage.setItem('language', lang)
  }

  // Get translation
  const t = (key) => {
    return translations[currentLanguage.value]?.[key] || translations.en[key] || key
  }

  // Format currency based on language
  const formatCurrency = (amount) => {
    const currency = currentLanguage.value === 'th' ? 'THB' : 'THB'
    const locale = currentLanguage.value === 'th' ? 'th-TH' : 'en-US'

    return new Intl.NumberFormat(locale, {
      style: 'currency',
      currency: currency,
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(amount || 0)
  }

  // Format date based on language
  const formatDate = (date, options = {}) => {
    const locale = currentLanguage.value === 'th' ? 'th-TH' : 'en-US'
    return new Date(date).toLocaleDateString(locale, options)
  }

  return {
    currentLanguage: computed(() => currentLanguage.value),
    setLanguage,
    t,
    formatCurrency,
    formatDate,
  }
}
