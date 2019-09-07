<template>
  <div class="home">
    <Header/>
      <!--左侧部分-->
    <div class="main-content-wrapper">
      <div class="left-wrapper">
        <Spin size="large" fix v-show="isDocContentLoading"></Spin>
        <div class="head">
          <h4>公司目录</h4>
          <Button @click="showAddModal" type="text" icon="md-add">新建用户</Button>
        </div>
        <ul class="cate-list">
<!--          <li>-->
<!--            <Icon type="md-folder" />-->
<!--            <span>国际商用机器公司(IBM)</span>-->
<!--          </li>-->
<!--          <li>-->
<!--            <Icon type="md-folder" />-->
<!--            <span>亚马逊公司{Amazon}</span>-->
<!--          </li>-->
<!--          <li>-->
<!--            <Icon type="md-folder" />-->
<!--            <span>阿里巴巴(Alibaba)</span>-->
<!--          </li>-->
          <li v-for="(item, index) in docContentsList" :key="index">
            <Icon type="md-folder" />
            <span>{{item}}</span>
          </li>
          <Button v-show="currentContentPage < totalContentPages" @click="loadMoreContents" style="margin:10px 0;" type="default" long>加载更多</Button>
          <p v-show="currentContentPage >= totalContentPages" style="margin:14px 0;text-align:center;color:#999">没有更多了</p>
        </ul>
      </div>
      <!--右侧部分-->
      <div class="right-wrapper">
        <Spin size="large" fix v-show="isDocsListLoading"></Spin>
        <div class="head">
          <h4>
            <Input v-model="searchInput"
                   placeholder="输入用户名"
                   style="width: auto">
              <Icon type="ios-search" slot="suffix" @click.prevent="getUserData1()" />
            </Input>
          </h4>

        </div>
        <ul class="doc-list">
          <Table :columns="userColumns" :data="userData"></Table>
<!--          <Button style="margin-top:20px;" type="default" long>加载更多</Button>-->
        </ul>
      </div>
    </div>
    <Modal
        v-model="isModalShow"
        :mask-closable="false"
        title="新建用户"
        @on-ok="submitAddUser"
        @on-cancel="closeModal">
        用户名：<Input v-model="userName" placeholder="输入用户名" />
        昵称：<Input v-model="nickName" placeholder="输入昵称" />
        公司名称：<Input v-model="companyName" placeholder="输入公司名称" />
        公司地址：<Input v-model="companyAddress" placeholder="输入公司地址" />
        公司电话：<Input v-model="companyPhone" placeholder="输入公司电话" />
        用户地址：<Input v-model="address" placeholder="输入用户地址" />
        用户电话：<Input v-model="phone" placeholder="输入用户电话" />
    </Modal>
  </div>
</template>

