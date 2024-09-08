const isProd = import.meta.env.PROD
const serverConfig = {
    isProduction: isProd,
    host: isProd ? '' : 'http://localhost:3300',
    version: isProd
        ? import.meta.env.VITE_VERSION || 'unversioned'
        : 'development',
    newPacstallSyntax: false,
}

export default serverConfig
;(window as any).config = serverConfig
