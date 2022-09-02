<template>
  <el-menu
      :default-active="activeIndex"
      class="lx-header"
      mode="horizontal"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      @select="handleSelect">
    <el-menu-item index="1" class="menuItem"><router-link to="/time" class="routerLink">时间转换</router-link></el-menu-item>
    <el-menu-item index="2" class="menuItem"><router-link to="/" class="routerLink">编码解码</router-link></el-menu-item>
    <el-menu-item index="3" class="menuItem"><router-link to="/about" class="routerLink">关于</router-link></el-menu-item>
    <el-menu-item index="4" class="menuItem"><router-link to="/demo" class="routerLink">示例</router-link></el-menu-item>
  </el-menu>
</template>

<script >
import { ref } from 'vue'

export default {
  setup() {
    const activeIndex = ref('1');
    const activeIndex2 = ref('1');
    const handleSelect = (key, keyPath) => {
      console.log(key, keyPath)
    };
    return {
      activeIndex2,
      activeIndex,
      handleSelect,
    }
  }
}

</script>

<style lang="scss" scoped>

.lx-header{
  height: 32px;
}
.menuItem{
  font-size: var(--el-font-size-small);
  height: 32px;
  padding: 0;
}
.menuItem:focus{
  border-bottom: 2px red solid;
}

.routerLink{
  text-decoration: none;
  display: inline-block;
  height: 32px;
  line-height: 32px;
  padding: 0 20px;
}

</style>
