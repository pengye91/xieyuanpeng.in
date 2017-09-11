<template>
  <span>
    <Button shape="circle" @click="showModal">登录</Button>
    <Modal1
      title="登录"
      :value="modal"
      :maskClosable="maskClosable"
      ok-text="登录"
      :loading="loading"
      :closable="closable"
      @on-cancel="modal = false"
      @on-ok="handleSubmit('formInline')">
      <Form @keydown.left.native.stop="" @keydown.right.native.stop="" @keyup.enter.native="handleSubmit('formInline')"
            ref="formInline"
            :model="formInline" :rules="ruleInline" style="width: 80%; margin-left: 10%">
        <Form-item prop="user">
        <Input type="text" v-model="formInline.user" placeholder="用户名或邮箱" :autofocus="true">
          <Icon type="ios-person-outline" slot="prepend"></Icon>
          </Input>
        </Form-item>
        <Form-item prop="password">
          <Input type="password" v-model="formInline.password" placeholder="密码">
          <Icon type="ios-locked-outline" slot="prepend"></Icon>
          </Input>
        </Form-item>
      </Form>
      <div style="padding-left: 88%; margin:0 0">
        <Button type="ghost" @click="handleReset('formInline')">重置</Button>
      </div>
    </Modal1>
  </span>
</template>
<script>
  import showdown from 'showdown'
  import {EventBus} from '../store/EventBus'
  import Modal1 from './Modal'
  //  import axios from 'axios'
  //  import router from '../router/index'
  import jwtDecode from 'jwt-decode'
  import {mapState, mapMutations} from 'vuex'
  import {config} from '../config/dev'

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
            config.HTTP.post(
              `/auth/login`,
              {
                username: this.formInline.user,
                password: this.formInline.password
              })
              .then((response) => {
                if (response.status === 200) {
                  this.modal = false
                  this.$Message.success('登录成功!')
                  let user = jwtDecode(response.data.token).user
                  // this must be the former one
                  localStorage.setItem('jwtToken', response.data.token)
                  this.login({user: user, isLogin: true})
                } else if (response.status === 401) {
                  this.$Message.error('用户名或密码不正确')
                }
              })
              .catch(error => {
                if (error.response.status === 401) {
                  this.$Message.error('用户名或密码不正确')
                }
              })
          } else {
            this.loading = false
            this.modal = true
            this.$Message.error('表单验证失败!')
          }
        })
      },

      handleReset (name) {
        this.$refs[name].resetFields()
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

