package Controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"repair/Database"
	"repair/Table_Struct"
	"repair/util"
)

func Maintenance_status(c *gin.Context) {
	db := Database.GetDB()
	//cid := c.PostForm("cid")
	cid, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	var repair Table_Struct.Repair_table
	var dis Table_Struct.Distribution
	db.Where("Rid = ?",cid).First(&repair)
	db.Where("Dtable_number = ?",repair.Rnumber).First(&dis)
	type duixiang struct {
		Dstatus string
		Rdistribution string
	}
	shuzu := duixiang{
		Dstatus: util.DeleteTailBlank(dis.Dstatus),
		Rdistribution: util.DeleteTailBlank(repair.Rdistribution),
	}
	c.JSON(200,shuzu)
}
func Engineer_imag(c *gin.Context) {
	db := Database.GetDB()
	cid, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	log.Println(cid)
	var user Table_Struct.Repair_table
	var fenpei Table_Struct.Distribution
	var engi Table_Struct.Engineer
	db.Where("Rid = ?",cid).First(&user)
	db.Where("Dtable_number = ?",user.Rnumber).First(&fenpei)
	db.Where("Enumber = ?",fenpei.Dengineer_number).First(&engi)
	noengi := Table_Struct.Engineer{
		Ename: "未分配",
		Eage: "未分配",
		Enumber: "未分配",
		Ecost: "未分配",
		Estatue: "未分配",
		Etelephone: "未分配",
		Eproblem: "未分配",
	}
	log.Println(engi)
	if util.DeleteTailBlank(user.Rdistribution) == "是" {
		c.JSON(200, engi)
	} else {
		c.JSON(200,noengi)
	}
}