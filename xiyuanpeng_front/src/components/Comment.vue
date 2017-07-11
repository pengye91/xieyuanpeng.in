<template>
  <div>
    <Row type="flex">
      <Col span="4">
      评论
      </Col>

    </Row>

    <li class="list-group-item">
      <!--{{ picture.comments[0].comment[].user.name }} said {{ comment.content }}-->
      {{ comments }}
      <!--<a @click="replyToComment = comment">Reply</a>-->
      <!--<comment-form v-if="replyToComment == comment" :comment="comment"></comment-form>-->
      <!--<comment-list v-if="collection[comment.id]" v-bind:comments="collection[comment.id]"-->
      <!--v-bind:collection="collection"></comment-list>-->
    </li>
  </div>
</template>
<script>
  //  import CommentList from './CommentList.vue'
  //  import CommentForm from './CommentForm.vue'

  export default{
    props: ['picture', 'comments'],
    data () {
      return {
        replyToComment: false
      }
    },
    mounted () {
    }
//    components: {
//      'comment-list': CommentList,
//      'comment-form': CommentForm
//    }
  }
</script>
<style>
  .list-group-item {
    position: relative;
    display: block;
    padding: 10px 15px;
    margin-bottom: -1px;
    background-color: #fff;
    border: 1px solid #ddd
  }
</style>
