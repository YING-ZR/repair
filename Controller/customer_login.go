package Controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"repair/util"

	//"github.com/sirupsen/logrus@v1.6.0"
	"log"
	"net/http"
	"repair/Database"
	"repair/Table_Struct"
	"repair/response"
)

//用户登陆
func Customer_Login(c *gin.Context) {
	//db := Database.GetDB()
	////获取参数
	//telephone := c.PostForm("telephone")
	//password := c.PostForm("password")
	db := Database.GetDB()
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	ID := m["CID"]
	password := m["Cpassword"]

	log.Printf(ID)
	log.Printf(password)
	//数据验证
	if len(ID) != 18 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "身份证号码输入错误"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//判断手机号是否存在
	var user Table_Struct.Customer
	db.Where("CID = ?", ID).First(&user)
	if len(user.Cid) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断密码是否正确
	Cpassword := util.DeleteTailBlank(user.Cpassword)
	log.Println(len(Cpassword))
	if Cpassword != password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码错误"})
		return
	}
	//返回结果
	response.SUCCESS(c, gin.H{}, "登陆成功")
}