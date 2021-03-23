<template>
  <div>
    根据当前浏览器窗口大小截取内容
    <el-input
      style="margin-top: 10px"
      placeholder="请输入图片文件名"
      v-model="modelValue.fileName"
    >
    </el-input>
  </div>
</template>

<script>
export default {
  name: "CaptureScreenshot",
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {
          fileName: ""
        };
      }
    }
  }
};
</script>
