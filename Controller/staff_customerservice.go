package Controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"repair/Database"
	"repair/Table_Struct"
	"repair/response"
	"repair/util"
	"strconv"
	"time"
)

//客户信息录入
func Customerservice_enter(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{"code": 1001, "msg": "身份证号必须18位"})
		return
	}
	if len(cpassword) < 6 {
		c.JSON(http.StatusOK, gin.H{"code": 1002, "msg": "密码不能少于6位"})
		return
	}
	if len(ctelephone) != 11{
		c.JSON(http.StatusOK, gin.H{"code":1003, "msg":"手机号不能少于11位"})
		return
	}
	if len(cemail) == 0{
		c.JSON(http.StatusOK, gin.H{"code":1004, "msg":"邮箱不能为空"})
		return
	}
	if len(caddress) == 0{
		c.JSON(http.StatusOK, gin.H{"code":1005, "msg":"地址不能为空"})
		return
	}
	if len(cpostcode) == 0{
		c.JSON(http.StatusOK, gin.H{"code":1006,"msg":"邮编不能为空"})
		return
	}
	//如果名称没有传，给一个10位的随机字符串
	if len(cname) == 0 {
		cname = util.RandomString(10)
	}
	//判断手机号是否存在
	if iscidExist(db, cid) {
		c.JSON(http.StatusOK, gin.H{"code": 1007, "msg": "用户已经存在"})
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
	//返回结果)
	c.JSON(200,gin.H{"code":200,"msg":"添加成功"})
}
//客户信息的编辑
func Customerservice_modify(c *gin.Context) {
	db := Database.GetDB()
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
	ccompany := m["Ccompany"]
	cnature := m["Cnature"]
	log.Println("321414142441")
	ctelephone = util.DeleteTailBlank(ctelephone)
	if len(ctelephone) != 11{
		c.JSON(http.StatusOK, gin.H{"code":3001, "msg":"手机号不能少于11位"})
		return
	}
	if len(cemail) == 0{
		c.JSON(http.StatusOK, gin.H{"code":3002, "msg":"邮编不能为空"})
		return
	}
	if len(caddress) == 0{
		c.JSON(http.StatusOK, gin.H{"code":3003, "msg":"地址不能为空"})
		return
	}
	if len(cpostcode) == 0{
		c.JSON(http.StatusOK, gin.H{"code":3004,"msg":"邮编不能为空"})
		return
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
	//var user Table_Struct.Customer
	//db.Where("Cid = ?", cid).Delete(&user)
	//db.Create(&mcustomer)
	var user Table_Struct.Customer
	db.Where("Cid = ?", cid).First(&user)
	if(util.DeleteTailBlank(user.Ctelephone) != util.DeleteTailBlank(ctelephone)) {
		db.Model(&user).Where("Cid = ?", cid).Update("Ctelephone",ctelephone)
	}
	if(util.DeleteTailBlank(user.Cpassword) != util.DeleteTailBlank(cpassword)) {
		db.Model(&user).Where("Cid = ?", cid).Update("Cpassword",cpassword)
	}
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
//客户信息的删除
func Customer_delete(c *gin.Context) {
	db := Database.GetDB()
	cid := c.Query("Cid")
	var user Table_Struct.Customer
	log.Println(cid)
	db.Where("Cid = ?", cid).First(&user)
	if user.Cname == "" {
		c.JSON(200, gin.H{
			"code":422,
			"msg" : "无此人",
		})
		return
	} else {
		db.Where("Cid = ?", cid).Delete(&user)
		c.JSON(200,gin.H{"code":200,"msg":"删除成功"})
	}
}
//客户信息的查看
//返回整个库存信息
func Customer_inquire(c *gin.Context) {
	db := Database.GetDB()
	var user[] Table_Struct.Customer
	err := db.Find(&user).Error
	log.Println(user)
	if err!=nil{
		c.JSON(200,gin.H{"error":err.Error()})
	}
	//.Where("Sname = ?", sname).First(&spare)
	c.JSON(200,user)
	log.Println(user)
}
//查询返回单个库存信息
func Customer_inquire_single(c *gin.Context){
	//db := Database.GetDB()
	//sname, ok:= c.Params.Get("Hno")
	//if !ok {
	//	c.JSON(200, gin.H{"error": "无效订单号"})
	//	return
	//}
	//log.Println(sname)
	////sname := c.PostForm("sname")
	////b, _ := c.GetRawData()
	///ar m map[string]string
	////_ = json.Unmarshal(b, &m)
	////sname := m["Sname"]
	//var spare Table_Struct.Sparepart
	//db.Where("Sname = ?", sname).First(&spare)
	//no := [] string{"0000","0000","0000","0000","0000","0000"}
	//if len(spare.Sname) == 0 {
	//	c.JSON(200,no)
	//	return
	//}
	////c.JSON(200,gin.H{"code":200,"msg":"查询成功"})
	//c.JSON(200,spare)
	db := Database.GetDB()
	cid, ok:= c.Params.Get("Cid")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	var user Table_Struct.Customer
	db.Where("Cid = ?", cid).First(&user)
	if len(user.Cid) == 0 {
		c.JSON(200,gin.H{"code":422,"msg":"系统中无该备件"})
		return
	}
	//c.JSON(200,gin.H{"code":200,"msg":"查询成功"})
	c.JSON(200,user)
}



//维护报修信息


//客服填写维修单
func Customer_Settable(c *gin.Context) {
	db := Database.GetDB()
	//id, ok:= c.Params.Get("Hno")
	//if !ok {
	//	c.JSON(200, gin.H{"error": "无效订单号"})
	//	return
	//}
	//获取参数
	//cid := c.PostForm("cid")
	//phenomenon := c.PostForm("phenomenon")
	//fault := c.PostForm("fault")
	//product := c.PostForm("product")
	//remarks := c.PostForm("remarks")
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	id := m["Cid"]
	phenomenon := m["Rphenomenon"]
	fault := m["Rfault"]
	product := m["Rproduct"]
	remarks := m["Rremarks"]
	//数据验证
	var table_ Table_Struct.Repair_table
	db.Where("RID=?",id).First(&table_)

	if len(table_.Rid)!=0{
		c.JSON(200, gin.H{"code": 10000, "msg": "该用户已填写过维修单"})
		return
	}
	if len(phenomenon) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "请输入机器故障现象"})
		return
	}
	if len(fault) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "请选择故障类型"})
		return
	}
	if len(product) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "请选择产品类型"})
		return
	}
	//自动生成订单号
	var number int
	var table []Table_Struct.Repair_table
	db.Find(&table)
	max := table[0].Rnumber
	log.Println(max)
	for i:=0; i < len(table)-1; i++{
		if (max < table[i+1].Rnumber){
			max = table[i+1].Rnumber
			log.Println("11111")
			log.Println(table[i].Rnumber)
			log.Println(table[i+1].Rnumber)
			log.Println(max)
		}
		log.Println(i)
	}
	log.Println("max")
	log.Println(max)
	if len(table) == 0 {
		number = 1
	}else{
		number__ := util.DeleteTailBlank(max)
		number_,_ := strconv.Atoi(number__)
		//log.Println(strconv.Atoi(number__))
		number = number_ + 1
	}
	//log.Println(number)

	//自动生成订单时间
	//year := time.Now().Year()
	//month := time.Now().Month()
	//day := time.Now().Day()
	//hour := time.Now().Hour()
	//minute := time.Now().Minute()
	time_ := time.Now().Unix()
	time := time.Unix(time_,0).Format("2006-01-02 15:04:05")

	//ID := 150203200109233123
	//计算预估时间,预估价格
	var problemtime[] Table_Struct.Problem
	var ptime string
	var price string
	log.Println("1111111111111111111111111111111")
	log.Println(fault)
	db.Where("Pname = ?",fault).Find(&problemtime)
	var i int
	for i=0;i<len(problemtime);i++ {
		if util.DeleteTailBlank(problemtime[i].Pproduct) == util.DeleteTailBlank(product) {
			log.Println("22222222")
			if util.DeleteTailBlank(problemtime[i].Pname) == util.DeleteTailBlank(fault) {
				log.Println("1111111111")
				num1,_ := strconv.Atoi(problemtime[i].Pcost)
				num1 = num1 + 400
				ptime = problemtime[i].Ptime
				price = strconv.Itoa(num1)
			}
		}
	}
	//创建维修单
	var newtable = Table_Struct.Repair_table{
		Rnumber: strconv.Itoa(number),
		Rtime : time,
		Rphenomenon : phenomenon,
		Rfault : fault,
		Rproduct : product,
		Rremarks : remarks,
		Rid : id,
		Rdistribution : "否",
		Predicttime: ptime,
		Predictprice: price,
	}
	db.Create(&newtable)
	//返回结果
	response.SUCCESS(c, gin.H{}, "填写成功")
}
//客服进行报修设备的维护
func Repair_information(c *gin.Context) {
	db := Database.GetDB()
	//rnumber := c.PostForm("number")
	//rphenomenon := c.PostForm("phenomenon")
	//rfault := c.PostForm("fault")
	//rproduct := c.PostForm("product")
	//rremarks := c.PostForm("remarks")
	//rid := c.PostForm("rid")
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	rnumber := m["rnumber"]
	rphenomenon := m["rphenomenon"]
	rfault := m["rfault"]
	rproduct := m["rproduct"]
	rremarks := m["rremarks"]
	rid := m["rid"]

	var user Table_Struct.Repair_table
	db.Where("Rnumber = ?", rnumber).First(&user)
	if len(user.Rnumber)==0 {
		c.JSON(200,gin.H{"code":400, "msg":"该订单不存在"})
		return
	}
	Ctime := util.DeleteTailBlank(user.Rtime)

	rorder := Table_Struct.Repair_table{
		Rnumber: rnumber,
		Rtime: Ctime,
		Rphenomenon: rphenomenon,
		Rfault: rfault,
		Rproduct: rproduct,
		Rremarks: rremarks,
		Rid: rid,
	}
	var order Table_Struct.Repair_table
	db.Where("Rnumber = ?", rnumber).Delete(&order)
	db.Create(&rorder)
	c.JSON(200, gin.H{
		"msg": "修改成功",
	})
}
//返回客户确认单
func Confirm_repair(c *gin.Context) {
	db := Database.GetDB()
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	cid := m["cid"]
	//cid := c.PostForm("cid")
	var user Table_Struct.Customer
	db.Where("Cid = ?",cid).First(&user)
	var repair Table_Struct.Repair_table
	db.Where("Rid = ?",cid).First(&repair)
	repair1 := [] string{repair.Rtime,repair.Rnumber,repair.Rphenomenon,repair.Rfault,repair.Rproduct,user.Caddress,user.Caddress}
	//comfirmerepair
	//comfirmrepair := confirm{time:repair.Rtime,number:repair.Rnumber,rphenomenon: repair.Rphenomenon,rfault: repair.Rfault,rproduct: repair.Rproduct,caddress: user.Caddress,ctelephone: user.Ctelephone}
	c.JSON(200,repair1)
}
//维修维修记录
func Update_Table1(c *gin.Context) {
	db := Database.GetDB()
	//rnumber := c.PostForm("number")
	//rtime := c.PostForm("time")
	//rphenomenon := c.PostForm("phenomenon")
	//rfault := c.PostForm("fault")
	//rproduct := c.PostForm("product")
	//rremarks := c.PostForm("remarks")
	//rid := c.PostForm("rid")
	//Predictprice := c.PostForm("Predictprice")
	//Predicttime := c.PostForm("Predicttime ")
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	rnumber := m["Rnumber"]
	rtime := m["Rtime"]
	rphenomenon := m["Rphenomenon"]
	rfault := m["Rfault "]
	rproduct := m["Rproduct"]
	rremarks := m["Rremarks"]
	rid := m["Rid"]
	Predictprice := m["Predictprice"]
	Predicttime := m["Predicttime"]
	var user Table_Struct.Repair_table
	db.Where("Rnumber = ?", rnumber).First(&user)
	if len(user.Rnumber)==0 {
		c.JSON(200,gin.H{"code":30001, "msg":"该订单不存在"})
		return
	}
	//Ctime := util.DeleteTailBlank(user.Rtime)
	//log.Println(Ctime)
	//log.Println(rtime)
	//if Ctime!=rtime {
	//	c.JSON(200,gin.H{"code":30002,"msg":"订单创建时间不能修改哦"})
	//	return
	//}
	rorder := Table_Struct.Repair_table{
		Rnumber: rnumber,
		Rtime: rtime,
		Rphenomenon: rphenomenon,
		Rfault: rfault,
		Rproduct: rproduct,
		Rremarks: rremarks,
		Rid: rid,
		Predicttime: Predicttime,
		Predictprice: Predictprice,
	}
	var order Table_Struct.Repair_table
	db.Where("Rnumber = ?", rnumber).Delete(&order)
	db.Create(&rorder)
	response.SUCCESS(c, gin.H{}, "修改成功")
}
//返回结算单
func Settlement_sheet(c *gin.Context) {
	db := Database.GetDB()
	//cid := c.PostForm("Cid")
	cid, ok:= c.Params.Get("Cid")
	if !ok {
		c.JSON(200, gin.H{"error": "无效身份证号"})
		return
	}
	log.Println(cid)
	var user Table_Struct.Repair_table
	db.Where("Rid = ?",cid).First(&user)
	var cost Table_Struct.Table_Cost
	db.Where("Tnumber = ?",user.Rnumber).First(&cost)
	type settable struct {
		Rnumber string
		Rtime string
		Rphenomenon string
		Rfault string
		Rproduct string
		Rremarks string
		Tproblem_cost string
		Tengineer_cost string
		Tsparepart_cost string
		Tcost string
	}
	settlement_sheet1 := settable{
		Rnumber : user.Rnumber,
		Rtime: user.Rtime,
		Rphenomenon: user.Rphenomenon,
		Rfault: user.Rfault,
		Rproduct: user.Rproduct,
		Rremarks: user.Rremarks,
		Tproblem_cost: cost.Tproblem_cost,
		Tengineer_cost: cost.Tengineer_cost,
		Tsparepart_cost: cost.Tsparepart_cost,
		Tcost: cost.Tcost,
	}
	settlement_sheet :=[]settable{settlement_sheet1}
	//var settlement_sheet = map[string] string{"Rnumber":user.Rnumber,"Rtime":user.Rtime,"Rphenomenon":user.Rphenomenon,"Rproduct":user.Rproduct,"Rfault":user.Rfault,"Rremarks":user.Rremarks,"Tproblem_cost":cost.Tproblem_cost,"Tengineer_cost":cost.Tengineer_cost,"Tsparepart_cost":cost.Tsparepart_cost,"Tcost":cost.Tcost}
	c.JSON(200,settlement_sheet)
	log.Println(settlement_sheet)
	//response.SUCCESS(c, gin.H{"data":settlement_sheet}, "修改成功")
}
//返回维修表的信息
func GetcustomerTable(c *gin.Context){
	DB := Database.GetDB()
	var table[] Table_Struct.Repair_table
	log.Println(table)
	if err := DB.Find(&table).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else{
		c.JSON(200, table)
		log.Println(table)
	}
}