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
    <Row type="flex" justify="center" align="bottom" class="slider">
          <div v-for="img in sliderImgs" :key="img" class="slider-img-div">
            <img :src="baseUrl + img.path" :alt="img.path"
            class="slider-img" :class="{'is-src': img.title == src}">
          </div>
    </Row>
    </div>
</template>
<style scoped>
  .slider {
    height: 10%;
    text-align: center;
  }
  .slider-img-div {
    -webkit-transition: width 0.4s, height 0.4s; /* For Safari 3.1 to 6.0 */
    transition: width 0.4s, height 0.4s;
    height: 100%;
    text-align: center;
    width: 30px;
    position: relative;
    margin:0 2px 0 2px;
  }
  .slider:hover div {
    max-height: 100%;
    height: 5vw;
    width: 5vw;
    margin: 0 3px 0 3px;
  }
  .slider:hover img {
    height: 100%;
    width: 100%;
    margin: 0 2px 0 2px;
  }
  .slider-img {
    width: 100%;
    max-height: 90%;
    height: auto;
    position: absolute;
    left: 0;
    bottom: 0;
  }
  .slider-img:hover {
    box-shadow: 6px 6px 4px #1f3c48;
  }
  .is-src {
    bottom: 5px;
    box-shadow: 4px 4px 3px #484848;
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
  .ivu-btn-large {
    font-size: 70px;
    transform: scale(0.8, 1);
  }
</style>
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
        baseUrl: 'http://192.168.0.103:8000/static/images/',
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
      axios.get('http://192.168.0.103:8000/v1/pictures')
        .then((response) => {
          this.imgs = response.data
          for (let i in this.imgs) {
            this.urls.push(this.baseUrl + this.imgs[i].path)
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
