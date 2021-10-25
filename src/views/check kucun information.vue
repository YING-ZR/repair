<template>
 
  <el-table
    ref="filterTable"
    :data="tableData"
    style="width: 100%">
    <el-table-column
      prop="date"
      label="日期"
      sortable
      width="180"
      column-key="date"
      :filters="[
       {text: '2021-09-12', value: '2021-09-12'},
       {text: '2021-09-12', value: '2021-09-12'}, 
       {text: '2021-09-13', value: '2021-09-13'}, 
       {text: '2021-09-13', value: '2021-09-13'}
       ]"
      :filter-method="filterHandler"
    >
    </el-table-column>
    <el-table-column
      prop="replacement_name"
      label="备件名称"
      width="180">
    </el-table-column>
    <el-table-column
      prop="model"
      label="型号"
       width="180">
    </el-table-column>
     <el-table-column
      prop="number"
      label="数量"
       width="180">
    </el-table-column>
     <el-table-column
      prop="unit_price"
      label="单价"
       width="180">
    </el-table-column>
 <el-table-column
      prop="warning_number"
      label="警戒数量"
      :formatter="formatter">
    </el-table-column>
    <el-table-column
      prop="statu"
      label="库存状态"
      width="100"
      :filters="[{ text: '正常', value: '正常' },
       { text: '临界', value: '临界' },
       { text: '警示', value: '警示' },
       { text: '缺货', value: '缺货' },
       ]"
      :filter-method="filterstatu"
      filter-placement="bottom-end">
      <template >
        <el-tag
          :type="scope.row.statu === '正常' ? 'primary' : 'success'"
          disable-transitions>{{scope.row.statu}}</el-tag>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
  export default {
    data() {
      return {
        tableData: [{
        date:'2021-09-12',
        replacement_name:'芯片',
        model:'01',
        number:'100',
        unit_price:'65',
        warning_number:'100',
        statu:'临界',
        }, {
        date:'2021-09-12',
        replacement_name:'屏幕',
        model:'07',
        number:'230',
        unit_price:'800',
        warning_number:'200',
        statu:'正常',
        }, {
       date:'2021-09-13',
        replacement_name:'制冷器',
        model:'11',
        number:'23',
        unit_price:'200',
        warning_number:'20',
        statu:'警示',
        }, {
         date:'2021-09-13',
        replacement_name:'芯片',
        model:'01',
        number:'109',
        unit_price:'65',
        warning_number:'100',
        statu:'正常',
        }]
      }
    },
    methods: {
      
      resetDateFilter() {
        this.$refs.filterTable.clearFilter('date');
      },
      clearFilter() {
        this.$refs.filterTable.clearFilter();
      },
      formatter(row, column) {
        return row.warning_number;
        
      },
      filterstatu(value, row) {
        return row.statu === value;
      },
      filterHandler(value, row, column) {
        const property = column['property'];
        return row[property] === value;
      }
    },
      mounted() {
    this.$http
      .get("api/admin/house_split")
      .then(response => (
        this.tableData = response.data));
  },
  }
</script>