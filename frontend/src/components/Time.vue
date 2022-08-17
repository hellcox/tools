<template>
  <div style="height: 100%;position: fixed;width: 100%;padding: 0;margin: 0;background: #ccc;float: right">

    <el-container style="height: 100%;padding: 0;margin: 0;position: absolute;width: 100%">
      <el-aside style="height:20%;width: 50%;padding: 0;margin: 0;position: fixed;left: 0">Aside</el-aside>
      <el-main style="height:20%;width: 50%;padding: 0;margin: 0;position: fixed;right: 0">
        Main
        <p class="p">1asdfasdjflka;jfdkljaslkjfalkjfl;kjaklsjdflkajl;kfjkljf;lk111</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>2</p>
        <p>2</p>
        <p>2</p>
        <p>2</p>
        <p>2</p>
        <p>3</p>
      </el-main>
    </el-container>

<!--    <el-row>-->
<!--      <el-col :span="12">-->
<!--        <div class="grid-content bg-purple"></div>-->
<!--        <el-button @click="getMessage">GET</el-button>-->
<!--        <el-input v-model="input" placeholder="请输入内容"></el-input>-->
<!--      </el-col>-->
<!--      <el-col :span="12" style="height: 100%">-->
<!--        <div class="grid-content bg-purple-light"></div>-->
<!--        <el-button @click="getMessage">GET</el-button>-->
<!--        <el-input v-model="input" placeholder="请输入内容"></el-input>-->
<!--      </el-col>-->
<!--    </el-row>-->


  </div>
</template>

<script>
export default {
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
        self.input = result;
      });
    }
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

</style>
