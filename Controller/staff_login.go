package Controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"repair/Database"
	"repair/Table_Struct"
	"repair/response"
	"repair/util"
)

func Staff_Login(c *gin.Context) {
	db := Database.GetDB()
	//获取参数
	identity := c.PostForm("identity")
	password := c.PostForm("password")
	//db := Database.GetDB()
	//b, _ := c.GetRawData()
	//var m map[string]string
	//_ = json.Unmarshal(b, &m)
	//ID := m["Identity"]
	//password := m["Password"]

	log.Printf("identity"+identity)
	log.Printf("password"+password)
	//判断身份进行登录
	var staff Table_Struct.Staff
	db.Where("IDENTITY =  ?", identity).Find(&staff)
	//判断密码是否正确
	Spassword := util.DeleteTailBlank(staff.Password)
	if Spassword != password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码错误"})
		return
	}
	//返回结果
	response.SUCCESS(c, gin.H{}, "登陆成功")
}