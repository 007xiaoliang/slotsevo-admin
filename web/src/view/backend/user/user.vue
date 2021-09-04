<template>
  <el-container>
    <el-main>
      <el-row>
        <div class="el-col-2 inputStyle">
          <el-input v-model="userID" placeholder="用户ID回车查询" autofocus=autofocus
                    @keyup.enter.native="getUserInfo"></el-input>
        </div>
        <div class="el-col-12">
          <button class="btnList" @click="change(0)" :class="{ newStyle:0===number}">基本信息</button>
          <button class="btnList" @click="change(1)" :class="{ newStyle:1===number}">主题信息</button>
          <button class="btnList" @click="change(2)" :class="{ newStyle:2===number}">活动信息</button>
        </div>
        <div class="el-col-3 inputStyle">
          <el-input v-model="showSpecifyID" placeholder="输入id或名称过滤" autofocus=autofocus  v-show="number===1"
                    @keyup.enter.native="getUserInfo"></el-input>
        </div>
        <div class="el-col-3 inputStyle">
          <el-input v-model="showSpecifyAct" placeholder="输入活动英文名过滤" autofocus=autofocus  v-show="number===2"
                    @keyup.enter.native="getUserInfo"></el-input>
        </div>
        <div class="el-col-4">
          <button class="sureBtn" @click="updateInfo(number)" v-show="number===0">提交修改</button>
        </div>
      </el-row>
      <el-row>
        <BaseInfo v-show="number===0" ref="baseInfo"/>
        <ThemeInfo v-show="number===1" :themeID="showSpecifyID"/>
        <ActInfo v-show="number===2" :actName="showSpecifyAct" ref="actInfo"/>
      </el-row>
    </el-main>
  </el-container>
</template>


<script>
import BaseInfo from '@/view/backend/user/baseInfo.vue'
import ActInfo from '@/view/backend/user/actInfo.vue'
import ThemeInfo from '@/view/backend/user/themeInfo.vue'
import {getUserInfo} from "@/api/userInfo";
import {mapGetters, mapMutations} from "vuex";

export default {
  name: "user",
  data() {
    return {
      number: 0,
      userID: "",
      showSpecifyID: null,
      showSpecifyAct: null,
      componentMap: {0: "baseInfo", 2: "actInfo"}
    }
  },
  computed: {
    ...mapGetters('rpcUser', ['rpcUserID']),
  },
  methods: {
    ...mapMutations("rpcUser", ["setRpcUserID", "setRpcBaseInfo", "setRpcThemeInfo", "setRpcActivityInfo"]),
    change: function (index) {
      this.number = index
    },
    getUserInfo: function () {
      if (!this.checkNumber(this.userID)) {
        this.$message({
          type: "error",
          message: "请正确填写用户id",
          showClose: true,
        });
        return
      }
      getUserInfo({"userID": this.userID}).then((ele) => {
        if (ele.code !== 0) {
          this.$message({
            type: "error",
            message: ele.msg,
            showClose: true,
          });
          return
        }
        this.setRpcUserID(this.userID)
        this.setRpcBaseInfo(ele.data.baseinfo)
        this.setRpcThemeInfo(ele.data.themeinfo)
        this.setRpcActivityInfo(ele.data.activityinfo)
      })
    },
    checkNumber(String) {
      const reg = /^[0-9]+.?[0-9]*$/;
      return reg.test(String);
    },
    updateInfo: function (index) {
        this.$refs[this.componentMap[index]].updateInfo()
    }
  },
  components: {
    BaseInfo,
    ActInfo,
    ThemeInfo
  },
  created() {
    this.userID = this.rpcUserID
  }
}
</script>
<style>
.btnList {
  width: 33%;
  padding: 5px 0;
  font-size: medium;
  font-weight: bold;
  border: 0 solid #fff;
  color: #000;
  outline: none;
  background: #fff;
}

.sureBtn {
  width: 33%;
  padding: 5px 0;
  font-size: medium;
  font-weight: bold;
  border: 0 solid #fff;
  color: #2e95f0;
  outline: none;
  background: #fff;
}

.inputStyle {
  margin: 0 30px 0 10px;
  font-size: medium;
  font-weight: bold;
  border: 0 solid #fff;
  color: #000;
  outline: none;
  background: #fff;
}

.newStyle {
  border-bottom: 2px solid #f0892e;
  color: #f0892e;
  font-size: medium;
  font-weight: bold;
}
</style>
