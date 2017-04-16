<template>
  <div style="height: 100%">
    <Row type="flex" style="height: 90%">
      <Col span="1" style="text-align: left">
      <Button type="text" icon="ios-arrow-left" :disabled="leftDisabled"
              size="large" @click="pre" class="pre-button"></Button>
      </Col>
      <Col span="22" style="height: 100%; text-align: center">
      <img :src="imgSrc" :alt="src" class="img" >
      </Col>
      <Col span="1">
      <Button type="text" icon="ios-arrow-right" size="large" :disabled="rightDisabled"
               @click="next" class="next-button"></Button>
      </Col>
    </Row>
    <Row type="flex" justify="center" align="bottom"
         class="grow">
      <Col span="1" v-for="img in sliderImgs" :key="img" style="max-height: 90%" >
      <img :src="baseUrl + img.path" :alt="img.path" class="slider-img"
           :class="{'is-src': img.title == src}">
      </Col>
    </Row>
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
        baseUrl: 'http://192.168.1.9:8000/static/images/',
        imgs: [],
        isHover: false
      }
    },
    computed: {
      leftDisabled () { return this.src === 1 },
      rightDisabled () { return this.imgs.length === this.src },
      imgSrc () {
        return this.baseUrl + this.src.toString() + '.jpg'
      },
      sliderImgs () {
        let base = 0
        if (this.src <= 4) {
          base = 4
        } else if (this.src >= this.imgs.length - 3) {
          base = this.imgs.length - 3
        } else {
          base = this.src
        }
        return this.imgs.slice(base - 4, base + 3)
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
    font-size: 70px;
    transform: scale(0.8, 1);
  }
  .img {
    box-shadow: 7px 7px 7px #484848;
    height: 100%;
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
  .is-src {
    margin: 5px;
    box-shadow: 6px 6px 4px #484848;
  }
  .grow div {
    -webkit-transition: width 0.5s, height 0.5s; /* For Safari 3.1 to 6.0 */
    transition: width 0.5s, height 0.5s;
  }
  .grow:hover div {
    height: 100%;
    padding-right: 100%;
    width: 0;
    position: relative;
    margin: 0 3px 0 3px;
  }
  .grow:hover img {
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    margin: 0 2px 0 2px;
  }
  .grow {
    height: 10%;
    text-align: center;
  }
  .slider-img {
    width: 65%;
    max-height: 90%;
  }
  .slider-img:hover {
    margin: 5px;
    box-shadow: 6px 6px 4px #1f3c48;
  }
</style>
