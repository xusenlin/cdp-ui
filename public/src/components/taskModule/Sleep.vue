<template>
  <div>
    <el-alert
      title="此子任务可以让任务等待一段时间，可以放在比如执行一些JS代码后"
      type="info"
      show-icon
    />
    <el-input
      style="margin-top: 20px"
      placeholder="请输入等待时间"
      v-model.number="modelValue.second"
    >
      <template #append>秒</template>
    </el-input>
  </div>
</template>

<script>
export default {
  name: "Sleep",
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {
          second: ""
        };
      }
    }
  }
};
</script>
