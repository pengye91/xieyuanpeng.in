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
    border-bottom: 0px;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item:hover {
    text-align: center;
    font-weight: bold;
    border-bottom: 0px;
    height: 59px;
    background: #f6f6f6 none repeat scroll 0 0;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item-selected {
    background: #e6e6e6 !important;
    font-weight: bold;
  }

  .ivu-menu-light.ivu-menu-vertical .ivu-menu-item-active:not(.ivu-menu-submenu) {
    background: #cbcbcb;
    border-right: 0px;
    font-weight: bold;
  }

  .ivu-menu-light.ivu-menu-vertical .ivu-menu-item:hover {
    font-weight: bold;
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
    padding: 70px 10px 10px 10px;
    overflow: hidden;
    background: transparent;
    border-radius: 4px;
  }

  .layout-content-main {
    padding: 10px;
  }

  .layout-copy {
    text-align: center;
    padding: 10px 0 20px;
    position: fixed;
    bottom: 0;
    width: 100%;
    color: #9ea7b4;
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
          @on-select="menuItemRoute" :activeName="currentPage">
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
          <Menu-item name="blog">
            技术博客

          </Menu-item>
          <Menu-item name="photography">
            摄影作品

          </Menu-item>
          <Menu-item name="contact-me">
            联系我

          </Menu-item>
        </div>
        </Col>
        <Col span="3">
        <div class="login">
          <Button shape="circle" @click="register">
            注册

          </Button>
          <Button shape="circle" @click="login">
            登录

          </Button>
        </div>
        </Col>
      </Row>
    </Menu>
    <div class="layout-content">
      <Row type="flex" style="height: 100%; background: #fff">
        <Col span="3">
        <Menu active-name="2-1" :open-names="['2']"
              style="height: 100%" width="100%">
          <!--<Submenu name="1">-->
          <template slot="title">
            <Icon type="ios-book" size="16"></Icon>
            技术博客

          </template>
          <Menu-item v-for="item in sideMenu[currentPage]" :name="item" :key="item">
            <Icon type="ios-book" size="16"></Icon>
            {{ item }}

          </Menu-item>
        </Menu>
        </Col>
        <Col span="18">
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
  export default {
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
//    computed: mapState(['searchText']),
    methods: {
//      ...mapActions(['setSearchText']),
      menuItemRoute (key) {
        this.$router.push(key)
        this.currentPage = key
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

