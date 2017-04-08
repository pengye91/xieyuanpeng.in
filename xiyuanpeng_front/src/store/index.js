import Vue from 'vue'
import Vuex from 'vuex'
import * as actions from './actions'
// import * as getters from './getters'
import mutations from './mutations'
// import * as defaults from './defaults'
import createLogger from 'vuex/dist/logger'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
  state: {
    searchText: null
  },
  mutations,
  actions,
  strict: debug,
  plugins: debug ? [createLogger()] : []
})

