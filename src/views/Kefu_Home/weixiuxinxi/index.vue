<template>
  <div class="about">
    <!-- 卡片视图区域 -->
    <el-card>
      <!-- 搜索与添加区域 -->
      
        <el-form :span="7">
          <el-form-item  label="搜索" label-width="80px"  >
            <el-input  placeholder="请输入用户身份证号码" prop="Cid" v-model="form.Cid" style="width:200px"></el-input>
            <el-button class = "kehu_require" @click="Cost(form.Cid);dialogFormVisibleSousuo = true" type="primary">查询</el-button>
            <el-button type="success" @click="showAddUserDia()">添加维修信息</el-button>
          </el-form-item>
        </el-form>
      
      <!-- 用户列表区域 -->
      <el-table height="400px" :data="rightlist" border stripe>
        <el-table-column type="index" label="#" width="100"></el-table-column>
        <el-table-column label="维修编号" prop="Rnumber"></el-table-column>
        <el-table-column label="维修时间" prop="Rtime"></el-table-column>
        <el-table-column label="机器故障现象" prop="Rphenomenon"></el-table-column>
        <el-table-column label="故障类型" prop="Rfault"></el-table-column>
        <el-table-column label="选择产品类型" prop="Rproduct"></el-table-column>
        <el-table-column label="备注" prop="Rremarks"></el-table-column>
        <el-table-column label="用户身份证号" prop="Rid"></el-table-column>
        <el-table-column label="是否分配工程师" prop="Rdistribution"></el-table-column>
        <el-table-column label="预估完成时间" prop="Predicttime"></el-table-column>
        <el-table-column label="预估最后价格" prop="Predictprice"></el-table-column>
        <el-table-column label="操作" width="200px">
          <template #default="scope">
            <!-- 修改按钮 -->
            <el-button
              size="mini"
              plain
              type="primary"
              icon="el-icon-edit"
              circle
              @click="showEditUserDia(scope.row.Rnumber,scope.row.Rtime,scope.row.Rphenomenon,scope.row.Rfault,scope.row.Rproduct,scope.row.Rremarks,scope.row.Rid,scope.row.Predictprice,scope.row.Predicttime)"
            ></el-button>
            <!-- 删除按钮 -->
            <!-- <el-button
              size="mini"
              plain
              type="danger"
              icon="el-icon-delete"
              circle
              @click="showDeleUserMsgBox(scope.$index,scope.row.Cid)"
            ></el-button> -->
            <!-- 客户结算单 -->
            <el-button type="primary" size="mini" plain @click="showJiesuan(scope.row.Rid);dialogFormVisibleJiesuan = true">结算清单</el-button>
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

      <!-- 对话框 -->
      <!-- 搜索对话框 -->
      <el-dialog title="查询维修信息" v-model="dialogFormVisibleSousuo">
        <el-form :model="form" ref="House_set" :rules="loginRules">
          <el-form-item label="客户姓名" label-width="100px" prop="Cname">
            <el-input v-model="form.Cname" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="客户电话" label-width="100px" prop="Ctelephone">
            <el-input v-model="form.Ctelephone" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="客户密码" label-width="100px" prop="Cpassword">
            <el-input v-model="form.Cpassword" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="身份证号" label-width="100px" prop="Cid">
            <el-input v-model="form.Cid" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="客户电子邮件" label-width="100px" prop="Cemail">
            <el-input v-model="form.Cemail" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="客户地址" label-width="100px" prop="Caddress">
            <el-input v-model="form.Caddress" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="客户邮编" label-width="100px" prop="Cpostcode">
            <el-input v-model="form.Cpostcode" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="客户公司" label-width="100px" prop="Ccompany">
            <el-input v-model="form.Ccompany" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="客户座机" label-width="100px" prop="Ctel">
            <el-input v-model="form.Ctel" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="客户性质" label-width="100px" prop="Cnature">
            <el-input v-model="form.Cnature" autocomplete="off" disabled></el-input>
          </el-form-item>
        </el-form>
      </el-dialog>
      <!-- 添加用户的对话框 -->
      <el-dialog title="添加维修信息" v-model="dialogFormVisibleAdd">
        <el-form :model="form" ref="House_set" :rules="loginRules">
          <el-form-item label="用户身份证号" label-width="100px" prop="Cid">
            <el-input v-model="form.Cid" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="机器故障现象" label-width="100px" prop="Rphenomenon">
            <el-input v-model="form.Rphenomenon" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item prop="Rfault" label="故障类型" label-width="120px">  
            <el-select v-model="form.Rfault" placeholder="请选择故障类型">
              <el-option v-for="item in options" :key="item.label" :label="item.label" :value="item.value"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item prop="Rproduct" label="选择产品类型" label-width="120px">  
            <el-select v-model="form.Rproduct" placeholder="请选择产品类型">
              <el-option v-for="item in option" :key="item.label" :label="item.label" :value="item.value"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="备注" label-width="100px" prop="Rremarks">
            <el-input v-model="form.Rremarks" autocomplete="off"></el-input>
          </el-form-item>
        </el-form>
        <template v-slot:footer>
          <div class="dialog-footer">
            <el-button @click="dialogFormVisibleAdd = false">取 消</el-button>
            <el-button type="primary" @click="asubmit">确 定</el-button>
          </div>
        </template>
      </el-dialog>

      <!-- 编辑用户对话框 -->
      <el-dialog title="修改维修信息" v-model="dialogFormVisibleBianji">
        <el-form :model="form" ref="House_set">
          <el-form-item label="维修编号" label-width="100px" prop="Rnumber">
            <el-input v-model="form.Rnumber" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="维修时间" label-width="100px" prop="Rtime">
            <el-input v-model="form.Rtime" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="机器故障现象" label-width="100px" prop="Rphenomenon">
            <el-input v-model="form.Rphenomenon" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item prop="Rfault" label="故障类型" label-width="120px">  
            <el-select v-model="form.Rfault" placeholder="请选择故障类型">
              <el-option v-for="item in options" :key="item.label" :label="item.label" :value="item.value"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item prop="Rproduct" label="选择产品类型" label-width="120px">  
            <el-select v-model="form.Rproduct" placeholder="请选择产品类型">
              <el-option v-for="item in option" :key="item.label" :label="item.label" :value="item.value"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="备注" label-width="100px" prop="Rremarks">
            <el-input v-model="form.Rremarks" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户身份证号" label-width="100px" prop="Rid">
            <el-input v-model="form.Rid" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="预估最后价格" label-width="100px" prop="Predictprice">
            <el-input v-model="form.Predictprice" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="预估完成时间" label-width="100px" prop="Predicttime">
            <el-input v-model="form.Predicttime" autocomplete="off"></el-input>
          </el-form-item>
        </el-form>
        <template v-slot:footer>
          <div class="dialog-footer">
            <el-button @click="dialogFormVisibleBianji = false">取 消</el-button>
            <el-button type="primary" @click="bianji">确 定</el-button>
          </div>
        </template>
      </el-dialog>

      <!-- 结算对话框 -->
      <el-dialog v-model="dialogFormVisibleJiesuan" height="700px">
        <!-- <el-table :data="tableData" border >
          <el-table-column prop="Rnumber" label="维修单号"></el-table-column>
          <el-table-column prop="Rtime" label="接修日期"></el-table-column>
          <el-table-column prop="Rphenomenon" label="机器故障现象"></el-table-column>
          <el-table-column prop="Rfault" label="故障类型"></el-table-column>
          <el-table-column prop="Rproduct" label="选择产品"></el-table-column>
          <el-table-column prop="Rremarks" label="备注"></el-table-column>
          <el-table-column prop="Tproblem_cost" label="问题费用"></el-table-column>
          <el-table-column prop="Tengineer_cost" label="人工费"></el-table-column>
          <el-table-column prop="Tsparepart_cost" label="备件费用"></el-table-column>
          <el-table-column prop="Tcost" label="总费用"></el-table-column>
        </el-table> -->
        <!-- <el-form :model="Form_" >
          <el-form-item label="维修单号" label-width="100px" prop="Rnumber">
            <el-input v-model="Form_.Rnumber" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="接修日期" label-width="100px" prop="Rtime">
            <el-input v-model="Form_.Rtime" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="机器故障现象" label-width="100px" prop="Rphenomenon">
            <el-input v-model="Form_.Rphenomenon" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="故障类型" label-width="100px" prop="Rfault">
            <el-input v-model="Form_.Rfault" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="选择产品" label-width="100px" prop="Rproduct">
            <el-input v-model="Form_.Rproduct" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="备注" label-width="100px" prop="Rremarks">
            <el-input v-model="Form_.Rremarks" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="问题费用" label-width="100px" prop="Tproblem_cost">
            <el-input v-model="Form_.Tproblem_cost" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="人工费" label-width="100px" prop="Tengineer_cost">
            <el-input v-model="Form_.Tengineer_cost" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="备件费用" label-width="100px" prop="Tsparepart_cost">
            <el-input v-model="Form_.Tsparepart_cost" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="总费用" label-width="100px" prop="Tcost">
            <el-input v-model="Form_.Tcost" autocomplete="off"></el-input>
          </el-form-item>
        </el-form> -->
        <el-table :data="tableData" border >
          <el-table-column prop="Rnumber" label="维修单号"></el-table-column>
          <el-table-column prop="Rtime" label="接修日期"></el-table-column>
          <el-table-column prop="Rphenomenon" label="机器故障现象"></el-table-column>
          <el-table-column prop="Rfault" label="故障类型"></el-table-column>
          <el-table-column prop="Rproduct" label="选择产品"></el-table-column>
          <el-table-column prop="Rremarks" label="备注"></el-table-column>
          <el-table-column prop="Tproblem_cost" label="问题费用"></el-table-column>
          <el-table-column prop="Tengineer_cost" label="人工费"></el-table-column>
          <el-table-column prop="Tsparepart_cost" label="备件费用"></el-table-column>
          <el-table-column prop="Tcost" label="总费用"></el-table-column>
        </el-table>
        <!-- <el-descriptions border :data="tableData" title="结算信息单">
          <el-descriptions-item label="维修单号" prop="Rnumber"></el-descriptions-item>
          <el-descriptions-item label="接修日期" prop="Rtime"></el-descriptions-item>
          <el-descriptions-item label="机器故障现象" prop="Rphenomenon"></el-descriptions-item>
          <el-descriptions-item label="故障类型" prop="Rfault"></el-descriptions-item>
          <el-descriptions-item label="选择产品" prop="Rproduct"></el-descriptions-item>
          <el-descriptions-item label="备注" prop="Rremarks"></el-descriptions-item>
          <el-descriptions-item label="问题费用" prop="Tproblem_cost"></el-descriptions-item>
          <el-descriptions-item label="人工费" prop="Tengineer_cost"></el-descriptions-item>
          <el-descriptions-item label="备件费用" prop="Tsparepart_cost"></el-descriptions-item>
          <el-descriptions-item label="总费用" prop="Tcost"></el-descriptions-item>
        </el-descriptions> -->
      </el-dialog>
    </el-card>

  </div>
