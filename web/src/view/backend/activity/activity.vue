<template>
  <el-container>
    <el-main>
      <el-row>
        <div class="el-col-4">
          <el-select v-model="ip" placeholder="请选择ip">
            <el-option
                v-for="item in ipList"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
        </div>
        <div class="el-col-2">
          <el-select v-model="port" placeholder="请选择端口">
            <el-option
                v-for="item in portList"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
        </div>
        <div class="el-col-2">
          <el-button type="primary" @click="getRpcThemeInfo">确定</el-button>
        </div>
      </el-row>
      <el-row>
        <div id="jsonEditor" style="width: 400px; height: 400px;"></div>
      </el-row>
    </el-main>
  </el-container>
</template>


<script>
import {getActInfo} from "@/api/actInfo";

export default {
  name: "theme",
  data() {
    return {
      portList: [{
        value: '7001',
        label: '7001'
      }, {
        value: '7003',
        label: '7003'
      }, {
        value: '7005',
        label: '7005'
      },{
        value: '7777',
        label: '7777'
      }, {
        value: '9527',
        label: '9527'
      }, {
        value: '8000',
        label: '8000'}],
      ipList: [{
        value: '192.168.2.34',
        label: '192.168.2.34'
      }],
      ip: '192.168.2.34',
      port: '7001',
      themeInfo: [],
    }
  },

  methods: {
    getRpcThemeInfo: function () {
      getActInfo({"rpcPort": this.port, "rpcHost": this.ip}).then((ele) => {
        if (ele.code !== 0) {
          this.$message({
            type: "error",
            message: ele.msg,
            showClose: true,
          });
          return
        }
        this.themeInfo = ele.data
        console.info(this.themeInfo)
      })
    },
  },
  components: {},
}
</script>
<style>
.el-row .el-button {
  margin: 0 15px 0 15px;
}
</style>
