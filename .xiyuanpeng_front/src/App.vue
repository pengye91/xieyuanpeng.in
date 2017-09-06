<template>
  <Layout></Layout>
</template>

<script>
  import Layout from './components/Layout'
  import {mapState} from 'vuex'

  export default {
    components: {Layout},
    computed: {
      ...mapState([
        'menuItems', 'sideMenuItems'
      ])
    },
    beforeCreate: function () {
      console.log('xixi')
      this.$store.dispatch('LOAD_MENU_ITEMS')
    }
  }
</script>
