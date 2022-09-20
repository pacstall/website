import ReactDOM from 'react-dom'

import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { FC } from 'react'
import { QueryParamProvider } from 'use-query-params'
import axios from 'axios'
import { setupCache } from 'axios-cache-adapter'
import { RecoilRoot } from 'recoil'
import {
    ChakraProvider,
    extendTheme,
    localStorageManager,
    StylesProvider,
    Text,
    useColorModeValue,
} from '@chakra-ui/react'

import serverConfig from './config/server'
import Packages from './pages/Packages'
import PackageDetails from './pages/PackageDetails'
import LoadingOverlay from './pages/LoadingOverlay'
import Pacstore from './pages/Pacstore'
import ScrollbarStyles from './components/ScrollbarStyles'
import { registerUpdateChecker } from './services/update-checker'

axios.defaults.adapter = setupCache({
    clearOnError: true,
    clearOnStale: true,
    maxAge: 1000 * 5 * 60,
}).adapter

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
registerUpdateChecker()
ReactDOM.render(
    <>
        <ChakraProvider theme={theme} colorModeManager={localStorageManager}>
            <StylesProvider value={{}}>
                <ScrollbarStyles />
                <RecoilRoot>
                    <QueryParamProvider>
                        <BrowserRouter>
                            <Routes>
                                <Route index element={<Packages />} />
                                <Route
                                    path='/packages'
                                    element={<Pacstore />}
                                />
                                <Route
                                    path='/packages/:name'
                                    element={<PackageDetails />}
                                />
                                <Route
                                    path='*'
                                    element={
                                        <LoadingOverlay
                                            timeout={1500}
                                            path='/packages'
                                        />
                                    }
                                />
                            </Routes>
                            <Footer />
                        </BrowserRouter>
                    </QueryParamProvider>
                </RecoilRoot>
            </StylesProvider>
        </ChakraProvider>
    </>,
    app,
)
