<template>
  <div>
    <div class="action">
      <el-input
        v-model="key"
        style="margin-right: 10px"
        placeholder="请输入key"
      ></el-input>
      <el-input
        v-model="val"
        style="margin-right: 10px"
        placeholder="请输入value"
      ></el-input>
      <el-button
        type="primary"
        icon="el-icon-circle-plus-outline"
        @click="addKey"
        >添加</el-button
      >
    </div>
    <div class="tags">
      {{ "{" }}
      <el-tag
        v-for="(v, k) in modelValue"
        class="tag"
        closable
        @close="delTag(k)"
        >{{ k }}:{{ v }}
      </el-tag>
      {{ "}" }}
    </div>
  </div>
</template>

<script>
import { clone } from "../../utils/common.js";
export default {
  name: "MapEdit",
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {};
      }
    }
  },
  emits: ["update:modelValue"],
  data() {
    return {
      key: "",
      val: ""
    };
  },
  methods: {
    addKey() {
      this.$emit("update:modelValue", {
        ...(this.modelValue || {}),
        ...{ [this.key]: this.val }
      });
      this.$nextTick(() => {
        this.key = "";
        this.val = "";
      });
    },
    delTag(k) {
      let newMap = clone(this.modelValue);
      delete newMap[k];
      this.$emit("update:modelValue", newMap);
    }
  }
};
</script>

<style scoped lang="scss">
.action {
  display: flex;
  margin-bottom: 20px;
}
.tags {
  .tag {
    margin-right: 10px;
    &:last-child {
      margin-right: 0;
    }
  }
}
</style>
