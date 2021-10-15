package Controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"repair/Database"
	"repair/Table_Struct"
	"repair/util"
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
	DB := Database.GetDB()
	var table []Table_Struct.Repair_table
	if err := DB.Where("RDISTRIBUTION = ?","否").Find(&table).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else {
		var engineer []Table_Struct.Engineer
		if err := DB.Where("Estatue = ?", "否").Find(&engineer).Error; err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
		} else {
			for j:=0; j < len(table); j++ {
				//log.Println("table")
				//log.Println(table[j].Rproduct)
				var problem Table_Struct.Problem
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
						}else {
							continue
						}
					}else {
						continue
					}
				}
			}
		}
	}
}
