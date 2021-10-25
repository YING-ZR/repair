<template>
  <div class="login-wrap">
    <div class="login_box">
      <el-form
        ref="loginFormRef"
        :model="loginForm"
        :rules="loginRules"
        class="login_form"
        label-width="100px"
      >
      <h2 class="title">设备管理系统</h2>
      <el-form-item prop="Identity">  
        <el-select v-model="loginForm.Identity" placeholder="请选择身份" >
            <el-option
            v-for="item in options"
            :key="item.label"
            :label="item.label"
            :value="item.value"
            >

            </el-option>
        </el-select>
        </el-form-item>
        <el-form-item prop="Number">
          <el-input
            placeholder="请输入账户名"
            v-model="loginForm.Number"
            prefix-icon="el-icon-user"
          ></el-input>
        </el-form-item>
        <el-form-item prop="Password">
          <el-input
            placeholder="请输入密码"
            v-model="loginForm.Password"
            prefix-icon="el-icon-lock"
            type="password"
          ></el-input>
        </el-form-item>
        <el-form-item class="button">
          <el-button @click="login" type="primary">登录</el-button>
          <el-button @click="logout" type="danger">退出</el-button>
          <el-button @click="resetLoginForm" type="info">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import Cookies from 'js-cookie'
export default {
  data() {
    return {
    //   labelee:'',
      options:[
          {
            value:'客服',
            label:'客服'
          },
          {
            value:'财务',
            label:'财务'
          },
          {
            value:'库管',
            label:'库管'
          },
          {
            value:'任务调度师',
            label:'任务调度师'
          },
          {
            value:'技术工程师',
            label:'技术工程师'
          },
      ],
      loginForm: {
        Identity: "",
        Number: "",
        Password: "",
      },
      loginRules: {
        Number: [
          {
            required: true,
            message: "请输入账户名",
            trigger: "blur",
          },
        //   {
        //     min: 5,
        //     max: 12,
        //     message: "长度在 5 到 12 个字符",
        //     trigger: "blur",
        //   },
        ],
        Password: [
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
    };
  },
  methods: {
    
    // selet(e){
    //     console.log(this.labelee)
    // },
    logout () {
      window.sessionStorage.clear()
      this.$router.push('/')
    },
    resetLoginForm() {
      this.$refs.loginFormRef.resetFields();
    },
    login() {
         this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return;
        const { data: res } = await this.$http.post(
          "repair/staff/login",
          this.loginForm
        );
        Cookies.set('number',this.loginForm.Number)
        if (res.code == 1001) {
          this.$message.error("没有此用户");
          return;
        }
        if (res.code == 1002) {
          this.$message.error("密码错误");
          return;
        }
        if (res.code == 200) {
          this.$message.success("登录成功");
        //   this.$router.push({ path: "/Home" });
          return;
        }
        this.$message.error("后台错误");
      });
      this.$refs.loginFormRef.validate(async (valid) => {
      if (this.loginForm.Identity == "客服"){
          this.$router.push({ path: "/Kefu_Home" });
    }
    else if(this.loginForm.Identity == "财务"){
          this.$router.push({ path: "/Caiwu_Home" });
    }else if(this.loginForm.Identity == "技术工程师"){

          this.$router.push({ path: "/engineerhome" });
      }else if(this.loginForm.Identity == "任务调度师"){
          this.$router.push({ path: "/renwudiaoduhome" });
       }else if(this.loginForm.Identity == "库管"){
           this.$router.push({ path: "/kuguanhome" });
      }
      });
    //   this.$refs.loginFormRef.validate(async (valid) => {
    //     if (!valid) return;
    //     const { data: res } = await this.$http.post(
    //       "repair/customer/login",
    //       this.loginForm
    //     );
    //     if (res.code == 1001) {
    //       this.$message.error("没有此用户");
    //       return;
    //     }
    //     if (res.code == 1002) {
    //       this.$message.error("密码错误");
    //       return;
    //     }
    //     if (res.code == 200) {
    //       this.$message.success("登录成功");
    //       this.$router.push({ path: "/Kefu_Home" });
    //       return;
    //     }
    //     this.$message.error("后台错误");
    //   });
    // if(this.index==1){
    //         this.$refs.loginFormRef.validate(async (valid) => {
    //         if (!valid) return;
    //         const { data: res } = await this.$http.post(
    //         "repair/staff/login",
    //         this.loginForm
    //         );
    //         if (res.code == 1001) {
    //         this.$message.error("没有此用户");
    //         return;
    //         }
    //         if (res.code == 1002) {
    //         this.$message.error("密码错误");
    //         return;
    //         }
    //         if (res.code == 200) {
    //         this.$message.success("登录成功");
    //         this.$router.push({ path: "/Kefu_Home" });
    //         return;
    //         }
    //         this.$message.error("后台错误");
    //         });
    //     }else if(this.labelee=="财务"){
    //         this.$refs.loginFormRef.validate(async (valid) => {
    //         if (!valid) return;
    //         const { data: res } = await this.$http.post(
    //         "repair/staff/login",
    //         this.loginForm
    //         );
    //         if (res.code == 1001) {
    //         this.$message.error("没有此用户");
    //         return;
    //         }
    //         if (res.code == 1002) {
    //         this.$message.error("密码错误");
    //         return;
    //         }
    //         if (res.code == 200) {
    //         this.$message.success("登录成功");
    //         this.$router.push({ path: "/Caiwu_Home" });
    //         return;
    //         }
    //         this.$message.error("后台错误");
    //         });
    //     }else if(this.labelee=="任务调度师"){
    //         this.$refs.loginFormRef.validate(async (valid) => {
    //         if (!valid) return;
    //         const { data: res } = await this.$http.post(
    //         "repair/staff/login",
    //         this.loginForm
    //         );
    //         if (res.code == 1001) {
    //         this.$message.error("没有此用户");
    //         return;
    //         }
    //         if (res.code == 1002) {
    //         this.$message.error("密码错误");
    //         return;
    //         }
    //         if (res.code == 200) {
    //         this.$message.success("登录成功");
    //         this.$router.push({ path: "/Renwudiaodu_Home" });
    //         return;
    //         }
    //         this.$message.error("后台错误");
    //         });
    //     }else if(this.labelee=="库管"){
    //         this.$refs.loginFormRef.validate(async (valid) => {
    //         if (!valid) return;
    //         const { data: res } = await this.$http.post(
    //         "repair/staff/login",
    //         this.loginForm
    //         );
    //         if (res.code == 1001) {
    //         this.$message.error("没有此用户");
    //         return;
    //         }
    //         if (res.code == 1002) {
    //         this.$message.error("密码错误");
    //         return;
    //         }
    //         if (res.code == 200) {
    //         this.$message.success("登录成功");
    //         this.$router.push({ path: "/Kuguan_Home" });
    //         return;
    //         }
    //         this.$message.error("后台错误");
    //         });
    //     }else if(this.labelee=="技术工程师"){
    //         this.$refs.loginFormRef.validate(async (valid) => {
    //         if (!valid) return;
    //         const { data: res } = await this.$http.post(
    //         "repair/staff/login",
    //         this.loginForm
    //         );
    //         if (res.code == 1001) {
    //         this.$message.error("没有此用户");
    //         return;
    //         }
    //         if (res.code == 1002) {
    //         this.$message.error("密码错误");
    //         return;
    //         }
    //         if (res.code == 200) {
    //         this.$message.success("登录成功");
    //         this.$router.push({ path: "/Jishugongchengshi_Home" });
    //         return;
    //         }
    //         this.$message.error("后台错误");
    //         });
    //     }
    },
  },
};
</script>

<style scoped>
.login-wrap {
    position: absolute;
    width: 100%;
    height: 100%;
    background-color: #303133;
}

.login-title {
    color: #303133;
    text-align: center;
}

.login_form {
    width: 650px;
    margin: 260px auto;
    background-color: rgb(255,255,255,0.8);
    padding: 30px;
    border-radius: 20px;
}

.button {
    margin-right: 50px;
    margin-bottom: 10px;
    margin-left: 80px;
}
</style>
