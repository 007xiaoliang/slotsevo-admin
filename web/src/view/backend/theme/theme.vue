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
        <div class="el-col-2">
          <el-input v-model="themeIndex" placeholder="查找主题" autofocus=autofocus v-show="this.themeInfo.length > 0"
                    @keyup.enter.native="setActiveItem"></el-input>
        </div>
        <div class="el-col-2">
          <el-button type="primary" @click="setActiveItem" v-show="this.themeInfo.length > 0">确定</el-button>
        </div>
      </el-row>
      <el-row>
        <div style="width: 80% ;margin: 40px">
          <el-carousel :autoplay="false" type="card" height="300px" trigger="click" ref="carousel">
            <el-carousel-item v-for="(item, count) in themeInfo" :key="count" :name="item[0]">
              <div @click="getConfig(item[0])">
                <div>{{ item[0] }}</div>
                <div>{{ item[1] }}</div>
              </div>
            </el-carousel-item>
          </el-carousel>
        </div>
      </el-row>
    </el-main>
  </el-container>
</template>


<script>
import {getThemeInfo} from "@/api/themeInfo";

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
      }, {
        value: '7777',
        label: '7777'
      }, {
        value: '9527',
        label: '9527'
      }, {
        value: '8000',
        label: '8000'
      }],
      ipList: [{
        value: '192.168.2.34',
        label: '192.168.2.34'
      }],
      ip: '192.168.2.34',
      port: '7003',
      themeInfo: [],
      themeIndex: "",
      themeIndexArr: []
    }
  },

  methods: {
    getRpcThemeInfo: function () {
      getThemeInfo({"rpcPort": this.port, "rpcHost": this.ip}).then((ele) => {
        if (ele.code !== 0) {
          this.$message({
            type: "error",
            message: ele.msg,
            showClose: true,
          });
          return
        }
        this.themeInfo = ele.data.sort().reverse()
        for (let i = 0; i < this.themeInfo.length; i++) {
          this.themeIndexArr.push(this.themeInfo[i].join("-"))
        }
      })
    },
    setActiveItem() {
      let index = ""
      for (let i = 0; i < this.themeIndexArr.length; i++) {
        if (this.themeIndexArr[i].indexOf(this.themeIndex) > -1) {
          index = this.themeIndexArr[i].split("-")[0]
          break
        }
      }
      this.$refs.carousel.setActiveItem(index)
    },
    getConfig(themeId) {
      alert(themeId)
    }
  },
  components: {},
}
</script>
<style>

.el-carousel__item div {
  color: #475669;
  font-size: 40px;
  text-align: center;
  margin: 10px;
  padding: 10px;
}

.el-row .el-button {
  margin: 0 15px 0 15px;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}
</style>
