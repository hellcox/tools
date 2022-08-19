<template>
  <div style="height: 100%;position: fixed;width: 100%;padding: 0;margin: 0;">

    <el-container style="height: 100%;padding: 0;margin: 0;min-width: 200px;">
      <el-aside style="height:100%;width: 50%;padding: 10px">

        <el-divider content-position="left"><span class="elline">当前时间</span></el-divider>
        <el-row class="elrow">
          <el-col :span="9">
            <el-input v-model="timeNow.timestamp" placeholder="" style="width: 96%;float: left"></el-input>
          </el-col>
          <el-col :span="6" style="text-align: center">
            <el-button type="primary" plain style="width: 100%" @click="ExecTimer">启停</el-button>
          </el-col>
          <el-col :span="9">
            <el-input v-model="timeNow.date" placeholder="" style="width: 96%;float: right"></el-input>
          </el-col>
        </el-row>

        <el-divider content-position="left"><span class="elline">戳值 <i class="el-icon-right"></i> 日期</span></el-divider>
        <el-row class="elrow">
          <el-col :span="9">
            <el-input v-model="time2Date.in" placeholder="" style="width: 96%;float: left"></el-input>
          </el-col>
          <el-col :span="6" style="text-align: center">
            <el-button type="primary" plain style="width: 100%" @click="TimeToDate">转换</el-button>
          </el-col>
          <el-col :span="9">
            <el-input v-model="time2Date.out" placeholder="" style="width: 96%;float: right"></el-input>
          </el-col>
        </el-row>

        <el-divider content-position="left"><span class="elline">日期 <i class="el-icon-right"></i> 戳值</span></el-divider>
        <el-row class="elrow">
          <el-col :span="9">
            <el-input v-model="date2Time.in" placeholder="" style="width: 96%;float: left"></el-input>
          </el-col>
          <el-col :span="6" style="text-align: center">
            <el-button type="primary" plain style="width: 100%" @click="DateToTime">转换</el-button>
          </el-col>
          <el-col :span="9">
            <el-input v-model="date2Time.out" placeholder="" style="width: 96%;float: right"></el-input>
          </el-col>
        </el-row>


      </el-aside>
      <el-main style="height:100%;width: 50%;">
        <p v-for="(item,i) in history" v-bind:key="i" style="font-size: 14px;color: #666">
          <span style="width: 25px;display: inline-block">{{history.length-i-1}}:</span>
          <span style=";width: 150px;display: inline-block" v-if="item[0]!=''">{{item[0]}}</span>
          <span style="">{{item[1]}}</span>
        </p>
      </el-main>
    </el-container>

  </div>
</template>

<script>
export default {
  data() {
    return {
      history: [['历史记录...','']],
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
      timeNow: {
        timestamp: '',
        date: '',
      },
      timer: '',
    };
  },
  methods: {
    GetMessage: function() {
      var self = this;
      window.backend.basic("xxx").then(result => {
        self.input = result;
      });
    },//end
    TimeToDate: function() {
      var self = this;
      window.backend.time2Date(self.time2Date.in).then(result => {
        self.time2Date.out = result;
        this.history.unshift([self.time2Date.in,result])
      });
    },//end
    DateToTime: function() {
      var self = this;
      window.backend.date2Time(self.date2Time.in).then(result => {
        self.date2Time.out = result;
        this.history.unshift([self.date2Time.in,result])
      });
    },//end
    NowTime: function() {
      var self = this;
      window.backend.getNowTimestamp().then(result => {
        self.timeNow.timestamp = result;
        window.backend.time2Date(result).then(result1 => {
          self.timeNow.date = result1;
        });
      });
    },//end
    ExecTimer: function() {
      var self = this;
      if (self.timer===''){
        self.NowTime()
        self.timer = setInterval(() => {
          self.NowTime()
        },1000)
      }else{
        clearTimeout(self.timer)
        self.timer = ''
      }
    },//end
  },//end
  created() {
    this.NowTime()
    this.ExecTimer()
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.el-aside {
  /*background-color: #D3DCE6;*/
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
  margin-bottom: 50px;
}

.elline{
  padding-right: 10px !important;
  padding-left: 0 !important;
  margin-left: 0 !important;
  color: #888 !important;
  font-size: 12px !important;
}

.el-divider__text.is-left{
  left: 0 !important;
}

.el-divider__text{
  padding: 0 !important;
}

</style>
