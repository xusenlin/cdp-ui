<template>
  <div class="app-bar">
    <div class="logo">
      <i class="el-icon-monitor"></i>
      <div class="title">CdpUi</div>
    </div>
    <div>
      <el-button type="primary" round size="small" @click="saveTask"
        >保存</el-button
      >
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "AppBar",
  computed: mapState(["tasks"]),
  methods: {
    saveTask() {
      this.$store.commit("saveJsonFile", this.tasks);
    }
  }
};
</script>

<style scoped lang="scss">
@import "../assets/style/variables";
.app-bar {
  color: $--color-primary;
  background-color: #fff;
  height: 52px;
  border-bottom: 1px solid #e5e5e5;
  z-index: 51;
  display: flex;
  padding: 0 20px;
  justify-content: space-between;
  align-items: center;
  .logo {
    display: flex;
    font-size: 24px;
    .title {
      margin-left: 10px;
    }
  }
}
</style>
