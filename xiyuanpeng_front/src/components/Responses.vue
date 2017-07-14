<template>
  <div>
    <Row v-for="(response, index) in responses" :key="response" class="list-group-item">
      <comment :picture="picture" :comment="response" :path="path +'.' + 'responses.' + index"></comment>
    </Row>
  </div>
</template>
<script>
//  import Comment from './Comment.vue'
  export default{
    name: 'responses',
    props: ['responses', 'path', 'picture'],
    computed: {},
    beforeCreate: function () {
      this.$options.components.Comment = require('./Comment.vue')
    }
  }
</script>
<style>
  .responses {
    position: relative;
    display: block;
    padding-left: 30px;
    padding-top: 30px;
    margin-bottom: -1px;
    background-color: #fff;
    border: 1px solid #ddd
  }
</style>
