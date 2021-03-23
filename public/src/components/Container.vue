<template>
  <div>
    <ul class="tasks-list">
      <li v-for="(task, i) in tasks">
        <Tasks
          :index="i"
          :task="task"
          @editTask="editTask"
          @delTask="delTask"
        />
      </li>
      <li class="row">
        <AddTaskPanel @addTask="addTask" />
      </li>
    </ul>
    <EditTask ref="EditTaskRef" @submitData="submit" />
  </div>
</template>

<script>
import { mapState } from "vuex";
import Tasks from "./taskPanel/Tasks.vue";
import AddTaskPanel from "./taskPanel/AddTaskPanel.vue";
import EditTask from "./taskPanel/EditTask.vue";

export default {
  name: "Container",
  components: { Tasks, AddTaskPanel, EditTask },
  computed: mapState(["tasks"]),
  methods: {
    editTask(index) {
      let task = this.tasks[index];
      this.$refs.EditTaskRef.Edit(task, index);
    },
    delTask(index) {
      this.$confirm("确定删除此任务吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          this.$store.commit("delTask", index);
          this.$notify({
            title: "成功",
            message: "删除任务成功",
            type: "success"
          });
        })
        .catch(() => {});
    },
    addTask() {
      this.$refs.EditTaskRef.Add();
    },
    submit(task, index) {
      index === -1
        ? this.$store.commit("addTask", task)
        : this.$store.commit("editTask", { task, index });
    }
  }
};
</script>

<style scoped lang="scss">
.tasks-list {
  min-width: 100vw;
  height: calc(100vh - 52px);
  padding: 20px;
  display: flex;
  overflow-x: auto;
  li {
    margin-right: 10px;
    list-style-type: none;
  }
}
</style>
