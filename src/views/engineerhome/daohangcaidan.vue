<template>
  <el-aside width="200px">
    <el-menu
      default-active="1"
      class="el-menu-vertical-demo" router>
      <!-- 无子级 -->
      <el-menu-item index="/index3">
        <i class="el-icon-menu"></i>
        <span >工程师首页</span>
      </el-menu-item>



      <!-- 无子级 -->
      <el-menu-item index="/all_function">
        <i class="el-icon-menu"></i>
        <span >操作维修信息</span>
      </el-menu-item>

      <!-- 无子级 -->
      <!-- <el-menu-item index="/maintance information">
        <i class="el-icon-menu"></i>
        <span >填写维修信息</span>
      </el-menu-item> -->

      <!-- 有子级 -->
      <!-- <el-submenu index="/baoxiu">
        <template >
          <i class="el-icon-menu"></i>信息管理</template>
        <el-menu-item index="/baoxiu/list">列表管理</el-menu-item>
      </el-submenu> -->

      <!-- <el-menu-item index="/view maintance task">
        <i class="el-icon-menu"></i>
        <span >查看维修任务</span>
      </el-menu-item> -->

    
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
