<!-- 添加个人信息-->


<template>
  <el-radio-group v-model="labelPosition" size="small">
  <el-radio-button label="left">左对齐</el-radio-button>
  <el-radio-button label="right">右对齐</el-radio-button>
  <el-radio-button label="top">顶部对齐</el-radio-button>
</el-radio-group>
<div style="margin: 20px;"></div>
<el-form :label-position="labelPosition" label-width="80px" :model="formLabelAlign" :rules="HoustRules" ref="House_split">
   <el-form-item prop="Sname" label="备件名称" >
    <el-input v-model="formLabelAlign.Sname"></el-input>
    </el-form-item>
 <el-form-item prop="Smodel" label="型号">
    <el-input v-model="formLabelAlign.Smodel"></el-input>
  </el-form-item>
 <el-form-item prop="Snumber" label="数量">
    <el-input v-model="formLabelAlign.Snumber"></el-input>
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
            Sname:'',
            Smodel:'',
            Snumber:'',
        },
        HoustRules:{
          Sname:[
            {
            required: true,
            message: "请输入备件名称",
            trigger: "blur",
          },
          ],
              Smodel:[
            {
            required: true,
            message: "请输入型号",
            trigger: "blur",
          },
          ],
          Snumber:[
            {
            required: true,
            message: "请输入数量",
            trigger: "blur",
          },
          ],
      //          unit_price:[
      //       {
      //       required: true,
      //       message: "请输入单价",
      //       trigger: "blur",
      //     },
      //     ],
      //       warning_number:[
      //       {
      //       required: true,
      //       message: "请输入警戒数量",
      //       trigger: "blur",
      //     },
      //     ],
      //     state:[
      //       {
      //       required: true,
      //       message: "请选择库存状态",
      //       trigger: "blur",
      //     },
      //     ],
       },
    };
  },
  methods: {
    submit() {
      this.$refs.House_split.validate(async (valid) => {
        if (!valid) return;
        const { data: res } = await this.$http.post(
          "repairock/in",
          this.formLabelAlign
        );
        if (res.code == 5001) {
          this.$message.error("备件名称不能为空");
          return;
        }
        if (res.code == 5002) {
          this.$message.error("数量不能为空");
          return;
        }
         if (res.code == 5003) {
          this.$message.error("备件数量不能为负数");
          return;
        }
         if (res.code == 5004) {
          this.$message.error("系统中无该备件，请先增加该备件");
          return;
        }
        if (res.code == 200) {
          this.$message.success("入库成功");
          return;
        }
        this.$message.error("后台错误");
      });
    },
  }
}
</script>