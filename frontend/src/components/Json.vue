<template>
  <div class="container">
      <el-container>
<!--        <el-header>Header</el-header>-->
        <el-container>
          <el-scrollbar style="height: 100%;width: 50%">
            <el-aside style="height: 100%;width: 100%">
              <p v-for="(item ,i) in list1" v-bind:key="i">{{i}}</p>
            </el-aside>
          </el-scrollbar>
          <el-main style="height: 100%;width: 50%">Main</el-main>
        </el-container>
      </el-container>

  </div>
</template>

<script>

export default {
  data() {
    return {
      list1:[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,],
      list2:[1,1,1,1,1,1,1,1,1],
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
  overflow: hidden;
}


.el-aside{
  background-color: #D3DCE6;
  overflow-y: hidden;
}

.bg-purple {
  background: #d3dce6;
}
.bg-purple-light {
  background: #e5e9f2;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
.el-header, .el-footer {
  background-color: #B3C0D1;
  color: #333;
  text-align: center;
  line-height: 60px;
}



body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}

</style>
