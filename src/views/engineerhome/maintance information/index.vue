<template>
  <el-radio-group v-model="labelPosition" size="small">
  <!-- <el-radio-button label="left">左对齐</el-radio-button>
  <el-radio-button label="right">右对齐</el-radio-button> -->
  <el-radio-button label="top">顶部对齐</el-radio-button>
</el-radio-group>
<div style="margin: 20px;"></div>
<el-form :label-position="labelPosition" label-width="80px" :model="formLabelAlign" :rules="HoustRules" ref="House_split">
  <el-form-item prop="detection_record" label="检修记录" >
    <el-input v-model="formLabelAlign.detection_record" ></el-input>
  </el-form-item>
     <el-form-item prop="maintance_record" label="维修记录" >
    <el-input v-model="formLabelAlign.maintance_record"></el-input>
  </el-form-item>
  <el-form-item prop="region" label="维修延迟程度">
   <el-cascader :options="region" v-model="formLabelAlign.region" ></el-cascader>
  </el-form-item>
   
  <el-form-item prop="state" label="维修状态">
   <el-cascader :options="state" v-model="formLabelAlign.state" ></el-cascader>
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
       region:[
          {
              value:'normal',
              label:'一般'
            },
           {
              value:'serious',
              label:'比较严重'
            },
           {
              value:'hard',
              label:'严重'
            },
       
        ],
        state:[
            {
              value:'undistributed',
              label:'未分配'
            },
            {
              value:'distributed',
              label:'已分配'
            },
            {
              value:'maintenanced',
              label:'已维修'
            },
            {
              value:'revocation',
              label:'撤销'
            },
          
        ],
        labelPosition: 'right',
        formLabelAlign: {
      detection_record:'',
      maintance_record:'',
      region:'',
      state:'',
      rremarks:'',
        },
    
        HoustRules:{
          detection_record:[
            {
            required: true,
            message: "请输入检测记录",
            trigger: "blur",
          },
          ],
              maintance_record:[
            {
            required: true,
            message: "请输入维修记录",
            trigger: "blur",
          },
          ],
              region:[
            {
            required: true,
            message: "请输入维修延迟程度",
            trigger: "blur",
          },
          ],
            state:[
            {
            required: true,
            message: "请输入维修状态",
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