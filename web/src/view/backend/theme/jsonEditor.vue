<template>
  <div style="width: 50%;">
    <b-code-editor v-model="jsonStr" :auto-format="true" :smart-indent="true" theme="dracula" :indent-unit="4" :line-wrap="false" ref="editor"></b-code-editor>
    <br>
    <el-button type="primary" @click="onSubmit">提交</el-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      jsonStr : this.themeConfig
    }
  },
  props:["themeConfig"],
  watch: {
    themeConfig: function (newVal) {
      console.info(newVal)
      this.jsonStr = newVal
    },
  },
  methods: {
    // 检测json格式
    isJSON(str) {
      if (typeof str == 'string') {
        try {
          const obj = JSON.parse(str);
          return !!(typeof obj == 'object' && obj);

        } catch(e) {
          return false;
        }
      }else if (typeof str == 'object'  && str) {
        return true;
      }
    },
    onSubmit(){
      if (!this.isJSON(this.jsonStr)){
        this.$message.error(`json格式错误`)
        return false
      }
      this.$message.success('json格式正确')
    }
  }
}
</script>

<style>

</style>
