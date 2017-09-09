/**
 * Created by xyp on 17-7-20.
 */
import {http} from './http-common'

const defaultAdminSideMenuItem = {
  'all': '所有',
  'upload': '上传',
  'with-selected': '选中操作'
}

const defaultSideMenuItem = {
  'blog': {'python': 'python', 'golang': 'golang', 'django': 'django', 'miscellaneous': '杂'},
  'photography': {'project-1': '项目1', 'project-2': '项目2'},
  'contact': {'github': 'github', 'wechat': 'wechat'}
}

const defaultMenuItem = {
  'blog': {
    ref: 'blog',
    name: '博客',
    sideMenuItems: {'python': 'python', 'golang': 'golang', 'django': 'django', 'miscellaneous': '杂'},
    adminSideMenuItems: defaultAdminSideMenuItem
  },
  'photography': {
    ref: 'photography',
    name: '摄影',
    sideMenuItems: {'project-1': '项目1', 'project-2': '项目2'},
    adminSideMenuItems: defaultAdminSideMenuItem
  },
  'contact': {
    ref: 'contact',
    name: '联系我',
    sideMenuItems: {'github': 'github', 'wechat': '微信'},
    adminSideMenuItems: defaultAdminSideMenuItem
  }
}

http.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${localStorage.getItem('jwtToken')}`
  return config
})

export const config = {
  HTTP: http,
  // BASE_URL: `http://localhost:8000`,
  BASE_URL: `https://www.xieyuanpeng.com`,
  IMAGE_BASE_URL: 'https://www.xieyuanpeng.com/static/images/',

// TODO: this should be put into store/state
  MENU_ITEMS: JSON.parse(localStorage.getItem('menuItems')) === null ? defaultMenuItem : JSON.parse(localStorage.getItem('menuItems')),
  SIDE_MENU_ITEMS: JSON.parse(localStorage.getItem('sideMenuItems')) === null ? defaultSideMenuItem : JSON.parse(localStorage.getItem('sideMenuItems'))
}

