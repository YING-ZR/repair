<template>
  <el-radio-group v-model="labelPosition" size="small">
  <!-- <el-radio-button label="left">左对齐</el-radio-button>
  <el-radio-button label="right">右对齐</el-radio-button> -->
  <el-radio-button label="top">顶部对齐</el-radio-button>
</el-radio-group>
<div style="margin: 20px;"></div>
<el-form :label-position="labelPosition" label-width="80px" :model="formLabelAlign" :rules="HoustRules" ref="House_split">
  <el-form-item prop="rphenomenon" label="机器故障现象" >
    <el-input v-model="formLabelAlign.rphenomenon"></el-input>
  </el-form-item>
    
  <el-form-item prop="rfault" label="故障类型">
   <el-cascader :options="rfault" v-model="formLabelAlign.rfault" ></el-cascader>
  </el-form-item>
   
  <el-form-item prop="rproduct" label="产品类型">
   <el-cascader :options="rproduct" v-model="formLabelAlign.rproduct" ></el-cascader>
  </el-form-item>
  <el-form-item prop="rremarks" label="备注">
    <el-input v-model="formLabelAlign.rremarks"></el-input>
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
        rfault:[
          {
            value:'间歇性故障',
            label:'间歇性故障'
          },
         {
            value:'固定性故障',
            label:'固定性故障'
          },
        ],
        rproduct:[
            {
              value:'台式机',
              label:'台式机'
            },
            {
              value:'笔记本电脑',
              label:'笔记本电脑'
            },
            {
              value:'投影仪',
              label:'投影仪'
            },
            {
              value:'打印机',
              label:'打印机'
            },
            {
              value:'其他',
              label:'其他'
            },
        ],
        labelPosition: 'right',
        formLabelAlign: {
       rphenomenon:'',
       rfualt:'',
       rproduct:'',
       rremarks:'',
        },
    
        HoustRules:{
           rphenomenon:[
            {
            required: true,
            message: "请输入故障现象",
            trigger: "blur",
          },
          ],
             rfault:[
            {
            required: true,
            message: "请输入故障类型",
            trigger: "blur",
          },
          ],
              rproduct:[
            {
            required: true,
            message: "请输入产品类型",
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
          "api/consumer/update",
          this.formLabelAlign
        );
        if (res.code == 2000) {
          this.$message.error("此面积您不可租赁");
          return;
        }
        if (res.code == 2001) {
          this.$message.error("您已填写过，请等待分房");
          return;
        }
        if (res.code == 200) {
          this.$message.success("提交成功");
          return;
        }
        this.$message.error("请等待月底分房");
      });
    },
  }
}
</script>