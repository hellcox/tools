<template>
  <div class="container">

    <div style="margin-left: 20px;margin-top: 5px;">
      <div style="font-size: 14px;margin-bottom: 20px">
        <span style="color: #666;display: inline-block;width: 70px;">作者:</span>
        <span style="color: #777">Hellcox</span>
      </div>

      <div style="font-size: 14px;margin-bottom: 20px">
        <span style="color: #666;display: inline-block;width: 70px;">版本:</span>
        <span style="color: #777">0.0.1</span>
      </div>

      <div style="font-size: 14px;margin-bottom: 20px">
        <span style="color: #666;display: inline-block;width: 70px;">Build:</span>
        <span style="color: #777">2022.08</span>
      </div>
    </div>

  </div>
</template>

<script>

export default {
  data() {
    return {
      input: '{"test":"hello world!"}',
      output: "",
    };
  },
  methods: {
    getMessage: function () {
      var self = this;
      window.backend.basic(self.message).then(result => {
        self.message = result;
      });//end method
    },//end func
    format: function () {
      var self = this;
      var data = unescape(self.input.replace(/\\u/g, '%u'))
      window.backend.jsonFormat(data).then(result => {
        self.output = result;
      });//end method
    },//end func
    compress: function () {
      var self = this;
      window.backend.jsonCompress(self.input).then(result => {
        self.output = result;
      });//end method
    },//end func
  }
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

/deep/ .el-scrollbar__wrap {
  overflow-x: hidden !important;
}

.container {
  margin: 10px;
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

/deep/ .el-textarea__inner {
  height: calc(100%);
  /*padding: 5px;*/
}


.el-aside {
  padding: 0;
  margin: 0;
  background-color: #D3DCE6;
  height: 100vh;
  width: 50% !important;
  position: relative;
}

.el-main {
  padding: 0;
  margin: 0;
  height: 100%;
  width: 100% !important;
  background-color: #c3DaE0;
}

pre {
  display: block;
  background: #fff;
  margin: 0;
  width: calc(100% - 22px) !important;
  height: 540px;
  border: 1px solid #DCDFE6;
  padding: 5px 10px;
  border-radius: 4px;
  overflow: scroll;
}

pre:hover {
  border-color: rgba(192, 196, 204);
}

</style>
