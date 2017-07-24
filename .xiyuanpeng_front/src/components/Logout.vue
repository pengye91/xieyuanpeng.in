<template>
  <span>
    <Button shape="circle" @click="logOut">登出</Button>
  </span>
</template>
<script>
  import Modal1 from './Modal'
  import {mapState, mapMutations} from 'vuex'
  import {HTTP} from '../config/dev'

  export default {
    components: {
      Modal1
    },
    data () {
      return {
        modal: false,
        loading: true,
        closable: false,
        maskClosable: false,
        searchMessage: '',
        formInline: {
          user: '',
          password: ''
        },
        ruleInline: {
          user: [
            {required: true, message: '请填写用户名', trigger: 'blur'}
          ],
          password: [
            {required: true, message: '请填写密码', trigger: 'blur'},
            {type: 'string', min: 6, message: '密码长度不能小于6位', trigger: 'blur'}
          ]
        }
      }
    },
    methods: {
      ...mapMutations([
        'logout'
      ]),
      showModal () {
        this.modal = true
      },
      logOut () {
        HTTP.get(
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

