import ReactDOM from "react-dom";

import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import { lazy, Suspense } from "react";
import Home from "./pages/Home";
import Showcase from "./pages/Showcase";
import NotFound from "./pages/NotFound";
import { QueryParamProvider } from "use-query-params";
import axios from "axios";
import { setupCache } from "axios-cache-adapter";
import { RecoilRoot } from "recoil";
import { Box, ChakraProvider, extendTheme, localStorageManager, Spinner } from '@chakra-ui/react'

import '@fontsource/raleway/400.css'
import '@fontsource/open-sans/700.css'

axios.defaults.adapter = setupCache({
    clearOnError: true,
    clearOnStale: true,
    maxAge: 1000 * 5 * 60
}).adapter

const Packages = lazy(() => import('./pages/Packages'))
const PackageDetails = lazy(() => import('./pages/PackageDetails'))

const theme = extendTheme({
    config: {
        initialColorMode: 'system',
        useSystemColorMode: false
    },
    fonts: {
        heading: 'Open Sans, sans-serif',
        body: 'Raleway, sans-serif',
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
        }
    }
})

const app = document.getElementById("app");
ReactDOM.render(<>
    <ChakraProvider theme={theme} colorModeManager={localStorageManager}>
        <RecoilRoot>
            <QueryParamProvider>

                <BrowserRouter>
                    <Suspense fallback={<Box textAlign='center' mt='20vh'><Spinner size='lg' /></Box>}>
                        <Routes>
                            <Route index element={<Home />} />
                            <Route path="/packages" element={<Packages />} />
                            <Route path="/packages/:name" element={<PackageDetails />} />
                            <Route path="/showcase" element={<Showcase />} />
                            <Route path="/not-found" element={<NotFound />} />
                            <Route path="*" element={<Navigate to="/not-found" />} />
                        </Routes>
                    </Suspense>
                </BrowserRouter>
            </QueryParamProvider>
        </RecoilRoot>
    </ChakraProvider>
</>, app);