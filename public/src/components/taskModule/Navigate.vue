<template>
  <div>
    <el-input
      style="margin-top: 20px"
      placeholder="请输入网址"
      v-model="modelValue.host"
    >
    </el-input>
  </div>
</template>

<script>
export default {
  name: "Navigate",
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {
          host: ""
        };
      }
    }
  }
};
</script>
