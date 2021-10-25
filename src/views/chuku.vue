<!-- 添加个人信息-->


<template>
  <el-radio-group v-model="labelPosition" size="small">
  <el-radio-button label="left">左对齐</el-radio-button>
  <el-radio-button label="right">右对齐</el-radio-button>
  <el-radio-button label="top">顶部对齐</el-radio-button>
</el-radio-group>
<div style="margin: 20px;"></div>
<el-form :label-position="labelPosition" label-width="80px" :model="formLabelAlign" :rules="HoustRules" ref="House_split">
   <el-form-item prop="replacement_name" label="备件名称" >
    <el-input v-model="formLabelAlign.replacement_name"></el-input>
    </el-form-item>
 <el-form-item prop="model" label="型号">
    <el-input v-model="formLabelAlign.model"></el-input>
  </el-form-item>
 <el-form-item prop="number" label="数量">
    <el-input v-model="formLabelAlign.number"></el-input>
  </el-form-item>
 <el-form-item prop="unit_price" label="单价">
    <el-input v-model="formLabelAlign.unit_price"></el-input>
  </el-form-item>
 <el-form-item prop="warning_number" label="警戒数量">
    <el-input v-model="formLabelAlign.warning_number"></el-input>
  </el-form-item>
   <el-form-item prop="state" label="库存状态">
    <el-cascader :options="state" v-model="formLabelAlign.state"></el-cascader>
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
  state:[
			{
				value:'正常：库存量>警戒数量',
				label:'正常：库存量>警戒数量'
			},
			{
				value:'临界：库存量=警戒数量',
				label:'临界：库存量=警戒数量'
			},
            {
				value:'警示：0<库存量<警戒数量',
				label:'警示：0<库存量<警戒数量'
			},
            {
				value:'缺货：库存量=0',
				label:'缺货：库存量=0'
			}
		],
		
        labelPosition: 'right',
        formLabelAlign: {
       replacement_name:'',
            model:'',
            number:'',
            unit_price:'',
            warning_number:'',
            statu:'',
        },
        HoustRules:{
          replacement_name:[
            {
            required: true,
            message: "请输入备件名称",
            trigger: "blur",
          },
          ],
              model:[
            {
            required: true,
            message: "请输入型号",
            trigger: "blur",
          },
          ],
          number:[
            {
            required: true,
            message: "请输入数量",
            trigger: "blur",
          },
          ],
               unit_price:[
            {
            required: true,
            message: "请输入单价",
            trigger: "blur",
          },
          ],
            warning_number:[
            {
            required: true,
            message: "请输入警戒数量",
            trigger: "blur",
          },
          ],
          state:[
            {
            required: true,
            message: "请选择库存状态",
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
          "api/consumer/register",
          this.formLabelAlign
        );
        if (res.code == 2000) {
          this.$message.error("此面积您不可租赁");
          return;
        }
        if (res.code == 2001) {
          this.$message.error("您已填写过，请等待分配");
          return;
        }
        if (res.code == 200) {
          this.$message.success("提交成功");
          // window.sessionStorage.setItem('username', res.data.user.StaffName);
          // window.sessionStorage.setItem('level', res.data.user.StaffLevel);
          // window.sessionStorage.setItem('salary', res.data.user.StaffSalary);
          // window.sessionStorage.setItem('remarks', res.data.user.StaffRemarks);
          // window.sessionStorage.setItem('telephone', res.data.user.StaffTelephone);
          // this.$router.push({ path: "/Home" });
          return;
        }
        this.$message.error("后台错误");
      });
    },
  }
}
</script>