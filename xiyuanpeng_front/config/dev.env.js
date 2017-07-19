var merge = require('webpack-merge')
var prodEnv = require('./prod.env.js')
var http = require('./http-common')


module.exports = merge(prodEnv, {
  BASE_URL: 'http://localhost:8000',
  NODE_ENV: '"development"',
})
