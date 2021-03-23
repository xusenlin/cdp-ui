import { ElNotification } from "element-plus";

export default {
  state() {
    return [];
  },
  mutations: {
    addTask(state, task) {
      state.push(task);
    },
    editTask(state, val) {
      let { task, index } = val;
      state[index] = task;
    },
    delTask(state, index) {
      state.splice(index, 1);
    },
    addSubTask(state, val) {
      let { task, parentIndex } = val;

      state[parentIndex].tasks.push(task);
    },
    editSubTask(state, val) {
      let { task, parentIndex, index } = val;
      state[parentIndex].tasks[index] = task;
    },
    delSubTask(state, val) {
      let { parentIndex, index } = val;
      state[parentIndex].tasks.splice(index, 1);
    },
    updateProgress(state, val) {
      let { taskIndex, progress } = val;
      state[taskIndex].progress = parseInt(progress);
    },
    initTask(state, t) {
      state.push(...t);
    }
  },
  actions: {
    initTaskByJsonFile({ commit }) {
      if (!window.hasOwnProperty("GetJsonTask")) {
        commit("initTask", []);
        return;
      }
      window
        .GetJsonTask()
        .then(r => {
          let task = JSON.parse(r);

          commit("initTask", task || []);
        })
        .catch(e => {
          ElNotification({
            message: JSON.stringify(e),
            type: "error",
            title: "json 初始化任务出错",
            duration: 4000
          });
          commit("initTask", []);
        });
    }
  }
};
