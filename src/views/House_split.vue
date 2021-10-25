<!-- 添加个人信息-->


<template>
  <el-radio-group v-model="labelPosition" size="small">
  <el-radio-button label="left">左对齐</el-radio-button>
  <el-radio-button label="right">右对齐</el-radio-button>
  <el-radio-button label="top">顶部对齐</el-radio-button>
</el-radio-group>
<div style="margin: 20px;"></div>
<el-form :label-position="labelPosition" label-width="80px" :model="formLabelAlign" :rules="HoustRules" ref="House_split">
  <el-form-item prop="CID" label="身份证号">
    <el-input v-model="formLabelAlign.CID"></el-input>
  </el-form-item>
<el-form-item prop="cname" label="姓名">
    <el-input v-model="formLabelAlign.cname"></el-input>
  </el-form-item>
 <el-form-item prop="cnature" label="客户性质">
    <el-cascader :options="cnature" v-model="formLabelAlign.cnature"></el-cascader>
  </el-form-item>
  <el-form-item prop="ccompany" label="单位名称">
    <el-input v-model="formLabelAlign.ccompany"></el-input>
  </el-form-item>
  <el-form-item prop="ctel" label="座机">
    <el-input v-model="formLabelAlign.ctel"></el-input>
  </el-form-item>
  <el-form-item prop="ctelephone" label="移动电话">
    <el-input v-model="formLabelAlign.ctelephone"></el-input>
  </el-form-item>
   <el-form-item prop="caddress" label="地址">
    <el-input v-model="formLabelAlign.caddress"></el-input>
  </el-form-item>
   <el-form-item prop="cpostcode" label="邮编">
    <el-input v-model="formLabelAlign.cpostcode"></el-input>
  </el-form-item>
   <el-form-item prop="cemail" label="邮箱">
    <el-input v-model="formLabelAlign.cemail"></el-input>
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
        cnature:[
			{
				value:'家庭用户',
				label:'家庭用户'
			},
			{
				value:'单位用户',
				label:'单位用户'
			},
            {
				value:'签约用户',
				label:'签约用户'
			},
            {
				value:'代理商',
				label:'代理商'
			}
		],
		
        labelPosition: 'right',
        formLabelAlign: {
         CID:'',
         cname:'',
         cnature:'',
         ccompany:'',
         ctel:'',
         ctelephone:'',
         caddress:'',
         cpostcode:'',
         cemail:'',
        },
        HoustRules:{
          CID:[
            {
            required: true,
            message: "请输入身份证号",
            trigger: "blur",
          },
          {
            min: 18,
            max: 18,
            message: "应为18个数字",
            trigger: "blur",
          },
          ],
          cname:[
            {
            required: true,
            message: "请输入姓名",
            trigger: "blur",
          },
          ],
         cnature: [
            {
            required: true,
            message: "请输入客户性质",
          trigger: "blur",
          },
          ],
         ccompany:[
            {
            required: true,
            message: "请输入单位名称",
            trigger: "blur",
          },
          ],
        
          ctelephone:[
            {
            required: true,
            message: "请输入移动电话",
            trigger: "blur",
          },
        ],
         caddress:[
            {
            required: true,
            message: "请输入地址",
            trigger: "blur",
          },
        ],
     
         cemail:[
            {
            required: true,
            message: "请输入邮箱",
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