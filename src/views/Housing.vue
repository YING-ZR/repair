<template>
  <el-table
    :data="tableData"
    stripe
    style="width: 100%">
    <el-table-column
      prop="Rnumber"
      label="单号"
      width="180">
    </el-table-column>
    <el-table-column
      prop="Name"
      label="工程师"
      width="180">
    </el-table-column>
    <el-table-column
      prop="Rtime"
      label="提交时间">
    </el-table-column>
    <el-table-column
      prop="Rphenomenon"
      label="机器故障现象">
    </el-table-column>
    <el-table-column
      prop="Rfault"
      label="故障类型">
    </el-table-column>
    <el-table-column
      prop="Rproduct"
      label="产品类型">
    </el-table-column>
    <el-table-column
      prop="Rremarks"
      label="备注">
    </el-table-column>
    <el-table-column
      prop=" Rid"
      label="身份证号">
    </el-table-column>
   
  </el-table>
</template>

<script>
  export default {
    data() {
      return {
         tableData:[
		{
         Rnumber:'',
         Name:'',
         Rtime:'',
         Rphenomenon:'',
         Rfault:'',
         Rproduct:'',
         Rremarks:'',
         Rid:'',
		},
		
      ],
      }
    },
    mounted() {
    this.$http
      .get("api/admin/house_split")
      .then(response => (
        this.tableData = response.data));
  },
   methods: {
   }
  }
</script>