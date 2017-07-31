<style scoped>
  .layout {
    border: 0 solid #d7dde4;
    background-color: white;
  }

  .ivu-menu-item {
    font-size: 12px;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item {
    font-size: 15px;
    padding-left: 45px;
    border-bottom: 0;
    text-align: center;
    line-height: 40px;
  }

  .login {
    line-height: 37px;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item:hover {
    font-weight: bold;
    border-bottom: 0;
    height: 39px;
    text-align: center;
    background: #f6f6f6 none repeat scroll 0 0;
  }

  .ivu-menu-horizontal {
    line-height: 38px;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item-selected {
    background: #c9c9c9 !important;
    font-weight: bold;
  }

  .ivu-menu-light.ivu-menu-vertical .ivu-menu-item-active:not(.ivu-menu-submenu) {
    background: #cbcbcb;
    border-right: 10px;
    font-weight: bold;
  }

  .ivu-menu-light.ivu-menu-vertical .ivu-menu-item:hover {
    font-weight: bold;
  }

  .ivu-menu-light.ivu-menu-vertical .ivu-menu-item {
    font-size: 13px;
    padding-left: 25%;
  }

  .layout-logo {
    position: relative;
    top: 1px;
    left: 15px;
  }

  .layout-nav li {
    left: 15%;
    font-size: 15px;
    float: left;
    padding: 0 40px;
  }

  .layout-assistant {
    width: 300px;
    margin: 0 auto;
    height: inherit;
  }

  .layout-breadcrumb {
    padding: 10px 15px 0;
  }

  .layout-content {
    height: 100%;
    padding: 42px 10px 30px 5px;
    background: transparent;
    border-radius: 4px;
  }

  .layout-content-main {
    height: 100%;
    padding: 10px;
    background: white;
  }

  .layout-copy {
    text-align: center;
    z-index: 10;
    background: #fff;
    padding: 5px 0 5px;
    position: fixed;
    bottom: 0;
    width: 100%;
    color: #828a97;
  }
</style>
<style>
  .ivu-input-wrapper-large .ivu-input-icon {
    line-height: 30px !important;
  }

  .ivu-input-large {
    height: 30px;
  }

  html, body, .layout {
    height: 100%;
  }
</style>

<template>
  <div class="layout">
    <Menu mode="horizontal" theme="light" style="width: 100%; position: fixed; z-index: 10; top: 0; height: 40px">
      <Row type="flex" style="height: 100%;" justify="space-between" align="middle">
        <Col span="1">
        <div class="layout-logo">
          <img src="../assets/logo.png" alt="logo" height="40" width="40">
        </div>
        </Col>
        <Col span="3" offset="1">
        <Input size="large" icon="search"
               placeholder="请输入搜索内容..." :value="searchText" @input="searchNotify"></Input>
        </Col>
        <Col :span="Math.floor(16/menuItems.length)" v-for="(menuItem, key) in menuItems" :key="key"
             class="layout-nav">
        <MyMenuItem :to="$route.path.startsWith('/admin') ? `/admin/${key}` : `${key}`">
          {{menuItem.name}}
        </MyMenuItem>
        </Col>
        <Col span="3">
        <div v-if="!isLogin" class="register-login">
          <Register></Register>
          <Login></Login>
        </div>
        <div v-if="isLogin">
          <User></User>
        </div>
        </Col>
      </Row>
    </Menu>
    <router-view></router-view>
    <div class="layout-copy">
      &copy; xieyuanpeng.com
    </div>
  </div>
</template>
<script>
  //  import { mapState, mapActions } from 'vuex'
  import {EventBus} from '../store/EventBus'
  import Login from './Login'
  import Register from './Register'
  import User from './User'
  import Logout from './Logout.vue'
  import MyMenuItem from './MyMenuItem'
  import {mapState} from 'vuex'
  import {config} from '@/config/dev'

  export default {
    components: {
      Login, Register, MyMenuItem, User, Logout
    },
    data () {
      return {
        currentPage: 'blog',
        sideMenu: config.SIDE_MENU_ITEMS,
        searchText: ''
      }
    },
    mounted: function () {
      this.$store.commit({
        type: 'check',
        jwtToken: localStorage.getItem('jwtToken')
      })

      let menu = {}
      config.HTTP.get('menu-items')
        .then(response => {
          if (response.status < 400) {
            menu = response.data
            this.$store.commit({
              type: 'loadMenuItems',
              menuItems: menu
            })
          } else {
            console.log(response.data)
          }
        })
        .catch(error => {
          this.$store.commit({
            type: 'loadMenuItems',
            menuItems: config.MENU_ITEMS
          })
          console.log(error.response.data)
          console.log(this.menuItems)
        })
    },
    computed: {
      ...mapState([
        'user', 'isLogin', 'menuItems'
      ])
    },
    methods: {
      menuItemRoute (key) {
        this.currentPage = key
        console.log(this.currentPage)
      },
      login () {
        this.$router.push('blog')
      },
      register () {
        this.$router.push('blog')
      },
      searchNotify (value) {
        this.searchText = value
        EventBus.$emit('search-text', this.searchText)
        console.log(this.searchText)
      }
    }
  }
</script>

