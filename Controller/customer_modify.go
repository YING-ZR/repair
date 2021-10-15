package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"repair/Database"
	"repair/Table_Struct"
)


func Customer_modify(c *gin.Context){
	db := Database.GetDB()
	cname := c.PostForm("cname")
	ctelephone := c.PostForm("ctelephone")
	cpassword := c.PostForm("cpassword")
	cid := c.PostForm("cid")
	ctel := c.PostForm("ctel")
	cemail := c.PostForm("cemail")
	caddress := c.PostForm("caddress")
	cpostcode := c.PostForm("cpostcode")
	ccompany := c.PostForm("ccompany")
	cnature := c.PostForm("cnature")
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
	mcustomer := Table_Struct.Customer{
		Cname: cname,
		Ctelephone: ctelephone,
		Cpassword: cpassword,
		Cid: cid,
		Ctel: ctel,
		Cemail: cemail,
		Ccompany: ccompany,
		Cpostcode: cpostcode,
		Cnature: cnature,
		Caddress: caddress,
	}
	var user Table_Struct.Customer
	db.Where("Cid = ?", cid).Delete(&user)
	db.Create(&mcustomer)
	c.JSON(200, gin.H{
		"msg": "修改成功",
	})
}
