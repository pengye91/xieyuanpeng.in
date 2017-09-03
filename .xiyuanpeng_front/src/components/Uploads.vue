<template>
  <div style="margin: 30px 10px">
    <div v-for="(item, index) in withSrcUploadList" :key="index" style="margin-bottom: 20px" class="dropbox">
      <Row type="flex" justify="space-around" align="middle" style="margin-top: 10px">
        <Col span="20">
        <img v-if="type==='pics'" :src="item.src" width="100%" style="max-height: 600px">
        <Card v-else :bordered="false">
          <p slot="title">{{uploadForms[index].title}}</p>
          <p class="card-description">{{uploadForms[index].description}}</p>
          <p class="card-foot">{{uploadForms[index].published_at}}</p>
        </Card>
        </Col>
        <Col span="1">
        <Button type="ghost" size="large" @click="deleteItemFromUploadList(index)">
          <Icon type="ios-trash" size="20" color="red"></Icon>
        </Button>
        </Col>
      </Row>
      <Row>
        <Form ref="uploadForm" :model="uploadForms[index]" :rules="uploadFormRules" inline style="margin-top: 5px">
          <Row type="flex" justify="space-around" align="middle">
            <Col span="6">
            <Form-item prop="title" class="form-item">
              <Input type="text" v-model="uploadForms[index].title" placeholder="标题" style="width: 100%">
              <p slot="prepend">标题</p>
              </Input>
            </Form-item>
            </Col>
            <Col span="16">
            <Form-item prop="description" class="form-item">
              <Input type="text" v-model="uploadForms[index].description" placeholder="简要描述">
              <p slot="prepend">描述</p>
              </Input>
            </Form-item>
            </Col>
          </Row>
        </Form>
      </Row>
    </div>

    <Row type="flex" justify="center" align="middle" v-if="uploadList.length!==0">
      <Col span="22">
      <Button style="width: 100%" type="primary" size="large" @click="submitAll">
        提交
      </Button>
      </Col>
    </Row>

    <Row>
      <Upload
        ref="upload"
        :multiple="true"
        :show-upload-list="false"
        :default-file-list="defaultList"
        :on-success="handleSuccess"
        :on-format-error="handleFormatError"
        :on-exceeded-size="handleMaxSize"
        :before-upload="handleBeforeUpload"
        type="drag"
        :action="`${baseUrl}/api/v1/${type}es`"
        class="upload">
        <Icon :type="this.type==='blogs'?'document-text':'camera'" size="80"></Icon>
      </Upload>
    </Row>
  </div>
