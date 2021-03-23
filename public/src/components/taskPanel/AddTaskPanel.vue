<template>
  <div class="row">
    <div class="top">
      <div class="add-task" @click="addTask">
        <i class="el-icon-plus"></i>
        <div class="title">新建任务</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "AddTaskPanel",
  emits: ["addTask"],
  methods: {
    addTask() {
      this.$emit("addTask");
    }
  }
};
</script>

<style scoped lang="scss">
.row {
  list-style-type: none;
  width: 300px;
  height: 100%;
  padding: 0 12px;

  .top {
    width: 100%;
    height: 48px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: #262626;
    font-size: 14px;

    .add-task {
      display: flex;
      justify-content: flex-start;
      align-items: center;
      color: #8c8c8c;
      cursor: pointer;

      &:hover {
        color: #3f51b5;
      }

      .title {
        margin-left: 6px;
      }
    }
  }
}
</style>
