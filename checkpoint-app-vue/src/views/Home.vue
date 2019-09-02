<template>
  <div class="home">
    <Header/>
    <div class="main-content-wrapper">
      <!--左侧部分-->
      <div class="left-wrapper">
        <Spin size="large" fix v-show="isDocContentLoading"></Spin>
        <div class="head">
          <h4>文档目录</h4>
          <Button @click="showAddModal" type="text" icon="md-add">新建</Button>
        </div>
        <ul class="cate-list">
          <li>
            <Icon type="md-folder" />
            <span>全部文档(不限目录)</span>
          </li>
          <li v-for="(item, index) in docContentsList" :key="index">
            <Icon type="md-folder" />
            <span>{{item.name}}</span>
          </li>
          <Button v-show="currentContentPage < totalContentPages" @click="loadMoreContents" style="margin:10px 0;" type="default" long>加载更多</Button>
          <p v-show="currentContentPage >= totalContentPages" style="margin:14px 0;text-align:center;color:#999">没有更多了</p>
        </ul>
      </div>
      <!--右侧部分-->
      <div class="right-wrapper">
        <Spin size="large" fix v-show="isDocsListLoading"></Spin>
        <div class="head">
          <h4>全部文档</h4>
        </div>
        <ul class="doc-list">
          <li>
            <a>标题比套题的温暖我你的空间瓦达瓦但您翁科技枪</a>
            <p>对不起框架的那位情况电脑是空前的五年前分开为单位群你离开我带你玩起来看到那位动能武器可能得蔚蓝群岛五千年的完全离开你带我去了带你玩起来多难为情的为</p>
            <div class="bottom">
              <span>创建时间：2018-12-20</span>
              <span>更新时间：2018-12-20</span>
              <span>用户：sangshaofeng</span>
            </div>
          </li>
          <li>
            <a>标题比套题的温暖我你的空间瓦达瓦但您翁科技枪</a>
            <p>对不起框架的那位情况电脑是空前的五年前分开为单位群你离开我带你玩起来看到那位动能武器可能得蔚蓝群岛五千年的完全离开你带我去了带你玩起来多难为情的为</p>
            <div class="bottom">
              <span>创建时间：2018-12-20</span>
              <span>更新时间：2018-12-20</span>
              <span>用户：sangshaofeng</span>
            </div>
          </li>
          <Button style="margin-top:20px;" type="default" long>加载更多</Button>
        </ul>
      </div>
    </div>
    <Modal
        v-model="isModalShow"
        :mask-closable="false"
        title="新建文档目录"
        @on-ok="submitCatalogName"
        @on-cancel="closeModal">
        <Input v-model="docContentName" placeholder="输入目录名称" />
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
    }
  },
  created() {
    this.getDocsContent(this.currentContentPage)
  },
  methods: {
    showAddModal () {
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

    // 新建文档目录
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
  },
}
</script>

<style lang="less" scoped>
.home {
  .main-content-wrapper {
    width: 1000px;
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
      width: 250px;
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
      width: 730px;
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

