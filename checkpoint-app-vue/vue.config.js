module.exports = {
    baseUrl: './',
    productionSourceMap: false,
    devServer: {
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:1001',
                changeOrigin: true,
                cookieDomainRewrite: "localhost"
            },
        }
    },
}