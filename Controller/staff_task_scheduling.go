package Controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"repair/Database"
	"repair/Table_Struct"
	"repair/response"
	"repair/util"
	"time"
)

//返回未分配维修单
func Get_Undistribution(c *gin.Context){
	DB := Database.GetDB()
	var table []Table_Struct.Repair_table
	log.Println(len(table))
	if err := DB.Where("RDISTRIBUTION = ?","否").Find(&table).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else {
		c.JSON(200, table)
	}
}
//分配工程师函数
func Distribution_Engineer(c *gin.Context)  {
	log.Println("1")
	DB := Database.GetDB()
	var table []Table_Struct.Repair_table
	if err := DB.Where("RDISTRIBUTION = ?","否").Find(&table).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
		} else {
			log.Println(table)
			for j:=0; j < len(table); j++ {
				//log.Println("table")
				//log.Println(table[j].Rproduct)
				var problem Table_Struct.Problem
				var engineer []Table_Struct.Engineer
				if err := DB.Where("Estatue = ?", "否").Find(&engineer).Error; err != nil {
					c.JSON(200, gin.H{"error": err.Error()})
				}
				for i := 0; i < len(engineer); i++ {
					//log.Println("Eproblem")
					//log.Println(engineer[i].Eproblem)
					DB.Where("PNUMBER = ?", engineer[i].Eproblem).First(&problem)
					//log.Println(problem.Pnumber)
					if util.DeleteTailBlank(problem.Pproduct) == util.DeleteTailBlank(table[j].Rproduct) {
						//log.Println(problem.Pproduct)
						//log.Println(table[j].Rproduct)
						if util.DeleteTailBlank(problem.Pname) == util.DeleteTailBlank(table[j].Rfault){
							DB.Model(Table_Struct.Repair_table{}).Where("Rnumber=?",table[j].Rnumber).Update("Rdistribution","是")
							DB.Model(Table_Struct.Engineer{}).Where("Enumber=?",engineer[i].Enumber).Update("Estatue","是")
							time_ := time.Now().Unix()
							time := time.Unix(time_,0).Format("2006-01-02 15:04:05")
							dis := Table_Struct.Distribution{
								Dsparepart_name : "无",
								Dtable_number : table[j].Rnumber,
								Dengineer_number : engineer[i].Enumber,
								Dtime : time, //分配技术工程师时间
								Dstatus : "正在维修",
							}
							DB.Create(&dis)
						}else {
							continue
						}
					}else {
						continue
					}
				}
			}
		}
	response.SUCCESS(c, gin.H{}, "分配成功")
}

//返回已分配订单的分配表
func Get_Distribution(c *gin.Context){
	DB := Database.GetDB()
	var dis []Table_Struct.Distribution
	if err := DB.Find(&dis).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else {
		c.JSON(200, dis)
	}
}

//根据订单号查找维修表单
func Get_Distribution_table(c *gin.Context){
	log.Println("11111")
	DB := Database.GetDB()
	var table Table_Struct.Repair_table
	number, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	if err := DB.Where("RNUMBER = ?", number).First(&table).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else {
		c.JSON(200, table)
		log.Println(table)
	}
}
