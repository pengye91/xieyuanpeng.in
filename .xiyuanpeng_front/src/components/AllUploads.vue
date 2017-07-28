<template>
  <div class="all-list-and-uploads">
    <Menu theme="primary" mode="horizontal" style="width: 100%; height: 40px">
      <Row type="flex" style="height: 100%;" justify="space-around" align="middle">
        <Col :span="Math.floor(24/menuItems[$route.params.post].adminSideMenuItems.length)"
             v-for="(menuItem, key) in menuItems[$route.params.post].adminSideMenuItems " :key="key">
        <MyMenuItem :to="{name: 'operation', params: {operation: `${key}`}}" :name="`${key}`">
          {{menuItem}}
        </MyMenuItem>
        </Col>
      </Row>
    </Menu>
    <router-view></router-view>
  </div>
</template>

<script>
  import MyMenuItem from './MyMenuItem'
  //  import {mapState} from 'vuex'
  import {config} from '@/config/dev'

  export default {
    components: {
      MyMenuItem
    },
    props: {
      type: String
    },
    data () {
      return {
        menuItems: config.MENU_ITEMS,
        searchText: ''
      }
    }
//    computed: {
//      ...mapState([
//        'menuItems'
//      ])
//    },
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

