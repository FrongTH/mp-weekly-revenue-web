import { createRouter, createWebHistory } from 'vue-router'
import { isMobileDevice } from '@/utils/deviceDetection'
import WelcomePage from '@/views/WelcomePage.vue'
import SignInPage from '@/views/SignInPage.vue'
import OTPVerificationPage from '@/views/OTPVerificationPage.vue'
import HomePage from '@/views/HomePage.vue'
import DeviceRestriction from '@/views/DeviceRestriction.vue'
import ManageMerchantPage from '@/views/ManageMerchantPage.vue'
import TrackingPage from '@/views/TrackingPage.vue'
import DailyTrackingPage from '@/views/DailyTrackingPage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Welcome',
      component: WelcomePage,
      meta: { requiresMobile: true }
    },
    {
      path: '/signin',
      name: 'SignIn',
      component: SignInPage,
      meta: { requiresMobile: true }
    },
    {
      path: '/verify-otp',
      name: 'OTPVerification',
      component: OTPVerificationPage,
      meta: { requiresMobile: true }
    },
    {
      path: '/home',
      name: 'Home',
      component: HomePage,
      meta: { requiresMobile: true }
    },
    {
      path: '/device-restriction',
      name: 'DeviceRestriction',
      component: DeviceRestriction,
      meta: { requiresMobile: false }
      // meta: { requiresMobile: true }
    },
    {
      path: '/merchant/manage/:merchantId',
      name: 'ManageMerchant',
      component: ManageMerchantPage,
      meta: { requiresMobile: true }
    },
    {
      path: '/merchant/tracking/:merchantId',
      name: 'TrackingPage',
      component: TrackingPage,
      meta: { requiresMobile: true }
    },
    {
      path: '/tracking/daily/:trackingId',
      name: 'DailyTrackingPage',
      component: DailyTrackingPage,
      meta: { requiresMobile: true }
    }
  ],
})

// Navigation guard to check for mobile device
router.beforeEach((to, from, next) => {
  const requiresMobile = to.matched.some(record => record.meta.requiresMobile)
  
  if (requiresMobile && !isMobileDevice()) {
    // Redirect to device restriction page if not mobile
    next({ name: 'DeviceRestriction' })
  } else if (to.name === 'DeviceRestriction' && isMobileDevice()) {
    // Redirect mobile users away from restriction page
    next({ name: 'Welcome' })
  } else {
    next()
  }
})

export default router
