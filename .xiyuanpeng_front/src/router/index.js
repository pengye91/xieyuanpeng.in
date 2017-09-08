import Vue from 'vue'
import Router from 'vue-router'
// import createSideMenuView from '../utils/createMenuView'
// import photo from '../components/Photo'
import sideMenuView from '../components/SideMenuView'
import postItems from '../components/PostItems'
// import blogs from '../components/Blogs'
import blog from '../components/Blog'
import store from '../store/index'
import {config} from '../config/dev'
import {adminRouter} from './admin.js'
import Meta from 'vue-meta'

Vue.use(Router)
Vue.use(Meta)

store.dispatch('LOAD_MENU_ITEMS')

// const keyComponentMap = {
//   'blog': blogs,
//   'photography': photo,
//   'contact': photo
// }

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

// let MI = config.MENU_ITEMS
let sideMI = config.SIDE_MENU_ITEMS

// store.watch(
//   (state) => { return [state.menuItems, state.sideMenuItems] },
//   (newMenu) => {
//     mi = newMenu[0]
//     smi = newMenu[1]
//   },
// )
//
// function secondRouters (key, sideMI) {
//   let secondRoutes = []
//   Object.keys(sideMI[key]).forEach((secondKey) => {
//     let secondRoute = {}
//     let secondRoutePlus = {}
//     if (key === 'blog') {
//       secondRoute.path = ':tag'
//       // secondRoute.name = sideMI[key][secondKey]
//       secondRoute.name = 'postItems'
//       secondRoute.component = keyComponentMap[key]
//       secondRoute.props = (route) => ({'tag': secondKey})
//       secondRoutePlus.path = `${secondKey}/:blogPath`
//       secondRoutePlus.name = `${secondKey}-blogPath`
//       secondRoutePlus.component = blog
//       secondRoutePlus.props = (route) => ({
//         'blogPath': route.params.blogPath,
//         'tag': secondKey
//       })
//       secondRoutes.push(secondRoute, secondRoutePlus)
//     } else {
//       secondRoute.path = secondKey
//       // secondRoute.name = sideMI[key][secondKey]
//       secondRoute.name = 'postItems'
//       secondRoute.component = keyComponentMap[key]
//       secondRoutes.push(secondRoute)
//     }
//   })
//   return secondRoutes
// }
//
// function firstRouters (MI, sideMI) {
//   let firstRoutes = []
//   Object.keys(MI).forEach((key) => {
//     let firstRoute = {}
//     firstRoute.path = `/${key}`
//     firstRoute.name = key
//     firstRoute.component = createSideMenuView(key)
//     firstRoute.redirect = {'name': Object.values(sideMI[key])[0], params: {'tag': Object.values(sideMI[key])[0]}}
//     firstRoute.children = secondRouters(key, sideMI)
//     firstRoutes.push(firstRoute)
//   })
//   return firstRoutes
// }

let secondRoute = {
  name: 'postItems',
  path: ':postItem',
  component: postItems,
  props: (route) => ({
    'type': route.params.post,
    'tag': route.params.postItem
  })
}

let secondRoutePlus = {
  name: 'postItemsPlus',
  path: ':postItem/:postItemPlus',
  component: blog,
  props: (route) => ({
    'tag': route.params.postItem
  })
}

let firstRoute = {
  name: 'posts',
  path: '/:post',
  component: sideMenuView,
  props: (route) => ({
    'type': route.params.post
  }),
  redirect: to => {
    return {
      name: 'postItems',
      params: {
        'postItem': Object.keys(sideMI[to.params.post])[0],
        'post': to.params.post
      }
    }
  },
  children: [secondRoute, secondRoutePlus]
}

// let rootRoute = {
//   name: 'rootRoute',
//   path: '/',
//   redirect: {name: 'posts', params: {post: Object.keys(sideMI)[0]}}
// }

const router = new Router({
  mode: 'history',
  scrollBehavior,
  routes: [
    ...adminRouter,
    // rootRoute,
    firstRoute
  ]
})

export default router
