<template>
  <div>
    <MyUpload v-if="$route.params.operation === 'upload'" :post="$route.params.post"
             :sideMenu="$route.params.sideMenu" :type="type">
    </MyUpload>
    <PostAll v-if="$route.params.operation === 'all'" :post="$route.params.post"
             :sideMenu="$route.params.sideMenu" :type="type">
    </PostAll>
  </div>
</template>

<script>
  import MyUpload from './Uploads'
  import PostAll from './PostAll'

  export default {
    components: {
      MyUpload, PostAll
    },
    props: {
      type: String
    }
  }
</script>