</template>
<script>
  import {config} from '@/config/dev'
  import ObjectId from 'bson-objectid'
  import moment from 'moment'

  export default {
    name: 'operation-upload',
    props: [
      'post', 'sideMenu', 'type'
    ],
    data () {
      return {
        baseUrl: config.BASE_URL,
        imgBaseUrl: config.IMAGE_BASE_URL,
        defaultList: [],
        imgName: '',
        visible: true,
        modal: false,
        withSrcUploadList: [],
        modalIsVisible: [],
        uploadList: [],
        uploadForm: {
          title: '',
          description: ''
        },
        uploadForms: [],
        uploadPicMetas: [],
        uploadFormRules: {
          title: [
            {required: true, message: '标题必须要写哦', trigger: 'blur'}
          ]
        }
      }
    },
    methods: {
      deleteItemFromUploadList (index) {
        this.withSrcUploadList.splice(index, 1)
        this.uploadList.splice(index, 1)
        this.uploadForms.splice(index, 1)
        this.uploadPicMetas.splice(index, 1)
        this.modalIsVisible.splice(index, 1)
      },
      createURL (item) {
        let blob = new Blob(item)
        return window.URL.createObjectURL(blob)
      },
      handleRemove (file) {
        // 从 upload 实例删除数据
        const fileList = this.$refs.upload.fileList
        this.$refs.upload.fileList.splice(fileList.indexOf(file), 1)
      },
      handleSuccess (res, file) {
      },
      handleFormatError (file) {
        this.$Notice.warning({
          title: '文件格式不正确',
          desc: '文件 ' + file.name + ' 格式不正确，请上传 jpg 或 png 格式的图片。'
        })
      },
      handleMaxSize (file) {
        this.$Notice.warning({
          title: '超出文件大小限制',
          desc: '文件 ' + file.name + ' 太大，不能超过 2M。'
        })
      },
      handleBeforeUpload (file) {
        let reader = new FileReader()
        // This is very tricky.
        reader.addEventListener('load', () => {
          file.src = reader.result
          this.withSrcUploadList.push(file)
          let form = {
            title: file.name.split('.')[file.name.split('.').length - 2],
            description: '',
            fileName: file.name
          }
          this.uploadForms.push(form)
          // Cool enough
          let realFile = new File([file], file.name, {type: file.type})
          delete realFile.src
          this.uploadList.push(realFile)
        }, false)
        if (file) {
          reader.readAsDataURL(file)
        }
        return false
      },
      submitAll () {
        this.uploadForms.forEach(i => {
          i.comments = []
          i.id = String(ObjectId())
          i.path = i.fileName
          if (this.type === 'blogs') {
            i.tags = []
            i.tags.push(this.$route.params.sideMenu)
            i.published_at = moment().format()
          } else {
            i.project = this.$route.params.sideMenu
          }
          this.uploadPicMetas.push(i)
        })
        let data = new FormData()
        // very tricky
        Array
          .from(Array(this.uploadList.length).keys())
          .map(i => {
            data.append('pics', this.uploadList[i])
            data.append('content-type', this.uploadList[i].type)
            data.append('size', this.uploadList[i].size)
          })
        config.HTTP.post(`/upload-${this.type}`, data)
          .then(response => {
            if (response.status === 201) {
              config.HTTP.post(`/${this.type}es`, this.uploadPicMetas)
                .then(response => {
                  if (response.status === 201) {
                    this.$Notice.success({
                      title: '提交成功',
                      desc: `所有图片成功提交至${this.$route.params.sideMenu}`
                    })
                    this.uploadPicMetas = []
                    this.withSrcUploadList = []
                    this.uploadList = []
                    // be very careful with the uploadForms and uploadForm.
                    // Very tricky.
                    // And this is the most important part, not this.uploadPicMetas.
                    this.uploadForms = []
                  }
                })
                .catch(error => {
                  this.$Notice.error({
                    title: '提交失败',
                    desc: error.response.data
                  })
//                  this.uploadPicMetas = []
                  // be very careful with the uploadForms and uploadForm.
                  // Very tricky.
                  this.uploadForms = []
                })
            }
          })
          .catch(error => {
            this.$Notice.error({
              title: '提交失败',
              desc: error.response.data
            })
          })
      }
    },
    mounted () {
      this.uploadList = this.$refs.upload.fileList
    }
  }
</script>
<style>
  .dropbox {
    outline: 2px dashed grey; /* the dash box */
    outline-offset: -10px;
    background: #f8f8f8;
    color: dimgray;
    padding: 10px 10px;
    min-height: 150px; /* minimum height */
    max-height: 730px; /* maximum height */
    position: relative;
    margin: 10px 40px 20px 40px;
    /*cursor: pointer;*/
  }

  .upload {
    outline: 2px dashed grey; /* the dash box */
    outline-offset: -10px;
    /*background: #09658c;*/
    color: dimgray;
    padding: 10px 10px;
    min-height: 100px; /* minimum height */
    position: relative;
    margin-bottom: 60px;
    margin-top: 30px;
    /*cursor: pointer;*/
  }

  .form-item {
    margin-bottom: 20px;
    width: 100%;
  }

  .input-file {
    opacity: 0; /* invisible but it's there! */
    width: 100%;
    height: 200px;
    position: absolute;
    cursor: pointer;
  }

  .upload:hover {
    background: #0986ba; /* when mouse over to the drop zone, change color */
  }

  .dropbox:hover {
    background: #dbdbdb; /* when mouse over to the drop zone, change color */
  }

  .upload p {
    background-color: transparent;
    /*text-align: center;*/
  }

  .demo-upload-list {
    display: inline-block;
    width: 60%;
    height: 60px;
    text-align: center;
    line-height: 60px;
    border: 1px solid transparent;
    border-radius: 4px;
    overflow: hidden;
    background: #fff;
    position: relative;
    box-shadow: 0 1px 1px rgba(0, 0, 0, .2);
    margin-right: 4px;
  }

  .demo-upload-list img {
    width: 100%;
    height: 100%;
  }

  .demo-upload-list-cover {
    display: none;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: rgba(0, 0, 0, .6);
  }

  .demo-upload-list:hover .demo-upload-list-cover {
    display: block;
  }

  .demo-upload-list-cover i {
    color: #fff;
    font-size: 20px;
    cursor: pointer;
    margin: 0 2px;
  }
</style>
