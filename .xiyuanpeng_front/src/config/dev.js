/**
 * Created by xyp on 17-7-20.
 */
import {http} from './http-common'

// TODO: this should be put into store/state
let adminSideMenuItem = {
  'blog': {'all': '所有', 'upload': '上传', 'with-selected': '选中操作'},
  'photography': {'all': '所有', 'upload': '上传', 'with-selected': '选中操作'},
  'contact': {'all': '所有', 'upload': '上传', 'with-selected': '选中操作'}
}

// TODO: this should be put into store/state
let sideMenuItem = {
  'blog': {'python': 'python', 'golang': 'golang', 'django': 'django', 'miscellaneous': 'miscellaneous'},
  'photography': {'project-1': '项目1', 'project-2': '项目2'},
  'contact': {'github': 'github', 'wechat': '微信'}
}

http.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${localStorage.getItem('jwtToken')}`
  return config
})

export const config = {
  HTTP: http,
  BASE_URL: `http://localhost:8000`,
  // BASE_URL: `https://www.xieyuanpeng.com`,
  IMAGE_BASE_URL: 'https://www.xieyuanpeng.com/static/images/',

// TODO: this should be put into store/state
  MENU_ITEMS: {
    'blog': {
      ref: 'blog',
      name: '博客',
      sideMenuItems: {'python': 'python', 'golang': 'golang', 'django': 'django', 'miscellaneous': '杂'},
      adminSideMenuItems: {
        'all': '所有',
        'upload': '上传',
        'with-selected': '选中操作'
      }
    },
    'photography': {
      ref: 'photography',
      name: '摄影',
      sideMenuItems: {'project-1': '项目1', 'project-2': '项目2'},
      adminSideMenuItems: {
        'all': '所有',
        'upload': '上传',
        'with-selected': '选中操作'
      }
    },
    'contact': {
      ref: 'contact',
      name: '联系我',
      sideMenuItems: {'github': 'github', 'wechat': 'wechat'},
      adminSideMenuItems: {
        'all': '所有',
        'upload': '上传',
        'with-selected': '选中操作'
      }
    }
  }
}
