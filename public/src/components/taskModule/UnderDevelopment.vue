<template>
  <div>
    功能正在开发中，此任务不可使用，否则任务无法执行
  </div>
</template>

<script>
export default {
  name: "UnderDevelopment",
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {};
      }
    }
  }
};
</script>
