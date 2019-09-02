// 检查是否登录
export function checkoutLogin () {
  var isLogin = localStorage.getItem("_USERNAME_")
  return isLogin ? true : false
}

// 保存用户名到localstorage
export function setUsername (username) {
  localStorage.setItem('_USERNAME_', username)
}

// 删除用户名从缓存
export function removeUsername () {
  localStorage.removeItem("_USERNAME_")
}

// 时间格式化
// 规则：时间为当天，显示时分，如“今天 10：20：24”
// 时间在当年，显示月-日，如“10-12”
// 每次跨年后，上一年信息时间显示年-月-日，如“2017-09-21”
export function timeFormat (timestamp) {
  if (!timestamp || timestamp === '') {
    return "00:00:00"
  }
  var nowTimestamp = Math.floor(Date.now() / 1000)
  var intThisYear = new Date().getFullYear()
  var intThisMonth = new Date().getMonth() + 1
  var intToday = new Date().getDate()
  var date = new Date(timestamp * 1000)
  var year = date.getFullYear()
  var month = (date.getMonth() + 1) > 10 ? date.getMonth() + 1 : "0" + (date.getMonth() + 1)
  var day = date.getDate() > 10 ? date.getDate() : "0" + date.getDate()
  var hour = date.getHours() > 10 ? date.getHours() : "0" + date.getHours()
  var minute = date.getMinutes() > 10 ? date.getMinutes() : "0" + date.getMinutes()
  var second = date.getSeconds() > 10 ? date.getSeconds() : "0" + date.getSeconds()
  // 超出当前时间
  if (timestamp > nowTimestamp) {
    return "00:00:00"
  }
  // 非今年
  if (intThisYear - year >= 1) {
    return year + "-" + month + "-" + day
  }
  // 今年，非本月
  if ((intThisYear - year == 0 && intThisMonth - month >= 1) ||
      (intThisYear - year == 0 && intThisMonth - month == 0 && intToday - day >= 1)) {
    return month + "-" + day
  }
  // 今天
  return hour + ":" + minute + ":" + second  
}