<template>
  <span>
    <Button shape="circle" @click="logOut">登出</Button>
  </span>
</template>
<script>
  import showdown from 'showdown'
  import {EventBus} from '../store/EventBus'
  import Modal1 from './Modal'
  import axios from 'axios'
  import router from '../router/index'
  import jwtDecode from 'jwt-decode'
  import {mapState, mapMutations} from 'vuex'
  import {HTTP} from '../config/dev'

  let converter = new showdown.Converter()
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
            localStorage.removeItem('jwtToken')
            this.logout()
          })
          .catch(error => {
            console.log(error)
            this.$Message.error('登出成功')
          })
      },
      handleSubmit (name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.value = true
            this.loading = true
            axios.post('http://localhost:8000/auth/login', {
              username: this.formInline.user,
              password: this.formInline.password
            }, {withCredentials: true})
              .then((response) => {
                if (response.status === 200) {
                  console.log(response.data)
                  this.modal = false
                  this.$Message.success('提交成功!')
                  let user
                  user = jwtDecode(response.data.token).user
                  this.login({user: user, isLogin: true})
                  localStorage.setItem('jwtToken', response.data.token)
                  router.push({name: 'wechat'})
                  console.log(this.$route.path)
                } else {
                  console.log('wrong')
                }
              })
          } else {
            this.loading = false
            this.modal = true
            this.$Message.error('表单验证失败!')
          }
        })
      }
    },
    computed: {
      search () {
        EventBus.$on('search-text', (searchText) => {
          this.searchMessage = converter.makeHtml('#' + searchText)
          console.log(this.searchMessage)
        })
        return this.searchMessage
      },
      ...mapState([
        'isLogin', 'user'
      ])
    }
  }
</script>

