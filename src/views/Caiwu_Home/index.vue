<template>
  <el-container class="home-container">
      <!-- 头部 -->
    <el-header class="toubu">
    <div>
        <!-- <img src="../../assets/toubu.png" alt=""> -->
        <span>设备管理系统</span>
    </div>
    <el-button type="info" @click="logout">退出</el-button>
    </el-header>
    <el-container>
        <!-- 左侧导航 -->
        <NavMenu />
        <!-- 主体 -->
        <el-container>
            <el-header class="mianbao">
                <Breadcrumb />
            </el-header>
            <el-main><router-view /></el-main>
        </el-container>
    </el-container>
  </el-container>
</template>

<script>
import NavMenu from './navMenu.vue'
import Breadcrumb from './breadcrumb.vue'
export default {
  data () {
    return {
    }
  },
  methods: {
    logout () {
      window.sessionStorage.clear()
      this.$router.push('/GuanliyuanLogin')
    }
  },
  components: {
    NavMenu,
    Breadcrumb
  }
}
</script>

<style scoped>
.home-container {
    height: 100%;
}

.toubu {
    background: rgba(75, 126, 235, 0.849);
    color: black;
    display: flex;
    justify-content: space-between;
    padding-left: 0;
    align-items: center;
    font-size: 30px;
}

.el-main {
    width: 100%;
    height: 100%;
}
.mianbao {
    display: flex;
    justify-content: space-between;
    padding-left: 0;
    align-items: center;
    font-size: 30px;
}
</style>