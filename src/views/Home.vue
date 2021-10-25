<!--工程师首页-->
<template>
  <el-container>
    <el-header>
      <div>
        <img class="logo" src="../assets/logo.png" alt="">
        <span>工程师首页</span>
      </div>
    </el-header>
    <el-container>
      <el-aside width="200px">
        <el-menu
        router
      class="el-menu-vertical-demo"

      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b">
      <el-menu-item v-for="i in A_routerList" :key="i.name" :index="i.path" >
        <i class="el-icon-setting"></i>
        <template #title>{{i.meta.title}}</template>
      </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main><router-view/></el-main>
    </el-container>
  </el-container>
  
</template>

<script>
import router from "../router/index.js"
import {useRouter} from 'vue-router'
import {reactive} from 'vue'
export default {
  name: "Home",
  setup() {
    const router = useRouter()
    const A_routerList = router.options.routes[3].children
    console.log(router.options.routes[3].children)
    console.log(router)
    // tablist = tablist.filter(v=>v.meta.isShow)  
    return {
      A_routerList,
    };
  },
};
</script>




<style scoped>
.logo {
  width: 50px;
  height: 50px;
}
.el-header {
  height: 10vh;
  height: 100px;
  background: rgb(167, 166, 165);
}
.el-container{
  height: 100vh;
}
.el-menu{
  height: 100%;
}
.el-main{
  background: white;
}
</style>