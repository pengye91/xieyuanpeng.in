/**
 * Created by xyp on 17-7-15.
 */
import axios from 'axios'
// import store from '../store/index'

// let jwtToken = store.state.jwtToken
// let jwtToken = localStorage.getItem('jwtToken')

export const HTTP = axios.create({
  baseURL: 'http://localhost:8000'
  // headers: {
  //   Authorization: `Bearer ${jwtToken}`
  // }
})
