/**
 * Created by xyp on 17-7-20.
 */
import {http} from './http-common'

// TODO: this should be put into store/state
let adminSideMenuItem = {
  'all': '所有',
  'upload': '上传',
  'with-selected': '选中操作'
}

// TODO: this should be put into store/state
let sideMenuItem = {
  'blog': ['python', 'golang', 'django', 'miscellaneous'],
  'photography': ['project-1', 'project-2'],
  'contact': ['github', 'wechat']
}

export const config = {
  HTTP: http,
  BASE_URL: `http://www.xieyuanpeng.com`,
  IMAGE_BASE_URL: 'http://www.xieyuanpeng.com/static/images',

// TODO: this should be put into store/state
  MENU_ITEMS: {
    'blog': {
      ref: 'blog',
      name: '博客',
      sideMenuItems: {'python': 'python', 'golang': 'golang', 'django': 'django', 'miscellaneous': '杂'},
      adminSideMenuItems: adminSideMenuItem
    },
    'photography': {
      ref: 'photography',
      name: '摄影',
      sideMenuItems: {'project-1': '项目1', 'project-2': '项目2'},
      adminSideMenuItems: adminSideMenuItem
    },
    'contact': {
      ref: 'contact',
      name: '联系我',
      sideMenuItems: {'github': 'github', 'wechat': 'wechat'},
      adminSideMenuItems: adminSideMenuItem
    }
  },
  SIDE_MENU_ITEMS: sideMenuItem
}
