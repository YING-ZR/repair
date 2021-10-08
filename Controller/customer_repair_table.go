package Controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"repair/Database"
	"repair/Table_Struct"
	"repair/util"
)
//type Repair_table struct {
//	Rnumber string
//	Rtime string
//	Rphenomenon string
//	Rfault string
//	Rproduct string
//	Rremarks string
//}
func Set_Table(c *gin.Context) {
	db := Database.GetDB()
	//获取参数
	phenomenon := c.PostForm("phenomenon")
	fault := c.PostForm("fault")
	product := c.PostForm("product")
	remarks := c.PostForm("remarks")
	//数据验证
	if len(phenomenon) != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请输入机器故障现象"})
		return
	}
	if len(fault) != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请选择故障类型"})
		return
	}
	if len(product) != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请选择产品类型"})
		return
	}
	//如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	//创建用户
	newcustomer := Table_Struct.Customer{
		Cname:      name,
		Ctelephone: telephone,
		Cpassword:  password,
	}
	db.Create(&newcustomer)
	//返回结果
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})
}
