const {override} = require('customize-cra');
const cspHtmlWebpackPlugin = require("csp-html-webpack-plugin");

const cspConfigPolicy = {
    'default-src':["'self'", "data:", "blob:", "filesystem:", "https://ikasmansituraja.org/", "http://ikasmansituraja.org/"],
    'img-src': ["*", "data:", "blob:", "filesystem:" /*, "'unsafe-inline'", "'unsafe-eval'"*/ ],
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
