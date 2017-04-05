import Vue from 'vue'
import Router from 'vue-router'
import Blogs from '../components/Login'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/blog',
      name: 'blog',
      component: Blogs
    }
  ]
})
