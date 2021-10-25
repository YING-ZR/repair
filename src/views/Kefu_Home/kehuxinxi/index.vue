<template>
  <div class="about">
    <!-- 卡片视图区域 -->
    <el-card>
      <!-- 搜索与添加区域 -->
      
        <el-form :span="7">
          <el-form-item  label="搜索" label-width="80px"  >
            <el-input  placeholder="请输入用户身份证号码" prop="Cid" v-model="form.Cid" style="width:200px"></el-input>
            <el-button class = "kehu_require" @click="Cost(form.Cid);dialogFormVisibleSousuo = true" type="primary">查询</el-button>
            <el-button type="success" @click="showAddUserDia()">添加用户</el-button>
          </el-form-item>
        </el-form>
      
      <!-- 用户列表区域 -->
      <el-table height="400px" :data="rightlist" border stripe>
        <el-table-column type="index" label="#" width="100"></el-table-column>
        <el-table-column label="客户姓名" prop="Cname"></el-table-column>
        <el-table-column label="客户电话" prop="Ctelephone"></el-table-column>
        <el-table-column label="客户密码" prop="Cpassword"></el-table-column>
        <el-table-column label="身份证号" prop="Cid"></el-table-column>
        <el-table-column label="客户电子邮件" prop="Cemail"></el-table-column>
        <el-table-column label="客户地址" prop="Caddress"></el-table-column>
        <el-table-column label="客户邮编" prop="Cpostcode"></el-table-column>
        <el-table-column label="客户公司" prop="Ccompany"></el-table-column>
        <el-table-column label="客户座机" prop="Ctel"></el-table-column>
        <el-table-column label="客户性质" prop="Cnature"></el-table-column>
        <el-table-column label="操作" width="170px">
          <template #default="scope">
            <!-- 修改按钮 -->
            <el-button
              size="mini"
              plain
              type="primary"
              icon="el-icon-edit"
              circle
              @click="showEditUserDia(scope.row.Cname,scope.row.Ctelephone,scope.row.Cpassword,scope.row.Cid,scope.row.Cemail,scope.row.Caddress,scope.row.Cpostcode,scope.row.Ccompany,scope.row.Ctel,scope.row.Cnature)"
            ></el-button>
            <!-- 删除按钮 -->
            <el-button
              size="mini"
              plain
              type="danger"
              icon="el-icon-delete"
              circle
              @click="showDeleUserMsgBox(scope.$index,scope.row.Cid)"
            ></el-button>
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
      <el-dialog title="添加用户" v-model="dialogFormVisibleSousuo">
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
      <el-dialog title="添加用户" v-model="dialogFormVisibleAdd">
        <el-form :model="form" ref="House_set" :rules="loginRules">
          <el-form-item label="客户姓名" label-width="100px" prop="Cname">
            <el-input v-model="form.Cname" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户电话" label-width="100px" prop="Ctelephone">
            <el-input v-model="form.Ctelephone" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户密码" label-width="100px" prop="Cpassword">
            <el-input v-model="form.Cpassword" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="身份证号" label-width="100px" prop="Cid">
            <el-input v-model="form.Cid" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户电子邮件" label-width="100px" prop="Cemail">
            <el-input v-model="form.Cemail" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户地址" label-width="100px" prop="Caddress">
            <el-input v-model="form.Caddress" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户邮编" label-width="100px" prop="Cpostcode">
            <el-input v-model="form.Cpostcode" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户公司" label-width="100px" prop="Ccompany">
            <el-input v-model="form.Ccompany" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户座机" label-width="100px" prop="Ctel">
            <el-input v-model="form.Ctel" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户性质" label-width="100px" prop="Cnature">
            <el-input v-model="form.Cnature" autocomplete="off"></el-input>
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
      <el-dialog title="修改用户" v-model="dialogFormVisibleBianji">
        <el-form :model="form" ref="House_set">
          <el-form-item label="客户姓名" label-width="100px" prop="Cname">
            <el-input v-model="form.Cname" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户电话" label-width="100px" prop="Ctelephone">
            <el-input v-model="form.Ctelephone" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户密码" label-width="100px" prop="Cpassword">
            <el-input v-model="form.Cpassword" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="身份证号" label-width="100px" prop="Cid">
            <el-input v-model="form.Cid" autocomplete="off"  disabled></el-input>
          </el-form-item>
          <el-form-item label="客户电子邮件" label-width="100px" prop="Cemail">
            <el-input v-model="form.Cemail" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户地址" label-width="100px" prop="Caddress">
            <el-input v-model="form.Caddress" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户邮编" label-width="100px" prop="Cpostcode">
            <el-input v-model="form.Cpostcode" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户公司" label-width="100px" prop="Ccompany">
            <el-input v-model="form.Ccompany" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户座机" label-width="100px" prop="Ctel">
            <el-input v-model="form.Ctel" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="客户性质" label-width="100px" prop="Cnature">
            <el-input v-model="form.Cnature" autocomplete="off"></el-input>
          </el-form-item>
        </el-form>
        <template v-slot:footer>
          <div class="dialog-footer">
            <el-button @click="dialogFormVisibleBianji = false">取 消</el-button>
            <el-button type="primary" @click="bianji">确 定</el-button>
          </div>
        </template>
      </el-dialog>
    </el-card>

  </div>
