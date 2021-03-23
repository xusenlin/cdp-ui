<template>
  <el-drawer
    :title="index === -1 ? '添加子任务' : '编辑子任务'"
    size="46%"
    v-model="show"
    direction="rtl"
    destroy-on-close
  >
    <div class="sub-task">
      <el-form
        label-position="top"
        :model="form"
        :rules="rules"
        ref="editSubTask"
      >
        <el-form-item label="选择子任务：" prop="name">
          <div style="display: flex;justify-content: space-between">
            <el-select
              @change="changeSelectTask"
              v-model="form.name"
              placeholder="请选择子任务"
            >
              <el-option
                v-for="item in tasks"
                :key="item.name"
                :label="item.name"
                :value="item.name"
              >
                <span style="float: left">{{ item.name }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">{{
                  item.title
                }}</span>
              </el-option>
            </el-select>
            <div class="sub-task-title">
              任务标题：{{ form.title ? form.title : "" }}
            </div>
          </div>
        </el-form-item>
        <el-form-item label="自定义描述：" prop="desc">
          <el-input type="textarea" v-model="form.desc"></el-input>
        </el-form-item>
        <el-form-item label="参数配置：" prop="args">
          <component
            v-if="form.name"
            :key="form.name"
            v-model="form.args"
            :is="tasks.filter(r => r.name === form.name)[0].argsModule"
          />
        </el-form-item>
      </el-form>
      <div class="dialog-footer">
        <el-button @click="show = false">取 消</el-button>
        <el-button type="primary" @click="submitForm()">确 定</el-button>
      </div>
    </div>
  </el-drawer>
</template>

<script>
import { markRaw } from "vue";
import MapEdit from "../base/ArrMapEdit.vue";
import tasks from "../../config/tasks.js";
import { Required, FillerFieldRules } from "../../utils/validateRules.js";
import { resetArgs, fillerLeft, clone } from "../../utils/common.js";

export default {
  name: "EditSubTask",
  components: { MapEdit },
  emits: ["submitData"],
  data() {
    return {
      tasks: markRaw(tasks),
      show: false,
      parentIndex: -1,
      index: -1,
      form: {
        name: "",
        desc: "",
        title: "",
        args: {}
      },
      rules: {
        ...FillerFieldRules(["name"], Required)
      }
    };
  },
  methods: {
    Add(parentIndex) {
      this.index = -1;
      this.parentIndex = parentIndex;
      resetArgs(this.form);
      this.$nextTick(() => {
        this.show = true;
      });
    },
    Edit(t, parentIndex, index) {
      this.index = index;
      this.parentIndex = parentIndex;
      fillerLeft(this.form, t);
      this.$nextTick(() => {
        this.show = true;
      });
    },
    submitForm() {
      this.$refs["editSubTask"].validate(valid => {
        if (valid) {
          this.$emit("submitData", clone(this.form), {
            parentIndex: this.parentIndex,
            index: this.index
          });
          this.$refs["editSubTask"].resetFields();
          this.$nextTick(() => {
            this.show = false;
          });
        } else {
          return false;
        }
      });
    },
    changeSelectTask(v) {
      if (v) {
        this.form.title = this.tasks.filter(r => r.name === v)[0].title;
      } else {
        this.form.title = "";
      }
    }
  }
};
</script>

<style scoped lang="scss">
.sub-task {
  padding: 20px;
  height: calc(100vh - 75px);
  overflow-x: auto;
  .sub-task-title {
    flex: 1;
    height: 40px;
    line-height: 40px;
    margin-left: 40px;
    background: #f7f7f7;
    padding: 0 20px;
    color: #8c8c8c;
    border-radius: 4px;
  }
}
</style>
