import Vue from 'vue'
import Router from 'vue-router'
import createSideMenuView from '../utils/createMenuView'
import photo from '../components/Photo'
import blogs from '../components/Blogs'
import blog from '../components/Blog'
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
let sideMI = config.SIDE_MENU_ITEMS

function firstRouters () {
  let firstRoutes = []
  Object.keys(MI).forEach((key) => {
    let firstRoute = {}
    firstRoute.path = `/${key}`
    firstRoute.name = key
    firstRoute.component = createSideMenuView(key)
    firstRoute.redirect = {'name': Object.values(sideMI[key])[0]}
    firstRoute.children = secondRouters(key)
    firstRoutes.push(firstRoute)
  })
  return firstRoutes
}

function secondRouters (key) {
  let secondRoutes = []
  Object.keys(sideMI[key]).forEach((secondKey) => {
    let secondRoute = {}
    let secondRoutePlus = {}
    if (key === 'blog') {
      secondRoute.path = secondKey
      secondRoute.name = sideMI[key][secondKey]
      secondRoute.component = keyComponentMap[key]
      secondRoute.props = {'tag': secondKey}
      secondRoutePlus.path = `${secondKey}/:blogPath`
      secondRoutePlus.name = `${secondKey}-blogPath`
      secondRoutePlus.component = blog
      secondRoutePlus.props = (route) => ({
        'blogPath': route.params.blogPath,
        'tag': secondKey
      })
      secondRoutes.push(secondRoute, secondRoutePlus)
    } else {
      secondRoute.path = secondKey
      secondRoute.name = sideMI[key][secondKey]
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
  ]
})

