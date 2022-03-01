const isProd = process.env.NODE_ENV === 'production'
const serverConfig = {
    isProduction: isProd,
    host: isProd ? '' : 'http://localhost:3300',
    version: isProd ? process.env.VERSION || 'local-prod' : 'development'
}


export default serverConfig;

(window as any).config = serverConfig;
