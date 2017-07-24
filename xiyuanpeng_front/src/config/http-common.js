/**
 * Created by xyp on 17-7-15.
 */
import axios from 'axios'
import store from '../store/index'

export const http = axios.create({
  baseURL: 'http://www.yukere.com',
  headers: {
    Authorization: `Bearer ${store.state.jwtToken}`
  },
  timeout: 5000,
  withCredentials: true
})
