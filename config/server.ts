const serverConfig = {
    host: process.env.NODE_ENV === 'production' ? 'https://pacstall.dev' : 'http://localhost:3000'
}

export default serverConfig