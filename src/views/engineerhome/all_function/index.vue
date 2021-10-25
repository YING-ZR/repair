<template>
  <el-radio-group v-model="labelPosition" size="small">
   <el-radio-button label="left">左对齐</el-radio-button>
  <el-radio-button label="right">右对齐</el-radio-button> 
  <el-radio-button label="top">顶部对齐</el-radio-button> 
 </el-radio-group> 
<div style="margin: 20px;"></div>

<el-form :label-position="labelPosition" label-width="150px" :model="formLabelAlign" :rules="HoustRules" ref="House_split">
  <el-form-item prop='Dtable_number' v-model="formLabelAlign.Dtable_number" label="维修单号" >{{Dtable_number}}

<!--  查看详细信息-->

    <el-button type="primary" @click="Chakan(formLabelAlign.Dtable_number),dialogTableVisible = true" plain>查看详细信息</el-button>

    <el-dialog title="查看详细信息" v-model="dialogTableVisible">
      
     <el-table :data="gridData"> 
 <el-table-column prop="Rnumber" v-model="formLabelAlign.table_number" label="维修单号" width="180"></el-table-column>
    <el-table-column prop="Rtime" v-model="formLabelAlign.Rtime" label="提交时间" width="180"></el-table-column>
     <el-table-column prop="Rid"  v-model="formLabelAlign.Rid" label="用户身份证号" width="180"></el-table-column>
      <el-table-column prop="Rphenomenon" v-model="formLabelAlign.Rphenomenon" label="机器故障现象" width="180"></el-table-column>
       <el-table-column prop="Rfault"  v-model="formLabelAlign.Rfault" label="故障类型" width="180"></el-table-column>
        <el-table-column prop="Rproduct"  v-model="formLabelAlign.Rproduct" label="选择产品类型" width="180"></el-table-column>
       <el-table-column prop="Rremarks"  v-model="formLabelAlign.Rremarks" label="备注" width="180"></el-table-column>
  </el-table>
</el-dialog>
           
  </el-form-item>

<!-- </el-form-item> -->

 <el-form-item prop='Dengineer_number' v-model="formLabelAlign.Dengineer_number" label="技术工程师员工号" >{{Dengineer_number}}  </el-form-item>
  
  <el-form-item prop='Dtime' v-model="formLabelAlign.Dtime" label="分配技术工程师时间" >{{Dtime}}  </el-form-item>

 <el-form-item prop='Dsparepart_name' v-model="formLabelAlign.Dsparepart_name" label="备件名称" >{{Dsparepart_name}}  
   <!--  填写备件信息按钮 -->
 <el-button type="primary" @click="beijian()" plain>填写备件信息</el-button>
<el-dialog title="填写备件信息" v-model="dialogFormVisible">
  <el-form :model="form" :rules="controlRules" ref="form">
    <el-form-item label="备件名称" label-width="100px" prop="Dsparepart_name">
      <el-input v-model="form.Dsparepart_name" autocomplete="off"></el-input>
    </el-form-item>
    <el-form-item label="数量" label-width="101px" prop="Snumber">
      <el-input v-model="form.Snumber" autocomplete="off"></el-input>
    </el-form-item>
  </el-form>
  <template #footer>
    <span class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="primary" @click="tianxie(formLabelAlign.Dtable_number)">提交</el-button>
    </span>
  </template>
</el-dialog>
 </el-form-item>

  <el-form-item prop='Dstatus' v-model="formLabelAlign.Dstatus" label="维修状态" >{{Dstatus}}  
    <!--维修完成-->
      <el-button @click="split(formLabelAlign.Dtable_number)" type="primary">维修完成<i class="el-icon-upload el-icon--right"></i></el-button>
  </el-form-item>

  
<el-form-item prop="Ddelay_state" label="维修延迟程度">
  <span>{{Ddelaystate}}</span>
   <el-cascader :options="Ddelay_state" v-model="formLabelAlign.Ddelay_state" ></el-cascader>
  </el-form-item>
  <el-form-item class="btn">
          <el-button class= "tijiao" @click="submit(formLabelAlign.Dtable_number)" type="primary">提交</el-button>
        </el-form-item>
      
</el-form>
</template>


