<template>
  <div style="height: 100%;">
    <Row type="flex" style="height: 89%">
      <Col span="1" style="text-align: left">
      <Button type="text" icon="ios-arrow-left" :disabled="leftDisabled"
              size="large" @click="pre" class="pre-button"></Button>
      </Col>
      <Col span="22" style="height: 100%; text-align: center">
      <img :src="imgSrc" :alt="src" class="img" @click="openImg">
      </Col>
      <Col span="1">
      <Button type="text" icon="ios-arrow-right" size="large" :disabled="rightDisabled"
              @click="next" class="next-button"></Button>
      </Col>
    </Row>
    <Row type="flex" justify="center" align="bottom" class="slider" style="height: 7%">
      <div v-for="img in sliderImgs" :key="img" class="slider-img-div">
        <img :src="baseUrl + img.path" :alt="img.path" @click="()=>{src=Number(img.title)}"
             class="slider-img" :class="{'is-src': img.title == src}">
      </div>
    </Row>
    <Row type="flex" justify="start" align="bottom" style="height: 0">
      <Col span="8" class="comment-box">
      <Input @keydown.left.native.stop="" @keydown.right.native.stop="" type="textarea"
             :autosize="{minRows:14, maxRows:14}" placeholder="点击评论此图片" v-model="newPicComment">
      </Input>
      <Button type="primary" @click="commentOnPic" :disabled="commentIsEmpty" class="comment-button">评论</Button>
      </Col>
    </Row>
    <Row type="flex" justify="space-between" align="top"
         style="border-bottom: 0.5px solid lightgray; height: 5.5%; min-height: 28px; max-height: 44px">
      <Col span="4" class="description-bar">
      <Button v-if="currentPic !== undefined" type="text"
              style="height: 26px; padding: 0 0 0 13px;">
        {{currentPic.title}}
      </Button>
      </Col>
      <Col span="2" offset="13" @click.native="()=>{liked=!liked}" class="description-bar">
      <Icon type="ios-eye" size="17" class="description-icon"></Icon>
      <span class="description-number">20</span>
      </Col>
      <Col span="2" @click.native="()=>{liked=!liked}" class="description-bar">
      <Icon :type="likeOrDislike" :color="Color" size="17" class="description-icon"></Icon>
      <span v-if="currentPic !== undefined" class="description-number">{{ currentPic.like }}</span>
      </Col>
      <Col span="2" class="description-bar">
      <router-link to="#picComments">
        <Icon id="picComments" type="android-textsms" size="15" class="description-icon"></Icon>
        <span v-if="cComments !== undefined" class="description-number">{{ cComments.length }}</span>
      </router-link>
      </Col>
    </Row>
    <Row type="flex" justify="start">
      <Col span="4" style="padding: 3% 20px 17px 20px; font-size: 28px">
      评论
      </Col>
    </Row>
    <Row type="flex" justify="center">
      <Col span="24">
      <comments v-if="currentPic" :picture="currentPic.id" :comments="cComments" :path="'comments.'"
                style="padding-bottom: 5%"></comments>
      </Col>
    </Row>
    <div style="height: 7%">

    </div>
  </div>
