<template>
  <div class="layout-content">
    <Row type="flex" style="height: 100%">
      <Col span="3">
      <Menu style="height: 100%; border-top: 1px solid lightgray; width: 100%;">
        <div style="position: fixed; width: 12.5%;" v-if="$route.path.startsWith('/admin')">
          <MyMenuItem v-for="(item, key) in sideMenuItems[$route.params.post]" :name="key"
                      :key="key" :to="{name: 'sideMenu', params: {'sideMenu': key, 'post': $route.params.post}}"
                      style="margin-top: 5%; width: 100%">
            {{item}}
          </MyMenuItem>
        </div>
        <div v-else>
          <MyMenuItem v-for="(item, key) in sideMenuItems[type]" :key="key"
                      :to="{name: 'postItems', params: {postItem: key, post: type}}" :name="item"
                      style="margin-top: 5%; width: 100%">
            <Icon type="ios-book" size="16"></Icon>
            {{item}}
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
  import MyMenuItem from './MyMenuItem'
  import photo from '../components/Photo'
  import blogs from '../components/Blogs'
  import {mapState} from 'vuex'
  //  import {config} from '@/config/dev'

  export default {
    components: {
      MyMenuItem
    },
    props: {
      type: String
    },
    data () {
      return {
        searchText: '',
        keyComponentMap: {
          'blog': blogs,
          'photography': photo,
          'contact': photo
        }
      }
    },
    computed: {
      ...mapState([
        'menuItems', 'sideMenuItems'
      ])
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

