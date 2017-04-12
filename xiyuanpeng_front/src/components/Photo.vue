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
      <Col span="2" >
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
  export default {
    data () {
      return {
        imgUrl: '1.jpg',
        src: 1,
        allImgs: () => {
          axios.get('http://localhost:8000/v1/pictures')
            .then(function (response) {
              console.log(response)
              return response
            })
        }
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
      axios.get('http://localhost:8000/v1/pictures')
        .then(function (response) {
          let images = []
          images = response.data
          console.log(response)
          return images
        })
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
