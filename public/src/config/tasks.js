import Pdf from "../components/taskModule/Pdf.vue";
import Click from "../components/taskModule/Click.vue";
import Sleep from "../components/taskModule/Sleep.vue";
import Evaluate from "../components/taskModule/Evaluate.vue";
import Navigate from "../components/taskModule/Navigate.vue";
import SetValue from "../components/taskModule/SetValue.vue";
import Screenshot from "../components/taskModule/Screenshot.vue";
import AddCookies from "../components/taskModule/AddCookies.vue";
import WaitVisible from "../components/taskModule/WaitVisible.vue";
import AddReqHeaders from "../components/taskModule/AddReqHeaders.vue";
import EmulateViewport from "../components/taskModule/EmulateViewport.vue";
import CaptureScreenshot from "../components/taskModule/CaptureScreenshot.vue";
import UnderDevelopment from "../components/taskModule/UnderDevelopment.vue";
export default [
  {
    name: "pdf",
    title: "生成PDF",
    argsModule: Pdf
  },
  {
    name: "click",
    title: "点击元素",
    argsModule: Click
  },
  {
    name: "sleep",
    title: "等待一会",
    argsModule: Sleep
  },
  {
    name: "evaluate",
    title: "执行脚本",
    argsModule: Evaluate
  },
  {
    name: "navigate",
    title: "打开网址",
    argsModule: Navigate
  },
  {
    name: "setValue",
    title: "往输入框输入内容",
    argsModule: SetValue
  },
  {
    name: "screenshot",
    title: "截取网页内容",
    argsModule: Screenshot
  },
  {
    name: "addCookies",
    title: "添加Cookies",
    argsModule: AddCookies
  },
  {
    name: "waitVisible",
    title: "等待元素加载完成",
    argsModule: WaitVisible
  },
  {
    name: "addReqHeaders",
    title: "添加请求头",
    argsModule: AddReqHeaders
  },
  {
    name: "emulateViewport",
    title: "改变窗口大小",
    argsModule: EmulateViewport
  },
  {
    name: "captureScreenshot",
    title: "捕获屏幕",
    argsModule: CaptureScreenshot
  },
  {
    name: "collectDataByExcel",
    title: "通过表格搜集数据",
    argsModule: UnderDevelopment
  },
  {
    name: "collectDataByWord",
    title: "通过文档搜集数据",
    argsModule: UnderDevelopment
  },
  {
    name: "packageFile",
    title: "打包目前生成的所有文件",
    argsModule: UnderDevelopment
  }
];
