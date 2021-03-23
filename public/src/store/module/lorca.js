import { clone } from "../../utils/common.js";
export default {
  state() {
    return {};
  },
  mutations: {
    runTask(state, val) {
      let { index, task } = val;
      let t = clone(task);
      t.index = index;
      window.RunTask(JSON.stringify(t));
    },
    stopTask(state, index) {
      window.StopTask(JSON.stringify(index));
    },
    openFile(state, index) {
      window.OpenFile(JSON.stringify(index));
    },
    saveJsonFile(state, tasks) {
      window.SaveJsonFile(JSON.stringify(tasks));
    }
  }
};
