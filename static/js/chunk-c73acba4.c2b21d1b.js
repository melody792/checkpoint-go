(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-c73acba4"],{"44be":function(s,t,e){},"73cf":function(s,t,e){"use strict";e.r(t);var a=function(){var s=this,t=s.$createElement,e=s._self._c||t;return e("div",{staticClass:"main-wrapper"},[e("div",{staticClass:"signin-wrapper"},[e("div",{staticClass:"tab"},[e("router-link",{staticClass:"link",attrs:{to:"/login",title:"登录"}},[e("h4",[s._v("登录")])]),e("span"),e("router-link",{staticClass:"active link",attrs:{to:"/register",title:"注册"}},[e("h4",[s._v("注册")])])],1),e("InputComponent",{staticClass:"input",attrs:{size:"large",placeholder:"请输入用户名"},model:{value:s.username,callback:function(t){s.username="string"===typeof t?t.trim():t},expression:"username"}}),e("InputComponent",{staticClass:"input",attrs:{size:"large",type:"password",placeholder:"请输入密码"},model:{value:s.password,callback:function(t){s.password="string"===typeof t?t.trim():t},expression:"password"}}),e("InputComponent",{staticClass:"input",attrs:{size:"large",type:"password",placeholder:"请再次输入密码"},model:{value:s.confirmPassword,callback:function(t){s.confirmPassword="string"===typeof t?t.trim():t},expression:"confirmPassword"}}),e("Button",{attrs:{type:s.buttonType,disabled:s.isButtonDisabled,long:""},on:{click:s.register}},[s._v(s._s(s.buttonText))])],1)])},n=[],i=e("4328"),r=e.n(i),o={data:function(){return{buttonType:"success",buttonText:"注册",isButtonDisabled:!1,username:"",password:"",confirmPassword:"",isOnRegister:!1}},methods:{register:function(){var s=this;if(!this.isOnRegister&&this.checkoutPassword()){this.isButtonDisabled=!0,this.buttonText="请求中...",this.isOnRegister=!0;var t=this;this.$axios.post("/api/register",r.a.stringify({username:t.username,password:t.password})).then(function(t){s.buttonText="注册",s.isButtonDisabled=!1,s.isOnRegister=!1,s.$Message.success(t.data.msg),1==t.data.code?setTimeout(function(){s.$router.push({path:"/login"})},1e3):0==t.data.code&&1062==t.data.data.Number&&s.$Message.success("用户名已存在")}).catch(function(t){s.buttonText="登录",s.isButtonDisabled=!1,s.isOnLogin=!1,t.response?s.$Message.warning(t.response.status+" "+t.response.statusText):s.$Message.warning(t.message)})}},checkoutPassword:function(){return""===this.username||""===this.password||""==this.confirmPassword?(this.$Message.warning("用户名和密码不能为空"),!1):this.password===this.confirmPassword||(this.$Message.warning("两次输入的密码不一致"),!1)}}},c=o,u=(e("c10a"),e("2877")),l=Object(u["a"])(c,a,n,!1,null,"35c82fc4",null);t["default"]=l.exports},c10a:function(s,t,e){"use strict";var a=e("44be"),n=e.n(a);n.a}}]);