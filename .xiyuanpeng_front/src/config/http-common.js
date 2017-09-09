/**
 * Created by xyp on 17-7-15.
 */
import axios from 'axios'

export const http = axios.create({
  baseURL: 'http://localhost:8000/api/v1',
  // baseURL: 'https://www.xieyuanpeng.com/api/v1',
  timeout: 5000,
  withCredentials: true
})
