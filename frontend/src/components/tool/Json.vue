<template>
  <div>

    <el-row style="margin-bottom: 10px">
      <el-button-group>
        <el-button type="primary" size="medium" @click="dialog.history=true">历史</el-button>
        <el-button type="primary" size="medium" @click="doRemoveEscaping">去转义</el-button>
        <el-button type="primary" size="medium" @click="compress">压缩</el-button>
        <el-button type="primary" size="medium" @click="format">格式化</el-button>
      </el-button-group>
    </el-row>

    <el-row :gutter="24">
      <el-col :span="12">
        <el-input type="textarea" style="width: 100%;height: 530px;" placeholder="请输入内容" v-model="input"></el-input>
      </el-col>
      <el-col :span="12">
        <el-input type="textarea" style="width: 100%;height: 530px;" v-model="output"></el-input>
      </el-col>
    </el-row>

    <el-dialog title="历史记录" :visible.sync="dialog.history" :fullscreen="false" width="90%" :lock-scroll="true" :modal-append-to-body="false">
      <el-table :data="history" stripe>
        <el-table-column property="action" label="动作" width="70px"></el-table-column>
        <el-table-column property="in" label="输入"></el-table-column>
        <el-table-column property="out" label="输出"></el-table-column>
        <el-table-column label="操作" width="70px">
          <template v-slot="scope">
            <el-button type="primary" size="mini" @click="fillInput(scope.row)">填充</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

  </div>
</template>

<script>

export default {
  data() {
    return {
      input: '{"key":"hello world!"}',
      output: "",
      dialog: {
        history: false,
      },
      history:[
        // {in:111,out:32,action:'解析'}
      ],
    };
  },
  methods: {
    getMessage: function () {
      var self = this;
      window.backend.basic(self.message).then(result => {
        self.message = result;
      });//end method
    },//end func
    doRemoveEscaping: function () {
      var self = this;
      window.backend.removeEscaping(self.input).then(result => {
        console.log(self.input,result)
        self.output = result;
        self.history.unshift({
          in:self.input,
          out:result,
          action:'去转义',
        })
      });//end method
    },//end func
    fillInput: function (row) {
      this.input = row.in;
      this.output = row.out;
      this.dialog.history = false
    },//end func
    format: function () {
      var self = this;
      var data = unescape(self.input.replace(/\\u/g, '%u'))
      window.backend.jsonFormat(data).then(result => {
        self.output = result;
        self.history.unshift({
          in:self.input,
          out:result,
          action:'格式化',
        })
      });//end method
    },//end func
    compress: function () {
      var self = this;
      window.backend.jsonCompress(self.input).then(result => {
        self.output = result;
        self.history.unshift({
          in:self.input,
          out:result,
          action:'格式化',
        })
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

/deep/ .el-textarea__inner {
  height: calc(100%);
}

</style>
