package Controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"repair/Database"
	"repair/Table_Struct"
	"repair/response"
	"repair/util"
)


func Customer_modify(c *gin.Context){
	db := Database.GetDB()
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
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	cname := m["Cname"]
	ctelephone := m["Ctelephone"]
//	cpassword := m["Cpassword"]
	cid := m["Cid"]
	ctel := m["Ctel"]
	cemail := m["Cemail"]
	caddress := m["Caddress"]
	cpostcode := m["Cpostcode"]
	ccompany := m["Ccompany"]
	cnature := m["Cnature"]
	if len(util.DeleteTailBlank(ctelephone)) != 11{
		c.JSON(http.StatusOK, gin.H{"code":50001, "msg":"手机号不能少于11位"})
	}
	if len(cemail) == 0{
		c.JSON(http.StatusOK, gin.H{"code":50002, "msg":"邮编不能为空"})
	}
	if len(caddress) == 0{
		c.JSON(http.StatusOK, gin.H{"code":50003, "msg":"地址不能为空"})
	}
	if len(cpostcode) == 0{
		c.JSON(http.StatusOK, gin.H{"code":50004,"msg":"邮编不能为空"})
	}
	//mcustomer := Table_Struct.Customer{
	//	Cname: cname,
	//	Ctelephone: ctelephone,
	//	Cpassword: cpassword,
	//	Cid: cid,
	//	Ctel: ctel,
	//	Cemail: cemail,
	//	Ccompany: ccompany,
	//	Cpostcode: cpostcode,
	//	Cnature: cnature,
	//	Caddress: caddress,
	//}
	var user Table_Struct.Customer
	db.Where("Cid = ?", cid).First(&user)
	if(util.DeleteTailBlank(user.Ctelephone) != util.DeleteTailBlank(ctelephone)) {
		db.Model(&user).Where("Cid = ?", cid).Update("Ctelephone",ctelephone)
	}
	//if(util.DeleteTailBlank(user.Cpassword) != util.DeleteTailBlank(cpassword)) {
	//	db.Model(&user).Where("Cid = ?", cid).Update("Cpassword",cpassword)
	//}
	if(util.DeleteTailBlank(user.Ctel) != util.DeleteTailBlank(ctel)) {
		db.Model(&user).Where("Cid = ?", cid).Update("Ctel",ctel)
	}
	if(util.DeleteTailBlank(user.Caddress) != util.DeleteTailBlank(caddress)) {
		db.Model(&user).Where("Cid = ?", cid).Update("Caddress",caddress)
	}
	if(util.DeleteTailBlank(user.Ccompany) != util.DeleteTailBlank(ccompany)) {
		db.Model(&user).Where("Cid = ?", cid).Update("Ccompany",ccompany)
	}
	if(util.DeleteTailBlank(user.Cname) != util.DeleteTailBlank(cname)) {
		db.Model(&user).Where("Cid = ?", cid).Update("Cname",cname)
	}
	if(util.DeleteTailBlank(user.Cemail) != util.DeleteTailBlank(cemail)) {
		db.Model(&user).Where("Cid = ?", cid).Update("Cemail",cemail)
	}
	if(util.DeleteTailBlank(user.Cnature) != util.DeleteTailBlank(cnature)) {
		db.Model(&user).Where("Cid = ?", cid).Update("Cnature",cnature)
	}
	if(util.DeleteTailBlank(user.Cpostcode) != util.DeleteTailBlank(cpostcode)) {
		db.Model(&user).Where("Cid = ?", cid).Update("Cpostcode",cpostcode)
	}
	//db.Create(&mcustomer)
	response.SUCCESS(c, nil,"修改完成")
}
