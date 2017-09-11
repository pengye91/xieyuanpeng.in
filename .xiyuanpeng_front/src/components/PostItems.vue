<template>
  <div style="height: 100%">
    <Blogs v-if="type === 'blog'" :tag="tag">
    </Blogs>
    <Photo v-if="type === 'photography'" :tag="tag">
    </Photo>
    <Contact v-if="type === 'contact'" :tag="tag">
    </Contact>
  </div>
</template>

<script>
  import Blogs from './Blogs'
  import Photo from './Photo'
  import Contact from './Contact'

  export default {
    components: {
      Blogs, Photo, Contact
    },
    props: {
      type: String,
      tag: String
    }
  }
</script>
