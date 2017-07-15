import Vue from 'vue'
import Vuex from 'vuex'
import * as actions from './actions'
// import * as getters from './getters'
// import mutations from './mutations'
// import * as defaults from './defaults'
import jwtDecode from 'jwt-decode'
import createLogger from 'vuex/dist/logger'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

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
    login (state, loginInfo) {
      state.user = loginInfo.user
      state.isLogin = loginInfo.isLogin
    },
    check (state, payload) {
      let jwtToken = payload.jwtToken
      console.log(jwtDecode(jwtToken))
      if ((jwtToken !== null) && (jwtToken !== undefined)) {
        try {
          let jwtInfo = jwtDecode(jwtToken)
          if (jwtInfo.exp > Math.floor(Date.now() / 1000)) {
            state.isLogin = true
            state.user = jwtInfo.user
          } else {
            console.log('expired')
            state.isLogin = false
          }
        } catch (e) {
          console.log(e)
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

