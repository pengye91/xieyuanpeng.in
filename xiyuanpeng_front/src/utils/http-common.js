/**
 * Created by xyp on 17-7-15.
 */
import axios from 'axios'
let jwtToken = localStorage.getItem('jwtToken')

export const HTTP = axios.create({
  baseURL: 'http://localhost:8000',
  headers: {
    Authorization: `Bearer ${jwtToken}`
  }
})
