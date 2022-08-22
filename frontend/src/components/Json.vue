<template>
  <div class="container">

    <el-row style="margin-bottom: 10px">
      <el-button-group>
        <el-button type="primary" size="medium" @click="dialog.history=true">历史记录</el-button>
        <el-button type="primary" size="medium">去除转义</el-button>
        <el-button type="primary" size="medium" @click="compress">压缩</el-button>
        <el-button type="primary" size="medium" @click="format">解析</el-button>
      </el-button-group>
    </el-row>

    <el-row :gutter="24">
      <el-col :span="12">
        <el-input type="textarea" style="width: 100%;height: 550px;" placeholder="请输入内容" v-model="input"></el-input>
      </el-col>
      <el-col :span="12">
        <el-input type="textarea" style="width: 100%;height: 550px;" v-model="output"></el-input>
      </el-col>
    </el-row>

    <el-dialog title="历史记录" :visible.sync="dialog.history" fullscreen="true" :modal-append-to-body="false">
      <el-table :data="history">
        <el-table-column property="in" label="输入"></el-table-column>
        <el-table-column property="out" label="输出"></el-table-column>
        <el-table-column label="操作">
          <el-button type="primary" size="mini">去除转义</el-button>
        </el-table-column>
      </el-table>
    </el-dialog>

  </div>
</template>

<script>

export default {
  data() {
    return {
      input: '{"test":"hello world!"}',
      output: "",
      dialog: {
        history: false,
      },
      history:[{in:111,out:222}],
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
