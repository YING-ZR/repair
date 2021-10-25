<template>
  <div class="about">
    <!-- 卡片视图区域 -->
    <el-card>
      <!-- 用户列表区域 -->
      <el-table height="400px" :data="rightlist" border stripe style="width:100%" :header-cell-style="rowClass">
        <!-- <el-table-column type="index" label="#" width="100"></el-table-column> -->
        <el-table-column label="备件名称" prop="Dsparepart_name"></el-table-column>
        <el-table-column label="维修单号" >
          <el-table-column prop="Dtable_number"></el-table-column>
          <el-table-column>
            <template #default="scope">
            <el-button type="success" size="mini" @click="Repair_table(scope.row.Dtable_number),dialogFormVisibleWeixiu = true">查看详情信息</el-button>
            </template>
          </el-table-column>
        </el-table-column>
        <el-table-column label="技术工程师员工号">
          <el-table-column label="显示工程师员工号" prop="Dengineer_number"></el-table-column>
          <el-table-column label="显示按钮">
            <template #default="scope">
            <el-button type="success" size="mini" @click="Engineer_table(scope.row.Dengineer_number),dialogFormVisibleJishu = true;">查看详情信息</el-button>
            </template>
          </el-table-column>
        </el-table-column>
        <el-table-column
          label="分配技术工程师时间"
          prop="Dtime"
          width=180
        ></el-table-column>
        <el-table-column label="维修状态" prop="Dstatus"></el-table-column>
        <el-table-column
          label="维修延迟程度"
          prop="Ddelay_state"
        ></el-table-column>
        <el-table-column label="总费用">
          <template #default="scope">
          <el-button type="success" @click="Cost(scope.row.Dtable_number);dialogFormVisibleCost = true" size="mini">计算费用</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页区域 -->
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="pagenum"
        :page-sizes="[1, 2, 5, 10]"
        :page-size="2"
        layout="total, sizes, prev, pager, next, jumper"
        :total="400"
      >
      </el-pagination>

      <!-- 维修表对话框 -->
      <el-dialog title="维修表的信息" v-model="dialogFormVisibleWeixiu">
        <el-form :model="Form_" >
          <el-form-item label="维修编号" :label-width="formLabelWidth" prop="Rnumber">
            <el-input v-model="Form_.Rnumber" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="维修时间" :label-width="formLabelWidth" prop="Rtime">
            <el-input v-model="Form_.Rtime" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="机器故障现象" :label-width="formLabelWidth" prop="Rphenomenon">
            <el-input v-model="Form_.Rphenomenon" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="故障类型" :label-width="formLabelWidth" prop="Rfault">
            <el-input v-model="Form_.Rfault" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="选择产品类型" :label-width="formLabelWidth" prop="Rproduct">
            <el-input v-model="Form_.Rproduct" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="备注" :label-width="formLabelWidth" prop="Rremarks">
            <el-input v-model="Form_.Rremarks" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="维修id" :label-width="formLabelWidth" prop="Rid">
            <el-input v-model="Form_.Rid" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="是否分配工程师" :label-width="formLabelWidth" prop="Rdistribution">
            <el-input v-model="Form_.Rdistribution" autocomplete="off" disabled></el-input>
          </el-form-item>

        </el-form>
      </el-dialog>

      <!-- 技术工程师表对话框 -->
      <el-dialog title="技术工程师表的信息" v-model="dialogFormVisibleJishu">
        <el-form :model="Form" >
          <el-form-item label="名字" :label-width="formLabelWidth" prop="Ename">
            <el-input v-model="Form.Ename" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="工龄" :label-width="formLabelWidth" prop="Eage">
            <el-input v-model="Form.Eage" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="工号" :label-width="formLabelWidth" prop="Enumber">
            <el-input v-model="Form.Enumber" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="电话" :label-width="formLabelWidth" prop="Etelephone">
            <el-input v-model="Form.Etelephone" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="人工费" :label-width="formLabelWidth" prop="Ecost">
            <el-input v-model="Form.Ecost" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="维修问题编号" :label-width="formLabelWidth" prop="Eproblem">
            <el-input v-model="Form.Eproblem" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="是否正在维修" :label-width="formLabelWidth" prop="Estatue">
            <el-input v-model="Form.Estatue" autocomplete="off" disabled></el-input>
          </el-form-item>

        </el-form>
      </el-dialog>

      <!-- 总费用表 -->
      <el-dialog title="总费用表的信息" v-model="dialogFormVisibleCost">
        <el-form :model="form">
          <el-form-item label="总费用" :label-width="formLabelWidth" prop="Tcost" >
            <el-input v-model="form.Tcost" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="维修信息单号" :label-width="formLabelWidth" prop="Tnumber">
            <el-input v-model="form.Tnumber" disabled></el-input>
          </el-form-item>
          <el-form-item label="备件费" :label-width="formLabelWidth" prop="Tsparepart_cost">
            <el-input v-model="form.Tsparepart_cost" disabled></el-input>
          </el-form-item>
          <el-form-item label="人工费" :label-width="formLabelWidth" prop="Tengineer_cost">
            <el-input v-model="form.Tengineer_cost" disabled></el-input>
          </el-form-item>
          <el-form-item label="问题费用" :label-width="formLabelWidth" prop="Tproblem_cost">
            <el-input v-model="form.Tproblem_cost" disabled></el-input>
          </el-form-item>

        </el-form>
      </el-dialog>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form:{
        Tcost:'',
        Tnumber:'',
        Tsparepart_cost:'',
        Tengineer_cost:'',
        Tproblem_cost:''
      },
      Form:{
        Ename:'',
        Eage:'',
        Enumber:'',
        Etelephone:'',
        Ecost:'',
        Eproblem:'',
        Estatue:''
      },
      Form_:{
        Rnumber:'',
        Rtime:'',
        Rphenomenon:'',
        Rfault:'',
        Rproduct:'',
        Rremarks:'',
        Rid:'',
        Rdistribution:''
      },
      rightlist: [
      ],
      
      dialogFormVisibleWeixiu: false,
      dialogFormVisibleJishu:false,
      dialogFormVisibleCost:false
    };
  },
  // created() {
  //   this.getRightlist();
  //   this.getRepairlist();
  // },
  mounted() {
    this.$http
      .get("repair/staff/finance/Finance_Get_Distribution")
      .then(response => (
        this.rightlist = response.data));
  },
  methods: {
    rowClass({row,column,rowIndex,columnIndex}){
      if(rowIndex===1){
        return{
          display:'none'
        }
      }
      
    },
    Cost(Dtable_number){
      this.$http
      .get("/repair/staff/finance/Finance_Settlement/"+Dtable_number)
      .then(response => (
        this.form = response.data));
    },

    // async getRightlist() {
    //   const res = await this.$http.get("repair/staff/finance/Finance_Get_Distribution");
    //   console.log(res);
    //   this.rightlist = res.data.data;
    // },
    // async getRepairlist() {
    //   const res = await this.$http.get("users?query=${this.query}");
    //   console.log(res);
    //   this.rightlist = res.data.data;
    // },
    Repair_table(Dtable_number) {
      this.$http
      .get("/repair/staff/finance/Finance_Get_table/"+Dtable_number)
      .then(response => (
        this.Form_ = response.data));
    },
    Engineer_table(Dengineer_number){
      this.$http
      .get("/repair/staff/finance/Finance_Get_Engineer/"+Dengineer_number)
      .then(response => (
        this.Form = response.data));
    },
  },
};
</script>