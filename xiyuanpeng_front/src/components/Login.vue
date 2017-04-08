<template>
  <div style="margin-left: 25%; margin-top: 4%">
  <Form ref="formInline" :model="formInline" :rules="ruleInline" labelWidth="100">
    <Form-item prop="user" label="用户名">
      <Row>
        <Col span="10">
        <Input type="text" v-model="formInline.user" placeholder="Username">
        <Icon type="ios-person-outline" slot="prepend"></Icon>
        </Input>
        </Col>
      </Row>
    </Form-item>
    <Form-item prop="password" label="密码">
      <Row>
        <Col span="10">
        <Input type="password" v-model="formInline.password" placeholder="Password">
        <Icon type="ios-locked-outline" slot="prepend"></Icon>
        </Input>
        </Col>
      </Row>
    </Form-item>
    <Form-item>
      <Button type="primary" @click="handleSubmit('formInline')">登录</Button>
    </Form-item>
  </Form>
  </div>
</template>
<script>
  export default {
    data () {
      return {
        formInline: {
          user: '',
          password: ''
        },
        ruleInline: {
          user: [
            { required: true, message: '请填写用户名', trigger: 'blur' }
          ],
          password: [
            { required: true, message: '请填写密码', trigger: 'blur' },
            { type: 'string', min: 6, message: '密码长度不能小于6位', trigger: 'blur' }
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
      }
    }
  }
</script>
