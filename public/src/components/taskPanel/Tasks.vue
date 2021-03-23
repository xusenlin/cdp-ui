<template>
  <div class="row">
    <div class="top">
      <div class="title">
        <!--        <el-tooltip-->
        <!--          class="item"-->
        <!--          effect="dark"-->
        <!--          popper-class="terminal"-->
        <!--          :content="task.terminal"-->
        <!--          placement="bottom"-->
        <!--        >-->
        <div class="v">
          {{ task.title }}
        </div>
        <span class="tasks-total">({{ task.tasks.length }})</span>
        <!--        </el-tooltip>-->
      </div>
      <div class="action" v-show="task.progress === -1">
        <el-tooltip effect="dark" content="添加子任务" placement="top">
          <i class="el-icon-circle-plus action-item" @click="addSubTask"></i>
        </el-tooltip>
        <el-tooltip effect="dark" content="任务设置" placement="top">
          <i class="el-icon-s-tools action-item" @click="editTask"></i>
        </el-tooltip>
        <el-tooltip effect="dark" content="删除任务" placement="top">
          <i class="el-icon-delete-solid action-item" @click="delTask"></i>
        </el-tooltip>
        <el-tooltip effect="dark" content="打开任务文件夹" placement="top">
          <i class="el-icon-folder-opened action-item" @click="openFile"></i>
        </el-tooltip>
      </div>
      <el-progress
        style="width: 120px"
        v-if="task.progress !== -1"
        :text-inside="true"
        :stroke-width="16"
        :percentage="
          Number(((task.progress / task.tasks.length) * 100).toFixed(2))
        "
      ></el-progress>
    </div>
    <div
      class="tasks"
      v-loading="task.progress !== -1"
      element-loading-text="正在执行任务"
    >
      <div class="cards">
        <TaskCard
          v-for="(t, i) in task.tasks"
          :index="i"
          :task="t"
          @editSubTask="editSubTask"
          @delSubTask="delSubTask"
        />
      </div>
    </div>
    <div class="bottom" v-if="task.tasks.length !== 0">
      <el-button
        v-if="task.progress === -1"
        style="width: 100%"
        type="primary"
        icon="el-icon-video-play"
        @click="runTask"
        >运行任务</el-button
      >
      <el-button
        v-else
        style="width: 100%"
        type="warning"
        @click="stopTask"
        icon="el-icon-video-pause"
        >停止任务({{ task.tasks[task.progress].title }})</el-button
      >
    </div>
  </div>
  <EditSubTask ref="EditSubTaskRef" @submitData="submitSubTaskData" />
</template>

<script>
import TaskCard from "./TaskCard.vue";
import EditSubTask from "./EditSubTask.vue";
export default {
  name: "Tasks",
  components: { TaskCard, EditSubTask },
  emits: ["editTask", "delTask"],
  props: {
    index: Number,
    task: {
      type: Object,
      default: () => {}
    }
  },
  methods: {
    editTask() {
      this.$emit("editTask", this.index);
    },
    delTask() {
      this.$emit("delTask", this.index);
    },
    openFile() {
      this.$store.commit("openFile", this.index);
    },
    addSubTask() {
      this.$refs["EditSubTaskRef"].Add(this.index);
    },
    editSubTask(t, index) {
      this.$refs["EditSubTaskRef"].Edit(t, this.index, index);
    },
    delSubTask(t, index) {
      this.$confirm("确定删除此子任务吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          let parentIndex = this.index;
          this.$store.commit("delSubTask", { parentIndex, index });
          this.$notify({
            title: "成功",
            message: "删除子任务成功",
            type: "success"
          });
        })
        .catch(() => {});
    },
    submitSubTaskData(task, pos) {
      let { parentIndex, index } = pos;
      if (index === -1) {
        console.log(task);
        this.$store.commit("addSubTask", { task, parentIndex });
      } else {
        this.$store.commit("editSubTask", { task, parentIndex, index });
      }
    },
    runTask() {
      this.$store.commit("runTask", { task: this.task, index: this.index });
    },
    stopTask() {
      this.$store.commit("stopTask", this.index);
    }
  }
};
</script>

<style scoped lang="scss">
@import "../../assets/style/variables";

.row {
  list-style-type: none;
  width: 300px;
  height: 100%;
  border-radius: 6px;
  background: #eff0f3;
  display: flex;
  flex-direction: column;
  .top {
    width: 100%;
    height: 48px;
    display: flex;
    padding: 0 12px;
    justify-content: space-between;
    align-items: center;
    color: $--color-primary;
    font-size: 16px;
    border-bottom: 1px solid #f7f7f7;
    .title {
      display: flex;
      align-items: center;
      overflow: hidden;
      flex: 1;
      margin-right: 10px;
      .v {
        flex: 1;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
      .tasks-total {
        display: inline-block;
        font-size: 12px;
        color: #595959;
        margin-left: 12px;
      }
    }

    .action {
      .action-item {
        cursor: pointer;
        margin-left: 12px;
        &:hover {
          color: $--color-primary;
        }
      }
    }
    .progress {
      flex: 1;
      height: 100%;
      display: flex;
      justify-content: center;
      align-items: center;
    }
  }

  .tasks {
    overflow: hidden;
    flex: 1;
    .cards {
      padding: 12px;
      height: 100%;
      overflow-x: hidden;
      overflow-y: auto;
      &::-webkit-scrollbar {
        display: none;
      }
    }
  }

  .bottom {
    width: 100%;
    border-top: 1px solid #f7f7f7;
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: #262626;
    padding: 8px 12px;
    font-size: 14px;
  }
}
</style>
