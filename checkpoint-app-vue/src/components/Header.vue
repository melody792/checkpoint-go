<template>
  <div class="header">
    <div class="content-wrapper">
      <Menu mode="horizontal" active-name="1">
        <h4>文档管理</h4>
        <Dropdown class="avatar-wrapper" @on-click="handleDropdownMenu">
          <a href="javascript:void(0)">
            <Avatar src="https://i.loli.net/2017/08/21/599a521472424.jpg" style="margin-right:6px;" />
          </a>
          <DropdownMenu slot="list">
            <DropdownItem name="toHome">
              <router-link to="/">首页</router-link>
            </DropdownItem>
            <DropdownItem name="toPersonalPage">
              <router-link to="/personalPage">个人主页</router-link>
            </DropdownItem>
            <DropdownItem name="toCreateDoc">
              <router-link to="/editDocument">创建文档</router-link>
            </DropdownItem>
            <DropdownItem name="toLogin">
              <span>登出</span>
            </DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </Menu>
    </div>
  </div>
</template>

<script>
import {removeUsername} from '../utils/utils'
export default {
  data () {
    return {

    }
  },
  methods: {
    handleDropdownMenu (name) {
      if (name === 'toLogin') {
        removeUsername()
        this.$router.push({ path: '/login' })
      }
    },
  },
}
</script>

<style lang="less" scoped>
a {
  color: #515a6e;
}
.header {
  width: 100%;
  height: 50px;
  text-align: left;
  background-color: #fff;
  box-shadow: 0 1px 3px rgba(26,26,26,.1);
  position: fixed;
  top: 0;
  left: 0;
  z-index: 98;
  .content-wrapper {
    width: 1000px;
    height: 100%;
    margin: 0 auto;
    .ivu-menu-horizontal {
      height: 50px;
      line-height: 50px;
      border-bottom: none;
    }
    .ivu-menu-horizontal.ivu-menu-light:after {
      background: #fff;
    }
    h4 {
      display: inline-block;
    }
    .avatar-wrapper {
      float: right;
      a {
        display: block;
        img {
          width: 30px;
          vertical-align: middle;
        }
      }
    }

  }
}
</style>