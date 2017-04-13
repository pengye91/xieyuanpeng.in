<template>
  <div>
    <Row>
      <Col span="1" offset="2">
      <Button type="text" icon="chevron-left" :disabled="leftDisabled"
              size="large" @click="pre" style="margin-top: 400%;"></Button>
      </Col>
      <Col span="15">
      <img :src="imgSrc" :alt="src" width="100%" height="430px"
           style="box-shadow: 7px 7px 7px #484848; margin: auto; display: block">
      </Col>
      <Col span="2">
      <Button type="text" icon="chevron-right" size="large" :disabled="rightDisabled"
              style="margin-top: 200%;" @click="next"></Button>
      </Col>
    </Row>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
  </div>
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
        baseUrl: 'http://localhost:8000/static/images/',
        imgs: []
      }
    },
    computed: {
      myImages () {
        return [
          {
            'imageUrl': this.baseUrl + this.imgUrl
          },
          {
            'imageUrl': 'http://localhost:8000/static/images/2.jpg'
          }
        ]
      },
      leftDisabled () { return this.src === 1 },
      rightDisabled () { return this.imgs.length === this.src },
      imgSrc () {
        return this.baseUrl + this.src.toString() + '.jpg'
      }
    },
    mounted () {
      axios.get('http://localhost:8000/v1/pictures')
        .then((response) => {
          this.imgs = response.data
          console.log(this.imgs)
          for (let i in this.imgs) {
            this.urls.push(this.baseUrl + this.imgs[i].path)
          }
          preloader.onProgress = function (info) {
            console.log('image with source %s is loaded with status %s', info.value.src, info.status)
          }
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
    font-size: 30px;
  }
</style>