<script>
  import qs from 'qs'
  import Header from '../components/Header'

  export default {
  components: {
    Header
  },
  data () {
    return {
      isModalShow: false,
      isDocContentLoading: true,
      isDocsListLoading: true,
      currentContentPage: 1,
      totalContentPages: '',
      docContentsList: [],
      docContentName: '',
      searchInput: '',
      id: '',
      userName: '',
      nickName: '',
      companyName: '',
      companyAddress: '',
      companyPhone: '',
      address: '',
      phone: '',

      userColumns: [{
        "type": "selection",
        "align": "center",
        "width": "30",
        "className": "border-r"
      }, {
      //   "title": "用户编号",
      //   "ellipsis": false,
      //   "align": "center",
      //   "key": "id"
      // }, {
        "title": "用户名",
        "align": "center",
        "key": "username"
      }, {
        "title": "昵称",
        "align": "center",
        "key": "nickname"
      }, {
        "title": "公司名称",
        "align": "center",
        "key": "company_name"
      }, {
        "title": "公司地址",
        "align": "center",
        "key": "company_address"
      }, {
        "title": "公司电话",
        "align": "center",
        "key": "company_phone"

      }, {
        "title": "用户地址",
        "align": "center",
        "key": "address"
      }, {
        "title": "用户电话",
        "align": "center",
        "key": "phone"
      }, {
        'title': '操作',
        'align': 'center',
        'key': 'action',
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
                  this.edit(params)
                }
              }
            }, 'Edit'),
            h('Button', {
              props: {
                type: 'error',
                size: 'small'
              },
              on: {
                click: () => {
                  this.remove(params)
                }
              }
            }, 'Delete')
          ]);
        }
      }],
      userData: []
    }
  },
  created() {
    this.getUserData()
  },
  methods: {
    getUserData1() {
      this.getUserData(this.searchInput)
    },
    getUserData (username) {
      if (username == undefined) {
        username = ''
      }
      this.$axios.get('/api/users?username='+ username).then(res => {
        this.userData = res.data.data
        res.data.data.forEach(item => {
          if (item.company_name != '') {
            this.docContentsList.push(item.company_name)
          }
          var list = this.docContentsList
          this.docContentsList = Array.from(new Set(list))
        });
        this.isDocContentLoading = false
        this.isDocsListLoading = false
      })
    },
    getUserDataById (id) {
      if (id == undefined) {
        id = ''
      }
      this.$axios.get('/api/getUserById?id='+ id).then(res => {
        //res.data.data.forEach(i => {})
        this.isModalShow = true
        this.id = res.data.data.Id
        this.userName = res.data.data.Username
        this.nickName = res.data.data.Nickname
        this.companyName = res.data.data.CompanyName
        this.companyAddress = res.data.data.CompanyAddress
        this.companyPhone = res.data.data.CompanyPhone
        this.address = res.data.data.Address
        this.phone = res.data.data.Phone
      })
    },
    deleteUser (id) {
      this.$axios.delete('/api/deleteUser?id='+ id).then(res => {
        if (res.data.code == 1) {
          this.$Message.destroy()
          this.$Message.success('删除成功')
        }
      })
    },
    showAddModal () {
      this.id = ''
      this.getUserDataById()
      this.isModalShow = true
    },

    submitCatalogName () {
      this.newDocContent(this.docContentName)
    },

    closeModal () {
      this.isModalShow = false
    },

    // 加载更多
    loadMoreContents () {
      if (this.currentContentPage == this.totalContentPages) {
        return
      }
      this.currentContentPage ++
      this.getDocsContent(this.currentContentPage)
    },

    // 获取全部文档目录
    getDocsContent (page) {
      this.$axios.get('/api/docContent?page=' + page).then(res => {
        this.isDocContentLoading = false
        res.data.data.forEach(item => {
          this.docContentsList.push(item)
        });
        this.totalContentPages = res.data.totalPages
      })
    },

    // 新建用户
    submitAddUser() {
      this.addUser(this.id,this.userName, this.nickName, this.companyName, this.companyAddress, this.companyPhone, this.address, this.phone)
    },
    addUser(id,userName, nickName, companyName, companyAddress, companyPhone, address, phone) {
      this.$Message.loading({
        content: '提交中',
        duration: 0
      })
      this.$axios.put('/api/addUser', qs.stringify({id:id, userName: userName,nickName: nickName, companyName: companyName, companyAddress: companyAddress
        , companyPhone: companyPhone, address: address, phone: phone})).then(res => {
        if (res.data.code == 1) {
          this.$Message.destroy()
          this.$Message.success('请求成功')
          this.getUserData()
        } else {
          this.$Message.destroy()
          this.$Message.success('请求失败')
        }
      })
    },
    newDocContent (name) {
      this.$Message.loading({
        content: '提交中',
        duration: 0
      })
      this.$axios.post('/api/docContent', qs.stringify({name: name})).then(res => {
        if (res.data.code == 1) {
          this.$Message.destroy()
          this.$Message.success('创建成功')
        }
      })
    },
    remove (params) {
      this.deleteUser(params.row.id)
      this.userData.splice(params.index, 1);
    },
    edit (params) {
      this.getUserDataById(params.row.id)
    }
  },
}
</script>

<style lang="less" scoped>
.home {
  .main-content-wrapper {
    width: 1400px;
    height: auto;
    padding: 10px 0;
    margin: 60px auto;
    text-align: left;
    display: flex;
    flex-flow: row nowrap;
    justify-content: flex-start;
    align-items: flex-start;
    .left-wrapper {
      display: inline-block;
      width: 350px;
      min-height: 280px;
      background: #fff;
      padding: 6px 14px 10px;
      border-radius: 4px;
      box-shadow: 0 1px 1px 0 rgba(0,0,0,.1);
      position: relative;
      .head {
        height: 34px;
        line-height: 34px;
        border-bottom: 1px solid #eee;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-between;
        align-items: center;
      }
      .cate-list {
        margin-top: 12px;
        max-height: 500px;
        overflow: auto;
        li {
          width: 100%;
          font-size: 14px;
          height: 28px;
          line-height: 28px;
          cursor: pointer;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
          transition: all 0.3s ease-in-out;
          &:hover {
            color: #2d8cf0;
          }
          span {
            margin-left: 8px;
          }
        }
      }
    }
    .right-wrapper {
      display: inline-block;
      margin-left: 20px;
      width: 1000px;
      min-height: 180px;
      background: #fff;
      padding: 6px 14px 20px;
      border-radius: 4px;
      box-shadow: 0 1px 1px 0 rgba(0,0,0,.1);
      position: relative;
      .head {
        height: 34px;
        line-height: 34px;
        border-bottom: 1px solid #eee;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-between;
        align-items: center;
      }
      .doc-list {
        width: 100%;
        margin-top: 6px;
        font-size: 14px;
        li {
          padding: 10px 0;
          border-bottom: 1px solid #eee;
          a {
            font-weight: bold;
            transition: all 0.3s ease-in-out;
            color: #515a6e;
            &:hover {
              color: #2d8cf0 !important;
            }
          }
          p {
            margin-top: 10px;
            color: #777777;
          }
          .bottom {
            margin-top: 10px;
            span {
              font-size: 12px;
              margin-right: 10px;
              color: #777777;
            }
          }
        }
      }
    }
  }
}
</style>

