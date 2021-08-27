<template>
  <div style="width:85%">
    <el-table :header-cell-style="{background:'#eef1f6',color:'#796130'}"
              :data="this.baseInfoData"
              height="630"
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
          label="值(单击修改)"
          min-width="180">
        <editable-cell slot-scope="{row}"
                       :can-edit="row.editable2"
                       v-model="row.value">
          <span slot="content">{{ row.value }}</span>
        </editable-cell>
      </el-table-column>
    </el-table>
  </div>
</template>


<script>
import {mapGetters} from "vuex";
import EditableCell from "@/view/backend/user/EditableCell.vue";

export default {
  name: "BaseInfo",
  components: {
    EditableCell
  },
  computed: {
    ...mapGetters('rpcUser', ['rpcBaseInfo']),
  },
  data() {
    return {
      baseInfoData: []
    }
  },
  watch: {
    rpcBaseInfo(newData) {
      this.fotmatBaseInfo(newData)
    },
  },
  methods: {
    fotmatBaseInfo: function (baseinfo) {
      this.baseInfoData = []
      for (let k in baseinfo) {
        let temp = new Map()
        temp["name"] = k
        temp["desc"] = baseinfo[k]["desc"]
        temp["editable"] = baseinfo[k]["editable"] === 0 ? "否" : "是"
        temp["editable2"] = temp["editable"] === "是"
        if(typeof(baseinfo[k]["data"]) === "object" || typeof(baseinfo[k]["data"]) === "number"){
          temp["value"] = JSON.stringify(baseinfo[k]["data"])
        }else{
          temp["value"] = baseinfo[k]["data"]
        }
        this.baseInfoData.push(temp)
      }
    },
    updateInfo: function(){
      this.$message({
        type: "success",
        message: "这里是baseinfo",
        showClose: true,
      })
    }
  },
  created() {
    if(this.rpcBaseInfo){
      this.fotmatBaseInfo(this.rpcBaseInfo)
    }
  }
}
</script>
<style>
</style>
