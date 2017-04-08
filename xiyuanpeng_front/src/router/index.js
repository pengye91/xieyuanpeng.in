import Vue from 'vue'
import Router from 'vue-router'
import Blog from '../components/Blog'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/blog',
      name: 'blog',
      component: Blog
    }
  ]
})
