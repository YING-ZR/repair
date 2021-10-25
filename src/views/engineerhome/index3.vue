<template>
    <div class="about">
  
<el-container>
<!--左侧导航-->
<daohangcaidan />
  <el-container>
      <!--右侧主体-->
      <!--头部-->
<el-container>
<el-header> 
  <!--右上角工程师信息-->
      <div class="title">您好，工程师{{showNumber}}号</div>
      <!-- <el-button class="exit" type="primary"  @click="goBack()">退出</el-button> -->
<mianbaoxie />
</el-header>

     <!--主体-->
    <el-main><router-view /></el-main>
  </el-container>
</el-container>
</el-container>
    </div>
</template>

<script>
import Cookie from 'js-cookie'
import daohangcaidan from './daohangcaidan.vue'
import mianbaoxie from './mianbaoxie'
export default{
  data() {
        return{

        }
    },
    computed:{
      showNumber() {
        return Cookie.get('number')
      }
    },
    methods:{

// goBack(){
//   this.$router.push("/GuanliyuanLogin");
// }
    },
    components:{
        daohangcaidan,
        mianbaoxie

    }
}
</script>




<style>

.el-header {
    background-color: #7dc6f0;
    color: rgb(10, 10, 9);
    line-height: 60px;

  }
.title{
 margin-left: 1150px;
}
  
</style>

