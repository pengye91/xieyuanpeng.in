<template>
  <div>
    <vimg :imgs="myImages" width="500px" height="500px"></vimg>
    <h1>test</h1>
    <Row>
      <Col span="1" offset="2">
      <Button type="text" icon="chevron-left" size="large" @click="pre"></Button>
      </Col>
      <Col span="15">
      <img :src="imgSrc" :alt="src" width="60%" height="60%"
           style="box-shadow: 10px 10px 10px #484848; margin: auto; display: block">
      </Col>
      <Col span="2" offset="1">
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
    <div>哈哈</div>
    <div>哈哈</div>
    <div>哈哈</div>
  </div>
</template>
<script>
  import vueImg from 'vue-images/dist/vue-images'

  export default {
    components: {
      vimg: vueImg
    },
    data () {
      return {
        imgUrl: '1.jpg',
        src: 1
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
<style>
  @import "../../node_modules/vue-images/dist/vue-images.css";
</style>
