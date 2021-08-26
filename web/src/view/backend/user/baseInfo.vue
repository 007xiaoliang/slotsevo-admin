<template>
  <div style="width:80%">
    <el-table
        :data="this.baseInfoData"
        height="600"
        border
        style="width: 100%">
      <el-table-column
          prop="name"
          label="字段"
          >
      </el-table-column>
      <el-table-column
          prop="desc"
          label="描述"
          >
      </el-table-column>
      <el-table-column
          prop="editable"
          label="可修改"
          >
      </el-table-column>
      <el-table-column
          prop="value"
          label="值">
      </el-table-column>
    </el-table>
  </div>
</template>


<script>
import {mapGetters} from "vuex";

export default {
  name: "BaseInfo",
  computed: {
    ...mapGetters('rpcUser', ['rpcBaseInfo']),
  },
  data() {
    return {
      baseInfoData: []
    }
  },
  watch:{
    rpcBaseInfo(newData){
      for(let k in newData){
        let temp = new Map()
        temp["name"] = k
        temp["desc"] = newData[k]["desc"]
        temp["editable"] = newData[k]["editable"] === 0 ? "否": "是"
        if(typeof(newData[k]["data"]) === "object"){
          temp["value"] = JSON.stringify(newData[k]["data"])
        }else{
          temp["value"] = newData[k]["data"]
        }
        this.baseInfoData.push(temp)
      }
    }
  }
};
</script>
<style>
</style>
