<template>
  <div class="content">
    <el-row class="main-row">
      <el-col :span="14">

        <el-divider content-position="left">当前时间</el-divider>
        <el-row class="input-row">
          <el-col :span="10">
            <el-input v-model="nowTime"/>
          </el-col>
          <el-col :span="4" style="text-align: center">
            <el-button type="primary" @click="OnOff">启停</el-button>
          </el-col>
          <el-col :span="10">
            <el-input v-model="nowDate"/>
          </el-col>
        </el-row>

        <el-divider content-position="left">戳值 → 时间</el-divider>
        <el-row class="input-row">
          <el-col :span="10">
            <el-input v-model="time2DateIn"/>
          </el-col>
          <el-col :span="4" style="text-align: center">
            <el-button type="primary" @click="Convert('time')">转换</el-button>
          </el-col>
          <el-col :span="10">
            <el-input v-model="time2DateOut"/>
          </el-col>
        </el-row>

        <el-divider content-position="left">时间 → 戳值</el-divider>
        <el-row class="input-row">
          <el-col :span="10">
            <el-input v-model="date2TimeIn"/>
          </el-col>
          <el-col :span="4" style="text-align: center">
            <el-button type="primary" @click="Convert('date')">转换</el-button>
          </el-col>
          <el-col :span="10">
            <el-input v-model="date2TimeOut"/>
          </el-col>
        </el-row>

      </el-col>

      <el-col :span="10" class="history">
        <el-scrollbar class="history-scrollbar">
          <p v-for="(item,i) in history" style="border-bottom: 1px dashed rgb(220,223,230);padding-bottom: 5px">
            <span style="width: 30px;display: inline-block;color: #bbb">{{ history.length - i - 1 }}</span>
            <span style="width: 150px;display: inline-block">{{ item[0] }}</span>
            <span style="width: 150px;display: inline-block">{{ item[1] }}</span>
          </p>
        </el-scrollbar>
      </el-col>

    </el-row>

  </div>
</template>

<script setup>
import {useI18n} from "vue-i18n";
import {onMounted, ref} from 'vue'
import {ElMessage} from 'element-plus'
import {DateToTime, GetNowTime, TimeToDate} from '../../../../wailsjs/go/main/Time'

const {t} = useI18n();
const nowTime = ref('')
const nowDate = ref('')
const time2DateIn = ref()
const time2DateOut = ref('')
const date2TimeIn = ref('')
const date2TimeOut = ref()
const history = ref([['历史记录', '']])
const ticker = ref(0)
//启停操作
const OnOff = () => {
  if (ticker.value > 0) {
    clearInterval(ticker.value);
    ticker.value = 0;
  } else {
    DoGetNowTime()
  }
};
//获取当前时间
const DoGetNowTime = () => {
  GetNowTime().then(res => {
    nowTime.value = res.timestamp
    nowDate.value = res.date
  })
  ticker.value = setInterval(function () {
    GetNowTime().then(res => {
      nowTime.value = res.timestamp
      nowDate.value = res.date
    })
  }, 1000);
};
//时间转换
const Convert = (type) => {
  if (type === 'time') {
    TimeToDate(time2DateIn.value).then(res => {
      if (res.code === 1) {
        time2DateOut.value = res.data.out
        history.value.unshift([res.data.in, res.data.out])
      } else {
        ElMessage.error({
          showClose: true,
          message: res.msg,
          type: 'error',
        })
      }
    })
  } else if (type === 'date') {
    DateToTime(date2TimeIn.value).then(res => {
      if (res.code === 1) {
        date2TimeOut.value = res.data.out
        history.value.unshift([res.data.in, res.data.out])
      } else {
        ElMessage.error({
          showClose: true,
          message: res.msg,
          type: 'error',
        })
      }
    })
  } else {
    ElMessage.error({
      showClose: true,
      message: '操作类型错误：' + type,
      type: 'error',
    })
  }
};

onMounted(() => {
  DoGetNowTime()
})

</script>

<style lang="scss">

.content {
  margin: 10px;
  height: calc(100% - 22px);
  overflow: hidden;
}

.main-row {
  height: 100%
}

.input-row {
  margin-bottom: 50px;
}

.history {
  height: 100%;
  padding-left: 10px;
}

.history-scrollbar {
  /*background: #dedede;*/
}

.el-divider__text.is-left {
  left: 0;
}

.el-divider__text {
  padding-left: 0;
}
</style>
