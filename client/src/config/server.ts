const isProd = location.port !== '1234'
const serverConfig = {
    isProduction: isProd,
    host: isProd ? '' : 'http://localhost:3300'
}

export default serverConfig