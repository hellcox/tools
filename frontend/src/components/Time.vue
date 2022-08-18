<template>
  <div style="height: 100%;position: fixed;width: 100%;padding: 0;margin: 0;">

    <el-container style="height: 100%;padding: 0;margin: 0;min-width: 200px;">
      <el-aside style="height:100%;width: 50%;padding: 10px">

        <el-row class="elrow">
          <el-col :span="9">
            <el-input v-model="input" placeholder="" style="width: 96%;float: left"></el-input>
          </el-col>
          <el-col :span="6" style="text-align: center">
            <el-button type="primary" style="width: 100%" @click="DateToTime">当前时间</el-button>
          </el-col>
          <el-col :span="9">
            <el-input v-model="input" placeholder="" style="width: 96%;float: right"></el-input>
          </el-col>
        </el-row>

        <el-row class="elrow">
          <el-col :span="9">
            <el-input v-model="date2Time.in" placeholder="" style="width: 96%;float: left"></el-input>
          </el-col>
          <el-col :span="6" style="text-align: center">
            <el-button type="primary" style="width: 100%" @click="DateToTime">TO 时间戳</el-button>
          </el-col>
          <el-col :span="9">
            <el-input v-model="date2Time.out" placeholder="" style="width: 96%;float: right"></el-input>
          </el-col>
        </el-row>

        <el-row class="elrow">
          <el-col :span="9">
            <el-input v-model="time2Date.in" placeholder="" style="width: 96%;float: left"></el-input>
          </el-col>
          <el-col :span="6" style="text-align: center">
            <el-button type="primary" style="width: 100%" @click="TimeToDate">TO 时间</el-button>
          </el-col>
          <el-col :span="9">
            <el-input v-model="time2Date.out" placeholder="" style="width: 96%;float: right"></el-input>
          </el-col>
        </el-row>


      </el-aside>
      <el-main style="height:100%;width: 50%;">
      </el-main>
    </el-container>

  </div>
</template>

<script>
export default {
  data() {
    return {
      message: "xxx",
      input: " ",
      date2Time: {
        in: '',
        out: '',
      },
      time2Date: {
        in: '',
        out: '',
      },
    };
  },
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.basic("xxx").then(result => {
        self.input = result;
      });
    },//end
    TimeToDate: function() {
      var self = this;
      window.backend.time2Date(self.time2Date.in).then(result => {
        self.time2Date.out = result;
      });
    },//end
    DateToTime: function() {
      var self = this;
      window.backend.basic("xxx").then(result => {
        self.input = result;
      });
    },//end
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.el-aside {
  background-color: #D3DCE6;
  color: #333;
}

.el-main {
  background-color: #E9EEF3;
  color: #333;
}

.p {
  word-break: break-all;
  white-space:normal;
  word-wrap: break-word;
}

.elrow{
  margin-bottom: 20px;
}

</style>
