<template>
  <span>
    <Button type="circle" @click="modal6 = true">登录</Button>
    <Modal
      v-model="modal6"
      title="登录"
      :visible="visible"
      :loading="loading"
      maskClosable="maskClosable"
      ok-text="登录"
      :closable="closable"
      @on-ok="handleSubmit('formInline')">
      <Form ref="formInline" :model="formInline" :rules="ruleInline" inline>
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
    </Modal>
  </span>
</template>
<script>
  import LoginForm from './Blog'
  import showdown from 'showdown'
  import {EventBus} from '../store/EventBus'

  let converter = new showdown.Converter()
  export default {
    components: {
      LoginForm
    },
    data () {
      return {
        modal6: false,
        loading: true,
        closable: false,
        visible: true,
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
      asyncOK () {
        setTimeout(() => {
          this.modal6 = false
        }, 2000)
      },
      handleSubmit (name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.$Message.success('提交成功!')
            this.loading = false
          } else {
            this.loading = false
            this.modal6 = true
            this.visible = true
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
      }
    }
  }
</script>

