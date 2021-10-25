<template>
  <div class="about">
    <!-- 走马灯 -->
    <el-carousel indicator-position="outside" type="card" :interval="3000">
    <el-carousel-item v-for="item in imgLists" :key="item">
      <img :src="item" />
      
    </el-carousel-item>
    </el-carousel>

    <!-- 卡片 -->
    <el-row :gutter="20">
  <el-col :span="18" v-for="o in cardLists" :key="o.img">
    <el-card :body-style="{ padding: '0px' }">
      <img :src="o.img" class="image">
      <div style="padding: 14px;">
        <span>{{o.title}}</span>
        <div class="bottom clearfix">
          <el-button type="text" class="button">操作按钮</el-button>
        </div>
      </div>
    </el-card>
  </el-col>
</el-row>
  </div>
</template>

<script>
import zouma1 from "@/assets/img/zouma1.png"
import zouma2 from "@/assets/img/zouma2.png"
import zouma3 from "@/assets/img/zouma3.png"
export default{
    data(){
        return{
            imgLists:[
                zouma1,
                zouma2,
                zouma3
            ],
            cardLists:[
                {title:'企业荣誉',img:require("@/assets/img/zouma1.png")},
                {title:'成长之路',img:require("@/assets/img/zouma2.png")},
                {title:'感恩有你',img:require("@/assets/img/zouma3.png")}
            ]
        }
    },
    methods:{

    }
}
</script>

<style scoped>
.el-card_body img{
    height: 200px;
    max-width: 100%;
}
</style>