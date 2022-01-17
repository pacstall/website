import ReactDOM from "react-dom";

import '../public/styles/dracula.css'
import '../public/styles/globals.css'
import '../public/styles/tailwind.pcss'
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { lazy, Suspense } from "react";
import Home from "./pages/Home";
import Showcase from "./pages/Showcase";

const Packages = lazy(() => import('./pages/Packages'))

const app = document.getElementById("app");
ReactDOM.render(<>
    <BrowserRouter>
        <Suspense fallback={<span>loading</span>}>
            <Routes>
                <Route index element={<Home />} />
                <Route path="/packages" element={<Packages />} />
                <Route path="/showcase" element={<Showcase />} />
            </Routes>
        </Suspense>

    </BrowserRouter>
</>, app);