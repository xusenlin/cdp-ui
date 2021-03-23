<template>
  <div>
    <el-alert title="在显示浏览器模式下无法运行此子任务" type="info" show-icon>
    </el-alert>
    <el-input
      style="margin-top: 20px"
      placeholder="请输入生成的文件名如:x.pdf"
      v-model="modelValue.fileName"
    >
      <template #prepend>文件名:</template>
    </el-input>
  </div>
</template>

<script>
export default {
  name: "Pdf",
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
