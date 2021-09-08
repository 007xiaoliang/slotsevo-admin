<template>
  <el-container>
    <el-main>
      <el-row>
        <div class="el-col-3 inputStyle">
          <el-input v-model="userID" placeholder="用户ID回车查询" autofocus=autofocus
                    @keyup.enter.native="getUserInfo"></el-input>
        </div>
        <div class="el-col-3 inputStyle">
          <el-input v-model="showSpecifyID" placeholder="输入id或名称过滤" autofocus=autofocus v-show="number==='theme'"></el-input>
        </div>
        <div class="el-col-4 inputStyle">
          <el-input v-model="showSpecifyAct" placeholder="输入活动英文名过滤" autofocus=autofocus v-show="number==='act'"></el-input>
        </div>
        <div class="el-col-4">
          <el-button type="primary" @click="updateInfo(number)" v-show="number==='base'">提交修改</el-button>
        </div>
      </el-row>
      <el-row>
        <div>
          <el-tabs v-model="number" type="card" @tab-click="change">
            <el-tab-pane label="基本信息" name="base">
              <BaseInfo ref="baseInfo"/>
            </el-tab-pane>
            <el-tab-pane label="主题信息" name="theme">
              <ThemeInfo :themeID="showSpecifyID"/>
            </el-tab-pane>
            <el-tab-pane label="活动信息" name="act">
              <ActInfo :actName="showSpecifyAct"/>
            </el-tab-pane>
          </el-tabs>
        </div>
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
      number: 'base',
      userID: "",
      showSpecifyID: null,
      showSpecifyAct: null
    }
  },
  computed: {
    ...mapGetters('rpcUser', ['rpcUserID']),
  },
  methods: {
    ...mapMutations("rpcUser", ["setRpcUserID", "setRpcBaseInfo", "setRpcThemeInfo", "setRpcActivityInfo"]),
    change: function (index) {
      this.number = index.name
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
      this.$refs["baseInfo"].updateInfo()
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
.inputStyle{
  margin: 0 20px 0 0;
}
</style>
