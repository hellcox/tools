<template>
  <div class="container">

      <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="JSON" name="first"><Time/></el-tab-pane>
        <el-tab-pane label="MD5" name="second"><Json/></el-tab-pane>
        <el-tab-pane label="Base64" name="third"><About/></el-tab-pane>
      </el-tabs>

  </div>
</template>

<script>

import Time from './tool/Time'
import Json from './tool/Json'
import About from './About'
export default {
  components: {
    Time,Json,About
  },
  data() {
    return {
      list:[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,],
      message: "tmp",
      activeName: 'first',
    };
  },
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.basic(self.message).then(result => {
        self.message = result;
      });//end method
    },//end func
    handleClick(tab, event) {
      console.log(tab, event);
    },
  }
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container{
  margin: 0 10px;
  display: block;
  overflow: hidden;
  /*background: rgba(250, 250, 250, 1);*/

  /*宽度撑破容器使用下列样式*/
  width: 98%; /*不支持calc()的浏览器*/
  width: -moz-calc(100% - 20px);
  width: -webkit-calc(100% - 20px);
  width: calc(100% - 20px);
  height: 90% !important; /*不支持calc()的浏览器*/
  height: -moz-calc(100% - 20px) !important;
  height: -webkit-calc(100% - 20px) !important;
  height: calc(100% - 20px) !important;
}


</style>
