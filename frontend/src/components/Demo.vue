<template>
  <div class="container">
    <div class="content">

      <p v-for="(item,i) in list" v-bind:key="i">{{i}}</p>
      <el-button @click="getMessage">BTN</el-button>
      <el-input v-model="message" style="width: 300px"></el-input>

    </div>
  </div>
</template>

<script>

export default {
  data() {
    return {
      list:[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,],
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
  }
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container{
  height: 100%;
  width: 100%;
  padding: 0 ;
  position: fixed;
  overflow: auto;
  background: red;
  top: 0;
}

.content{
  background: #ccc;
  position: absolute;
  width: 100%;

  /*宽度撑破容器使用下列样式*/
  /*width: 98%;!*不支持calc()的浏览器*!*/
  /*width:-moz-calc(100% - 20px);*/
  /*width:-webkit-calc(100% - 20px);*/
  /*width: calc(100% - 20px);*/
  /*padding: 10px;*/
}

</style>
