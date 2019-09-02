<template>
  <button class="vue-btn" type="button"
   :class="classes"
   @click="$emit('click')">
   <!--这里的emit事件可以是任何自定义事件名，
       父组件监听的事件名应和此名相同-->
    <slot>Button</slot>
  </button>
</template>  

<script>
export default {
  props: {
    // 参数：默认, primary, info, success, warning, error, ghost, dashed
    type: {
      type: String,
      default: 'default'
    },
    // 参数： 默认，large,small
    size: {
      type: String,
      default: 'normal'
    },
    // 参数：默认，circle
    shape: {
      type: String,
      default: 'default'
    },
    long: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    classes () {
      return [
        {
          ['btn-long']: this.long,
          ['btn-' + this.type]: this.type !== 'default',
          ['btn-' + this.size]: this.size !== 'default',
          ['btn-' + this.shape]: this.shape !== 'default',
        }
      ]
    }
  }
}
</script>

<style scoped lang="less">
@import '../style/common.less';
button {
  text-transform: none;
}
.vue-btn {
  display: inline-block;
  margin-bottom: 0;
  padding: 6px 15px;
  font-weight: 400;
  font-size: 12px;
  text-align: center;
  vertical-align: middle;
  -ms-touch-action: manipulation;
  touch-action: manipulation;
  cursor: pointer;
  background-image: none;
  border: 1px solid transparent;
  white-space: nowrap;
  line-height: 1.5;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  border-radius: 4px;
  transition: color .2s linear, background-color .2s linear,border .2s linear, box-shadow .2s linear;
  color: #495060;
  background-color: @color-default;
  border-color: #dddee1;
  -webkit-appearance: button;
}
.vue-btn:hover{
  color: @color-primary-hover;
  background: @color-default-hover;
  border-color: @color-primary-hover;
}
.vue-btn:focus {
  outline: 0;
}
.btn-primary {
  color: #fff;
  background: @color-primary;
  border-color: @color-primary;
}
.btn-info {
  color: #fff;
  background: @color-info;
  border-color: @color-info;
}
.btn-success {
  color: #fff;
  background: @color-success;
  border-color: @color-success;
}
.btn-warning {
  color: #fff;
  background: @color-warning;
  border-color: @color-warning;
}
.btn-error {
  color: #fff;
  background: @color-error;
  border-color: @color-error;
}
.btn-ghost {
  background: #fff;
  border-color: #dddee1;
}
.btn-dashed {
  background: #fff;
  border-style: dashed;
}
.btn-large {
  padding: 6px 15px 7px;
  font-size: 14px;
  border-radius: 4px;
}
.btn-small {
  padding: 2px 7px;
  font-size: 12px;
  border-radius: 3px;
}
.btn-circle {
  border-radius: 32px;
}

.btn-primary:hover {
  color: #fff;
  background: @color-primary-hover;
  border-color: @color-primary-hover;
}
.btn-info:hover {
  color: #fff;
  background: @color-info-hover;
  border-color: @color-info-hover;
}
.btn-success:hover {
  color: #fff;
  background: @color-success-hover;
  border-color: @color-success-hover;
}
.btn-warning:hover {
  color: #fff;
  background: @color-warning-hover;
  border-color: @color-warning-hover;
}
.btn-error:hover {
  color: #fff;
  background: @color-error-hover;
  border-color: @color-error-hover;
}

.btn-primary:active {
  color: #fff;
  background: @color-primary-active;
  border-color: @color-primary-active;
}
.btn-info:active {
  color: #f2f2f2;
  background-color: @color-info-active;
  border-color: @color-info-active;
}
.btn-success:active {
  color: #fff;
  background: @color-success-active;
  border-color: @color-success-active;
}
.btn-warning:active {
  color: #f2f2f2;
  background-color: @color-warning-active;
  border-color: @color-warning-active;
}
.btn-error:active {
  color: #fff;
  background: @color-error-active;
  border-color: @color-error-active;
}

.btn-primary:focus, .vue-btn:focus {
  box-shadow: 0 0 0 2px rgba(45,140,240,.2);
}
.btn-info:focus {
  box-shadow: 0 0 0 2px rgba(45,183,245,.2);
}
.btn-success:focus {
  box-shadow: 0 0 0 2px rgba(25,190,107,.2);
}
.btn-warning:focus {
  box-shadow: 0 0 0 2px rgba(255,153,0,.2);
}
.btn-error:focus {
  box-shadow: 0 0 0 2px rgba(237,63,20,.2);
}

.btn-long {
  width: 100%;
}
</style>


