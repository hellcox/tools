<template>
  <div>

    <el-row style="margin-bottom: 10px">
      <el-button-group>
        <el-button type="primary" size="medium" @click="dialog.history=true">历史</el-button>
        <el-button type="primary" size="medium" @click="doDecode" v-if="show.encode">解密</el-button>
        <el-button type="primary" size="medium" @click="doEncode">加密</el-button>
      </el-button-group>
      <el-input type="input" v-model="input" v-if="show.password" placeholder="请输入秘钥" style="width: 280px;margin-left: 30px;" size="medium"></el-input>
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
        <el-table-column property="action" label="动作" width="120px"></el-table-column>
        <el-table-column property="in" label="输入" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column property="out" label="输出"  :show-overflow-tooltip="true"></el-table-column>
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
      input: '',
      output: '',
      password: '',
      dialog: {
        history: false,
      },
      history:[
        // {in:111,out:32,action:'解析'}
      ],
      show:{
        encode:true,
        password:false,
      }
    };
  },
  props: {
    action: String
  },
  created() {
    if (this.action==='md5'){
      this.show.encode = false
    }
    if (this.action==='crc32'){
      this.show.encode = false
    }
  },
  methods: {
    getMessage: function () {
      var self = this;
      window.backend.basic(self.message).then(result => {
        self.message = result;
      });//end method
    },//end func
    fillInput: function (row) {
      this.input = row.in;
      this.output = row.out;
      this.dialog.history = false
    },//end func
    doEncode: function () {
      var self = this;
      window.backend.encode(self.action,self.input,self.password).then(result => {
        self.output = result;
        self.history.unshift({
          in:self.input,
          out:result,
          action:self.action+': 加密',
        })
      });//end method
    },//end func
    doDecode: function () {
      var self = this;
      window.backend.decode(self.action,self.input,self.password).then(result => {
        self.output = result;
        self.history.unshift({
          in:self.input,
          out:result,
          action:self.action+': 解密',
        })
      });//end method
    },//end func
  },
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
