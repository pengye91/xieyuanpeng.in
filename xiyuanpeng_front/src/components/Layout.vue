<style scoped>
  .layout {
    border: 1px solid #d7dde4;
    background: #f5f7f9;
  }

  .ivu-menu-item {
    font-size: 12px;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item {
    font-size: 15px;
    padding-left: 45px;
    border-bottom: 0;
    text-align: center;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item:hover {
    font-weight: bold;
    border-bottom: 0;
    height: 59px;
    text-align: center;
    background: #f6f6f6 none repeat scroll 0 0;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item-selected {
    background: #e6e6e6 !important;
    font-weight: bold;
  }

  .ivu-menu-light.ivu-menu-vertical .ivu-menu-item-active:not(.ivu-menu-submenu) {
    background: #cbcbcb;
    border-right: 0;
    font-weight: bold;
  }

  .ivu-menu-light.ivu-menu-vertical .ivu-menu-item:hover {
    font-weight: bold;
  }

  .ivu-menu-light.ivu-menu-vertical .ivu-menu-item{
    font-size: 13px;
    padding-left: 25%;
  }

  .layout-logo {
    position: relative;
    top: 10px;
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
    padding: 70px 10px 58px 10px;
    background: transparent;
    border-radius: 4px;
  }

  .layout-content-main {
    padding: 10px;
    background: white;
  }

  .layout-copy {
    text-align: center;
    z-index: 10;
    background: #fff;
    padding: 10px 0 20px;
    position: fixed;
    bottom: 0;
    width: 100%;
    color: #828a97;
  }
</style>
<style>
  .ivu-input-wrapper-large .ivu-input-icon {
    line-height: 56px !important;
  }

  html, body, .layout {
    height: 100%;
  }
</style>

<template>
  <div class="layout">
    <Menu mode="horizontal" theme="light" style="width: 100%; position: fixed; z-index: 10; top: 0"
          @on-select="menuItemRoute">
      <Row>
        <Col span="1" offset="1">
        <div class="layout-logo">
          <img src="../assets/logo.png" alt="logo" height="40" width="40">
        </div>
        </Col>
        <Col span="3" offset="1">
        <Input size="large" icon="search"
               placeholder="请输入搜索内容..." :value="searchText" @input="searchNotify"></Input>
        </Col>
        <Col span="14" offset="1">
        <div class="layout-nav">
          <MyMenuItem name="blog" to="/blog">
            技术博客
          </MyMenuItem>
          <MyMenuItem name="photography" to="/photography">
            摄影作品
          </MyMenuItem>
          <MyMenuItem name="contact-me" to="/contact-me">
            联系我
          </MyMenuItem>
        </div>
        </Col>
        <Col span="3">
        <div class="login">
          <Register></Register>
          <Login></Login>
        </div>
        </Col>
      </Row>
    </Menu>
    <div class="layout-content">
      <Row type="flex" style="height: 100%;">
        <Col span="3">
        <Menu active-name="2-1" :open-names="['2']"
              style="height: 100%" width="100%">
          <!--<Submenu name="1">-->
          <template slot="title">
            <Icon type="ios-book" size="16"></Icon>
            技术博客
          </template>
          <MyMenuItem v-for="item in sideMenu[currentPage]" :name="item" :key="item" :to="item">
            <Icon type="ios-book" size="16"></Icon>
            {{ item }}
          </MyMenuItem>
        </Menu>
        </Col>
        <Col span="21" style="overflow: auto">
        <div class="layout-content-main">
          <router-view></router-view>
        </div>
        </Col>
      </Row>
    </div>
    <div class="layout-copy">
      &copy; XieYuanpeng.in
    </div>
  </div>
</template>
<script>
//  import { mapState, mapActions } from 'vuex'
  import {EventBus} from '../store/EventBus'
  import Login from './Login'
  import Register from './Register'
  import MyMenuItem from './MyMenuItem'
  export default {
    components: {
      Login, Register, MyMenuItem
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

