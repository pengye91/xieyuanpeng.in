<template>
    <Row type="flex" style="height: 100%">
      <Col span="1" style="text-align: left">
      <Button type="text" icon="chevron-left" :disabled="leftDisabled"
              size="large" @click="pre" class="pre-button"></Button>
      </Col>
      <Col span="22" style="height: 100%; text-align: center">
      <img :src="imgSrc" :alt="src" class="img" >
      </Col>
      <Col span="1">
      <Button type="text" icon="chevron-right" size="large" :disabled="rightDisabled"
               @click="next" class="next-button"></Button>
      </Col>
    </Row>
</template>
<script>
  import axios from 'axios'
  import ImagePreloader from 'image-preloader'
  let preloader = new ImagePreloader()
  export default {
    data () {
      return {
        imgUrl: '1.jpg',
        src: 1,
        images: [],
        urls: [],
        baseUrl: 'http://192.168.1.9:8000/static/images/',
        imgs: []
      }
    },
    computed: {
      leftDisabled () { return this.src === 1 },
      rightDisabled () { return this.imgs.length === this.src },
      imgSrc () {
        return this.baseUrl + this.src.toString() + '.jpg'
      }
    },
    mounted () {
      axios.get('http://192.168.1.9:8000/v1/pictures')
        .then((response) => {
          this.imgs = response.data
          for (let i in this.imgs) {
            this.urls.push(this.baseUrl + this.imgs[i].path)
          }
//          preloader.onProgress = function (info) {
//            console.log('image with source %s is loaded with status %s', info.value.src, info.status)
//          }
          preloader.preload(this.urls)
            .then(function (status) {
              console.log('all done!', status)
            })
        }
        )
    },
    methods: {
      next () {
        this.src = this.src + 1
      },
      pre () {
        this.src = this.src - 1
      }
    }
  }
</script>
<style scoped>
  .ivu-btn-large {
    font-size: 50px;
  }
  .img {
    box-shadow: 7px 7px 7px #484848;
    height: 95%;
    width: auto;
    max-width: 100%;
  }
  .pre-button{
    height: 100%;
    width: 100%;
  }
  .next-button{
    height: 100%;
    width: 100%;
  }
</style>
