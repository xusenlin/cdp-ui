import { createStore } from "vuex";
import tasks from "./module/tasks.js";
import lorca from "./module/lorca.js";
import "./module/LorcaToJs.js";

// Create a new store instance.
export const store = createStore({
  modules: {
    tasks,
    lorca
  }
});
