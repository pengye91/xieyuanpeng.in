<template>
  <div>
    <Row type="flex" align="middle" justify="start">
      <Col span="2" style="display: flex; align-items: center; align-self: flex-start; flex-direction: column">
      <div>
        <router-link :to="`/users/${comment.byId}`">
          {{comment.byName}}
        </router-link>
      </div>
      <div>
        {{commentTime}}
      </div>
      </Col>
      <Col :span="isTheUser?15:20"
           style="font-size: 13px; padding: 6px 6px 6px 1%; border-right: 1px solid lightgray; border-left: 1px solid lightgray">
      {{ clickedEdit ? '' : comment.wordContent }}
      <div v-if="clickedEdit">
        <Row type="flex" align="bottom" justify="center">
          <Col span="20">
          <Input v-model="comment.wordContent" type="textarea" :autosize="{minRows:3, maxRows:400}"
                 :autofocus="true" @keydown.left.native.stop="" @keydown.right.native.stop="">
          </Input>
          </Col>
          <Col span="2">
          <Button type="primary" style="font-size: 13px" @click="realEdit"
                  :disabled="!comment.wordContent">修改
          </Button>
          </Col>
        </Row>
      </div>
      </Col>
      <Col span="2" align="center">
      <Button type="ghost" @click="clickedReply=!clickedReply">回复</Button>
      </Col>
      <Col v-if="isTheUser" span="2" align="center">
      <Button type="ghost" @click="clickedEdit=!clickedEdit">编辑</Button>
      </Col>
      <Col v-if="isTheUser" span="2" align="center">
      <Button type="error" @click="deleteComment">删除</Button>
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
    <comments v-if="(comment.comments !== null) && comment.comments.length" :post="post"
              :path="path + '.comments.'"
              class="comment-comments" :comments="comment.comments" :type="type">
    </comments>
  </div>
</template>
<script>
  //  import axios from 'axios'
  import {EventBus} from '../store/EventBus'
  import {mapState} from 'vuex'
  import moment from 'moment'
  import {config} from '../config/dev'

  moment.locale('zh-cn')

  export default{
    name: 'comment',
    data () {
      return {
        deleted: false,
        clickedReply: false,
        clickedEdit: false,
        replyWordContent: '',
        commentId: this.comment.id
      }
    },
    beforeCreate () {
      this.$options.components.Comments = require('./Comments.vue')
    },
    props: ['comment', 'path', 'post', 'type'],
    computed: {
      isTheUser () {
        return this.user.id === this.comment.byId
      },
      commentTime () {
        return moment(this.comment.publishedAt).fromNow()
      },
      ...mapState([
        'user', 'isLogin'
      ])
    },
    methods: {
      deleteComment () {
        config.HTTP.delete(
          `/${this.type}/${this.post}/comments?id=${this.commentId}&internalPath=${this.path.slice(0, -2)}`
        )
          .then(response => {
            EventBus.$emit('delete-comment', this.commentId)
            this.$Message.success('删除成功')
          })
          .catch(error => {
            this.$Message.error('删除失败:' + error.response.data)
          })
      },
      realEdit () {
        config.HTTP.put(
          `/${this.type}/${this.post}/comments`,
          {
            'wordContent': this.comment.wordContent,
            'comments': [],
            'byName': this.user.name,
            'byId': this.user.id,
            'internalPath': this.path
          })
          .then(response => {
            let changedComment = response.data
            EventBus.$emit('edit-comment', changedComment)
            this.clickedEdit = false
            this.$Message.success('编辑成功')
          })
          .catch(error => {
            console.log(error)
            this.$Message.error('编辑失败')
          })
      },
      realReply () {
        config.HTTP.post(
          `/${this.type}/${this.post}/comments`,
          {
            'wordContent': this.replyWordContent,
            'comments': [],
            'byName': this.user.name,
            'byId': this.user.id,
            'internalPath': this.path
          })
          .then(response => {
            this.comment.comments.push(response.data)
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
    display: flex;
    padding: 0.1% 0.2% 0.1% 0.5%;
    background-color: #fff;
    border: 2px solid #ddd;
    margin-bottom: 1px;
  }

  .reply-input {
    padding-left: 1%;
  }

  .comment-comments {
    position: relative;
    /*display: block;*/
    padding: 0.1%;
    margin-bottom: -1px;
    background-color: #fff;
    /*border: 1px solid #ddd*/
  }

</style>
