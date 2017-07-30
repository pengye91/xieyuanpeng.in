<template>
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
        <MyMenuItem v-for="item in sideMenu[currentPage]" :name="item" :key="item" :to="item">
          <Icon type="ios-book" size="16"></Icon>
          {{ item }}
          </MyMenuItem>
      </Menu>
      </Col>
      <Col span="21">
      <div class="layout-content-main">
        <router-view></router-view>
      </div>
      </Col>
    </Row>
  </div>
</template>

<script>

</script>
