<template>
  <div>
    <div class="crumbs">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>
          <i class="el-icon-lx-cascadesss"></i> 查看修改个人信息
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="container">
      <div class="handle-box">
     
     
      </div>
      <el-table
          :data="tableData"
          style="width: 100%"
          class="table"
          ref="multipleTable"
          header-cell-class-name="table-header"
          @selection-change="handleSelectionChange">
        
         
        <el-table-column prop="CID"
                         label="身份证号"

                         align="center">

        </el-table-column>
		<el-table-column prop="cname"
						label="姓名"
		
						align="center">
		
		</el-table-column>
		<el-table-column prop="cnature"
						label="客户性质"
		
						align="center">
		
		</el-table-column>
		<el-table-column prop="ccompany"
						label="单位名称"
		
						align="center">
		
		</el-table-column>
		<el-table-column prop="ctel"
						label="座机"
		
						align="center">
		
		</el-table-column>
        <el-table-column prop="ctelephone"
                         label="移动电话"

                         align="center">
        </el-table-column>
		<el-table-column prop="caddress"
						label="地址"
		
						align="center">
		
		</el-table-column>
		<el-table-column prop="cpostcode"
						label="邮编"
		
						align="center">
		
		</el-table-column>
			<el-table-column prop="cemail"
						label="邮箱"
		
						align="center">
		
		</el-table-column>


        <el-table-column label="操作" width="180" align="center">
          <template #default="scope">
            <el-button
                type="text"
                icon="el-icon-edit"
                @click="handleEditClick(scope.$index,scope.row)"
            >编辑</el-button>
			<el-button
				type="text"
				icon="el-icon-delete"
				class="red"
				@click="handleDelete(scope.$index, scope.row)"
			>删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- <div class="pagination">
        <el-pagination
            background
            layout="total, prev, pager, next"
            :current-page="sex"
            :page-size="55"
            :total="pageTotal"
            @current-change="handlePageChange"
        ></el-pagination>
      </div> -->
    </div>

    <!-- 编辑弹出框 -->
    <el-dialog title="修改" v-model="editVisible" width="30%">
      <el-form ref="form" :model="form" label-width="70px">
        <el-form-item label="身份证号">
			<el-input v-model="form.CID"></el-input>
        </el-form-item>
        <el-form-item label="姓名">
			<el-input v-model="form.cname"></el-input>
        </el-form-item>
        <el-form-item label="客户性质">
			<el-input v-model="form.cnature"></el-input>
        </el-form-item>
        <el-form-item label="单位名称">
			<el-input v-model="form.ccompany"></el-input>
        </el-form-item>
        <el-form-item label="座机">
			<el-input v-model="form.ctel"></el-input>
        </el-form-item>
        <el-form-item label="移动电话">
			<el-input v-model="form.ctelephone"></el-input>
        </el-form-item>
        <el-form-item label="地址">
			<el-input v-model="form.caddress"></el-input>
        </el-form-item>
        <el-form-item label="邮编">
			<el-input v-model="form.cpostcode"></el-input>
        </el-form-item>
         <el-form-item label="邮箱">
			<el-input v-model="form.cemail"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
            <el-button @click="editVisible = false">取 消</el-button>
            <el-button type="primary" @click="saveEdit">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
//import { fetchData } from "../api/index";
export default {
  name: "tableData",
  data() {
    return {
		eno:'',
		
		editVisible: false,
		pageTotal: 0,
		
      tableData:[
		{
		CID:"1234y",
cname:"hh",
cnature:"单位用户",
ccompany:"xx",
ctel:"123111",
ctelephone:"1234542",
caddress:"fgfd",
cpostcode:"23456",
cemail:"1234567u432",
		},
		
      ],
    }
    

  },
   mounted() {
    this.$http
      .get("api/admin/Score")
      .then(response => (
        this.tableData_ = response.data));
  },
    
  created() {
    this.getData();
  },
  methods: {
		// 获取 easy-mock 的模拟数据
		getData() {
			// fetchData(this.query).then(res => {
			// console.log(res);
			// this.tableData = res.list;
			// this.pageTotal = res.pageTotal || 50;
			// });
		},
		// 删除操作
		handleDelete(index) {
			// 二次确认删除
			this.$confirm("确定要删除吗？", "提示", {
				type: "warning"
			})
				.then(() => {
					this.$message.success("删除成功");
					this.tableData.splice(index, 1);
				})
				.catch(() => {});
		},
	
    //   // 编辑操作
    handleEditClick(index, row) {
      this.idx = index;
      this.form = row;
      this.editVisible = true;
      //this.$message.success('提交成功！');
    },

    // 保存编辑
    saveEdit() {
      this.editVisible = false;
      this.$message.success(`修改成功`);
      this.$set(this.tableData, this.idx, this.form);
    },
  
  }
};
</script>

<style scoped>
.handle-box {
  margin-bottom: 20px;
}

.handle-select {
  width: 120px;
}

.handle-input {
  width: 300px;
  display: inline-block;
}
.table {
  width: 100%;
  font-size: 14px;
}
.red {
  color: #ff0000;
}
.mr10 {
  margin-right: 10px;
}
.table-td-thumb {
  display: block;
  margin: auto;
  width: 40px;
  height: 40px;
}
</style>
