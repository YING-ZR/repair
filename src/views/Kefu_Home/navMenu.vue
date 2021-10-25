<template>
  <el-aside width="200px">
    <el-menu
      default-active="1"
      class="el-menu-vertical-demo" router>
      <!-- 无子级 -->
      <el-menu-item index="/Kefu_index">
        <i class="el-icon-menu"></i>
        <span >客服首页</span>
      </el-menu-item>

      <!-- 无子级 -->
      <el-menu-item index="/Kefu_kehuxinxi">
        <i class="el-icon-menu"></i>
        <span >客户信息管理</span>
      </el-menu-item>

      <!-- 有子级 -->
      <!-- <el-submenu index="/baoxiu">
        <template >
          <i class="el-icon-menu"></i>信息管理</template>
        <el-menu-item index="/baoxiu/list">列表管理</el-menu-item>
      </el-submenu> -->

      <el-menu-item index="/Kefu_weixiuxinxi">
        <i class="el-icon-menu"></i>
        <span >客户维修信息管理</span>
      </el-menu-item>

      <el-menu-item index="/Kefu_hesuanxinxi">
        <i class="el-icon-menu"></i>
        <span >客户确认单</span>
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
