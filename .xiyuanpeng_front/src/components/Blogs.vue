<template>
  <div style="height: 100%;">
    <div class="blog-item" v-for="b in blogsWithTag" :key="b.title">
      <router-link :to="{path: `${$route.path}/${b.title}`, params: {'blogPath': b.title}}" class="active-link">
        <Card :bordered="false">
          <p slot="title">{{b.title}}</p>
          <p class="card-description">{{b.description}}</p>
          <p class="card-foot">{{b.published_at}}</p>
        </Card>
      </router-link>
    </div>
    <div style="padding-bottom: 10%">
    </div>
  </div>
</template>
<script>
  import {config} from '../config/dev'
  import moment from 'moment'

  moment.locale('zh-cn')

  export default{
    name: 'blogs',
    props: ['tag'],
    metaInfo: {
      title: 'Blogs',
      titleTemplate: 'xieyuanpeng.com|%s'
    },
    watch: {
      '$route' (to, from) {
        config.HTTP.get('/blogs/', {
          params: {
            tag: to.path.split('/')[to.path.split('/').length - 1]
          }}
        )
          .then(response => {
            if (response.status === 200) {
              response.data.forEach(item => {
                item.published_at = moment(item.published_at).format('LLL')
              })
              this.blogsWithTag = response.data
            }
          })
          .catch(error => {
            console.log(error.data)
          })
      }
    },
    data () {
      return {
        replyToComment: false,
        blogsWithTag: [],
        show: true,
        publishedTime: ''
      }
    },
    methods: {
    },
    mounted () {
      config.HTTP.get('/blogs/', {
        params: {
          tag: this.tag
        }}
      )
        .then(response => {
          if (response.status === 200) {
            response.data.forEach(item => {
              item.published_at = moment(item.published_at).format('LLL')
            })
            this.blogsWithTag = response.data
          }
        })
        .catch(error => {
          console.log(error.data)
        })
    }
  }
</script>
<style scoped>
  .blog-item {
    background: #eee;
    padding: 5px 10px;
    margin-top: 3px;
  }

  .card-description {
    border-bottom: dashed 1px #e9eaec;
  }

  .card-foot {

  }

  .active-link {
    color: inherit;
  }
</style>
