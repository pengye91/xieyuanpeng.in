<template>
  <div>
    <Row>
      <Col span="1" offset="2">
      <Button type="text" icon="chevron-left" size="large" @click="pre" style="margin-top: 400%;"></Button>
      </Col>
      <Col span="15">
      <img :src="imgSrc" :alt="src" width="100%" height="430px"
           style="box-shadow: 7px 7px 7px #484848; margin: auto; display: block">
      </Col>
      <Col span="2">
      <Button type="text" icon="chevron-right" size="large" style="margin-top: 200%;" @click="next"></Button>
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
  export default {
    data () {
      return {
        imgUrl: '1.jpg',
        src: 1,
        images: []
      }
    },
    computed: {
      myImages () {
        return [
          {
            'imageUrl': 'http://localhost:8000/static/images/' + this.imgUrl
          },
          {
            'imageUrl': 'http://localhost:8000/static/images/2.jpg'
          }
        ]
      },
      imgSrc () {
        return 'http://localhost:8000/static/images/' + this.src.toString() + '.jpg'
      }
    },
    mounted: () => {
      let preloader = new ImagePreloader()
      let urls = []
      let baseUrl = 'http://localhost:8000/static/images/'
      let baseTestUrl = 'http://localhost:8000/test/images/第064期美女图片第'
      let baseTestUrl1 = 'http://localhost:8000/test/images/第67期第'
      let baseTestUrl2 = 'http://localhost:8000/test/images/第70期第'
      let baseTestUrl3 = 'http://localhost:8000/test/images/'
      let imgs = []
      axios.get('http://localhost:8000/v1/pictures')
        .then(function (response) {
          imgs = response.data
          for (let i in imgs) {
            urls.push(baseUrl + imgs[i].path)
          }
          let j = 1
          for (j = 1; j < 210; j++) {
            let u = ''
            let u2 = ''
            let u1 = ''
            let u3 = ''
            u = j.toString() + '张.jpg'
            u1 = j.toString() + '张.jpg'
            u2 = j.toString() + '张.jpg'
            u3 = j.toString() + '.jpg'

            urls.push(baseTestUrl + u)
            urls.push(baseTestUrl1 + u1)
            urls.push(baseTestUrl2 + u2)
            urls.push(baseTestUrl3 + u3)
          }
          preloader.onProgress = function (info) {
            console.log('image with source %s is loaded with status %s', info.value.src, info.status)
          }
          preloader.preload(urls)
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
