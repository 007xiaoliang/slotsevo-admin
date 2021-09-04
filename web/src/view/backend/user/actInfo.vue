<template>
  <div>
    <div class="actDiv">
      <div class="actName" v-for="(actInfo, actName) in this.rpcActivityInfo" :key="actName" :id="actName"
           v-show="!showActName || actName.indexOf(showActName) > -1">
        <el-row :gutter="20">
          <el-col :span="4">
            <div class="grid-content bg-purple"><p>{{ actName }}</p></div>
          </el-col>
          <el-col :span="5">
            <div class="grid-content bg-purple"><a>提交修改</a></div>
          </el-col>
        </el-row>
        <el-row :gutter="20" v-for="(v, k) in actInfo">
          <el-col :span="6">
            <div class="grid-content bg-purple"><span>{{ k }}</span></div>
          </el-col>
          <el-col :span="18">
            <div class="grid-content bg-purple"><span contenteditable="true">{{ v }}</span></div>
          </el-col>
        </el-row>
      </div>

    </div>
  </div>
</template>


<script>
import {mapGetters} from "vuex";

export default {
  name: "ActInfo",
  computed: {
    ...mapGetters('rpcUser', ['rpcActivityInfo']),
  },
  data() {
    return {
      showActName: null,
    }
  },
  props: ['actName'],
  watch: {
    actName: function (newVal) {
      this.showActName = newVal
    },
  },
};
</script>
<style>
.actDiv {
  height: 660px;
  width: 100%;
  overflow: scroll;
}

.actName {
  width: 99%;
  margin-top: 8px;
  background: #f6f6f6;
  text-align: center;
}

.actName:hover {
  -webkit-transform: translateY(-3px);
  -ms-transform: translateY(-3px);
  transform: translateY(-3px);
  -webkit-box-shadow: 0 0 6px #999;
  box-shadow: 0 0 6px #999;
  -webkit-transition: all .5s ease-out;
  transition: all .1s ease-out;
  border: #000 solid;
}

.bg-purple {
  background: #ececec;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}

.bg-purple p {
  font-size: medium;
  color: #f0892e;
  display: inline-block;
  vertical-align: middle;
  text-align: center;
  padding: 10px;
}

.bg-purple a {
  font-size: medium;
  color: #2ed3f0;
  display: inline-block;
  vertical-align: middle;
  text-align: center;
  padding: 10px;
}

.bg-purple span {
  font-size: medium;
  color: #0e0d0d;
  display: inline-block;
  vertical-align: middle;
  text-align: center;
  padding: 10px;
}
</style>
