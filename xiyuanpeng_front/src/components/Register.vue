<template>
  <span>
    <Button type="primary" shape="circle" @click="modal = true">注册</Button>
    <Modal
      v-model="modal"
      title="注册"
      :loading="loading"
      @on-cancel=""
      @on-ok="">

    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
        <Form-item label="用户名" prop="username">
            <Input v-model="formValidate.username" placeholder="请输入用户名"></Input>
        </Form-item>
        <Form-item label="邮箱" prop="mail">
            <Input v-model="formValidate.email" placeholder="请输入邮箱"></Input>
        </Form-item>
        <Form-item label="密码" prop="password1">
            <Input v-model="formValidate.password1" type="password"
                   placeholder="输入密码"></Input>
        </Form-item>
        <Form-item label="确认密码" prop="password2">
            <Input v-model="formValidate.password2" type="password"
                   placeholder="请再次输入密码"></Input>
        </Form-item>
    </Form>
      <div slot="footer">
            <Button type="primary" @click="handleSubmit('formValidate')">注册</Button>
            <Button type="ghost" @click="handleReset('formValidate')" style="margin-left: 8px">重置</Button>
      </div>
    </Modal>
  </span>
</template>


<script>
  import {HTTP} from '../../config/http-common'

  export default {
    data () {
//      const passwordCheck = (rule, password2, callback) => {
//        if (password2 !== this.formValidate.password1) {
//          callback(new Error('两次密码输入不一致!'))
//        }
//      }

      const usernameCheck = (rule, username, callback) => {
        HTTP.get(
          `/users/auto-search?username=${username}`
        )
          .then(response => {
            if (response.status === 404) {
              callback()
            } else if (response.status === 200) {
              callback(new Error('该用户名已被注册，换一个吧'))
            }
          })
      }

      return {
        modal: false,
        loading: true,
        closable: false,
        maskClosable: false,

        formValidate: {
          username: '',
          email: '',
          password1: '',
          password2: ''
        },

        ruleValidate: {
          username: [
            {required: true, message: '用户姓名不能为空', trigger: 'blur'},
            {validator: usernameCheck, trigger: 'blur'}
          ],
          email: [
            {required: true, message: '邮箱不能为空', trigger: 'blur'},
            {type: 'email', message: '请填写正确的邮箱格式哦', trigger: 'blur'}
          ]
        }
      }
    },
    methods: {
      handleSubmit (name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.$Message.success('提交成功!')
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

