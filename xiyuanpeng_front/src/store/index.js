import Vue from 'vue'
import Vuex from 'vuex'
import * as actions from './actions'
import jwtDecode from 'jwt-decode'
import createLogger from 'vuex/dist/logger'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'
const anonUser = {
  created_at: '2017-07-08T14:24:53.091+08:00',
  email: 'anon@xyp.com',
  id: '000000000000000000',
  name: '匿名用户',
  pass: 'anon',
  updated_at: '2017-07-08T14:24:53.091+08:00'
}

export default new Vuex.Store({
  state: {
    user: {
      created_at: '2017-07-08T14:24:53.091+08:00',
      email: 'anon@xyp.com',
      id: '000000000000000000',
      name: '匿名用户',
      pass: 'anon',
      updated_at: '2017-07-08T14:24:53.091+08:00'
    },
    isLogin: false,
    jwtToken: localStorage.getItem('jwtToken')
  },
  mutations: {
    jwtTokenChange (state) {
      state.jwtToken = localStorage.getItem('jwtToken')
    },
    logout (state) {
      state.user = anonUser
      state.isLogin = false
    },
    login (state, loginInfo) {
      state.user = loginInfo.user
      state.isLogin = loginInfo.isLogin
    },
    check (state, payload) {
      let jwtToken = payload.jwtToken
      console.log(jwtToken)
      if (jwtToken) {
        let jwtPayload
        try {
          jwtPayload = jwtDecode(jwtToken)
        } catch (e) {
          console.log('decode wrong')
          return
        }
        if (jwtPayload.exp > Math.floor(Date.now() / 1000)) {
          state.isLogin = true
          state.user = jwtPayload.user
        } else {
          console.log('expired')
          state.isLogin = false
        }
      } else {
        console.log('jwtToken === null')
      }
    }
  },
  actions,
  strict: debug,
  plugins: debug ? [createLogger()] : []
})

