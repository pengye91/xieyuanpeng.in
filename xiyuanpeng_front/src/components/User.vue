<template>
  <span>
    <Button shape="circle" @click="modal = true">{{ user.name }}</Button>
    <Modal
      v-model="modal"
      title="对话框标题"
      :loading="loading"
      @on-ok="asyncOK">
      <p>点击确定后，对话框将在 2秒 后关闭。</p>
    <logout></logout>
    </Modal>
  </span>
</template>
<script>
  import {mapState} from 'vuex'
  import Logout from './Logout'

  export default {
    data () {
      return {
        modal: false,
        loading: true
      }
    },
    computed: {
      ...mapState([
        'user', 'isLogin'
      ])

    },
    methods: {
      asyncOK () {
        setTimeout(() => {
          this.modal = false
        }, 2000)
      }
    },
    components: {
      Logout
    }
  }
</script>

