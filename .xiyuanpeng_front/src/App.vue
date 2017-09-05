<template>
  <Layout></Layout>
</template>

<script>
  import Layout from './components/Layout'

  export default {
    components: {Layout},
    mounted: function () {
      this.$store.dispatch('LOAD_MENU_ITEMS')
    }
  }
</script>
