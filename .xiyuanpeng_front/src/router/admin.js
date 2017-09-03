/**
 * Created by xyp on 17-8-17.
 */
import Operations from '../components/Operations'
import SideMenuView from '../components/SideMenuView'
import AllUploads from '../components/AllUploads'
import {config} from '@/config/dev'

export const adminRouter = [
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
