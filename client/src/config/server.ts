const isProd = process.env.NODE_ENV === 'production'
const serverConfig = {
    isProduction: isProd,
    host: isProd ? '' : 'http://localhost:3300',
    version: isProd ? process.env.VERSION || 'unversioned' : 'development',
}

export default serverConfig
;(window as any).config = serverConfig
