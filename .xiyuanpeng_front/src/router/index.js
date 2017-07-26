import Vue from 'vue'
import Router from 'vue-router'
import createSideMenuView from '../utils/createMenuView'
import photo from '../components/Photo'
import postAll from '../components/PostAll'
import SideMenuView from '../components/SideMenuView.vue'
Vue.use(Router)

const scrollBehavior = (to, from, savedPosition) => {
  if (savedPosition) {
    return savedPosition
  } else {
    let position = {}
    if (to.hash) {
      console.log(to, from)
      position.selector = to.hash
    }

    if (to.matched.some(m => m.meta.scrollToTop)) {
      position.x = 0
      position.y = 0
    }
    console.log(position)
    return position
  }
}

export default new Router({
  mode: 'history',
  scrollBehavior,
  routes: [
    {
      path: '/admin/:post',
      name: 'admin',
      component: SideMenuView,
      redirect: '/admin/:post/all',
      children: [
        {
          path: ':option',
          component: postAll
        }
      ]
    },
    {
      path: '/blog',
      name: 'blog',
      component: createSideMenuView('blog'),
      redirect: {'name': 'python'},
      children: [
        {
          path: 'python',
          name: 'python',
          component: photo
        },
        {
          path: 'golang',
          name: 'golang',
          component: photo
        },
        {
          path: '杂',
          name: '杂',
          component: photo
        },
        {
          path: 'django',
          name: 'django',
          component: photo
        }
      ]
    },
    {
      path: '/photography',
      name: 'photography',
      component: createSideMenuView('photography'),
      redirect: {'name': '项目1'},
      children: [
        {
          path: '1',
          name: '项目1',
          component: photo
        },
        {
          path: '2',
          name: '项目2',
          component: photo
        }
      ]
    },
    {
      path: '/contact',
      name: 'contact',
      component: createSideMenuView('contact'),
      redirect: {'name': 'github'},
      children: [
        {
          path: 'wechat',
          name: 'wechat',
          component: photo
        },
        {
          path: 'github',
          name: 'github',
          component: photo
        }
      ]
    }
  ]
})
