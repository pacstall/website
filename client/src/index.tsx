import ReactDOM from "react-dom";

import '../public/styles/dracula.css'
import '../public/styles/globals.css'
import '../public/styles/tailwind.pcss'
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import { lazy, Suspense } from "react";
import Home from "./pages/Home";
import Showcase from "./pages/Showcase";
import NotFound from "./pages/NotFound";
import { QueryParamProvider } from "use-query-params";
import axios from "axios";
import { setupCache } from "axios-cache-adapter";
import Notifications from "./components/Notifications";
import { RecoilRoot } from "recoil";

axios.defaults.adapter = setupCache({
    clearOnError: true,
    clearOnStale: true,
    maxAge: 1000 * 5 * 60
}).adapter

const Packages = lazy(() => import('./pages/Packages'))
const PackageDetails = lazy(() => import('./pages/PackageDetails'))

const app = document.getElementById("app");
ReactDOM.render(<>
    <RecoilRoot>
        <QueryParamProvider>
            <Notifications />

            <BrowserRouter>
                <Suspense fallback={<span>loading</span>}>
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
</>, app);