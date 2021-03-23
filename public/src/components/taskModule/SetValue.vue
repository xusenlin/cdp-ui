<template>
  <div>
    <SelTip />
    <el-input
      style="margin-top: 20px"
      placeholder="请输入节点匹配选择器，如:#id > div.mask"
      v-model="modelValue.sel"
    >
    </el-input>
    <el-input
      style="margin-top: 20px"
      placeholder="请输入要输入的内容"
      v-model="modelValue.val"
    >
    </el-input>
  </div>
</template>

<script>
import SelTip from "../base/SelTip.vue";
export default {
  name: "SetValue",
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {
          sel: "",
          val: ""
        };
      }
    }
  },
  components: { SelTip }
};
</script>
