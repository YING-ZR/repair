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
        <el-form-item prop="CID">
          <el-input
            placeholder="请输入技术工程师身份证号码"
            v-model="loginForm.CID"
            prefix-icon="el-icon-user"
          ></el-input>
        </el-form-item>
        <el-form-item prop="Cpaaaword">
          <el-input
            placeholder="请输入密码"
            v-model="loginForm.Cpaaaword"
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
export default {
  data() {
    return {
      loginForm: {
        CID: "",
        Cpaaaword: "",
      },
      loginRules: {
        CID: [
          {
            required: true,
            message: "请输入技术工程师身份证号码",
            trigger: "blur",
          },
          {
            min: 5,
            max: 12,
            message: "长度在 5 到 12 个字符",
            trigger: "blur",
          },
        ],
        Cpaaaword: [
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
    logout () {
      window.sessionStorage.clear()
      this.$router.push('/Shenfen')
    },
    resetLoginForm() {
      this.$refs.loginFormRef.resetFields();
    },
    login() {
      this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return;
        const { data: res } = await this.$http.post(
          "repair/customer/login",
          this.loginForm
        );
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
          this.$router.push({ path: "/Kefu_Home" });
          return;
        }
        this.$message.error("后台错误");
      });
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
