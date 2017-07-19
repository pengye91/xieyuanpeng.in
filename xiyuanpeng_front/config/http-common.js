/**
 * Created by xyp on 17-7-15.
 */
import axios from 'axios'

export const HTTP = axios.create({
  baseURL: 'http://localhost:8000',
  headers: {
    Authorization: `Bearer ${localStorage.getItem('jwtToken')}`
  },
  timeout: 5000,
  withCredentials: true,
})
