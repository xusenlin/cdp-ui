import { store } from "../index.js";
import { ElNotification } from "element-plus";
window.Notice = (message, type) => {
  ElNotification({ message, type, title: type.toUpperCase(), duration: 4000 });
};

window.UpdateProgress = (taskIndex, progress) => {
  store.commit("updateProgress", { taskIndex, progress });
};
