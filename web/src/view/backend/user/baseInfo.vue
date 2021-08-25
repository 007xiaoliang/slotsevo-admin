<template>
  <div style="width:80%">
    <h1>{{ this.rpcBaseInfo }}</h1>
    <el-row :gutter="20">
      <el-col :span="6"><div class="grid-content bg-purple">字段</div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple">描述</div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple">可修改</div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple">值</div></el-col>
    </el-row>
  </div>
</template>


<script>
import {mapGetters} from "vuex";

export default {
  name: "BaseInfo",
  computed: {
    ...mapGetters('rpcUser', ['rpcBaseInfo']),
  }
};
</script>
<style>
.el-row {
  margin-bottom: 20px;
}
.el-col {
  border-radius: 4px;
}
.bg-purple-dark {
  background: #99a9bf;
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
  /*height:36px;*/
  line-height:36px;
  overflow:hidden;
  font-size: large;
  text-align: center;
}
.row-bg {
  padding: 10px 0;
  background-color: #f9fafc;
}
</style>
