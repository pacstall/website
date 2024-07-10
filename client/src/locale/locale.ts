import { Resource } from 'i18next'
import bnLocale from './bn-IN.locale'
import angLocale from './en-ANG.locale'
import enLocale from './en-US.locale'
import roLocale from './ro-RO.locale'
import esLocale from './es-ES.locale'
import ptbrLocale from './pt-BR.locale'
import ptptLocale from './pt-PT.locale'
import plLocale from './pl-PL.locale'
import svLocale from './sv-SE.locale'
import itLocale from './it-IT.locale'
import frLocale from './fr-FR.locale'
import deLocale from './de-DE.locale'
import nlLocale from './nl-NL.locale'
import idLocale from './id-ID.locale'
import trLocale from './tr-TR.locale'
import { NumericDisplayHandler } from '../hooks/useNumericDisplay'
import { arabic } from './numeric-systems/arabic'
import { bengali } from './numeric-systems/bengali'

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
    cookieConsent: {
        title: string
        paragraphs: [string, string, string, string]
        accept: string
    }
    packageSearch: {
        dropdown: {
            package: string
            packageTooltip: string
            maintainer: string
            maintainerTooltip: string
        }
        orphaned: string
        table: {
            name: string
            maintainer: string
            version: string
            install: string
        }
        versionTooltip: {
            notInRegistry: string
            latest: string
            patch: string
            minor: string
            major: string
            isGit: string
        }
        noResults: string
        search: string
        maintainerTooltip: {
            maintainedBy: string
            noMaintainer: string
        }
        pagination: {
            previous: string
            next: string
        }
    }
    packageDetails: {
        table: {
            name: string
            version: string
            maintainer: string
            dependencies: string
            requiredBy: string
            lastUpdatedAt: string
        }
        noResults: string
        orphaned: string
        view: string
        dependenciesModal: {
            title: string
            buildDependencies: string
            optionalDependencies: string
            runtimeDependencies: string
            pacstallDependencies: string
            name: string
            version: string
            close: string
            provider: string
            noDescription: string
        }
        requiredByModal: {
            title: string
            name: string
            provider: string
            close: string
            noDescription: string
        }
        openInGithub: string
        howToInstall: {
            title: string
            step1: string
            step2: string
        }
    }
}

export const translations = {
    'bn-IN': {
        translation: bnLocale,
    },
    'en-ANG': {
        translation: angLocale,
    },
    'en-US': {
        translation: enLocale,
    },
    'ro-RO': {
        translation: roLocale,
    },
    'es-ES': {
        translation: esLocale,
    },
    'pt-BR': {
        translation: ptbrLocale,
    },
    'pt-PT': {
        translation: ptptLocale,
    },
    'pl-PL': {
        translation: plLocale,
    },
    'sv-SE': {
        translation: svLocale,
    },
    'it-IT': {
        translation: itLocale,
    },
    'nl-NL': {
        translation: nlLocale,
    },
    'fr-FR': {
        translation: frLocale,
    },
    'de-DE': {
        translation: deLocale,
    },
    'id-ID': {
        translation: idLocale,
    },
    'tr-TR': {
        translation: trLocale,
    },
} as const satisfies Resource

export const localeNumericDisplay = {
    'en-US': arabic,
    'en-AN': arabic,
    'ro-RO': arabic,
    'es-ES': arabic,
    'pt-BR': arabic,
    'pt-PT': arabic,
    'pl-PL': arabic,
    'sv-SE': arabic,
    'it-IT': arabic,
    'nl-NL': arabic,
    'fr-FR': arabic,
    'de-DE': arabic,
    'id-ID': arabic,
    'tr-TR': arabic,
    bn_IN: bengali,
} satisfies Record<keyof typeof translations, NumericDisplayHandler>

export const localeFlags: Record<keyof typeof translations, string> = {
    'en-US': 'US 🇺🇸',
    'en-ANG': 'ᚩᛝ 🏴󠁧󠁢󠁥󠁮󠁧󠁿',
    'bn-IN': 'BN 🇮🇳',
    'de-DE': 'DE 🇩🇪',
    'es-ES': 'ES 🇪🇸',
    'fr-FR': 'FR 🇫🇷',
    'id-ID': 'ID 🇮🇩',
    'it-IT': 'IT 🇮🇹',
    'nl-NL': 'NL 🇳🇱',
    'pl-PL': 'PL 🇵🇱',
    'pt-BR': 'PT 🇧🇷',
    'pt-PT': 'PT 🇵🇹',
    'ro-RO': 'RO 🇷🇴',
    'sv-SE': 'SV 🇸🇪',
    'tr-TR': 'TR 🇹🇷',
}

export const locales = Object.keys(
    translations,
) as (keyof typeof translations)[]
export const flags = Object.values(localeFlags)
export const localeEntries = Object.entries(localeFlags) as unknown as [
    keyof typeof translations,
    string,
][]
