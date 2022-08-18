<template>
  <div class="container">

    <el-tabs type="border-card" tab-position="left" style="height: 100%;width:100%;position: fixed">
      <el-tab-pane label="时间转换"><Time/></el-tab-pane>
      <el-tab-pane label="Json转换"><Json/></el-tab-pane>
      <el-tab-pane label="角色管理">角色管理</el-tab-pane>
      <el-tab-pane label="定时任务">定时任务补偿</el-tab-pane>
    </el-tabs>

  </div>
</template>

<script>
import Time from './Time'
import Json from './Json'
export default {
  components: {
    Time,Json
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
  height: 100%;
  /*display: block;*/
  /*position: relative;*/
}
</style>
