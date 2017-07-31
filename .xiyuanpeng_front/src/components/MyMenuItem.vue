<template>
  <router-link tag="li" activeClass="ivu-menu-item-active ivu-menu-item-selected" :to="to"
               :class="classes" @click.stop="handleClick">
    <slot></slot>
  </router-link>
</template>
<script>
  import Emitter from 'iview/src/mixins/emitter'
  const prefixCls = 'ivu-menu'

  export default {
    name: 'MyMenuItem',
    mixins: [Emitter],
    props: {
      name: {
        type: [String, Number],
        required: false
      },
      to: {
        type: [String, Object],
        default: '/'
      },
      disabled: {
        type: Boolean,
        default: false
      }
    },
    data () {
      return {
        active: false
      }
    },
    computed: {
      classes () {
        return [
          `${prefixCls}-item`,
          {
            [`${prefixCls}-item-active`]: this.active,
            [`${prefixCls}-item-selected`]: this.active,
            [`${prefixCls}-item-disabled`]: this.disabled
          }
        ]
      }
    },
    methods: {
      handleClick () {
        if (this.disabled) return

        let parent = this.$parent
        let name = parent.$options.name
        while (parent && (!name || name !== 'Submenu')) {
          parent = parent.$parent
          if (parent) name = parent.$options.name
        }
        console.log(1)
        if (parent) {
          this.dispatch('Submenu', 'on-menu-item-select', this.name)
        } else {
          this.dispatch('Menu', 'on-menu-item-select', this.name)
        }
      }
    },
    mounted () {
      this.$on('on-update-active-name', (name) => {
        if (this.name === name) {
          this.active = true
          this.dispatch('Submenu', 'on-update-active-name', true)
        } else {
          this.active = false
        }
      })
    }
  }
</script>
