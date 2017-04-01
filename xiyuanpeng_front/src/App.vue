<style scoped>
  .layout {
    border: 1px solid #d7dde4;
    background: #f5f7f9;
  }

  .layout-logo {
    /*width: 100px;*/
    /*height: 30px;*/
    /*background: #5b6270;*/
    border-radius: 3px;
    float: left;
    position: relative;
    top: 12px;
    left: 40px;
  }

  .layout-search {
    width: 200px;
    /*height: 10px;*/
    /*background: #5b6270;*/
    /*border-radius: 10px;*/
    float: right;
    position: relative;
    top: 12px;
    right: 30px;
  }

  .layout-nav {
    width: 420px;
    margin: 0 auto;
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
    min-height: 200px;
    margin: 15px;
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
<template>
  <div class="layout">
    <Menu mode="horizontal" theme="light" active-name="1">
      <div class="layout-logo">
        <img src="./assets/logo.png" alt="logo" height="40" width="40">
      </div>
      <div class="layout-nav">
        <Menu-item name="1">
          技术博客
        </Menu-item>
        <Menu-item name="2">
          摄影作品
        </Menu-item>
        <Menu-item name="4">
          联系我
        </Menu-item>
      </div>
      <Input size="large" class="layout-search"
             placeholder="请输入搜索内容..." v-model="search">
        <Button slot="append" icon="search"></Button>
      </Input>
    </Menu>
    <div class="layout-content">
      <Row>
        <i-col span="3">
          <Menu active-name="1-2" :open-names="['1']">
            <Submenu name="1">
              <template slot="title">
                <Icon type="document-text"></Icon>
                技术博客
              </template>
              <Menu-item name="1-1">Python</Menu-item>
              <Menu-item name="1-2">Django</Menu-item>
              <Menu-item name="1-3">Golang</Menu-item>
              <Menu-item name="1-4">杂</Menu-item>
            </Submenu>
            <Submenu name="2">
              <template slot="title">
                <Icon type="ios-keypad"></Icon>
                摄影作品
              </template>
              <Menu-item name="2-1">项目 1</Menu-item>
              <Menu-item name="2-2">项目 2</Menu-item>
            </Submenu>
            <Submenu name="3">
              <template slot="title">
                <Icon type="ios-analytics"></Icon>
                联系我
              </template>
              <Menu-item name="3-1">
                <Icon type="social-github" size="18"></Icon>
              </Menu-item>
              <Menu-item name="3-2">wechat</Menu-item>
            </Submenu>
          </Menu>
        </i-col>
        <i-col span="19">
          <div class="layout-content-main">{{ search }}</div>
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
    }
  }
</script>

