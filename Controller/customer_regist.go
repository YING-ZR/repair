package Controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"repair/Database"
	"repair/Table_Struct"
	"repair/util"
)

//用户注册
func Register(c *gin.Context) {
	db := Database.GetDB()
	//获取参数
	//cname := c.PostForm("cname")
	//ctelephone := c.PostForm("ctelephone")
	//cpassword := c.PostForm("cpassword")
	//cid := c.PostForm("cid")
	//ctel := c.PostForm("ctel")
	//cemail := c.PostForm("cemail")
	//caddress := c.PostForm("caddress")
	//cpostcode := c.PostForm("cpostcode")
	//ccompany := c.PostForm("ccompany")
	//cnature := c.PostForm("cnature")
	//数据验证
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	cname := m["Cname"]
	ctelephone := m["Ctelephone"]
	cpassword := m["Cpassword"]
	cid := m["Cid"]
	ctel := m["Ctel"]
	cemail := m["Cemail"]
	caddress := m["Caddress"]
	cpostcode := m["Cpostcode"]
	ccompany := m["Company"]
	cnature := m["Cnature"]
	if len(cid) != 18 {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "身份证号必须18位"})
		return
	}
	if len(cpassword) < 6 {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	if len(ctelephone) != 11{
		c.JSON(http.StatusOK, gin.H{"code":422, "msg":"手机号不能少于11位"})
	}
	if len(cemail) == 0{
		c.JSON(http.StatusOK, gin.H{"code":422, "msg":"邮编不能为空"})
	}
	if len(caddress) == 0{
		c.JSON(http.StatusOK, gin.H{"code":422, "msg":"地址不能为空"})
	}
	if len(cpostcode) == 0{
		c.JSON(http.StatusOK, gin.H{"code":422,"msg":"邮编不能为空"})
	}
	//如果名称没有传，给一个10位的随机字符串
	if len(cname) == 0 {
		cname = util.RandomString(10)
	}
	//判断手机号是否存在
	if iscidExist(db, cid) {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	//创建用户
	newcustomer := Table_Struct.Customer{
		Cname:      cname,
		Ctelephone: ctelephone,
		Cpassword:  cpassword,
		Cid: cid,
		Ctel: ctel,
		Cemail: cemail,
		Caddress: caddress,
		Cpostcode: cpostcode,
		Ccompany: ccompany,
		Cnature: cnature,
	}
	db.Create(&newcustomer)
	//返回结果
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})
}
//判断数据库中是否有这个客户
func iscidExist(db *gorm.DB, id string) bool {
	var user Table_Struct.Customer
	db.Where("Cid = ?", id).First(&user)
	if len(user.Cid) != 0 {
		return true
	}
	return false
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user Table_Struct.Customer
	db.Where("CTELEPHONE = ?", telephone).First(&user)
	if len(user.Ctelephone) != 0 {
		return true
	}
	return false
}

