export function isMobileDevice() {
  const userAgent = navigator.userAgent.toLowerCase()
  const isMobile = /android|webos|iphone|ipad|ipod|blackberry|iemobile|opera mini/i.test(userAgent)
  
  // Also check screen width as fallback
  const isSmallScreen = window.innerWidth <= 768
  
  // Check for touch support
  const isTouchDevice = 'ontouchstart' in window || navigator.maxTouchPoints > 0
  
  return isMobile || (isSmallScreen && isTouchDevice)
}

export function getDeviceType() {
  const userAgent = navigator.userAgent.toLowerCase()
  
  if (/android/i.test(userAgent)) {
    return 'Android'
  } else if (/iphone|ipod/i.test(userAgent)) {
    return 'iPhone'
  } else if (/ipad/i.test(userAgent)) {
    return 'iPad'
  } else if (/blackberry/i.test(userAgent)) {
    return 'BlackBerry'
  } else if (/webos/i.test(userAgent)) {
    return 'WebOS'
  } else if (/iemobile/i.test(userAgent)) {
    return 'Windows Phone'
  } else {
    return 'Desktop'
  }
}