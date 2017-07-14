<template>
  <div>
    <Row v-for="(comment, index) in comments" :key="comment">
      <comment :picture="picture" :comment="comment" :path="path + index"></comment>
    </Row>
  </div>
</template>
<script>
  import Comment from './Comment.vue'

  export default{
    name: 'comments',
    props: ['picture', 'comments', 'path'],
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
<style>
  .list-group-item {
    position: relative;
    display: block;
    padding: 10px 15px;
    margin-bottom: -1px;
    background-color: #b8fffa;
    border: 1px solid #ddd
  }
</style>
