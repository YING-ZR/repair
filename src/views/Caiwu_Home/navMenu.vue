<template>
  <el-aside width="200px">
    <el-menu
      default-active="1"
      class="el-menu-vertical-demo" router>
      <!-- 无子级 -->
      <el-menu-item index="/Caiwu_index">
        <i class="el-icon-menu"></i>
        <span >财务首页</span>
      </el-menu-item>

      <!-- 无子级 -->
      <el-menu-item index="/Caiwu_hesuanxinxi">
        <i class="el-icon-menu"></i>
        <span >结算信息管理</span>
      </el-menu-item>

    </el-menu>
  </el-aside>
</template>

<script>
export default {
  data () {
    return {
    }
  },
  methods: {
  }
}
</script>

<style scoped>
/* .el-aside {
  background: rgba(75, 126, 235, 0.849);
}
.el-menu-vertical-demo {
  background: rgba(75, 126, 235, 0.849);
} */
</style>