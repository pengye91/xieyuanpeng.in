<template>
  <div>
    <Row v-for="(comment, index) in comments" :key="comment" class="list-group-item">
      <comment :post="post" :comment="comment" :path="path + index" :type="type"></comment>
    </Row>
  </div>
</template>
<script>
  import Comment from './Comment.vue'

  export default{
    name: 'comments',
    props: ['post', 'comments', 'path', 'type'],
    data () {
      return {
        replyToComment: false
      }
    },
    mounted () {
    },
    components: {
      'comment': Comment
    }
  }
</script>
