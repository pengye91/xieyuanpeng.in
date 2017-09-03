<template>
  <div style="height: 100%;">
    <div style="height: 94.5%; width: 100%;">
      <Row type="flex" style="height: 100%;" justify="space-between">
        <Col span="20" style="height: 100%;">
        <iframe :src="blogSrc" height="100%" width="100%" frameborder="0"></iframe>
        </Col>
        <Col span="4" class="comment-main">
        <Row type="flex">
          <Input @keydown.left.native.stop="" @keydown.right.native.stop="" type="textarea"
                 :autosize="{minRows:19, maxRows:22}" placeholder="点击评论此博文" v-model="newBlogComment" class="comment-box">
          </Input>
        </Row>
        <Row type="flex" justify="end">
          <Button type="primary" @click="commentOnBlog" :disabled="commentIsEmpty" class="comment-button">评论</Button>
        </Row>
        </Col>
      </Row>
    </div>
    <Row type="flex" justify="space-between" align="top"
         style="border-bottom: 0.5px solid lightgray; height: 5.5%; min-height: 28px; max-height: 44px">
      <Col span="4" class="description-bar">
      <Button v-if="curBlog !== undefined" type="text"
              style="height: 26px; padding: 0 0 0 13px;">
        {{curBlog.title}}
      </Button>
      </Col>
      <Col span="2" offset="13" @click.native="" class="description-bar">
      <Icon type="ios-eye" size="17" class="description-icon"></Icon>
      <span class="description-number">20</span>
      </Col>
      <Col span="2" @click.native="likeBlog" class="description-bar">
      <Icon v-if="curBlog" :type="likeIconType" :color="color" size="17" class="description-icon"></Icon>
      <span v-if="curBlog" class="description-number">{{ curBlog.like }}</span>
      </Col>
      <Col span="2" class="description-bar">
      <router-link to="#Comments">
        <Icon id="Comments" type="android-textsms" size="15" class="description-icon"></Icon>
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
      <comments v-if="curBlog" :post="curBlog.id" :comments="cComments" :path="'comments.'"
                type="blogs" style="padding-bottom: 5%"></comments>
      </Col>
    </Row>
  </div>
</template>
<script>
  import Comments from './Comments.vue'
  import {EventBus} from '../store/EventBus'
  import {config} from '../config/dev'
  import {mapState} from 'vuex'

  export default {
    data () {
      return {
        src: 0,
        iframeSrc: `${config.BASE_URL}/api/v1/html/`,
        baseUrl: `${config.IMAGE_BASE_URL}`,
        cComments: [],
        isHover: false,
        newBlogComment: '',
        likeColor: '',
        blogs: [],
        iframeType: '.html'
      }
    },
    props: ['blogPath', 'tag'],
    computed: {
      likeIconType () {
        let userIdName = {}
        userIdName[this.user.id] = this.user.name
        return this.curBlog.likedBy.some(i => {
          return i[this.user.id] !== undefined
        }) ? 'ios-heart' : 'ios-heart-outline'
      },
      color () {
        return this.curBlog.likedBy.some(i => {
          return i[this.user.id] !== undefined
        }) ? '#CE0000' : ''
      },
      leftDisabled () {
        return this.src === 0
      },
      curBlog () {
        if (this.blogs[this.src] !== undefined) {
          this.cComments = this.blogs[this.src].comments
          return this.blogs[this.src]
        }
      },
      rightDisabled () {
        return this.blogs.length === this.src + 1
      },
      blogSrc () {
        if (this.blogs[this.src] !== undefined) {
          return `${this.iframeSrc}${this.blogs[this.src].path}`
        }
      },
      commentIsEmpty () {
        return this.newBlogComment === ''
      },
      ...mapState([
        'user', 'isLogin', 'jwtToken'
      ])
    },
    mounted () {
      config.HTTP.get('/blogs/', {
        params: {
          tag: this.tag
        }}
      )
        .then(response => {
          if (response.status === 200) {
            this.blogs = response.data
          }
        })

      this.blogs.forEach((blog) => {
        if (blog.title === this.blogPath) {
          this.src = this.blogs.indexOf(blog)
        }
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
      likeBlog () {
        let userIdName = {}
        userIdName[this.user.id] = this.user.name
        if (!this.curBlog.likedBy.some(i => {
          return i[this.user.id] !== undefined
        })) {
          this.curBlog.like++
          this.curBlog.likedBy.push(userIdName)
          console.log(this.curBlog.likedBy)
          config.HTTP.put(
            `/blogs/${this.curBlog.id}/like`,
            {
              'likeType': '$push',
              'increase': 1,
              'likedBy': userIdName
            })
        } else {
          this.curBlog.like--
          this.curBlog.likedBy.splice(this.curBlog.likedBy.findIndex(i => i[this.user.id] !== undefined), 1)
          config.HTTP.put(
            `/blogs/${this.curBlog.id}/like`,
            {
              'likeType': '$pull',
              'increase': -1,
              'likedBy': userIdName
            })
        }
      },
      next () {
        this.$router.push({
          name: 'blogPath',
          params: {
            'blogPath': this.blogs[this.src + 1].title
          }
        })
      },
      pre () {
        this.$router.push({
          name: 'blogPath',
          params: {
            'blogPath': this.blogs[this.src - 1].title
          }
        })
      },
      commentOnBlog () {
        config.HTTP.post(
          `/blogs/${this.curBlog.id}/comments`,
          {
            'wordContent': this.newBlogComment,
            'byId': this.user.id,
            'byName': this.user.name,
            'comments': [],
            'internalPath': ''
          },
          {
            headers: {
              Authorization: `Bearer ${this.jwtToken}`
            }
          }
        )
          .then(response => {
            this.cComments.push(response.data)
            this.newBlogComment = ''
            this.$Message.success('评论成功')
          })
          .catch(error => {
            if (error.response.status === 401) {
              this.$Message.error(`评论失败, 你手动清理localStorage了吗？\n请登录或者刷新页面重新评论`)
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

<style scoped>
  .blog-main {
    height: 94.5%;
    display: flex;

  }

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

  .comment-main {
    display: flex;
    flex-direction: column;
    /*align-items: flex-end;*/
    margin-top: 20px;
  }

  .comment-box {
    margin-left: 10px;
  }

  .comment-button {
    margin-top: 5px;
  }

</style>
