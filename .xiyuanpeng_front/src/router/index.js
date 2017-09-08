import Vue from 'vue'
import Router from 'vue-router'
import createSideMenuView from '../utils/createMenuView'
import photo from '../components/Photo'
import blogs from '../components/Blogs'
import blog from '../components/Blog'
import store from '../store/index'
import {config} from '../config/dev'
import {adminRouter} from './admin.js'
import Meta from 'vue-meta'

Vue.use(Router)
Vue.use(Meta)

store.dispatch('LOAD_MENU_ITEMS')

const keyComponentMap = {
  'blog': blogs,
  'photography': photo,
  'contact': photo
}

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

let mi = config.MENU_ITEMS
let smi = config.SIDE_MENU_ITEMS

store.watch(
  (state) => { return [state.menuItems, state.sideMenuItems] },
  (newMenu) => {
    mi = newMenu[0]
    smi = newMenu[1]
  },
)

function secondRouters (key, sideMI) {
  let secondRoutes = []
  Object.keys(sideMI[key]).forEach((secondKey) => {
    let secondRoute = {}
    let secondRoutePlus = {}
    if (key === 'blog') {
      secondRoute.path = ':tag'
      secondRoute.name = sideMI[key][secondKey]
      // secondRoute.name = 'blogs'
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
      // secondRoute.name = 'postItem'
      secondRoute.component = keyComponentMap[key]
      secondRoutes.push(secondRoute)
    }
  })
  return secondRoutes
}

function firstRouters (MI, sideMI) {
  let firstRoutes = []
  Object.keys(MI).forEach((key) => {
    let firstRoute = {}
    firstRoute.path = `/${key}`
    firstRoute.name = key
    firstRoute.component = createSideMenuView(key)
    firstRoute.redirect = {'name': Object.values(sideMI[key])[0], params: {'tag': Object.values(sideMI[key])[0]}}
    firstRoute.children = secondRouters(key, sideMI)
    firstRoutes.push(firstRoute)
  })
  return firstRoutes
}

// function fR () {
//   function proceed () {
//     if (store.state.menuItems) {
//       MI = store.state.menuItems
//       let firstRoutes = []
//       sideMI = config.SIDE_MENU_ITEMS
//       Object.keys(MI).forEach((key) => {
//         let firstRoute = {}
//         firstRoute.path = `/${key}`
//         firstRoute.name = key
//         firstRoute.component = createSideMenuView(key)
//         firstRoute.redirect = {'name': Object.values(sideMI[key])[0]}
//         firstRoute.children = secondRouters(key)
//         firstRoutes.push(firstRoute)
//       })
//       return firstRoutes
//     }
//   }
//
//   if (store.state.menuItems.blog === undefined) {
//     store.watch(
//       (state) => state.menuItems,
//       (value) => {
//         if (value.blog !== undefined) {
//           return proceed()
//         }
//       }
//     )
//   } else {
//     return proceed()
//   }
// }
const router = new Router({
  mode: 'history',
  scrollBehavior,
  routes: [
    ...adminRouter,
    ...firstRouters(mi, smi)
  ]
})

export default router
