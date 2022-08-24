<template>
  <div class="container">

    <el-tabs type="border-card" tab-position="left" style="height: 100%;width:100%;position: fixed">
      <el-tab-pane label="时间转换"> <Time/> </el-tab-pane>
      <el-tab-pane label="加密解密"> <Decode/> </el-tab-pane>
      <el-tab-pane label="格式化"> <Format/> </el-tab-pane>
      <el-tab-pane label="编码转换"> 编码转换 </el-tab-pane>
      <el-tab-pane label="数据计算"> 数据计算 </el-tab-pane>
      <el-tab-pane label="我的收藏"> 我的收藏 </el-tab-pane>
      <el-tab-pane label="关于"> <About/> </el-tab-pane>
    </el-tabs>

  </div>
</template>

<script>
import Time from './tool/Time'
import About from './About'
import Format from './Format'
import Decode from './Decode'
export default {
  components: {
    Time,About,Format,Decode
  },
  data() {
    return {
      message: "xxx",
      input: " "
    };
  },
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.basic().then(result => {
        self.message = result;
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container /deep/ .el-tabs--left.el-tabs--border-card .el-tabs__item.is-left.is-active {
  /*border-color: #d1dbe5 transparent !important;*/
  border-bottom-color: #d1dbe5 !important;
}

/deep/ .el-tabs--left.el-tabs--border-card .el-tabs__item.is-left{
  border-right-color: #DCDFE6 !important;
}

.container /deep/ .el-tabs--left .el-tabs__header.is-left{
  margin-right: 0;
}

.container /deep/ .el-tabs--border-card>.el-tabs__content {
  margin: 0;
  padding: 0;
  height: 100vh;
  /*display: block;*/
  /*position: relative;*/
}

/deep/ .el-tab-pane{
  height: 100vh;
}
</style>
