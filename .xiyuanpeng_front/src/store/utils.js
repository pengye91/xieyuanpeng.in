/**
 * Created by xyp on 17-7-26.
 */
import {config} from '../config/dev'

let menu = {}

config.HTTP.get('/menu-items')
  .then(response => {
    if (response.status < 400) {
      menu = response.data
      console.log(menu)
    } else {
      console.log('not')
      menu = config.MENU_ITEMS
    }
  })
  .catch(error => {
    menu = config.MENU_ITEMS
    console.log(error.data)
  })

export const menuItems = menu
