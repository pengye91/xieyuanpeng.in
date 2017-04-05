<style scoped>
  .layout {
    border: 1px solid #d7dde4;
    background: #f5f7f9;
  }

  .ivu-menu-item {
    font-size: 12px;
    margin-left: 6px;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item {
    font-size: 15px;
    padding-left: 45px;
    border-bottom: 0px;
  }

  .ivu-input-wrapper {
    width: 15%;
    float: right;
    top: auto;
    right: 3%;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item:hover {
    border-bottom: 0px;
    height: 59px;
    background: #f6f6f6 none repeat scroll 0 0;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item-selected {
    background: #e6e6e6;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item-active {
    background: #e6e6e6 !important;
  }

  /*vertical*/
  .ivu-menu-light.ivu-menu-vertical .ivu-menu-item-active:not(.ivu-menu-submenu) {
    background: #cbcbcb;
    border-right: 0px;
  }

  .layout-logo {
    border-radius: 3px;
    float: left;
    position: relative;
    top: 12px;
    left: 40px;
  }

  .layout-search {
    width: 15%;
    float: right;
    top: auto;
    right: 3%;
  }

  .layout-nav li {
    left: 25%;
    font-size: 15px;
    float: left;
    padding: 0 40px;
  }

  .layout-assistant {
    width: 300px;
    margin: 0 auto;
    height: inherit;
  }

  .layout-left {
    left: 0;
  }

  .layout-breadcrumb {
    padding: 10px 15px 0;
  }

  .layout-content {
    min-height: 200px;
    margin: 10px;
    overflow: hidden;
    background: #fff;
    border-radius: 4px;
  }

  .layout-content-main {
    padding: 10px;
  }

  .layout-copy {
    text-align: center;
    padding: 10px 0 20px;
    color: #9ea7b4;
  }
</style>
<style>
  .ivu-input-wrapper-large .ivu-input-icon {
    line-height: 58px !important;
  }
</style>

<template>
  <div class="layout">
    <Menu mode="horizontal" theme="light" @on-select="menuItemRoute">
      <div class="layout-logo">
        <img src="../assets/logo.png" alt="logo" height="40" width="40">
      </div>
      <div class="layout-nav">
        <Menu-item name="blog">
          技术博客
        </Menu-item>
        <Menu-item name="2">
          摄影作品
        </Menu-item>
        <Menu-item name="3">
          联系我
        </Menu-item>
      </div>
      <Input size="large" icon="search"
             placeholder="请输入搜索内容..." v-model="search"></Input>
    </Menu>
    <div class="layout-content">
      <Row>
        <i-col span="5">
          <Menu active-name="2-1" :open-names="['2']"
                class="layout-left" width="70%">
            <Submenu name="1">
              <template slot="title">
                <Icon type="ios-book" size="16"></Icon>
                技术博客
              </template>
              <Menu-item name="python">Python</Menu-item>
              <Menu-item name="django">Django</Menu-item>
              <Menu-item name="golang">Golang</Menu-item>
              <Menu-item name="杂">杂</Menu-item>
            </Submenu>
            <Submenu name="2">
              <template slot="title">
                <Icon type="images" size="15"></Icon>
                摄影作品
              </template>
              <Menu-item name="2-1">项目 1</Menu-item>
              <Menu-item name="2-2">项目 2</Menu-item>
            </Submenu>
            <Submenu name="3">
              <template slot="title">
                <Icon type="android-contact" size="16"></Icon>
                联系我
              </template>
              <Menu-item name="3-1">
                <Icon type="social-github" size="14"></Icon>
                github
              </Menu-item>
              <Menu-item name="3-2">wechat</Menu-item>
            </Submenu>
          </Menu>
        </i-col>
        <i-col span="18">
          <div class="layout-content-main">
            <router-view></router-view>
          </div>
        </i-col>
      </Row>
    </div>
    <div class="layout-copy">
      &copy; XieYuanpeng.in
    </div>
  </div>
</template>
<script>
  export default {
    data () {
      return {
        search: ''
      }
    },
    methods: {
      menuItemRoute (key) {
        this.$router.push(key)
      }
    }
  }
</script>

