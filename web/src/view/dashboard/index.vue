<template>
  <div class="big">
    <el-row>
      <div class="card">
        <el-col :xs="24" :lg="16" :md="16">
          <div class="car-left">
            <el-row>
              <div>
                  <div class="text">
                    <h4>Hi,亲爱的 {{userInfo.nickName}}! 欢迎访问slotsevo后台管理系统</h4>
                  </div>
              </div>
            </el-row>
          </div>
        </el-col>
      </div>
    </el-row>
    <div class="shadow">
      <el-row :gutter="20">
        <el-col
            :span="4"
            v-for="(card, key) in toolCards"
            :key="key"
            @click.native="toTarget(card.name)"
            :xs="8"
        >
          <el-card shadow="hover" class="grid-content">
            <i :class="card.icon" :style="{ color: card.color }"></i>
            <p>{{ card.label }}</p>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>

import {mapGetters} from "vuex";

export default {
  name: "dashboard",
  data() {
    return {
      toolCards: [
        {
          label: "用户信息",
          icon: "el-icon el-icon-s-custom",
          name: "user",
          color: "#ff9c6e",
        },
        {
          label: "主题信息",
          icon: "el-icon el-icon-s-grid",
          name: "theme",
          color: "#69c0ff",
        },
        {
          label: "活动信息",
          icon: "el-icon el-icon-menu",
          name: "activity",
          color: "#b37feb",
        },
        {
          label: "产品文案",
          icon: "el-icon el-icon-folder",
          name: "product",
          color: "#6eff9e",
        },
        {
          label: "数值控制",
          icon: "el-icon el-icon-edit",
          name: "data",
          color: "#d769ff",
        }
      ],
    };
  },
  methods: {
    toTarget(name) {
      this.$router.push({name});
    },
  },
  computed: {
    ...mapGetters('user', ['userInfo']),
  }
};
</script>

<style lang="scss" scoped>
.big {
  margin: 100px 0 0 0;
  background-color: rgb(243, 243, 243);
  padding-top: 15px;

  .top {
    width: 100%;
    height: 360px;
    margin-top: 20px;
    overflow: hidden;

    .chart-container {
      position: relative;
      width: 100%;
      height: 100%;
      padding: 20px;
      background-color: #fff;
    }
  }

  .mid {
    width: 100%;
    height: 380px;

    .chart-wrapper {
      height: 340px;
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }

  .bottom {
    width: 100%;
    height: 300px;
    // margin: 20px 0;
    .el-row {
      margin-right: 4px !important;
    }

    .chart-player {
      width: 100%;
      height: 270px;
      padding: 10px;
      background-color: #fff;
    }
  }
}
</style>
