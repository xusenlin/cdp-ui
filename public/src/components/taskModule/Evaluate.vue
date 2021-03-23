<template>
  <div>
    <el-alert
      title="你最好先将你的JS代码在控制台跑一下看它是否能正常运行"
      type="info"
      show-icon
    />
    <el-input
      style="margin-top: 20px"
      placeholder="文件名，会帮你保存Js执行结果"
      v-model="modelValue.fileName"
    >
    </el-input>
    <el-input
      style="margin-top: 20px"
      type="textarea"
      :rows="5"
      v-model="modelValue.script"
      placeholder="Js Code"
    ></el-input>
  </div>
</template>

<script>
export default {
  name: "Evaluate",
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {
          fileName: "",
          script: ""
        };
      }
    }
  }
};
</script>
