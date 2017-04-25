<template>
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
          <MyMenuItem v-for="item in sideMenu[type]" :name="item" :key="item" :to="item">
            <Icon type="ios-book" size="16"></Icon>
            {{ item }}
          </MyMenuItem>
        </Menu>
        </Col>
        <Col span="21" style="overflow: auto; height: 100%">
        <div class="layout-content-main">
          <router-view></router-view>
        </div>
        </Col>
      </Row>
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
