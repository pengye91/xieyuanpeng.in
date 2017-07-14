<template>
  <div>
    <Row type="flex" class="list-group-item" align="middle" justify="space-between">
      <Col span="20" style="font-size: 13px;">
      {{ comment.wordContent }}
      {{ path }}
      </Col>
      <Col span="2">
      <Button type="text" style="font-size: 13px" @click="deleteComment">删除</Button>
      </Col>
      <Col span="2">
      <Button type="text" style="font-size: 13px" @click="clickedReply=!clickedReply">点击回复</Button>
      </Col>
    </Row>
    <Row v-if="clickedReply" type="flex" justify="space-between" align="bottom">
      <Col span="22">
      <Input v-model="replyWordContent" type="textarea" :autosize="{minRows:3, maxRows:400}" placeholder="请输入回复内容..."
             class="reply-input" :autofocus="true" @keydown.left.native.stop="" @keydown.right.native.stop=""></Input>
      </Col>
      <Col span="2">
      <Button type="primary" style="font-size: 13px" @click="realReply" :disabled="!replyWordContent">回复</Button>
      </Col>
    </Row>
    <comments v-if="comment.comments !== undefined" :picture="picture" :path="path + '.comments.'"
              class="comment-comments" :comments="comment.comments"></comments>
  </div>
</template>
<script>
  import axios from 'axios'
  import {EventBus} from '../store/EventBus'
  export default{
    name: 'comment',
    data () {
      return {
        deleted: false,
        clickedReply: false,
        replyWordContent: ''
      }
    },
    beforeCreate () {
      this.$options.components.Comments = require('./Comments.vue')
    },
    props: ['comment', 'path', 'picture'],
    computed: {},
    methods: {
      deleteComment () {
        axios.delete(`http://localhost:8000/pics/${this.picture}/comments?id=${this.comment.id}&internalPath=${this.path.slice(0, -2)}`
        )
          .then(response => {
            EventBus.$emit('delete-comment', this.comment.id)
            console.log(response.data)
            this.$Message.success('删除成功')
          })
          .catch(error => {
            console.log(error)
            this.$Message.error('删除失败')
          })
      },
      realReply () {
        axios.post(`http://localhost:8000/pics/${this.picture}/comments`, {
          'wordContent': this.replyWordContent,
          'comments': [],
          'internalPath': this.path + '.comments'
        })
          .then(response => {
            this.comment.comments.push(response.data)
            console.log(response.data)
            this.replyWordContent = ''
            this.clickedReply = false
            this.$Message.success('评论成功')
          })
          .catch(error => {
            console.log(error)
            this.$Message.error('评论失败')
          })
      }
    }
  }
</script>
<style>
  .list-group-item {
    position: relative;
    display: block;
    padding: 0.3% 0.8%;
    background-color: #fff;
    border: 1px solid #ddd
  }

  .reply-input {
    padding-left: 1%;
  }

  .comment-comments {
    padding: 0.1% 1% 0.1% 1%;
  }

</style>
