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
      [
      <el-tag
        v-for="(v, i) in modelValue"
        class="tag"
        closable
        @close="delTag(i)"
        >{{ v }}
      </el-tag>
      ]
    </div>
  </div>
</template>

<script>
export default {
  name: "ArrMapEdit",
  props: {
    modelValue: {
      type: Array,
      default: () => []
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
      let k = null;
      if (this.val === "true") {
        k = { [this.key]: true };
      } else if (this.val === "false") {
        k = { [this.key]: false };
      } else if (/^\d+\b$/.test(this.val)) {
        k = { [this.key]: Number(this.val) };
      } else {
        k = { [this.key]: this.val };
      }
      this.$emit("update:modelValue", [...this.modelValue, k]);
      this.$nextTick(() => {
        this.key = "";
        this.val = "";
      });
    },
    delTag(i) {
      let newMap = [];
      this.modelValue.forEach((v, index) => {
        if (index === i) return;
        newMap.push(v);
      });
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
