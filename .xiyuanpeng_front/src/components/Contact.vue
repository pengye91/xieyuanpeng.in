<template>
  <div style="height: 100%;">
    <Row style="height: 80%">
      <Col span="2" style="height: 100%">
      </Col>
      <Col span="20" style="height: 100%; display: flex; align-items: center; justify-content: center;">
      <img :src="`${imgBase}${tag}_2D.jpg`" :alt="`${tag} 2D image`" class="img" @click="open2DImg">
      </Col>
      <Col span="2" style="height: 100%">
      </Col>
    </Row>
    <Row style="height: 16%; display: flex; align-items: center; justify-content: center;">
      <img :src="`${imgBase}${tag}_avatar.jpg`" class="avatar" :alt="`${tag} avatar image`" @click="open2DImg">
    </Row>
  </div>
</template>
<style scoped>
  .img {
    box-shadow: 15px 20px 15px #484848;
    /*height: 100%;*/
    /*width: auto;*/
    max-height: 100%;
    width: auto;
    max-width: 100%;
  }

  .avatar {
    width: 80px;
    height: 80px;
  }
</style>
<script>
  import {config} from '../config/dev'

  export default {
    data () {
      return {
        imgBase: config.IMAGE_BASE_URL
      }
    },
    metaInfo: {
      title: 'xieyuanpeng.com|contact'
    },
    props: [
      'tag'
    ],
    watch: {
      '$route' (to, from) {
      }
    },
    methods: {
      openImg () {
        window.open()
      },
      open2DImg () {
        window.open('https://github.com/pengye91')
      }
    }
  }
</script>