</template>

<script>
export default {
  data() {
    return {
      loginRules: {
        Cname:[
          {
            required: true,
            message: "请输入用户姓名",
            trigger: "blur",
          }
        ],
        Cemail:[
          {
            required: true,
            message: "请输入用户电子邮箱",
            trigger: "blur",
          }
        ],
        Caddress:[
          {
            required: true,
            message: "请输入用户地址",
            trigger: "blur",
          }
        ],
        Cpostcode:[
          {
            required: true,
            message: "请输入用户邮编",
            trigger: "blur",
          }
        ],
        Ctelephone:[
          {
            required: true,
            message: "请输入用户电话",
            trigger: "blur",
          }
        ],
        Cid: [
          {
            required: true,
            message: "请输入用户身份证号码",
            trigger: "blur",
          },
          {
            min: 18,
            max: 18,
            message: "长度必须为18位",
            trigger: "blur",
          },
        ],
        Cpassword: [
          {
            required: true,
            message: "请输入密码",
            trigger: "blur",
          },
          {
            min: 6,
            max: 12,
            message: "密码为 6~12 位",
            trigger: "blur",
          },
        ],
      },

      query: "",
      // 分页相关数据
      total: -1,
      pagenum: 1,
      pagesize: 2,

      // 添加对话框的属性
      dialogFormVisibleSousuo:false,
      dialogFormVisibleAdd: false,
      dialogFormVisibleBianji:false,
      dialogFormVisibleEdit:false,
      //
      // 添加用户的表单数据
      form: {
          Cname:'',
          Ctelephone:'',
          Cpassword:'',
          Cid:'',
          Cemail:'',
          Caddress:'',
          Cpostcode:'',
          Ccompany:'',
          Ctel:'',
          Cnature:''
      },
      // 获取用户列表的参数对象
      rightlist: [],

    };
  },
  mounted() {
    this.$http
      .get("repair/service/inquire")
      .then(response => (
        this.rightlist = response.data));
  },
  // created() {
  //   this.getRightlist();
  // },
  methods: {
    Cost(Cid){
      this.$http
      .get("/repair/service/inquire/single/"+Cid)
      .then(response => (
        this.form = response.data));
    },
    
    // 添加用户
    asubmit() {
      this.$refs.House_set.validate(async (valid) => {
        if (!valid) return;
        const { data: res } = await this.$http.post(
          "repair/service/enter",
          this.form
        );
       console.log(res.code)
        if(res.code == 1001){
          this.$message.error("身份账号必须18位");
          return;
        }
        if(res.code == 1002){
          this.$message.error("密码不能少于6位");
          return;
        }
        if(res.code == 1003){
          this.$message.error("手机号不能少于6位");
          return;
        }
        if(res.code == 1004){
          this.$message.error("电子邮件不能少于11位");
          return;
        }
        if(res.code == 1005){
          this.$message.error("地址不能为空");
          return;
        }
        if(res.code == 1006){
          this.$message.error("邮编不能为空");
          return;
        }
        if(res.code == 1007){
          this.$message.error("用户已存在");
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
          "repair/service/modify",
          this.form
        );
       console.log(res.code)
        if(res.code == 3001){
          this.$message.error("手机号不能少于11位");
          return;
        }
        if(res.code == 3002){
          this.$message.error("密码不能少于6位");
          return;
        }
        if(res.code == 3003){
          this.$message.error("地址不能为空");
          return;
        }
        if(res.code == 3004){
          this.$message.error("邮编不能为空");
          return;
        }

       console.log(res.code)
        if (res.code == 200) {
          this.$message.success("修改成功");
          console.log("gxugewuewudefiwe")
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
    showEditUserDia(Cname,Ctelephone,Cpassword,Cid,Cemail,Caddress,Cpostcode,Ccompany,Ctel,Cnature) {
        this.form = {}
        // 获取用户数据
        this.dialogFormVisibleBianji = true;
        this.form.Cname = Cname;
        this.form.Ctelephone = Ctelephone;
        this.form.Cpassword = Cpassword;
        this.form.Cid = Cid;
        this.form.Cemail = Cemail;
        this.form.Caddress = Caddress;
        this.form.Cpostcode = Cpostcode;
        this.form.Ccompany = Ccompany;
        this.form.Ctel = Ctel;
        this.form.Cnature = Cnature;


    }

  }
}
</script>
<style scoped>
.kehu_require{
  margin: 10px;
}
</style>