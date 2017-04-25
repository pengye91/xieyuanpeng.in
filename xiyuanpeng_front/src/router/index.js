import Vue from 'vue'
import Router from 'vue-router'
// import Blog from '../components/Blog'
// import Photo from '../components/Photo'
import createMenuView from '../components/createMenuView'
Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/blog',
      name: 'blog',
      component: createMenuView('blog')
    },
    {
      path: '/photography',
      name: 'photographs',
      component: createMenuView('photography')
    }
  ]
})
