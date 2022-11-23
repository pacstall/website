/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VERSION: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
