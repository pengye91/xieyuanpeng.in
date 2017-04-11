<template>
  <div>
    <vimg :imgs="images" width="500px" height="500px"></vimg>
    <h2>test</h2>
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
        images: [
          {
            'imageUrl': 'http://localhost:8000/static/images/1.jpg'
          },
          {
            'imageUrl': 'http://localhost:8000/static/images/2.jpg'
          }
        ]
      }
    },
    computed: {
      myImages () {
        this.images = [
          {
            'imageUrl': 'http://localhost:8000/static/images/' + this.imgUrl
          },
          {
            'imageUrl': 'http://localhost:8000/static/images/2.jpg'
          }
        ]
      }

    }
  }
</script>
<style>
  @import "../../node_modules/vue-images/dist/vue-images.css";
</style>
