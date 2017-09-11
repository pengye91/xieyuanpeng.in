<template>
  <Layout></Layout>
</template>

<script>
  import Layout from './components/Layout'
  import {mapState} from 'vuex'

  export default {
    components: {Layout},
    name: 'App',
    metaInfo: {
      title: 'xiyuanpeng.com',
      titleTemplate: null
    },
    computed: {
      ...mapState([
        'menuItems', 'sideMenuItems'
      ])
    }
  }
</script>
