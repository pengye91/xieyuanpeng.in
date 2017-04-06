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
                 placeholder="请输入搜索内容..." v-model="search"></Input>
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
        <Col span="3" style="border-right: ">
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
  export default {
    data () {
      return {
        search: '',
        currentPage: 'blog',
        sideMenu: {
          'blog': ['python', 'golang', 'django', '杂'],
          'photography': ['项目1', '项目2'],
          'contact-me': ['github', 'wechat']
        }
      }
    },
    methods: {
      menuItemRoute (key) {
        this.$router.push(key)
        this.currentPage = key
      },
      login () {
        this.$router.push('blog')
      },
      register () {
        this.$router.push('blog')
      }
    }
  }
</script>

<!--<template>-->
<!--<el-row class="tac">-->
  <!--<el-col :span="8">-->
    <!--<el-menu default-active="2" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose">-->
      <!--<el-submenu index="1">-->
        <!--<template slot="title"><i class="el-icon-message"></i>导航一</template>-->
        <!--<el-menu-item-group>-->
          <!--<template slot="title">分组一</template>-->
          <!--<el-menu-item index="1-1">选项1</el-menu-item>-->
          <!--<el-menu-item index="1-2">选项2</el-menu-item>-->
        <!--</el-menu-item-group>-->
        <!--<el-menu-item-group title="分组2">-->
          <!--<el-menu-item index="1-3">选项3</el-menu-item>-->
        <!--</el-menu-item-group>-->
        <!--<el-submenu index="1-4">-->
          <!--<template slot="title">选项4</template>-->
          <!--<el-menu-item index="1-4-1">选项1</el-menu-item>-->
        <!--</el-submenu>-->
      <!--</el-submenu>-->
      <!--<el-menu-item index="2"><i class="el-icon-menu"></i>导航二</el-menu-item>-->
      <!--<el-menu-item index="3"><i class="el-icon-setting"></i>导航三</el-menu-item>-->
    <!--</el-menu>-->
  <!--</el-col>-->
<!--</el-row>-->
<!--</template>-->
<!---->
<!--<script>-->
  <!--export default {-->
    <!--methods: {-->
      <!--handleOpen (key, keyPath) {-->
        <!--console.log(key, keyPath)-->
      <!--},-->
      <!--handleClose (key, keyPath) {-->
        <!--console.log(key, keyPath)-->
      <!--}-->
    <!--}-->
  <!--}-->
<!--//</script>-->