</template>
<style scoped>
  .description-bar {
    margin-top: 0.5%;
    height: 4%;
    border-top: 0.5px solid lightgray;
  }

  .description-icon {
    padding-left: 15%;
  }

  .description-number {
    padding-left: 3px;
    font-size: 16px;
  }

  .comment-box {
    position: fixed;
    left: 2px;
    width: 12.2%;
    height: 50%;
    bottom: 50px;
  }

  .comment-button {
    position: fixed;
    left: 7%;
    margin-top: 4px;
  }

  .slider {
    height: 6%;
    text-align: center;
  }

  .slider-img-div {
    -webkit-transition: width 0.3s, height 0.4s; /* For Safari 3.1 to 6.0 */
    transition: width 0.2s, height 0.4s;
    height: 100%;
    text-align: center;
    width: 25px;
    position: relative;
    margin: 0 2px 0 2px;
  }

  .slider:hover div {
    max-height: 95%;
    height: 5vw;
    width: 3vw;
    margin: 0 3px 0 3px;
  }

  .slider:hover img {
    height: 95%;
    width: 100%;
    margin: 0 2px 0 2px;
  }

  .slider-img {
    width: 100%;
    max-height: 80%;
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

  .pre-button {
    height: 100%;
    width: 100%;
  }

  .next-button {
    height: 100%;
    width: 100%;
  }

  .ivu-btn-large {
    font-size: 70px;
    transform: scale(0.8, 1);
  }
</style>
<script>
  import ImagePreloader from 'image-preloader'
  import Comments from './Comments.vue'
  import {EventBus} from '../store/EventBus'
  import {HTTP} from '../../config/http-common'
  let preloader = new ImagePreloader()
  import {mapState} from 'vuex'

  export default {
    data () {
      return {
        imgUrl: '1.jpg',
        src: 1,
        images: [],
        urls: [],
        baseUrl: 'https://s3.ap-northeast-2.amazonaws.com/xyp-s3/public/images/',
        imgs: [],
        cComments: [],
        isHover: false,
        newPicComment: '',
        liked: false,
        likeColor: ''
      }
    },
    computed: {
      likeOrDislike () {
        return this.liked ? 'ios-heart' : 'ios-heart-outline'
      },
      Color () {
        return this.liked ? '#CE0000' : ''
      },
      leftDisabled () {
        return this.src === 1
      },
      currentPic () {
        if (this.imgs[this.src - 1] !== undefined) {
          this.cComments = this.imgs[this.src - 1].comments
          return this.imgs[this.src - 1]
        }
      },
      rightDisabled () {
        return this.imgs.length === this.src
      },
      imgSrc () {
        return this.baseUrl + this.src.toString() + '.jpg'
      },
      commentIsEmpty () {
        return this.newPicComment === ''
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
      },
      ...mapState([
        'user', 'isLogin'
      ])
    },
    mounted () {
      HTTP.get('/pics/')
        .then((response) => {
          this.imgs = response.data
          for (let i in this.imgs) {
            this.urls.push(this.baseUrl + this.imgs[i].path)
          }
          preloader.preload(this.urls)
            .then(function (status) {
              console.log('all done!', status)
            })
        })
        .catch((error) => {
          console.log(error)
        })
    },
    created: function () {
      window.addEventListener('keydown', this.keyDown)
      EventBus.$on('delete-comment', (commentId) => {
        console.log(this.cComments)
        this.cComments = JSON.parse(JSON.stringify(this.cComments, (key, value) => {
          if (value.id !== commentId) {
            return value
          }
          return undefined
        }).replace(/,?null/g, '').replace(/\[,/g, '['))
        console.log(this.cComments)
      })
      EventBus.$on('edit-comment', (changedComment) => {
        this.cComments = JSON.parse(JSON.stringify(this.cComments, (key, value) => {
          if (value.id !== changedComment.id) {
            return value
          } else {
            return changedComment
          }
        }))
        console.log(this.cComments)
      })
    },
    beforeDestroy: function () {
      window.removeEventListener('keydown', this.keyDown)
    },
    methods: {
      next () {
        this.rightDisabled ? null : (this.src = this.src + 1)
      },
      pre () {
        this.leftDisabled ? null : (this.src = this.src - 1)
      },
      commentOnPic () {
        HTTP.post(
          `/pics/${this.currentPic.id}/comments`,
          {
            'wordContent': this.newPicComment,
            'byId': this.user.id,
            'byName': this.user.name,
            'comments': [],
            'internalPath': ''
          })
          .then(response => {
            this.cComments.push(response.data)
            this.newPicComment = ''
            this.$Message.success('评论成功')
          })
          .catch(error => {
            console.log(error)
            this.$Message.error('评论失败')
          })
      },
      openImg () {
        window.open(this.baseUrl + this.src + '.jpg')
      },
      keyDown (e) {
        if (e.keyCode === 39) {
          this.next()
        } else if (e.keyCode === 37) {
          this.pre()
        }
      }
    },
    components: {
      'comments': Comments
    }
  }
</script>
