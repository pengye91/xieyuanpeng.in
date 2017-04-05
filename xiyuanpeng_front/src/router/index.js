import Vue from 'vue'
import Router from 'vue-router'
import Layout from '../components/Layout'
import Blogs from '../components/Blogs'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Layout',
      component: Layout,
      children: [
        {
          path: 'blog',
          name: 'blog',
          component: Blogs
        }
      ]
    }
  ]
})
