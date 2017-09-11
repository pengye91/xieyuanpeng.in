<template>
  <div class="layout-content">
    <Row type="flex" style="height: 100%">
      <Col span="3">
      <Menu style="height: 100%; border-top: 1px solid lightgray; width: 100%;">
        <div style="position: fixed; width: 12.5%;">
          <MyMenuItem v-for="item in sideMenu[type]" :name="item" :key="item" :to="{name: item}"
                      style="margin-top: 5%; width: 100%">
            <Icon type="ios-book" size="16"></Icon>
            {{ item }}
          </MyMenuItem>
        </div>
      </Menu>
      </Col>
      <Col span="21" style="height: 100%">
      <div class="layout-content-main">
        <router-view></router-view>
      </div>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {EventBus} from '../store/EventBus'
  import Login from './Login'
  import Register from './Register'
  import MyMenuItem from './MyMenuItem'
  export default {
    components: {
      Login, Register, MyMenuItem
    },
    props: {
      type: String
    },
    data () {
      return {
        currentPage: 'blog',
        sideMenu: {
          'blog': ['python', 'golang', 'django', '杂'],
          'photography': ['项目1', '项目2'],
          'contact-me': ['github', 'wechat']
        },
        searchText: ''
      }
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
<style scoped>
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
    background: #e6e6e6 !important;
    font-weight: bold;
  }

  .ivu-menu-light.ivu-menu-vertical .ivu-menu-item-active:not(.ivu-menu-submenu) {
    background: #eaeaea;
    border-right: 0;
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
    position: relative;
    height: 100%;
    padding: 42px 0 30px 0;
    background: transparent;
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
    height: 20px;
  }

  html, body, .layout {
    height: 100%;
  }
</style>

