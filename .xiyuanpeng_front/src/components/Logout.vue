<template>
  <span>
    <Button shape="circle" @click="logOut">登出</Button>
  </span>
</template>
<script>
  import {mapState, mapMutations} from 'vuex'
  import {config} from '../config/dev'

  export default {
    data () {
      return {
      }
    },
    methods: {
      ...mapMutations([
        'logout'
      ]),
      logOut () {
        config.HTTP.get(
          `/auth/logout`
        )
          .then(response => {
            console.log(response.data)
            this.$Message.success('登出成功')
            this.logout()
          })
          .catch(error => {
            console.log(error)
            this.$Message.error('登出失败')
          })
      }
    },
    computed: {
      ...mapState([
        'isLogin', 'user'
      ])
    }
  }
</script>

