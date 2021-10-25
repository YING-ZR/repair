<template>
<!-- 客户确认单格式 -->
<table border="1">
  <tr>
    <td colspan="4">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;********辽宁市大连民族大学有限公司取机凭证********&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td>
  </tr>
  <tr>
    <td>接修日期</td>
    <td></td>
    <td>维修编号</td>
    <td></td>
  </tr>
  <tr>
    <td>产品类型</td>
    <td></td>
    <td>机器品牌</td>
    <td></td>
  </tr>
  <tr>
    <td>机器型号</td>
    <td></td>
    <td>系列号</td>
    <td></td>
  </tr>
  <tr>
    <td class="s">单位名称</td>
    <td>   </td>
    <td>联系人</td>
    <td></td>
  </tr>
  <tr>
    <td colspan="4">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;机器故障现象</td>
  </tr>
  <tr>
    <td colspan="4"></td>
  </tr>
  <tr>
    <td colspan="2">缺少零件</td>
    <td colspan="2">随机附件</td>
  </tr>
  <tr>
    <td colspan="2"></td>
    <td colspan="2"></td>
  </tr>
  <!-- <el-button type="success" class="btn">打印客户确认单</el-button> -->
</table>
<el-button type="success" class="btn">打印客户确认单</el-button>


  
</template>

<script>
  export default {
    data() {
      return {
        form: {
          time:'',
          bianhao:'',
          leixing:'',
          pinpai:'',
          xinghao:'',
          liehao:'',
          mingcheng:'',
          xiren:'',
          xianxiang:'',
          lingjian:'',
          fujian:'',
        }
      }
    }
  }
</script>

<style scoped>
.btn{
  margin-left: 500px ;
}
.s{
  margin: 80px;
}
body{
  width: 340px;
  height: 800px;
}
table{
  border-collapse: collapse;
}
td{
  width: 300px;
  height: 50px;
}
</style>
