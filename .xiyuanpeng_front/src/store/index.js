import Vue from 'vue'
import Vuex from 'vuex'
import * as actions from './actions'
import jwtDecode from 'jwt-decode'
import createLogger from 'vuex/dist/logger'
import ObjectId from 'bson-objectid'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'
const anonUser = {
  'email': 'anonymous@xyp.com',
  'id': ObjectId(),
  'name': '匿名用户'
}
const anonUserJwtToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzI3NzMwNzEsImlkIjoiNTk3MDU3N2FkNmFlMjUwNWYxODRlYThmIiwib3JpZ19pYXQiOjE1MDE2NjkwNzEsInVzZXIiOnsiaWQiOiI1OTcwNTc3YWQ2YWUyNTA1ZjE4NGVhOGYiLCJuYW1lIjoi5Yy_5ZCN55So5oi3IiwiZW1haWwiOiJhbm9ueW1vdXNAeHlwLmNvbSJ9fQ.AwGsNWxzl6kN8yUXNhvUS1sZpJUOzOEyCop_UZX4AHs'

export default new Vuex.Store({
  // TODO: Add all pics here.
  state: {
    user: {
      'email': 'anonymous@xyp.com',
      'id': ObjectId(),
      'name': '匿名用户'
    },
    isLogin: false,
    jwtToken: anonUserJwtToken,
    anonUserJwtToken: anonUserJwtToken,
    menuItems: {}
  },
  mutations: {
    logout (state) {
      state.user = anonUser
      state.isLogin = false
      state.jwtToken = anonUserJwtToken
      localStorage.setItem('jwtToken', anonUserJwtToken)
    },
    login (state, loginInfo) {
      state.user = loginInfo.user
      state.isLogin = loginInfo.isLogin
      state.jwtToken = localStorage.getItem('jwtToken')
    },
    loadMenuItems (state, payload) {
      state.menuItems = payload.menuItems
    },
    check (state, payload) {
      let jwtToken = payload.jwtToken
      if (jwtToken) {
        let jwtPayload
        try {
          jwtPayload = jwtDecode(jwtToken)
        } catch (e) {
          console.log('decode wrong')
          return
        }
        if (jwtPayload.exp > Math.floor(Date.now() / 1000)) {
          if (jwtPayload.user.name !== '匿名用户') {
            state.isLogin = true
            state.user = jwtPayload.user
          } else {
            state.isLogin = false
            state.user = anonUser
            state.jwtToken = anonUserJwtToken
          }
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
