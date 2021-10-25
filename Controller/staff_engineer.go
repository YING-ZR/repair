package Controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"repair/Database"
	"repair/Table_Struct"
	"repair/response"
	"repair/util"
)

//查询distribution表
func Engineer_Get_Distribution(c *gin.Context){
	DB := Database.GetDB()
	//number := c.PostForm("number")//接收技术工程师的number
	number, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	log.Println("!11111111111111111111111111111111111111111")
	log.Println(number)
	var dis Table_Struct.Distribution
	number = util.DeleteTailBlank(number)
	if err := DB.Where("Dengineer_number = ?", number).Find(&dis).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else{
		c.JSON(200, dis)
	}
}
//修改维修状态
func Engineer_Update_Distribution_status(c *gin.Context){
	DB := Database.GetDB()
	number, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	//status := c.PostForm("dstatus")//维修状态
	if err := DB.Model(Table_Struct.Distribution{}).Where("DTABLE_NUMBER=?",number).Update("Dstatus","维修完成").Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
		return
	}else{
		var dis Table_Struct.Distribution
		DB.Where("DTABLE_NUMBER=?",number).First(&dis)
		num := util.DeleteTailBlank(dis.Dengineer_number)
		log.Println(num)
		log.Println("11111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
		if err := DB.Model(Table_Struct.Engineer{}).Where("Enumber=?",num).Update("Estatue", "否").Error; err != nil{
			c.JSON(200,gin.H{"error": err.Error()})
		}else{
			log.Println("fou")
			response.SUCCESS(c, nil,"维修完成")
		}
	}
}
//修改维修延迟程度状态
func Engineer_Update_Distribution_ddelay_state(c *gin.Context){
	DB := Database.GetDB()
	number, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	sname := m["Ddelay_state"]
	log.Println(sname)
	//delay_state := c.PostForm("ddelay_state")//维修延迟程度
	if err :=DB.Model(Table_Struct.Distribution{}).Where("DTABLE_NUMBER=?",number).Update("Ddelay_state",sname).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else{
		response.SUCCESS(c, nil,"修改成功")
	}
}
//查询维修单表
func Engineer_Get_Repair_Table(c *gin.Context){
	DB := Database.GetDB()
	//number := c.PostForm("number")//接收技术工程师的number
	number, ok:= c.Params.Get("Number")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	var dis Table_Struct.Distribution
	var table[] Table_Struct.Repair_table
	number = util.DeleteTailBlank(number)
	DB.Where("Dtable_number = ?", number).Find(&dis)
	DB.Where("Rnumber = ?" , dis.Dtable_number).First(&table)
	log.Println(table)
	c.JSON(200, table)
}

//是否需要备件
func Engineer_Get_Sparepart(c *gin.Context){
	number, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	db := Database.GetDB()
	//获取数据
	//number := c.PostForm("number")//获取维修单编号
	//sname := c.PostForm("sname")
	//snumber := c.PostForm("snumber")
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	sname := m["Dsparepart_name"]
	snumber := m["Snumber"]
	log.Println(snumber)
	//将数量转化为数字型
	Ceshi(sname,snumber,c)
	db.Model(Table_Struct.Distribution{}).Where("DTABLE_NUMBER=?",number).Update("Dsparepart_name",sname)
	//var dis Table_Struct.Distribution
	//db.Where("Dtable_number=?",number).First(&dis)
	//if util.DeleteTailBlank(dis.Dsparepart_name)=="无" {
	//	db.Model(Table_Struct.Distribution{}).Where("DTABLE_NUMBER=?",number).Update("Dsparepart_name",sname)
	//}else {
	//	distribution := Table_Struct.Distribution{
	//		Dsparepart_name : dis.Dsparepart_name,//备件名称
	//		Dtable_number : dis.Dtable_number, //维修单号
	//		Dengineer_number : dis.Dengineer_number,//技术工程师员工号
	//		Dtime : dis.Dtime, //分配技术工程师时间
	//		Dstatus : dis.Dstatus, //维修状态
	//		Ddelay_state : sname, //维修延迟程度
	//	}
	//	db.Create(&distribution)
	//}
	//response.SUCCESS(c, gin.H{}, "备件成功")
}

//维修完成状态改变

