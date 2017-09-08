<template>
  <div style="height: 100%">
    <Blogs v-if="type === 'blog'" :tag="tag">
    </Blogs>
    <Photo v-if="type === 'photography'" :tag="tag">
    </Photo>
  </div>
</template>

<script>
  import Blogs from './Blogs'
  import Photo from './Photo'

  export default {
    components: {
      Blogs, Photo
    },
    props: {
      type: String,
      tag: String
    }
  }
</script>
