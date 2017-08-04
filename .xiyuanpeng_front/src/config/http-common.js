/**
 * Created by xyp on 17-7-15.
 */
import axios from 'axios'
import store from '../store/index'

export const http = axios.create({
  baseURL: 'http://localhost:8000/api/v1',
  // baseURL: 'https://www.xieyuanpeng.com/api/v1',
  headers: {
    Authorization: `Bearer ${store.state.jwtToken}`
  },
  timeout: 5000,
  withCredentials: true
})
