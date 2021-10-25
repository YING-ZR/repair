<template>
  <el-aside width="200px">
    <el-menu
      default-active="1"
      class="el-menu-vertical-demo" router>
      <!-- 无子级 -->
      <el-menu-item index="/index5">
        <i class="el-icon-menu"></i>
        <span >库管首页</span>
      </el-menu-item>

      <!-- 无子级 -->
      <el-menu-item index="/add_kucun_information">
        <i class="el-icon-menu"></i>
        <span >入库</span>
      </el-menu-item>
       <el-menu-item index="/update">
        <i class="el-icon-menu"></i>
        <span >添加备件</span>
      </el-menu-item>
        <el-menu-item index="/search_kucun_information">
        <i class="el-icon-menu"></i>
        <span >查看库存信息</span>
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
