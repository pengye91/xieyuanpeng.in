<template>
  <span>
    <Button shape="circle" @click="showModal">登录{{user.name}}{{isLogin?"true":"false"}}</Button>
    <Modal1
      title="登录"
      :value="modal"
      :maskClosable="maskClosable"
      ok-text="登录"
      :loading="loading"
      :closable="closable"
      @on-cancel="modal = false"
      @on-ok="handleSubmit('formInline')">
      <Form @keyup.enter.native="handleSubmit('formInline')" ref="formInline" :model="formInline" :rules="ruleInline"
            inline>
        <Form-item prop="user">
        <Input type="text" v-model="formInline.user" placeholder="Username">
          <Icon type="ios-person-outline" slot="prepend"></Icon>
          </Input>
        </Form-item>
        <Form-item prop="password">
          <Input type="password" v-model="formInline.password" placeholder="Password">
          <Icon type="ios-locked-outline" slot="prepend"></Icon>
          </Input>
        </Form-item>
        <div v-html="search"></div>
      </Form>
    </Modal1>
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
        'login'
      ]),
      showModal () {
        this.modal = true
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
                  localStorage.setItem('jwt_token', response.data.token)
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

