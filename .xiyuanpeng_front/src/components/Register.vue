<template>
  <span>
    <Button type="primary" shape="circle" @click="modal = true">注册</Button>
    <Modal
      v-model="modal"
      title="注册"
      :loading="loading"
      :closable="closable"
      :maskClosable="maskClosable"
      @on-cancel="modal = false"
      @on-ok="handleSubmit('registerForm')">

    <Form @keydown.left.native.stop="" @keydown.right.native.stop="" @keyup.enter.native="handleSubmit('registerForm')"
          ref="registerForm" :model="registerForm" :rules="registerRule" :label-width="80">
        <Form-item label="用户名" prop="username">
            <Input v-model="registerForm.username" placeholder="请输入用户名" :autofocus="true"></Input>
        </Form-item>
        <Form-item label="邮箱" prop="email">
            <Input v-model="registerForm.email" placeholder="请输入邮箱"></Input>
        </Form-item>
        <Form-item label="密码" prop="password1">
            <Input v-model="registerForm.password1" type="password"
                   placeholder="输入密码"></Input>
        </Form-item>
        <Form-item label="确认密码" prop="password2">
            <Input v-model="registerForm.password2" type="password"
                   placeholder="请再次输入密码"></Input>
        </Form-item>
    </Form>
      <div style="padding-left: 88%; margin:0 0">
        <Button type="ghost" @click="handleReset('registerForm')">重置</Button>
      </div>
    </Modal>
  </span>
</template>


<script>
  import {config} from '../config/dev'

  export default {
    data () {
      const passwordCheck = (rule, password2, callback) => {
        if (password2 !== this.registerForm.password1) {
          callback(new Error('两次密码输入不一致!'))
        } else {
          callback()
        }
      }

      const emailCheck = (rule, email, callback) => {
        config.HTTP.get(
          `/users/auto-search?email=${email}`
        )
          .then(response => {
            if (response.status === 204) {
              // add a validation pass callback
              callback()
            } else if (response.status === 200) {
              callback(new Error('该邮箱已被注册，换一个吧'))
            }
          })
          .catch(error => {
            callback(error)
          })
      }
      const usernameCheck = (rule, username, callback) => {
        config.HTTP.get(
          `/users/auto-search?username=${username}`
        )
          .then(response => {
            if (response.status === 204) {
              // add a validation pass callback
              callback()
            } else if (response.status === 200) {
              callback(new Error('该用户名已被注册，换一个吧'))
            } else {
              callback(response.data)
            }
          })
          .catch(error => {
            callback(error)
          })
      }

      return {
        modal: false,
        loading: true,
        closable: true,
        maskClosable: false,

        registerForm: {
          username: '',
          email: '',
          password1: '',
          password2: ''
        },

        registerRule: {
          username: [
            {required: true, message: '用户姓名不能为空', trigger: 'blur'},
            {validator: usernameCheck, trigger: 'blur'}
          ],
          email: [
            {required: true, message: '邮箱不能为空', trigger: 'blur'},
            {type: 'email', message: '请填写正确的邮箱格式哦', trigger: 'blur'},
            {validator: emailCheck, trigger: 'blur'}
          ],
          password1: [
            {required: true, message: '密码不能为空', trigger: 'blur'},
            {type: 'string', min: 6, message: '密码长度要大于等于6位哦', trigger: 'blur'}

          ],
          password2: [
            {required: true, message: '确认密码不能为空', trigger: 'blur'},
            {validator: passwordCheck, trigger: 'blur'}
          ]
        }
      }
    },
    methods: {
      handleSubmit (name) {
        this.loading = true
        this.$refs[name].validate((valid) => {
          if (valid) {
            config.HTTP.post(
              '/auth/register',
              {
                name: this.registerForm.username,
                pass: this.registerForm.password2,
                email: this.registerForm.email
              })
              .then(response => {
                if (response.status === 201) {
                  this.modal = false
                  this.$Message.success('注册成功!')
                  this.$refs[name].resetFields()
                }
              })
              .catch(error => {
                console.log(error)
                this.modal = false
                this.$Message.error('注册失败!')
              })
          } else {
            this.$Message.error('表单验证失败!')
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields()
      }
    }
  }
</script>

