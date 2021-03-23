<template>
  <el-dialog
    :title="index === -1 ? '添加任务' : '编辑任务'"
    v-model="show"
    width="60%"
  >
    <div>
      <el-form
        label-position="left"
        :model="form"
        :rules="rules"
        ref="editTask"
        label-width="160px"
      >
        <el-form-item label="名称：" prop="title">
          <el-input v-model="form.title"></el-input>
        </el-form-item>
        <el-form-item label="超时：" prop="timeOut">
          <el-input v-model.number="form.timeOut"></el-input>
        </el-form-item>
        <el-form-item label="是否显示浏览器：" prop="showBrowser">
          <el-switch
            v-model="form.showBrowser"
            active-text="显示"
            inactive-text="隐藏"
          >
          </el-switch>
        </el-form-item>
        <el-form-item label="浏览器宽：" prop="width">
          <el-input v-model.number="form.width"></el-input>
        </el-form-item>
        <el-form-item label="浏览器高：" prop="height">
          <el-input v-model.number="form.height"></el-input>
        </el-form-item>
        <el-form-item label="浏览器userAgent：" prop="userAgent">
          <el-input v-model.number="form.userAgent"></el-input>
        </el-form-item>
        <el-form-item label="浏览器启动flags：" prop="flags">
          <MapEdit v-model="form.flags" />
        </el-form-item>
      </el-form>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="show = false">取 消</el-button>
        <el-button type="primary" @click="submitForm()">确 定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
import MapEdit from "../base/ArrMapEdit.vue";
import {
  Required,
  FillerFieldRules,
  StrLen3To18
} from "../../utils/validateRules.js";
import { resetArgs, fillerLeft, clone } from "../../utils/common.js";

export default {
  name: "EditTask",
  components: { MapEdit },
  emits: ["submitData"],
  data() {
    return {
      show: false,
      index: -1, //-1是添加，否则就是修改
      form: {
        progress: -1,
        title: "",
        timeOut: 30,
        showBrowser: false,
        width: 0,
        height: 0,
        userAgent: "",
        tasks: [],
        flags: []
      },
      rules: {
        ...FillerFieldRules(["timeOut", "showBrowser"], Required),
        title: [StrLen3To18, Required]
      }
    };
  },
  methods: {
    Add() {
      this.index = -1;
      resetArgs(this.form);
      this.$nextTick(() => {
        this.show = true;
      });
    },
    Edit(t, index) {
      this.index = index;
      fillerLeft(this.form, t);
      this.$nextTick(() => {
        this.show = true;
      });
    },
    submitForm() {
      this.$refs["editTask"].validate(valid => {
        if (valid) {
          this.form.progress = -1;
          this.$emit("submitData", clone(this.form), this.index);
          this.$refs["editTask"].resetFields();
          this.$nextTick(() => {
            this.show = false;
          });
        } else {
          return false;
        }
      });
    }
  }
};
</script>
