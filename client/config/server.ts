const isProd = process.env.NODE_ENV === 'production'
const serverConfig = {
    isProduction: isProd,
    host: isProd ? '' : 'http://localhost:3000'
}

export default serverConfig