
<template>
<el-button @click="split" type="primary">本月分房<i class="el-icon-upload el-icon--right"></i></el-button>
  <el-table
    :data="tableData"
    stripe
    style="width: 100%">
    <el-table-column
      prop="No"
      label="用户电话"
      width="180">
    </el-table-column>
    <el-table-column
      prop="Name"
      label="用户姓名"
      width="180">
    </el-table-column>
    <el-table-column
      prop="Department"
      label="用户部门">
    </el-table-column>
    <el-table-column
      prop="Title"
      label="用户职称">
    </el-table-column>
    <el-table-column
      prop="Num_family"
      label="用户家庭人口">
    </el-table-column>
    <el-table-column
      prop="Hno"
      label="房号">
    </el-table-column>
    <el-table-column
      prop="Rent"
      label="房租/月">
    </el-table-column>
  </el-table>
</template>

<script>
  export default {
    data() {
      return {
       tableData:[],
            addForm: {
					No: '',
					Name: '',
					Department: '',
					Title: '',
                    Num_family: '',
                    Hno: '',
                    Rent: '',
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
       split() {   
				this.$http
      .get("api/admin/House_split")
      .then(response => (
        this.tableData = response.data));
			},
   }
  }
</script>