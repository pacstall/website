const isPacstore = (window as any).$__PACSTALL.isElectron;
const isWeb = !isPacstore

const isProd = process.env.NODE_ENV === 'production'
const serverConfig = {
    isProduction: isProd,
    host: isWeb ? (isProd ? '' : 'http://localhost:3300') : 'https://pacstall.dev',
    version: isProd ? process.env.VERSION || 'unversioned' : 'development',
    isPacstore,
    isWeb
}

export default serverConfig
;(window as any).config = serverConfig
