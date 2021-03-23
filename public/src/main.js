import { createApp } from "vue";
import App from "./App.vue";
import ElementPlus from "element-plus";
import { store } from "./store/index";
import "./assets/style/style.scss";
createApp(App)
  .use(store)
  .use(ElementPlus)
  .mount("#app");

store.dispatch("initTaskByJsonFile");
