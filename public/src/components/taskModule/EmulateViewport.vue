<template>
  <div class="emulate-viewport">
    <el-input placeholder="请输入浏览器宽度" v-model.number="modelValue.width">
    </el-input>
    <div style="width: 40px;height: 1px"></div>
    <el-input placeholder="请输入浏览器高度" v-model.number="modelValue.height">
    </el-input>
  </div>
</template>

<script>
export default {
  name: "EmulateViewport",
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {
          width: 0,
          height: 0
        };
      }
    }
  }
};
</script>
<style scoped lang="scss">
.emulate-viewport {
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
}
</style>
