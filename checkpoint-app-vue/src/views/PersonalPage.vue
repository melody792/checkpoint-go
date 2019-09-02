<template>
  <div class="personal-page">
    <Header />
    <div class="main-content-wrapper">
      <div class="left-wrapper">
        <div class="user-info">
          <img src="../assets/logo.png" alt="">
          <p class="nickname">昵称：sangshaofeng</p>
          <Button style="margin-top:12px;">修改资料</Button>
        </div>
        <div class="user-doc">
          <ul>
            <li><Icon type="md-add" />我创建的文档</li>
            <li><Icon type="md-heart-outline" />我收藏的文档</li>
          </ul>
        </div>
      </div>
      <div class="right-wrapper">
        
      </div>
    </div>
  </div>
</template>

<script>
import Header from '../components/Header'
export default {
  components: {
    Header
  }
}
</script>

<style lang="less" scoped>
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
    .user-info {
      width: 100%;
      min-height: 200px;
      padding: 6px 14px 20px;
      border-radius: 4px;
      background: #fff;
      box-shadow: 0 1px 1px 0 rgba(0,0,0,.1);
      position: relative;
      text-align: center;
      img {
        width: 60px;
        border-radius: 50%;
        background: #eee; 
        margin-top: 20px;
      }
      .nickname {
        margin-top: 16px;
      }
    }
    .user-doc {
      margin-top: 10px;
      width: 100%;
      border-radius: 4px;
      background: #fff;
      box-shadow: 0 1px 1px 0 rgba(0,0,0,.1);
      position: relative;
      ul {
        width: 100%;
        padding: 10px 0;
        li {
          width: 100%;
          height: 36px;
          line-height: 36px;
          padding: 0 10px;
          cursor: pointer;
          &:hover {
            background: rgb(243, 243, 243);
          }
          i {
            margin-right: 5px;
          }
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
  }
}
</style>
