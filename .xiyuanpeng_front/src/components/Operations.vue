<template>
  <div>
    <Uploads v-if="$route.params.operation === 'upload'" :post="$route.params.post"
             :sideMenu="$route.params.sideMenu">
    </Uploads>
    <PostAll v-if="$route.params.operation === 'all'" :post="$route.params.post"
             :sideMenu="$route.params.sideMenu"></PostAll>
  </div>
</template>

<script>
  import Uploads from './Uploads'
  import PostAll from './PostAll'

  export default {
    components: {
      Uploads, PostAll
    },
    props: {
      type: String
    }
  }
</script>

