<template>
  <div>
    <div class="themeDiv">
      <div class="themeID" v-for="(themeInfo, themeID) in this.rpcThemeInfo" :key="themeID" :id="themeID"
           v-show="!specifyID || specifyID === themeID || specifyID === themeID.slice(5) || themeInfo.desc.indexOf(specifyID) > -1">
        <el-row :gutter="20">
          <el-col :span="4">
            <div class="grid-content bg-purple"><p>{{ themeID }}</p></div>
          </el-col>
          <el-col :span="4">
            <div class="grid-content bg-purple"><p>{{ themeInfo.desc }}</p></div>
          </el-col>
          <el-col :span="5">
            <div class="grid-content bg-purple"><a>提交修改</a></div>
          </el-col>
          <el-col :span="5">
            <div class="grid-content bg-purple"><a>指定信息</a></div>
          </el-col>
        </el-row>
        <el-row :gutter="20" v-for="(v, k) in themeInfo" v-if="k!=='desc'">
          <el-col :span="6">
            <div class="grid-content bg-purple"><span>{{ k }}</span></div>
          </el-col>
          <el-col :span="14">
            <div class="grid-content bg-purple"><span contenteditable="true">{{ v }}</span></div>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="20">
            <div class="grid-content bg-purple"><a>点击加载更多...</a></div>
          </el-col>
        </el-row>
      </div>

    </div>
  </div>

</template>


<script>
import {mapGetters} from "vuex";

export default {
  name: "ThemeInfo",
  data() {
    return {
      specifyID: null,
    }
  },
  props: ['themeID'],
  watch: {
    themeID: function (newVal) {
      this.specifyID = newVal
    },
  },
  computed: {
    ...mapGetters('rpcUser', ['rpcThemeInfo']),
  }
}
</script>

<style>
.themeDiv {
  height: 660px;
  width: 100%;
  overflow: scroll;
}

.themeID {
  width: 99%;
  margin: 8px;
  background: #f6f6f6;
  text-align: center;
}

.themeID:hover {
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
