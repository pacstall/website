import { Resource } from 'i18next'
import enLocale from './en.locale'
import roLocale from './ro.locale'

export default interface Locale {
    home: {
        title: string
        subtitle: string
        cards: {
            whyDifferent: {
                title: string
                description: string
            }
            howItWorks: {
                title: string
                description: string
            }
        }
        installationInstructions: string
        showcase: {
            title: string
            packageSearch: string
        }
    }
    navbar: {
        title: string
        contribute: {
            title: string
            workOnFeatures: string
            helpTranslate: string
            becomeAMaintainer: string
        }
        social: {
            title: string
            discord: string
            matrix: string
            reddit: string
            mastodon: string
        }
        browse: {
            title: string
        }
        privacy: {
            title: string
        }
        install: string
    }
}

export const translations = {
    'en-US': {
        translation: enLocale,
    },
    'ro-RO': {
        translation: roLocale,
    },
} as const satisfies Resource

export const localeFlags: Record<keyof typeof translations, string> = {
    "en-US": "US 🇺🇸",
    "ro-RO": "RO 🇷🇴",
};

export const locales = Object.keys(translations) as (keyof typeof translations)[]
export const flags = Object.values(localeFlags)
export const localeEntries = Object.entries(localeFlags) as unknown as [keyof typeof translations, string][]
