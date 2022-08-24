<template>
  <div class="container">

      <el-tabs @tab-click="handleClick">
        <el-tab-pane label="MD5"><EncodeDecode action="md5"/></el-tab-pane>
        <el-tab-pane label="Base64"><EncodeDecode action="base64"/></el-tab-pane>
        <el-tab-pane label="CRC32"><EncodeDecode action="crc32"/></el-tab-pane>
      </el-tabs>

  </div>
</template>

<script>

import EncodeDecode from './tool/EncodeDecode'
export default {
  components: {
    EncodeDecode
  },
  data() {
    return {
      message: "tmp",
    };
  },
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.basic(self.message).then(result => {
        self.message = result;
      });//end method
    },//end func
    handleClick() {
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

/deep/ .el-tabs__item{
  padding: 0 10px;
}


</style>
