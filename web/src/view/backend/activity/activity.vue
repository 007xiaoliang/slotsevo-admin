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
          <el-button type="primary" @click="getRpcActInfo">确定</el-button>
        </div>
        <div class="el-col-3">
          <el-input v-model="themeIndex" placeholder="输入活动英文名查询" autofocus=autofocus v-show="this.actInfo.length > 0"
                    @keyup.enter.native="setActiveItem"></el-input>
        </div>
        <div class="el-col-2">
          <el-button type="primary" @click="setActiveItem" v-show="this.actInfo.length > 0">确定</el-button>
        </div>
        <div class="el-col-4">
          <el-button type="danger" ref="serverTimeBtn" v-show="this.serverTime>0"></el-button>
        </div>
      </el-row>
      <el-row>
        <el-row>
          <div style="width: 60% ;margin: 40px 0 40px 0">
            <el-carousel :autoplay="false" type="card" height="100px" trigger="click" @change='change'
                         indicator-position="none" ref="carousel">
              <el-carousel-item v-for="item in actInfo" :key="item" :name="item">
                <div @click="getConfig(item)">
                  <div>{{ item }}</div>
                </div>
              </el-carousel-item>
            </el-carousel>
          </div>
        </el-row>
        <el-row>
          <jsonEditor :actConfig="this.actConfig"></jsonEditor>
        </el-row>
      </el-row>
    </el-main>
  </el-container>
</template>


<script>
import {getActInfo} from "@/api/actInfo";
import {getActWeight} from "@/api/actInfo";
import jsonEditor from '@/view/backend/activity/jsonEditor.vue'

export default {
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
      port: '7001',
      actInfo: [],
      serverTime: 0,
      actConfig: {},
      themeIndex: ""
    }
  },

  methods: {
    getRpcActInfo: function () {
      getActInfo({"rpcPort": this.port, "rpcHost": this.ip}).then((ele) => {
        if (ele.code !== 0) {
          this.$message({
            type: "error",
            message: ele.msg,
            showClose: true,
          });
          return
        }
        this.actInfo = ele.data.activity_type_list
        this.serverTime = parseInt(ele.data.server_time)
      })
    },
    setActiveItem() {
      let index = ""
      for (let i = 0; i < this.actInfo.length; i++) {
        if (this.actInfo[i].indexOf(this.themeIndex.toUpperCase()) > -1) {
          index = this.actInfo[i]
          break
        }
      }
      this.$refs.carousel.setActiveItem(index)
    },
    getConfig(actType) {
      getActWeight({"rpcPort": this.port, "rpcHost": this.ip, "actType": actType}).then((ele) => {
        if (ele.code !== 0) {
          this.$message({
            type: "error",
            message: ele.msg,
            showClose: true,
          });
          return
        }
        if (ele.data.err_code !== 0) {
          this.$message({
            type: "error",
            message: ele.data.err_code,
            showClose: true,
          });
          return
        }
        this.actConfig = ele.data.activityData
      })
    },
    change() {
      this.actConfig = {}
    },
    timestampToTime(timestamp) {
      // 时间戳转日期时间--从年到分
      let time = new Date(timestamp * 1000)
      let y = time.getFullYear();
      let m = time.getMonth() + 1;
      let d = time.getDate();
      let h = time.getHours();
      let mm = time.getMinutes();
      let s = time.getSeconds();
      return y + '-' + this.add0(m) + '-' + this.add0(d) + ' ' + this.add0(h) + ':' + this.add0(mm) + ':' + this.add0(s);
    },
    add0(m) {
      return m < 10 ? '0' + m : m
    },
    serverDate() {
      if (this.serverTime > 0) {
        this.serverTime++
        this.$refs.serverTimeBtn.$el.innerText = this.timestampToTime(this.serverTime)
      }
    }
  },
  components: {
    jsonEditor
  },
  mounted() {
    setInterval(this.serverDate, 1000);
  },
}
</script>
<style>
.el-carousel__item div {
  color: #475669;
  font-size: larger;
  text-align: center;
  /*margin: 10px;*/
  padding: 20px;
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
