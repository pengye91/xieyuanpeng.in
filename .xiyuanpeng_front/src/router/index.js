import Vue from 'vue'
import Router from 'vue-router'
import createSideMenuView from '../utils/createMenuView'
import photo from '../components/Photo'
import blogs from '../components/Blogs'
import blog from '../components/Blog'
// import store from '../store/index'
// import Operations from '../components/Operations'
// import Uploads from '../components/Uploads'
// import SideMenuView from '../components/SideMenuView'
// import AllUploads from '../components/AllUploads'
import {config} from '@/config/dev'
import {adminRouter} from './admin.js'
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

const keyComponentMap = {
  'blog': blogs,
  'photography': photo,
  'contact': photo
}

let MI = config.MENU_ITEMS

function firstRouters () {
  let firstRoutes = []
  console.log(MI)
  Object.keys(MI).forEach((key) => {
    let firstRoute = {}
    firstRoute.path = `/${key}`
    firstRoute.name = key
    firstRoute.component = createSideMenuView(key)
    // firstRoute.redirect = {'name': MI[Object.keys(MI)[0]]}
    firstRoute.children = secondRouters(key)
    firstRoutes.push(firstRoute)
  })
  return firstRoutes
}

function secondRouters (key) {
  let sideMI = config.SIDE_MENU_ITEMS
  let secondRoutes = []
  console.log(sideMI)
  Object.keys(sideMI).forEach((secondKey) => {
    let secondRoute = {}
    // TODO: remove this by restructuring blogs.vue or blog.vue
    let secondRoutePlus = {}
    switch (key) {
      case 'blog':
        secondRoute.path = secondKey
        secondRoute.name = secondKey
        secondRoute.component = keyComponentMap[key]
        secondRoute.props = {'tag': secondKey}
        secondRoutePlus.path = `${secondKey}/:blogPath`
        secondRoutePlus.name = 'blogPath'
        secondRoutePlus.component = blog
        secondRoutePlus.props = (route) => ({
          'blogPath': route.params.blogPath,
          'tag': secondKey
        })
        secondRoutes.push(secondRoute, secondRoutePlus)
        break
      default:
        secondRoute.path = secondKey
        secondRoute.name = sideMI[secondKey]
        secondRoute.component = keyComponentMap[key]
        secondRoutes.push(secondRoute)
    }
  })
  return secondRoutes
}

export default new Router({
  mode: 'history',
  scrollBehavior,
  routes: [
    ...adminRouter,
    ...firstRouters()
    // {
    //   path: '/blog',
    //   name: 'blog',
    //   component: createSideMenuView('blog'),
    //   redirect: {'name': 'python'},
    //   children: [
    //     {
    //       path: 'python',
    //       name: 'python',
    //       component: blogs,
    //       props: {'tag': 'python'}
    //     },
    //     {
    //       path: 'python/:blogPath',
    //       name: 'blogPath',
    //       component: blog,
    //       props: (route) => ({
    //         'blogPath': route.params.blogPath,
    //         'tag': 'python'
    //       })
    //     },
    //     {
    //       path: 'golang',
    //       name: 'golang',
    //       component: photo
    //     },
    //     {
    //       path: '杂',
    //       name: '杂',
    //       component: photo
    //     },
    //     {
    //       path: 'django',
    //       name: 'django',
    //       component: photo
    //     }
    //   ]
    // },
    // {
    //   path: '/photography',
    //   name: 'photography',
    //   component: createSideMenuView('photography'),
    //   redirect: {'name': '项目1'},
    //   children: [
    //     {
    //       path: '1',
    //       name: '项目1',
    //       component: photo
    //     },
    //     {
    //       path: '2',
    //       name: '项目2',
    //       component: photo
    //     }
    //   ]
    // },
    // {
    //   path: '/contact',
    //   name: 'contact',
    //   component: createSideMenuView('contact'),
    //   redirect: {'name': 'github'},
    //   children: [
    //     {
    //       path: 'wechat',
    //       name: 'wechat',
    //       component: photo
    //     },
    //     {
    //       path: 'github',
    //       name: 'github',
    //       component: photo
    //     }
    //   ]
    // }
  ]
})
