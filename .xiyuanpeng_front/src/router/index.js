import Vue from 'vue'
import Router from 'vue-router'
import createSideMenuView from '../utils/createMenuView'
import photo from '../components/Photo'
import blog from '../components/Blog'
import Operations from '../components/Operations'
// import Uploads from '../components/Uploads'
import SideMenuView from '../components/SideMenuView'
import AllUploads from '../components/AllUploads'
import {config} from '@/config/dev'
Vue.use(Router)

const scrollBehavior = (to, from, savedPosition) => {
  if (savedPosition) {
    return savedPosition
  } else {
    let position = {}
    if (to.hash) {
      position.selector = to.hash
    }

    if (to.matched.some(m => m.meta.scrollToTop)) {
      position.x = 0
      position.y = 0
    }
    return position
  }
}

export default new Router({
  mode: 'history',
  scrollBehavior,
  routes: [
    {
      path: '/admin',
      name: 'admin',
      redirect: {
        name: 'operation',
        params: {
          post: Object.keys(config.SIDE_MENU_ITEMS)[0],
          sideMenu: config.SIDE_MENU_ITEMS[Object.keys(config.SIDE_MENU_ITEMS)[0]][0],
          operation: 'all'
        }
      }
    },
    {
      path: '/admin/:post',
      name: 'post',
      component: SideMenuView,
      redirect: to => {
        return {
          name: 'operation',
          params: {
            post: to.params.post,
            sideMenu: config.SIDE_MENU_ITEMS[to.params.post][0],
            operation: 'all'
          }
        }
      },
      children: [
        {
          path: ':sideMenu',
          component: AllUploads,
          name: 'sideMenu',
          redirect: to => {
            return {
              name: 'operation',
              params: {
                post: to.params.post,
                sideMenu: to.params.sideMenu,
                operation: 'all'
              }
            }
          },
          children: [
            {
              path: ':operation',
              name: 'operation',
              component: Operations
            }
          ]
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
          component: blog
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