<script>
import Cookie from 'js-cookie'
  export default {
    data() {
      return {
        Number:Cookie.get("number"),
      Ddelay_state:[
          {
              value:'一般',
              label:'一般'
            },
           {
              value:'比较严重',
              label:'比较严重'
            },
           {
              value:'严重',
              label:'严重'
            },
       
        ],
       gridData: [
        // {
        //  Rnumber:'',
        //  Rtime:'',
        //  Rid:'',
        //  Rphenomenon:'',
        //  Rfault:'',
        //  Rproduct:'',
        //  Rremarks:'',
        // },
         ],
       tableData:[],
            addForm: {
					Hno: '',
					Harea: '',
					Hrent_m: '',
					Hdistribute: '',
				},
        labelPosition: 'right',
         dialogTableVisible: false,
        dialogFormVisible: false,

//填写备件信息 
         form: {
          Dsparepart_name:'',
          Snumber:'',
        },


        formLabelAlign: {
           Dtable_number:'',
           Dengineer_number:'',
           Dtime:'',
           Dsparepart_name:'',
           Dstatus:'',
           Ddelay_state:'',
        },
    
        HoustRules:{
       
          Ddelay_state:[
            {
            required: true,
            message: "请选择维修延迟程度",
            trigger: "blur",
          },
          ],
          
      },
    };
  },
  computed:{
      showNumber() {
        // this.Number = Cookie.get("number")
        return Cookie.get('number')
      },
      Dtable_number(){
        return this.formLabelAlign.Dtable_number
      },
      Dengineer_number(){
        return this.formLabelAlign.Dengineer_number
      },
      Dtime(){
        return this.formLabelAlign.Dtime
      },
      Dsparepart_name(){
        return this.formLabelAlign.Dsparepart_name
      },
      Dstatus(){
         return this.formLabelAlign.Dstatus
      },
      table_number(){
         return this.formLabelAlign.table_number
      },
Rtime(){
         return this.formLabelAlign.Rtime
      },
Rid(){
         return this.formLabelAlign.Rid
      },
Rphenomenon(){
         return this.formLabelAlign.Rphenomenon
      },
Rfault(){
         return this.formLabelAlign.Rfault
      },
Rproduct(){
         return this.formLabelAlign.Rproduct
      },
Rremarks(){
         return this.formLabelAlign.Rremarks
      },
      Ddelaystate(){
         return this.formLabelAlign.Ddelay_state
      },
  },
//     watch: {
//   showNumber: {
//     deep: true,
//     handler: function (newVal) {
//       console.log(newVal)
//       this.Number = newVal
//     }
//   }
// },
  // computed:{
  //       number=Cookie.get('number'),
  //   },
  //查看详细信息
   mounted() {
    this.$http
      .get("repair/staff/engineer/GetDistribution/"+this.Number)
      .then(response => (
        this.formLabelAlign = response.data));
        // console.log(this.Number)
          },
  methods: {
    beijian(){
      console.log(this.formLabelAlign.Dsparepart_name)
      if(this.formLabelAlign.Dsparepart_name !='无                   '){
        this.dialogFormVisible = false;
        this.$message.error("您已填写备件");
      }else{
        this.dialogFormVisible = true;
      }
    },
    
//填写备件信息按钮
       handleClick(row) {
        console.log(row);
      },
      submit(Dtable_number){
        if (this.formLabelAlign.Ddelay_state == '一般'){
          this.formLabelAlign.Ddelay_state = '一般';
        }else if(this.formLabelAlign.Ddelay_state == '比较严重'){
          this.formLabelAlign.Ddelay_state = '比较严重';
        }else if(this.formLabelAlign.Ddelay_state == '严重'){
          this.formLabelAlign.Ddelay_state = '严重';
        }
        this.$refs.House_split.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.put("repair/staff/engineer/Update_Distribution_ddelay_state/"+Dtable_number, this.formLabelAlign);
					if (res.code == 200) {
						// this.dialogEditVisible = false;
						return this.$message.success("编辑成功");
					}
					this.$message.error("编辑失败");
				})
      },

    //维修完成按钮
     split(Dtable_number) {   
        this.$http.put("repair/staff/engineer/Update_Distribution_status/"+Dtable_number);
					// if (res.code == 200) {
					// 	this.dialogEditVisible = false;
					// 	return this.$message.success("编辑成功");
					// }
					// this.$message.error("编辑失败");
			},
      Chakan(Dtable_number) {
        this.$http.get("repair/staff/engineer/GetTable/"+Dtable_number)
        .then(response =>(
          this.gridData = response.data
        ));
      },
      tianxie(Dtable_number){
           this.$refs.form.validate(async (valid) => {
        if (!valid) return;
        const { data: res } = await this.$http.post(
          "repair/staff/engineer/Get_Sparepart/"+Dtable_number,
          this.form
        );
         if (res.code == 60004) {
          this.$message.error("系统中无该备件");
          return;
        }
         if (res.code == 60005) {
          this.$message.error("库存数量不足");
          return;
        }
     
       //this.$message.success("备件成功");
       console.log(res.code)
        if (res.code == 200) {
          this.$message.success("备件成功");
          return;
        }
         this.$message.error("后台错误");
      });
      }
  }
}
</script>

<style>
.el-input{
    width:600px;
}
.el-cascader{
    width:300px;
}
.el-button{
  margin-left:60px;
}

</style>