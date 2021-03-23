<template>
  <div class="cookies">
    <div class="tags">
      [
      <el-tag
        v-for="(v, i) in modelValue.cookies"
        class="tag"
        closable
        @close="delTag(i)"
        >{{ v }}
      </el-tag>
      ]
    </div>

    <div class="action">
      <el-input
        v-model="cookie.name"
        style="margin-right: 10px"
        placeholder="请输入cookie name"
      ></el-input>
      <el-input
        v-model="cookie.value"
        style="margin-right: 10px"
        placeholder="请输入cookie value"
      ></el-input>
      <!--      <el-input-->
      <!--        v-model="cookie.url"-->
      <!--        style="margin-right: 10px"-->
      <!--        placeholder="请输入cookie url"-->
      <!--      ></el-input>-->
      <el-input
        v-model="cookie.domain"
        style="margin-right: 10px"
        placeholder="请输入cookie domain"
      ></el-input>
      <el-button
        type="primary"
        icon="el-icon-circle-plus-outline"
        @click="addKey"
        >添加</el-button
      >
    </div>
    <!--    【name，value，url，domain】 url和domain可为空-->
  </div>
</template>

<script>
import { clone } from "../../utils/common.js";
export default {
  name: "AddCookies",
  emits: ["update:modelValue"],
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {
          cookies: []
        };
      }
    }
  },
  data() {
    return {
      cookie: {
        name: "",
        value: "",
        // url: "xusenlin.com",
        domain: "" //坑爹，chromedp说可以为空，但是源代码必须要传
        // path   : "",
        // secure    : "",
        // httpOnly   : false,
      }
    };
  },
  methods: {
    addKey() {
      let { name, value, domain } = this.cookie;
      if (!name || !value || !domain) {
        this.$message.info("请填写完整cookie需要的字段");
        return;
      }

      let newVal = {
        cookies: [...(this.modelValue.cookies || []), clone(this.cookie)]
      };
      this.$emit("update:modelValue", newVal);
      this.$nextTick(() => {
        this.cookie.name = "";
        this.cookie.value = "";
      });
    },
    delTag(i) {
      let newVal = [];
      this.modelValue.cookies.forEach((v, index) => {
        if (index === i) return;
        newVal.push(v);
      });
      this.$emit("update:modelValue", { cookies: newVal });
    }
  }
};
</script>
<style scoped lang="scss">
.cookies {
  .tags {
    .tag {
      margin-right: 10px;
      &:last-child {
        margin-right: 0;
      }
    }
  }
  .action {
    margin-top: 40px;
    display: flex;
    margin-bottom: 20px;
  }
}
</style>
