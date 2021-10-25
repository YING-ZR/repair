<template>
  <el-table
    :data="tableData"
    stripe
    style="width: 100%">
    <el-table-column
      prop="No"
      label="单号"
      width="150">
    </el-table-column>
    <el-table-column
      prop="Name"
      label="工程师"
      width="150">
    </el-table-column>
    <el-table-column
      prop="submittime"
      label="提交时间"
       width="180">
    </el-table-column>
    <el-table-column
      prop="equipment"
      label="设备"
       width="150">
    </el-table-column>
    <el-table-column
      prop="rphenomenon"
      label="设备故障现象"
       width="280">
    </el-table-column>
    <el-table-column
      prop="rfault"
      label="故障类型"
       width="180">
    </el-table-column>
     <el-table-column
      prop="rremarks"
      label="客人备注"
      >
    </el-table-column>
   
  </el-table>
</template>

<script>
  export default {
    data() {
      return {
       tableData:[],
        addForm: {
				No:'',
        Name:'',
        submittime:'',
        equipment:'',
      
     
        estimate_time:'',
				},
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