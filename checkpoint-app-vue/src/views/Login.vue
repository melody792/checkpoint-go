<template>
  <div class="main-wrapper">
    <div class="signin-wrapper">
      <div class="tab">
        <router-link class="active link" to="/login" title="登录">
          <h4>登录</h4>
        </router-link>
        <span></span>
        <router-link class="link" to="/register" title="注册">
          <h4>注册</h4>
        </router-link>
      </div>
      <InputComponent class="input" size="large" v-model.trim="username" placeholder="请输入用户名" />
      <InputComponent class="input" size="large" v-model.trim="password" type="password" placeholder="请输入密码" />
      <Button @click="login" :type="buttonType"  :disabled="isButtonDisabled" long>{{buttonText}}</Button>
    </div>
  </div>
</template>

<script>
import qs from 'qs'
import {setUsername} from '../utils/utils'
export default {
  data() {
    return {
      buttonType: 'primary',
      buttonText: '登录',
      isButtonDisabled: false,
      username: '',
      password: '',
      confirmPassword: '',
      isOnLogin: false, 
    }
  },
  methods: {
    login () {
      if (this.isOnLogin) return
      if (!this.checkout()) return
      this.buttonText = '登录中...'
      this.isButtonDisabled = true
      this.isOnLogin = true
      var me = this
      this.$axios.post('/api/login', qs.stringify({
        username: me.username,
        password: me.password
      })).then(res => {
        this.buttonText = '登录'
        this.isButtonDisabled = false
        this.isOnLogin = false
        this.$Message.success(res.data.msg);
        if (res.data.code == 1) {
          setUsername(this.username)
          setTimeout(() => {
            this.$router.push({ path: '/' })
          }, 1000)
        }
      }).catch(error => {
        this.buttonText = '登录'
        this.isButtonDisabled = false
        this.isOnLogin = false
        if (error.response) {
          this.$Message.warning(error.response.status + ' ' + error.response.statusText)
        } else {
          this.$Message.warning(error.message)
        }
      })
    },

    checkout () {
      if (this.username === '' || this.password === '') {
        this.$Message.warning('用户名和密码不能为空');
        return false
      }
      return true
    }
  },
}
</script>

<style scoped>
.main-wrapper {
  width: 100%;
  height: 100%;
  position: absolute;
  background-color: rgb(247, 247, 247);
  display: flex;
  display: -webkit-flex;
  justify-content: center;
  align-items: center;
  -webkit-box-align: center;
}
.signin-wrapper {
  width: 400px;
  padding: 30px 40px;
  background: #fff;
  border-radius: 6px;
  box-shadow: 0px 10px 40px -2px rgba(0, 64, 128, 0.2);
}
.signin-wrapper .tab {
  margin-bottom: 10px;
  font-size: 18px;
  color: #2c3e50;
  display: flex;
  flex-flow: row nowrap;
  justify-content: center;
  align-items: center;
}
.signin-wrapper .tab .link {
  color: #999;
  margin: 0 14px;
}
.signin-wrapper .tab .active {
  color: #2db7f5;
}
.signin-wrapper .tab span {
  width: 2px;
  height: 17px;
  background: #999;
}
.signin-wrapper .ivu-input-wrapper {
  margin: 6px 0;
}
.signin-wrapper .ivu-btn {
  margin: 6px 0;
}
.signin-wrapper .password-text {
  margin: 10px 0;
  float: right;
}
.signin-wrapper .password-text a {
  color: #ed4014;
}
.input {
  width: 100%;
  margin: 0 0 10px;
}
.submit-button {
  margin: 16px 0 0;
}
</style>