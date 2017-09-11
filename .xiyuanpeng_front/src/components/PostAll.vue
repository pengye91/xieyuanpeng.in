<template>
  <div>
    <Row type="flex" justify="center" align="middle" style="margin-top: 10px;">
      <Col span="23">
      <Table border :columns="columns" :data="allTable" height="500"></Table>
      </Col>
    </Row>
  </div>
</template>
<script>
  import {config} from '@/config/dev'
  import moment from 'moment'

  moment().locale('zh-cn')

  export default {
    name: 'operation-all',
    props: [
      'post', 'sideMenu', 'type'
    ],
    watch: {
      '$route' (to, from) {
        this.innerSideMenu = to.params.sideMenu
        this.innerPost = to.params.post
        this.initTable = []
        if (this.innerPost !== from.params.post) {
          this.metaData = this.getMetadata(this.type)
        }
        if (this.metaData !== undefined) {
          this.metaData.forEach(i => {
            if (i.hasOwnProperty('tags')) {
              if (i.tags.includes(this.innerSideMenu)) { this.initTable.push(i) }
            } else if (i.project === this.innerSideMenu) { this.initTable.push(i) }
          })
        }
      }
    },
    computed: {
      allTable () {
        if ((this.initTable.length === 0) && (this.metaData !== undefined)) {
          this.metaData.forEach(i => {
            if (i.hasOwnProperty('tags')) {
              if (i.tags.includes(this.sideMenu)) { this.initTable.push(i) }
            } else if (i.project === this.sideMenu) { this.initTable.push(i) }
          })
        }
        return this.initTable
      }
    },
    data () {
      return {
        innerSideMenu: '',
        innerPost: '',
        initTable: [],
        typeMap: {},
        columns: [
          {
            type: 'selection',
            width: 60,
            align: 'center'
          },
          {
            title: '标题',
            width: 150,
            align: 'center',
            key: 'title'
          },
          {
            title: '被赞次数',
            width: 100,
            align: 'center',
            key: 'like'
          },
          {
            title: '赞同人',
            align: 'center',
            key: 'likedBy'
          },
          {
            title: '创建时间',
            align: 'center',
            width: 150,
            key: 'created_at'
          },
          {
            title: '操作',
            align: 'center',
            key: 'action',
            width: 150,
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'primary',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px'
                  },
                  on: {
                    click: () => {
                      this.show(params.index)
                    }
                  }
                }, '查看'),
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  on: {
                    click: () => {
                      this.remove(params.index)
                    }
                  }
                }, '删除')])
            }
          }
        ],
        metaData: []
      }
    },
    mounted () {
      config.HTTP.get(`/${this.type}/`)
        .then(response => {
          if (response.status === 200) {
            this.metaData = JSON.parse(JSON.stringify(response.data, (key, value) => {
              if (key === 'created_at') {
                return moment(value).format('YYYY-MM-D')
              }
              if (key === 'likedBy') {
                let newArray = []
                value.forEach(o => {
                  newArray.push(Object.values(o)[0])
                })
                return newArray
              }
              return value
            }))
            this.typeMap[this.type] = this.metaData
          }
        })
        .catch(error => {
          console.log(error.response.data)
        })
    },
    methods: {
      getMetadata (type) {
        if (this.typeMap[type] === undefined) {
          config.HTTP.get(`/${type}/`)
            .then(response => {
              if (response.status === 200) {
                this.metaData = JSON.parse(JSON.stringify(response.data, (key, value) => {
                  if (key === 'created_at') {
                    return moment(value).format('YYYY-MM-D')
                  }
                  if (key === 'likedBy') {
                    let newArray = []
                    value.forEach(o => {
                      newArray.push(Object.values(o)[0])
                    })
                    return newArray
                  }
                  return value
                }))
                this.typeMap[type] = this.metaData
              }
            })
            .catch(error => {
              console.log(error.response.data)
            })
        }
        return this.typeMap[type]
      },
      remove (index) {
        // TODO: delete in database
        this.metaData.splice(index, 1)
      },
      show (index) {
        this.$Modal.info({
          title: '用户信息'
        })
      }

    }
  }
</script>