</template>

<script>
export default { 
  data() {
    return {
      // Form_:{
      //   Rnumber:'',
      //   Rtime:'',
      //   Rphenomenon:'',
      //   Rfault:'',
      //   Rproduct:'',
      //   Rremarks:'',
      //   Tproblem_cost:'',
      //   Tengineer_cost:'',
      //   Tsparepart_cost:'',
      //   Tcost:''
      // },
      options: [{
          value: '间歇性故障',
          label: '间歇性故障'
        }, {
          value: '固定性故障',
          label: '固定性故障'
        }],
      option:[
        {
          value: '台式机',
          label: '台式机'
        },
        {
          value: '手机',
          label: '手机'
        },
        {
          value: '笔记本电脑',
          label: '笔记本电脑'
        },
        {
          value: '投影仪',
          label: '投影仪'
        },
        {
          value: '打印机',
          label: '打印机'
        },
        {
          value: '其他',
          label: '其他'
        }
      ],
      Form:{
        Rnumber:''
      },
      tableData:[],
  
      loginRules: {
        Cid:[
          {
            required: true,
            message: "请输入用户身份证号",
            trigger: "blur",
          }
        ],
        Rphenomenon:[
          {
            required: true,
            message: "请输入机器故障现象",
            trigger: "blur",
          }
        ],
        Rfault:[
          {
            required: true,
            message: "请输入故障类型",
            trigger: "blur",
          }
        ],
        Rproduct:[
          {
            required: true,
            message: "请输入选择产品类型",
            trigger: "blur",
          }
        ],
        Rremarks: [
          {
            required: true,
            message: "请输入备注",
            trigger: "blur",
          }
        ],
      },

      query: "",
      // 分页相关数据
      total: -1,
      pagenum: 1,
      pagesize: 2,

      // 添加对话框的属性
      dialogFormVisibleJiesuan:false,
      dialogFormVisibleSousuo:false,
      dialogFormVisibleAdd: false,
      dialogFormVisibleBianji:false,
      dialogFormVisibleEdit:false,
      //
      // 添加用户的表单数据
      form: {
          Rnumber:'',
          Rtime:'',
          Rid:'',
          Predictprice:'',
          Predicttime:'',
          Cid:'',
          Rphenomenon:'',
          Rfault:'',
          Rproduct:'',
          Rremarks:'',
      },
      // 获取用户列表的参数对象
      rightlist: [],

    };
  },
  
  mounted() {
    this.$http
      .get("repair/service/GetcustomerTable")
      .then(response => (
        this.rightlist = response.data));
  },
  // created() {
  //   this.getRightlist();
  // },
  methods: {
    // 结算清单
    showJiesuan(Cid){
      this.$http
      .get("/repair/service/settablement_sheet/"+Cid)
      .then(response => (
       this.tableData = response.data));
        console.log(this.tableData)
    },

    Cost(Cid){
      this.$http
      .get("/repair/service/inquire/single/"+Cid)
      .then(response => (
        this.tableData = response.data));
    },
    
    // 添加维修信息
    asubmit() {
      this.$refs.House_set.validate(async (valid) => {
        if (!valid) return;
        const { data: res } = await this.$http.post(
          "repair/service/setrepair",
          this.form
        );
       console.log(res.code)
         if(res.code == 10000){
          this.$message.error("该用户已填写过维修单");
          return;
        }
        if(res.code == 10001){
          this.$message.error("请输入机器故障现象");
          return;
        }
        if(res.code == 10002){
          this.$message.error("请选择故障类型");
          return;
        }
        if(res.code == 10003){
          this.$message.error("请选择产品类型");
          return;
        }

       console.log(res.code)
        if (res.code == 200) {
          this.$message.success("添加成功");
          console.log("gxugewuewudefiwe")
          return;
        }
        this.$message.error("后台错误");
      });
    },
    
    // 修改用户
    bianji() {
      this.$refs.House_set.validate(async (valid) => {
        if (!valid) return;
        const { data: res } = await this.$http.put(
          "repair/service/modify/record",
          this.form
        );
       console.log(res.code)
        if(res.code == 30001){
          this.$message.error("该订单不存在");
          return;
        }
        if(res.code == 30002){
          this.$message.error("订单创建时间不能修改哦");
          return;
        }

       console.log(res.code)
        if (res.code == 200) {
          this.$message.success("修改成功");
          console.log("gxugewuewudefiwe");
          return;
        }
        this.$message.error("后台错误");
      });
    },
    // 搜索用户
    searchUser() {
      this.getRightlist();
    },
    // 清空搜索框，重新获取数据
    loadUserList() {
      this.getRightlist();
    },
    // 分页相关的方法
    handleSizeChange(val) {
      console.log("每页 ${val}条");
    },
    handleCurrentChange(val) {
      console.log("当前页 ${val}");
    },
    // 添加用户-显示对话框
    showAddUserDia() {
        this.form={}
      this.dialogFormVisibleAdd = true;
    },
    // 添加用户 -发送请求
     async addUser(){
        // 2.关闭对话框
        this.dialogFormVisibleAdd = false
        const res = await this.$http.post('',this.form)
        console.log(res)
        const {meta:{status,msg},data} = res.data
        if(status===201){
            // 1.提示成功
            this.$message.success(msg)
            
            // 3.更新视图
            this.getRightlist()
            // 4.清空文本框
            this.form ={}
        }else{
            this.$message.warning(msg)
        }
    },
    // 删除用户-打开消息盒子（config）
    showDeleUserMsgBox(idx,Cid){
        this.$confirm('删除用户？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async() => {
            // 发送删除的请求:id--->用户id
            // 1.data中找user ID
            // 2.把userID以showDeleUserMsgBox参数形式传进来
          const res = await this.$http.delete('repair/service/delete?Cid='+Cid)
          console.log(res)
          if(res.code == 200){
             
            // 提示
            this.$message({
            type: 'success',
            message: '删除成功!'
          })
          }
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });          
        });
    },
    // 编辑用户
    showEditUserDia(Rnumber,Rtime,Rphenomenon,Rfault,Rproduct,Rremarks,Rid,Rdistribution,Predicttime,Rpredictprice) {
        this.form = {}
        // 获取用户数据
        this.dialogFormVisibleBianji = true;
        this.form.Rnumber = Rnumber;
        this.form.Rtime = Rtime;
        this.form.Rphenomenon = Rphenomenon;
        this.form.Rfault = Rfault;
        this.form.Rproduct = Rproduct;
        this.form.Rremarks = Rremarks;
        this.form.Rid = Rid;
        this.form.Rdistribution = Rdistribution;
        this.form.Predicttime = Predicttime;
        this.form.Rpredictprice = Rpredictprice;


    }

  }
}
</script>
<style scoped>
.kehu_require{
  margin: 10px;
}
</style>