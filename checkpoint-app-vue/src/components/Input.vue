<template>
  <div class="input-wrapper">
    <input
      :class="['input-' + size]"
      :placeholder="placeholder"
      :value="value"
      :type="type"
      @input="$emit('input', $event.target.value)">
    <i
      v-if="value !== '' && clearable"
      @click="clearInput" title="清空">x</i>
  </div>
</template>

<script>
export default {
  props: {
    type: {
      type: String,
      default: 'text'
    },
    value: {
      type: [String, Number],
      default: ''
    },
    placeholder: {
      type: String,
      default: ''
    },
    clearable: {
      type: Boolean,
      default: false
    },
    // 参数： 默认，large,small
    size: {
      type: String,
      default: 'normal'
    },
  },
  methods: {
    clearInput: function () {
      this.$emit('input', '');
    }
  }
}
</script>

<style lang="less" scoped>
@import '../style/common.less';
.input-wrapper {
  display: inline-block;
  position: relative;
  vertical-align: middle;
  line-height: normal;
  position: relative;

  input {
    display: inline-block;
    width: 100%;
    height: 30px;
    line-height: 1.5;
    padding: 0 7px;
    font-size: 12px;
    box-sizing: border-box;
    border: 1px solid #dddee1;
    border-radius: 4px;
    color: #495060;
    background-color: #fff;
    background-image: none;
    position: relative;
    cursor: text;
    outline: none;
    transition: all .2s ease-in-out;
    &:hover {
      border: 1px solid @color-primary-hover;
    }
    &:hover + i {
      display: inline-block;
    }
    &:focus {
      outline: 0;
      border: 1px solid @color-primary-hover;
      box-shadow: 0 0 0 2px rgba(45,140,240,.2);
    }
  }

  .input-large {
    height: 40px;
  }

  i {
    display: none;
    width: 16px;
    height: 16px;
    line-height: 13px;
    text-align: center;
    font-size: 10px;
    font-style: normal;
    color: #eee;
    position: absolute;
    top: 7px;
    right: 8px;
    cursor: pointer;
    background: #999;
    border-radius: 50%;
    &:hover {
      background: #666;
      display: inline-block;
    }
  }
}
</style>
