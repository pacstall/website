const isProd = location.port !== '1234'
const serverConfig = {
    isProduction: isProd,
    host: isProd ? '' : 'http://localhost:3000'
}

console.log(isProd, process.env)

export default serverConfig