<template>
  <el-radio-group v-model="labelPosition" size="small">
  <el-radio-button label="left">左对齐</el-radio-button>
  <el-radio-button label="right">右对齐</el-radio-button>
  <el-radio-button label="top">顶部对齐</el-radio-button>
</el-radio-group>
<div style="margin: 20px;"></div>
<el-form :label-position="labelPosition" label-width="80px" :model="formLabelAlign" :rules="HoustRules" ref="House_split">
  <el-form-item prop="No" label="电话号码">
    <el-input v-model="formLabelAlign.No"></el-input>
  </el-form-item>
  <el-form-item class="btn">
          <el-button @click="submit" type="primary">提交</el-button>
        </el-form-item>
</el-form>
</template>

<script>
  export default {
    data() {
      return {
        labelPosition: 'right',
        formLabelAlign: {
          No: '',
        },
        HoustRules:{
          No:[
            {
            required: true,
            message: "请输入电话号码",
            trigger: "blur",
          },
          {
            min: 11,
            max: 11,
            message: "应为11个数字",
            trigger: "blur",
          },
          ],
      },
    };
  },
  methods: {
    submit() {
      this.$refs.House_split.validate(async (valid) => {
        if (!valid) return;
        const { data: res } = await this.$http.post(
          "api/consumer/delete",
          this.formLabelAlign
        );
        this.$message.success("删除成功");
      });
    },
  }
}
</script>