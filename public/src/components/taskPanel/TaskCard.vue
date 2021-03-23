<template>
  <div class="card">
    <div class="desc">
      <div class="item">
        <el-tooltip effect="dark" content="关键字" placement="top">
          <i class="el-icon-key"></i>
        </el-tooltip>
        <div class="key">{{ task.name }}</div>
      </div>
      <div class="item">
        <el-tooltip effect="dark" content="操作" placement="top">
          <i class="el-icon-price-tag"></i>
        </el-tooltip>
        <div class="title">{{ task.title }}</div>
      </div>
      <div class="item" v-if="task.desc">
        <el-tooltip effect="dark" content="描述" placement="top">
          <i class="el-icon-chat-dot-square"></i>
        </el-tooltip>
        <div class="desc">{{ task.desc }}</div>
      </div>
    </div>
    <div class="action">
      <!--      <i class="el-icon-circle-check done"></i>-->
      <i class="el-icon-s-tools" style="color: #67C23A" @click="edit"></i>
      <i class="el-icon-delete-solid" style="color: #F56C6C" @click="del"></i>
    </div>
    <div class="index">{{ index + 1 }}</div>
  </div>
</template>

<script>
export default {
  name: "TaskCard",
  emits: ["editSubTask", "delSubTask"],
  props: {
    index: Number,
    task: {
      type: Object,
      default: () => {}
    }
  },
  methods: {
    edit() {
      this.$emit("editSubTask", this.task, this.index);
    },
    del() {
      this.$emit("delSubTask", this.task, this.index);
    }
  }
};
</script>

<style scoped lang="scss">
@import "../../assets/style/variables";
.card {
  min-height: 62px;
  position: relative;
  width: 100%;
  margin-bottom: 12px;
  white-space: normal;
  background-color: #fff;
  border-radius: 4px;
  box-shadow: 0 1px 2px 0 rgb(0 0 0 / 10%);
  padding: 12px;
  font-size: 14px;
  display: flex;
  &:last-child {
    margin-bottom: 0;
  }
  .index {
    position: absolute;
    right: 6px;
    font-size: 30px;
    transform: rotate(20deg);
    color: #eff0f3;
    top: -5px;
  }
  .desc {
    flex: 1;
    overflow: hidden;
    i {
      color: $--color-primary;
      margin-right: 8px;
      cursor: pointer;
    }
    .item {
      display: flex;
      align-items: center;
      margin-bottom: 10px;
      &:last-child {
        margin-bottom: 0;
      }
      .key,
      .title {
        background: #f7f7f7;
        border-radius: 2px;
        padding: 2px 6px;
        color: #8c8c8c;
      }
      .desc {
        width: 100%;
        color: #8c8c8c;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
    }
  }
  .action {
    width: 20%;
    display: flex;
    color: $--color-primary;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
    font-size: 16px;
    cursor: pointer;
    .done {
      color: $--color-success;
      font-size: 28px;
    }
  }
}
</style>
