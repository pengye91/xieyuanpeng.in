<template>
  <div style="height: 100%;">
    <Row type="flex" style="height: 80%" justify="space-between" align="middle">
      <Col span="2" style="text-align: left">
      <Button type="text" icon="ios-arrow-left" :disabled="leftDisabled"
              size="large" @click="pre" class="pre-button"></Button>
      </Col>
      <Col span="20">
        <img :src="imgSrc" :alt="src" class="img" @click="openImg">
      </Col>
      <Col span="2">
      <Button type="text" icon="ios-arrow-right" size="large" :disabled="rightDisabled"
              @click="next" class="next-button"></Button>
      </Col>
    </Row>
    <Row type="flex" justify="center" align="middle" style="height: 16%">
      <Col v-for="img in sliderImgs" :key="img" span="2" style="margin: 0 0.5% 0 0.5%">
      <img :src="baseUrl + img.path" :alt="img.path" @click="()=>{src=Number(img.title)}"
           class="slider-img" :class="{'is-src': img.title == src}">
      </Col>
    </Row>
    <Row type="flex" justify="start" align="bottom" style="height: 0">
      <Col span="8" class="comment-box">
      <Input @keydown.left.native.stop="" @keydown.right.native.stop="" type="textarea"
             :autosize="{minRows:14, maxRows:14}" placeholder="点击评论此图片" v-model="newPicComment">
      </Input>
      <div>
        <Button type="primary" @click="commentOnPic" :disabled="commentIsEmpty" class="comment-button">评论</Button>
      </div>
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
      <Col span="2" offset="13" @click.native="" class="description-bar">
      <Icon type="ios-eye" size="17" class="description-icon"></Icon>
      <span class="description-number">20</span>
      </Col>
      <Col span="2" @click.native="likePic" class="description-bar">
      <Icon v-if="currentPic" :type="likeIconType" :color="color" size="17" class="description-icon"></Icon>
      <span v-if="currentPic" class="description-number">{{ currentPic.like }}</span>
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
    left: 6%;
    margin-top: 4px;
  }

  .slider {
    height: 16%;
  }

  .slider-img-div {
    -webkit-transition: width 0.3s, height 0.4s; /* For Safari 3.1 to 6.0 */
    transition: width 0.2s, height 0.4s;
    height: 100%;
    width: 50px;
    position: relative;
    margin: 0 4px 0 4px;
  }

  .slider-img {
    width: 98%;
    height: auto;
  }

  .is-src {
    box-shadow: 8px 8px 8px #484848;
    margin-bottom: 80%;
    width: 100%;
  }

  .img {
    box-shadow: 15px 20px 15px #484848;
    /*height: 100%;*/
    /*width: auto;*/
    width: 100%;
    max-width: 100%;
    height: auto;
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
  import {config} from '../config/dev'
  let preloader = new ImagePreloader()
  import {mapState} from 'vuex'

  export default {
    data () {
      return {
        imgUrl: '1.jpg',
        src: 1,
        images: [],
        urls: [],
        baseUrl: 'http://www.xieyuanpeng.com/static/images/',
        imgs: [],
        cComments: [],
        isHover: false,
        newPicComment: '',
        likeColor: ''
      }
    },
    computed: {
      likeIconType () {
        return this.currentPic.likedBy.includes(this.user.id) ? 'ios-heart' : 'ios-heart-outline'
      },
      color () {
        return this.currentPic.likedBy.includes(this.user.id) ? '#CE0000' : ''
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
      config.HTTP.get('/pics/')
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
      likePic () {
        if (!this.currentPic.likedBy.includes(this.user.id)) {
          this.currentPic.like++
          this.currentPic.likedBy.push(this.user.id)
          config.HTTP.put(
            `/pics/${this.currentPic.id}/like`,
            {
              'likeType': '$push',
              'increase': 1,
              'likedBy': this.user.id
            })
        } else {
          this.currentPic.like--
          this.currentPic.likedBy.splice(this.currentPic.likedBy.indexOf(this.user.id), 1)
          config.HTTP.put(
            `/pics/${this.currentPic.id}/like`,
            {
              'likeType': '$pull',
              'increase': -1,
              'likedBy': this.user.id
            })
        }
      },
      next () {
        this.rightDisabled ? null : (this.src = this.src + 1)
      },
      pre () {
        this.leftDisabled ? null : (this.src = this.src - 1)
      },
      commentOnPic () {
        config.HTTP.post(
          `/pics/${this.currentPic.id}/comments`,
          {
            'wordContent': this.newPicComment,
            'byId': this.user.id,
            'byName': this.user.name,
            'comments': [],
            'internalPath': ''
          }
        )
          .then(response => {
            this.cComments.push(response.data)
            this.newPicComment = ''
            this.$Message.success('评论成功')
          })
          .catch(error => {
            if (error.response.status === 401) {
              this.$Message.error(`评论失败\n请登录或者刷新页面重新评论`)
            }
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
