<template>
  <div>
    <SelTip />
    <el-switch
      style="margin-top: 20px"
      v-model="modelValue.fullPage"
      active-text="截取整个页面"
      inactive-text="截取某个元素"
    >
    </el-switch>
    <el-input
      style="margin-top: 20px"
      placeholder="请输入图片文件名"
      v-model="modelValue.fileName"
    >
    </el-input>
    <el-input
      v-if="!modelValue.fullPage"
      style="margin-top: 20px"
      placeholder="请输入元素节点匹配选择器，如:#id > div.mask"
      v-model="modelValue.sel"
    >
    </el-input>
  </div>
</template>

<script>
import SelTip from "../base/SelTip.vue";
export default {
  name: "Screenshot",
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {
          sel: "",
          fullPage: true,
          fileName: ""
        };
      }
    }
  },
  components: { SelTip }
};
</script>
