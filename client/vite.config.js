import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

/** @type {import('vite').UserConfig} */
const commonConfig = {
    plugins: [react()],
    publicDir: 'public',
    appType: 'spa',
}

/** @type {import('vite').UserConfig} */
const buildConfig = {
    ...commonConfig,
    build: {
        chunkSizeWarningLimit: 750,
        rollupOptions: {
            output: {
                compact: true,
            },
            input: {
                app: './index.production.html',
            },
        },
    },
}

/** @type {import('vite').UserConfig} */
const serveConfig = {
    ...commonConfig,
    build: {
        rollupOptions: {
            input: {
                app: './index.html',
            },
        },
    },
}

export default defineConfig(({ command }) => {
    switch (command) {
        case 'build': {
            return buildConfig
        }

        case 'serve': {
            return serveConfig
        }

        default: {
            throw 'unreachable'
        }
    }
})
