<template>
  <span>
    <Button shape="circle" @click="modal = true">{{ user.name }}</Button>
    <Modal
      v-model="modal"
      title="对话框标题"
      :loading="loading"
      @on-ok="asyncOK">
    <logout></logout>
    <Button v-if="user.name==='xyp'" shape="circle" @click="goToAdmin">Admin</Button>
    </Modal>
  </span>
</template>
<script>
  import {mapState} from 'vuex'

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
      },
      goToAdmin () {
        this.$router.push('/admin')
        this.modal = false
      }
    }
  }
</script>

