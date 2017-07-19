/**
 * Created by xyp on 17-7-15.
 */
import axios from 'axios'
// import store from '../store/index'

export const HTTP = axios.create({
  baseURL: 'http://localhost:8000'
})
