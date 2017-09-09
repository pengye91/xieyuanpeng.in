/**
 * Created by xyp on 17-8-17.
 */
import Operations from '../components/Operations'
import SideMenuView from '../components/SideMenuView'
import AllUploads from '../components/AllUploads'
import store from '../store/index'
// import {config} from '@/config/dev'

store.commit('setAllMenu')

let defaultDest = {
  name: 'operation',
  params: {
    // post: Object.keys(JSON.parse(localStorage.getItem('sideMenuItems')))[0],
    // sideMenu: Object.keys(JSON.parse(localStorage.getItem('sideMenuItems'))[Object.keys(JSON.parse(localStorage.getItem('sideMenuItems')))[0]])[0],
    operation: 'all'
  }
}

export const adminRouter = [
  {
    path: '/admin',
    name: 'admin',
    redirect: defaultDest
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
          sideMenu: Object.keys(JSON.parse(localStorage.getItem('sideMenuItems'))[to.params.post])[0],
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
            component: Operations,
            props: (route) => ({
              'type': route.params.post === 'blog' ? 'blogs' : 'pics'
            })
          }
        ]
      }
    ]
  }
]
