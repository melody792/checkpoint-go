<template>
  <ul class="page-list">

    <li 
      :class="{'disabled': currentPage === 1}"
      @click="prev">
      <button>&lt;</button>
    </li>

    <li
      v-if="(currentPage - Math.floor(showNumber / 2)) > 2"
      @click="toFirst">
      <button>1</button>
    </li>

    <li class="ellipsis"
      v-if="(currentPage - Math.floor(showNumber / 2)) > 2">
      <button>...</button>
    </li>

    <li 
      v-for="index in showPages" 
      :class="{'active': currentPage === index}"
      :key="index"
      @click="indexClick(index)">
      <button>{{ index }}</button>
    </li>

    <li class="ellipsis"
      v-if="(totalPages - currentPage -1) > Math.floor(showNumber / 2) && totalPages !== showNumber">
      <button>...</button>
    </li>

    <li 
      v-if="(totalPages - currentPage) > Math.floor(showNumber / 2) && totalPages !== showNumber"
      @click="toLast">
      <button>{{ totalPages }}</button>
    </li>

    <li 
      :class="{'disabled': totalPages === currentPage}"
      @click="next">
      <button>&gt;</button>
    </li>

  </ul>
</template>

<script>
export default {
  props: {
    current: Number,
    total: Number,
    showItems: {
      type: Number,
      default: 5
    }
  },
  data () {
    // props数据本地化，遵循vue单向数据流动原则，
    // 避免子组件修改父组件数据
    return  {
      currentPage: this.current,
      totalPages: this.total,
      showNumber: this.showItems
    }
  },
  computed: {
    showPages: function () {
      var pages = [];
      if (this.currentPage < this.showNumber) {                           // 如果当前页小于显示的条目
        var i = Math.min(this.showNumber, this.totalPages);               // 取总页数和显示条目中小的
        while (i) pages.unshift(i--);
      } else {                                                            // 如果当前页数大于要显示的条数
        var start = this.currentPage - Math.floor(this.showNumber / 2)    // 起始索引
        var i = this.showNumber;
        if (start > (this.totalPages - this.showNumber)) {                // 如果start大于总页数减显示条目
          start = (this.totalPages - this.showNumber) + 1
        }
        while (i--) pages.push(start++);
      }
      return pages;
    }
  },
  methods: {
    // 上一页
    prev: function () {
      if (this.currentPage === 1) return;
      this.currentPage --;
      this.$emit('on-change', this.currentPage);
    },

    // 首页
    toFirst: function () {
      this.currentPage = 1
      this.$emit('on-change', this.currentPage);
    },

    // 索引点击
    indexClick: function (index) {
      if (index === this.currentPage) return;
      this.currentPage = index;
      this.$emit('on-change', this.currentPage);
    },

    // 末页
    toLast: function () {
      this.currentPage = this.totalPages;
      this.$emit('on-change', this.currentPage);
    },

    // 下一页
    next: function () {
      if (this.currentPage === this.totalPages) return;
      this.currentPage ++;
      this.$emit('on-change', this.currentPage);
    },
  }
}
</script>

<style lang="less" scoped>
@import '../style/common.less';
.page-list {
  margin: 0;
  padding: 0;
  list-style: none;

  li {
    display: inline-block;
    font-size: 14px;
    margin: 0;
    padding: 0;

    button {
      display: inline-block;
      vertical-align: middle;
      padding: 0;
      width: 32px;
      height: 32px;
      line-height: 31px;
      margin-right: 4px;
      text-align: center;
      list-style: none;
      color: #777;
      background-color: #fff;
      -webkit-user-select: none;
      -moz-user-select: none;
      -ms-user-select: none;
      user-select: none;
      cursor: pointer;
      font-family: Arial;
      border: 1px solid #dddee1;
      border-radius: 4px;
      transition: border .2s ease-in-out,color .2s ease-in-out;
      outline: none;
      &:hover {
        color: @color-primary-hover;
        border-color: @color-primary-hover;
      }
    }
  }

  li.ellipsis {
    button {
      cursor: not-allowed;
      &:hover{
        border: 1px solid #dddee1;
      }
    }
  }

  li.active {
    button {
      color: #fff;
      background: @color-primary-hover;
      border-color: @color-primary-hover;
    }
  }

  li.disabled {
    button {
      cursor: not-allowed;
      color: #bbbec4;
      background-color: #f7f7f7;
      border-color: #dddee1;
      &:hover {
        color: #bbbec4;
        background-color: #f7f7f7;
        border-color: #dddee1;
      }
    }
  }
}
</style>
