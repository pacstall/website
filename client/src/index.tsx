import { createRoot } from 'react-dom/client'

import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { FC, lazy, Suspense } from 'react'
import { QueryParamProvider } from 'use-query-params'
import { ReactRouter6Adapter } from 'use-query-params/adapters/react-router-6'
import { RecoilRoot } from 'recoil'
import {
    ChakraProvider,
    createStylesContext,
    extendTheme,
    localStorageManager,
    Text,
} from '@chakra-ui/react'

import serverConfig from './config/server'
import Navigation from './components/Navigation'
import CookieBanner from './components/CookieBanner'
import i18n from 'i18next'
import { initReactI18next } from 'react-i18next'
import detector from 'i18next-browser-languagedetector'
import { translations } from './locale/locale'

i18n.use(detector)
    .use(initReactI18next) // passes i18n down to react-i18next
    .init({
        // the translations
        // (tip move them in a JSON file and import them,
        // or even better, manage them via a UI: https://react.i18next.com/guides/multiple-translation-files#manage-your-translations-with-a-management-gui)
        resources: translations,
        fallbackLng: 'en-US',
        interpolation: {
            escapeValue: false, // react already safes from xss => https://www.i18next.com/translation-function/interpolation#unescape
        },
    })

const Home = lazy(() => import('./pages/Home'))
const NotFound = lazy(() => import('./pages/NotFound'))
const Packages = lazy(() => import('./pages/Packages'))
const PackageDetails = lazy(() => import('./pages/PackageDetails'))
const PrivacyPolicy = lazy(() => import('./pages/PrivacyPolicy'))

const theme = extendTheme({
    config: {
        initialColorMode: 'system',
        useSystemColorMode: false,
    },
    fonts: {
        heading: 'Open Sans, sans-serif',
        body: 'Titillium Web, sans-serif',
    },
    colors: {
        brand: {
            50: '#E6FFFA',
            100: '#B2F5EA',
            200: '#81E6D9',
            300: '#4FD1C5',
            400: '#38B2AC',
            500: '#319795',
            600: '#2C7A7B',
            700: '#285E61',
            800: '#234E52',
            900: '#1D4044',
        },
    },
})

const Footer: FC = () => (
    <>
        <div style={{ marginTop: '50px' }}></div>
        <Text
            position='fixed'
            right='15px'
            bottom='15px'
            padding='25px'
            color='gray.300'
            fontWeight='600'
            fontSize='md'
            opacity='0%'
            transition='all 0.3s ease-in-out'
            _hover={{
                opacity: '100%',
            }}
        >
            {serverConfig.version}
        </Text>
    </>
)

const app = document.getElementById('app')
const root = createRoot(app)

const App = () => {
    const [StylesProvider, useStyles] = createStylesContext('App')

    return (
        <>
            <ChakraProvider
                theme={theme}
                colorModeManager={localStorageManager}
            >
                <StylesProvider value={{}}>
                    <RecoilRoot>
                        <BrowserRouter>
                            <QueryParamProvider adapter={ReactRouter6Adapter}>
                                <Navigation />
                                <Suspense fallback={<></>}>
                                    <Routes>
                                        <Route index element={<Home />} />
                                        <Route
                                            path='/packages'
                                            element={<Packages />}
                                        />
                                        <Route
                                            path='/packages/:name'
                                            element={<PackageDetails />}
                                        />
                                        <Route
                                            path='/privacy'
                                            element={<PrivacyPolicy />}
                                        />
                                        <Route
                                            path='*'
                                            element={<NotFound />}
                                        />
                                    </Routes>
                                </Suspense>
                                <Footer />
                                <CookieBanner />
                            </QueryParamProvider>
                        </BrowserRouter>
                    </RecoilRoot>
                </StylesProvider>
            </ChakraProvider>
        </>
    )
}

root.render(<App />)
