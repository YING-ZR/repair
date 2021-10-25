package Controller

import (
	"github.com/gin-gonic/gin"
	"repair/Database"
	"repair/Table_Struct"
	"repair/util"
	"strconv"
	_ "time"
)

//查询distribution表
func Finance_Get_Distribution(c *gin.Context){
	DB := Database.GetDB()
	var dis []Table_Struct.Distribution
	if err := DB.Find(&dis).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else {
		c.JSON(200, dis)
	}
}
//查询维修单表
func Finance_Get_table(c *gin.Context){
	DB := Database.GetDB()
	number, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	var table Table_Struct.Repair_table
	if err := DB.Where("Rnumber=?",number).Find(&table).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else {
		c.JSON(200, table)
	}
}
//查询工程师表
func Finance_Get_Engineer(c *gin.Context){
	DB := Database.GetDB()
	number, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	var engineer Table_Struct.Engineer
	if err := DB.Where("Enumber=?",number).Find(&engineer).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else {
		c.JSON(200, engineer)
	}
}
//结算单计算
func Finance_Settlement(c *gin.Context){
	DB := Database.GetDB()
	number, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	var dis Table_Struct.Distribution
	DB.Where("Dtable_number=?",number).First(&dis)
	sparepart_name := util.DeleteTailBlank(dis.Dsparepart_name)
	var sparepart Table_Struct.Sparepart
	DB.Where("Sname=?",sparepart_name).First(&sparepart)
	//sparepart_cost := util.DeleteTailBlank(sparepart.Sinventorystatus)
	sparepart_cost:=sparepart.Sunivalence
	engineer_number := util.DeleteTailBlank(dis.Dengineer_number)
	var engineer Table_Struct.Engineer
	DB.Where("Enumber=?",engineer_number).First(&engineer)
	tengineer_cost := util.DeleteTailBlank(engineer.Ecost)
	var problem Table_Struct.Problem
	eproblem := util.DeleteTailBlank(engineer.Eproblem)
	DB.Where("Pnumber=?",eproblem).First(&problem)
	pcost := util.DeleteTailBlank(problem.Pcost)
	tengineer_cost_,_ := strconv.ParseFloat(tengineer_cost,64)
	pcost_,_ := strconv.ParseFloat(pcost,64)
	cost := sparepart_cost + tengineer_cost_ + pcost_
	table_cost := Table_Struct.Table_Cost{
		Tcost : strconv.FormatFloat(cost,'f',-1,64), //总费用
		Tnumber : number, //维修信息单单号
		Tsparepart_cost : strconv.FormatFloat(sparepart_cost,'f',-1,64) ,//备件费
		Tengineer_cost : tengineer_cost, //人工费
		Tproblem_cost : eproblem, //问题费用
	}
	DB.Create(&table_cost)
	c.JSON(200, table_cost)
}
