import { useTranslation } from 'react-i18next'
import { localeNumericDisplay } from '../locale/locale'

export type NumericDisplayHandler = (value: number) => string

const useNumericDisplay = (): NumericDisplayHandler => {
    const { i18n } = useTranslation()
    const numericSystem =
        localeNumericDisplay[
            i18n.language as keyof typeof localeNumericDisplay
        ] ?? localeNumericDisplay['en-US']

    return numericSystem
}

export default useNumericDisplay
