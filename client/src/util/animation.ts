import browser from './browser'

export const PREFERS_REDUCED_MOTION = window.matchMedia(
    '(prefers-reduced-motion: reduce)',
).matches
export const ANIMATIONS_DISABLED = browser.isFirefox || PREFERS_REDUCED_MOTION
