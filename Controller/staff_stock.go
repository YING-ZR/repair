package Controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"repair/Database"
	"repair/Table_Struct"
	"repair/util"
	"strconv"
)
//新增备件
func Setsparepart(c *gin.Context){
	db := Database.GetDB()
	//sname := c.PostForm("sname")
	//smodel := c.PostForm("smodel")
	//sunivalence := c.PostForm("sunivalence")
	//snumber := c.PostForm("snumber")
	//salertquantity := c.PostForm("salertquantity")
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	sname := m["sname"]
	smodel := m["smodel"]
	sunivalence := m["sunivalence"]
	snumber := m["snumber"]
	salertquantity := m["salertquantity"]
	var univalence,err = strconv.ParseFloat(sunivalence,64)
	if err!=nil{
		fmt.Println("转化失败",sunivalence)
	}
	var number, err1 = strconv.Atoi(snumber)
	if err1 != nil{
		fmt.Println("转化失败",snumber)
	}
	var alertquantity, err2 = strconv.Atoi(salertquantity)
	if err2 != nil{
		fmt.Println("转化失败",salertquantity)
	}
	if len(sname) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "备件名不能为空"})
		return
	}
	if len(sunivalence) == 0 {
		c.JSON(200,gin.H{"code":422,"msg":"备件单价不能为空"})
		return
	}
	if len(snumber) == 0{
		c.JSON(200,gin.H{"code":422,"msg":"备件数量不能为空"})
		return
	}
	if number<0 {
		c.JSON(200,gin.H{"code":422,"msg":"备件数量不能为负数"})
		return
	}
	if len(salertquantity) == 0 {
		c.JSON(200,gin.H{"code":422,"msg":"警戒数量不能为空"})
		return
	}
	//sinventorystatus :=util.Updatesinventorystatus(number,alertquantity)
	sinventorystatus := util.Updatesinventorystatus(number,alertquantity)
	newsparepart := Table_Struct.Sparepart{
		Sname: sname,
		Smodel: smodel,
		Sunivalence: univalence,
		Snumber: number,
		Salertquantity: alertquantity,
		Sinventorystatus: sinventorystatus,
	}
	db.Create(&newsparepart)
	//返回结果
	c.JSON(200, gin.H{
		"msg": "增加成功",
	})
}
//入库操作
func Stock_in(c *gin.Context) {
	db := Database.GetDB()
	sname := c.PostForm("sname")
	snumber := c.PostForm("snumber")
	//b, _ := c.GetRawData()
	//var m map[string]string
	//_ = json.Unmarshal(b, &m)
	//sname := m["sname"]
	//snumber := m["snumber"]
	var number, err1 = strconv.Atoi(snumber)
	if err1 != nil{
		fmt.Println("转化失败",snumber)
	}
	if len(sname) == 0 {
		c.JSON(200, gin.H{"code":422,"msg":"备件名称不能为空"})
		return
	}
	if len(snumber) == 0{
		c.JSON(200, gin.H{"code":422,"msg":"数量不能为空"})
		return
	}
	if number<0 {
		c.JSON(200,gin.H{"code":422,"msg":"备件数量不能为负数"})
		return
	}
	var spare Table_Struct.Sparepart
	var newnumber int
	db.Where("Sname = ?", sname).First(&spare)
	if len(spare.Sname) == 0 {
		c.JSON(200,gin.H{"code":422,"msg":"系统中无该备件，请先增加该备件"})
		return
	} else {
		newnumber = spare.Snumber + number
	}
	sinventorystatus :=util.Updatesinventorystatus(newnumber,spare.Salertquantity)
	//db.Model(&spare).Update("Snumber", newnumber)
	//db.Model(&spare).Update("Sinventorystatus", sinventorystatus)
	db.Model(&spare).Where("Sname = ?", sname).Update("Sinventorystatus", sinventorystatus)
	db.Model(&spare).Where("Sname = ?", sname).Update("Snumber", newnumber)

	c.JSON(200, gin.H{
		"msg": "入库成功",
	})
}

func StockOut(c *gin.Context){
	db := Database.GetDB()
	//获取数据
	sname := c.PostForm("sname")
	snumber := c.PostForm("snumber")
	//b, _ := c.GetRawData()
	//var m map[string]string
	//_ = json.Unmarshal(b, &m)
	//sname := m["sname"]
	//snumber := m["snumber"]
	//将数量转化为数字型
	var number, err1 = strconv.Atoi(snumber)
	if err1 != nil{
		fmt.Println("转化失败",snumber)
	}
	//给前端错误提示信息
	if len(sname) == 0 {
		c.JSON(200, gin.H{"code":422,"msg":"备件名称不能为空"})
		return
	}
	if len(snumber) == 0{
		c.JSON(200, gin.H{"code":422,"msg":"数量不能为空"})
		return
	}
	if number<0 {
		c.JSON(200,gin.H{"code":422,"msg":"备件数量不能为负数"})
		return
	}
	//出库操作
	var spare Table_Struct.Sparepart
	var newnumber int
	db.Where("Sname = ?", sname).First(&spare)
	if len(spare.Sname) == 0 {
		c.JSON(200,gin.H{"code":422,"msg":"系统中无该备件，请先增加该备件"})
		return
	} else {
		newnumber = spare.Snumber - number
	}
	//"code":422,spare.Snumber,"msg":"库存数量不足,库存数量只有"}
	if newnumber < 0 {
		c.JSON(200,gin.H{"code":422,"msg":"库存数量不足,库存数量只有"})
		c.JSON(200,spare.Snumber)
		return
	}
	//修改备件状态
	sinventorystatus :=util.Updatesinventorystatus(newnumber,spare.Salertquantity)
	db.Model(&spare).Where("Sname = ?", sname).Update("Sinventorystatus", sinventorystatus)
	db.Model(&spare).Where("Sname = ?", sname).Update("Snumber", newnumber)
	c.JSON(200, gin.H{
		"msg": "出库成功",
	})
}
//返回整个库存信息
func Stock_inquire(c *gin.Context) {
	db := Database.GetDB()
	var spare[] Table_Struct.Sparepart
	db.Find(&spare)
	//.Where("Sname = ?", sname).First(&spare)
	c.JSON(200,spare)
}
//查询返回单个库存信息
func Stock_inquire_single(c *gin.Context){
	db := Database.GetDB()
	sname := c.PostForm("sname")
	var spare Table_Struct.Sparepart
	db.Where("Sname = ?", sname).First(&spare)
	if len(spare.Sname) == 0 {
		c.JSON(200,gin.H{"code":422,"msg":"系统中无该备件"})
		return
	}
	c.JSON(200,gin.H{"code":200,"msg":"查询成功"})
	c.JSON(200,spare)
}