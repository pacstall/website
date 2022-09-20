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
import CookieBanner from './components/CookieBanner'
import Navigation from './components/Navigation'

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
    <Text
        position='fixed'
        right='15px'
        bottom='15px'
        color='gray.500'
        fontSize='md'
    >
        {serverConfig.version}
    </Text>
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
