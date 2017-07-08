<template>
  <span>
    <Button shape="circle" @click="modal = true">注册</Button>
    <Modal
      v-model="modal"
      title="对话框标题"
      :loading="loading"
      @on-ok="asyncOK">
      <p>点击确定后，对话框将在 2秒 后关闭。</p>
    </Modal>
  </span>
</template>
<script>
  export default {
    data () {
      return {
        modal: false,
        loading: true
      }
    },
    methods: {
      asyncOK () {
        setTimeout(() => {
          this.modal = false
        }, 2000)
      }
    }
  }
</script>

