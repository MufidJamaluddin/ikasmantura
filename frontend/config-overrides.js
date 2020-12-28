const {override} = require('customize-cra');
const cspHtmlWebpackPlugin = require("csp-html-webpack-plugin");

const cspConfigPolicy = {
    'default-src':["'self'", "data:", "blob:", "filesystem:", "'unsafe-inline'", "https://ikasmansituraja.org/", "http://ikasmansituraja.org/"],
    'img-src': ["*", "data:", "blob:", "filesystem:"],
    'script-src': ["'self'", "data:"],
    'font-src': ["'self'", "data:"],
    'style-src': ["'self'", "data:"],
};

function addCspHtmlWebpackPlugin(config) {
    if(process.env.NODE_ENV === 'production') {
        config.plugins.push(new cspHtmlWebpackPlugin(cspConfigPolicy));
    }

    return config;
}

module.exports = {
    webpack: override(addCspHtmlWebpackPlugin),
};
