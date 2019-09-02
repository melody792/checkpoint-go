<template>
  <div class="main-wrapper">
    <div class="signin-wrapper">
      <div class="tab">
        <router-link class="link" to="/login" title="登录">
          <h4>登录</h4>
        </router-link>
        <span></span>
        <router-link class="active link" to="/register" title="注册">
          <h4>注册</h4>
        </router-link>
      </div>
      <InputComponent class="input" size="large" v-model.trim="username" placeholder="请输入用户名" />
      <InputComponent class="input" size="large" v-model.trim="password" type="password" placeholder="请输入密码" />
      <InputComponent class="input" size="large" v-model.trim="confirmPassword" type="password" placeholder="请再次输入密码" />
      <Button @click="register" :type="buttonType"  :disabled="isButtonDisabled" long>{{buttonText}}</Button>
    </div>
  </div>
</template>

<script>
import qs from 'qs'
export default {
  data() {
    return {
      buttonType: 'success',
      buttonText: '注册',
      isButtonDisabled: false,
      username: '',
      password: '',
      confirmPassword: '',
      isOnRegister: false,    // 是否正在请求中
    }
  },
  methods: {
    register() {
      if (this.isOnRegister) return
      if (!this.checkoutPassword()) return
      this.isButtonDisabled = true
      this.buttonText = "请求中..."
      this.isOnRegister = true
      var me = this
      this.$axios.post('/api/register', qs.stringify({
        username: me.username,
        password: me.password
      })).then(res => {
        this.buttonText = '注册'
        this.isButtonDisabled = false
        this.isOnRegister = false
        this.$Message.success(res.data.msg)
        if (res.data.code == 1) {
          setTimeout(() => {
            this.$router.push({ path: '/login' })
          }, 1000)
        } else if (res.data.code == 0) {
          if (res.data.data.Number == 1062) {
            this.$Message.success('用户名已存在')
          }
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

    checkoutPassword () {
      if (this.username === '' || this.password === '' || this.confirmPassword == '') {
        this.$Message.warning('用户名和密码不能为空');
        return false
      }
      if (this.password !== this.confirmPassword) {
        this.$Message.warning('两次输入的密码不一致');
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