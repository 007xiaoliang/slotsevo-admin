<template>
  <el-container>
    <el-main>
      <el-row>
        <div class="el-col-3 inputStyle">
          <el-input v-model="userID" placeholder="用户ID" autofocus=autofocus
                    @keyup.enter.native="getUserInfo"></el-input>
        </div>
        <div class="el-col-4">
          <button class="sureBtn" @click="getUserInfo()">查询</button>
        </div>
        <div class="el-col-14">
          <button class="btnList" @click="change(0)" :class="{ newStyle:0===number}">基本信息</button>
          <button class="btnList" @click="change(1)" :class="{ newStyle:1===number}">主题信息</button>
          <button class="btnList" @click="change(2)" :class="{ newStyle:2===number}">活动信息</button>
        </div>
      </el-row>
      <el-row>
        <BaseInfo v-show="number===0"/>
        <ActInfo v-show="number===1"/>
        <ThemeInfo v-show="number===2"/>
      </el-row>
    </el-main>
  </el-container>
</template>


<script>
import BaseInfo from '@/view/backend/user/baseInfo.vue'
import ActInfo from '@/view/backend/user/actInfo.vue'
import ThemeInfo from '@/view/backend/user/themeInfo.vue'
import {getUserInfo} from "@/api/userInfo";
import {mapMutations} from "vuex";

export default {
  name: "user",
  data() {
    return {
      number: 0,
      userID: "",
    }
  },
  methods: {
    ...mapMutations("rpcUser", ["setRpcBaseInfo", "setRpcThemeInfo", "setRpcActivityInfo"]),
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
      const rpcUserInfo = getUserInfo({"userID": this.userID})
      this.setRpcBaseInfo(rpcUserInfo.data.baseinfo)
      this.setRpcThemeInfo(rpcUserInfo.data.themeinfo)
      this.setRpcActivityInfo(rpcUserInfo.data.activityinfo)
    },
    checkNumber(String) {
      const reg = /^[0-9]+.?[0-9]*$/;
      return reg.test(String);
    }
  },
  components: {
    BaseInfo,
    ActInfo,
    ThemeInfo
  }
}
</script>
<style>
.btnList {
  width: 33%;
  padding: 25px 0;
  font-size: large;
  font-weight: bold;
  border: 0 solid #fff;
  color: #000;
  outline: none;
  background: #fff;
}

.sureBtn {
  width: 33%;
  padding: 25px 0;
  font-size: large;
  font-weight: bold;
  border: 0 solid #fff;
  color: rgba(18, 33, 15, 0.54);
  outline: none;
  background: #fff;
}

.inputStyle {
  padding: 20px 2px;
  font-size: large;
  font-weight: bold;
  border: 0 solid #fff;
  color: #000;
  outline: none;
  background: #fff;
}

.newStyle {
  border-bottom: 2px solid #f0892e;
  color: #f0892e;
  font-size: larger;
  font-weight: bold;
}
</style>
