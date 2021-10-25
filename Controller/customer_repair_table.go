package Controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"repair/Database"
	"repair/Table_Struct"
	"repair/response"
	"repair/util"
	"strconv"
	"time"
)
//type Repair_table struct {
//	Rnumber string
//	Rtime string
//	Rphenomenon string
//	Rfault string
//	Rproduct string
//	Rremarks string
//}

//填写维修单
func Set_Table(c *gin.Context) {
	db := Database.GetDB()
	id, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	//获取参数
	//cid := c.PostForm("cid")
	//phenomenon := c.PostForm("phenomenon")
	//fault := c.PostForm("fault")
	//product := c.PostForm("product")
	//remarks := c.PostForm("remarks")
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	//cid := m["cid"]
	phenomenon := m["rphenomenon"]
	fault := m["rfault"]
	product := m["rproduct"]
	remarks := m["rremarks"]
	var table_ Table_Struct.Repair_table
    db.Where("RID=?",id).First(&table_)
	if len(table_.Rid)!=0{
		c.JSON(200, gin.H{"code": 100010, "msg": "您已填写过维修单"})
		return
	}
	//数据验证
	if len(phenomenon) == 0 {
		c.JSON(200, gin.H{"code": 100001, "msg": "请输入机器故障现象"})
		return
	}
	if len(fault) == 0 {
		c.JSON(200, gin.H{"code": 100002, "msg": "请选择故障类型"})
		return
	}
	if len(product) == 0 {
		c.JSON(200, gin.H{"code": 100003, "msg": "请选择产品类型"})
		return
	}
	//自动生成订单号
	var number int
	var table []Table_Struct.Repair_table
	db.Find(&table)
	if len(table) == 0 {
		number = 1
	}else{
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
		number__ := util.DeleteTailBlank(max)
		number_,_ := strconv.Atoi(number__)
		//log.Println(strconv.Atoi(number__))
		number = number_ + 1
	}

	//自动生成订单时间
	//year := time.Now().Year()
	//month := time.Now().Month()
	//day := time.Now().Day()
	//hour := time.Now().Hour()
	//minute := time.Now().Minute()
	time_ := time.Now().Unix()
	time := time.Unix(time_,0).Format("2006-01-02 15:04:05")
	log.Println(time)

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
	log.Println(problemtime)
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
//获取信息
func GetTable(c *gin.Context){
	DB := Database.GetDB()
	id, ok:= c.Params.Get("Hno")
	if !ok {
		c.JSON(200, gin.H{"error": "无效订单号"})
		return
	}
	log.Println(id)
	var table Table_Struct.Repair_table
	log.Println(table)
	if err := DB.Where("Rid=?",id).Find(&table).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else{
		c.JSON(200, table)
		log.Println(table)
	}
}

//查询信息
func Select_Table(c *gin.Context){
	db := Database.GetDB()
	product := c.PostForm("product")
	var table []Table_Struct.Repair_table
	//db.Where("RPRODUCT =  ?", product).Find(&table)
	if err := db.Where("RPRODUCT =  ?", product).Find(&table).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else{
		c.JSON(200, table)
	}
}

//删除表单
func Delete_Table(c *gin.Context){
	db := Database.GetDB()
	rnumber := c.Query("id")
	if rnumber!=""{
		c.JSON(200, gin.H{
			"code":200,
			"msg" : "删除成功",
		})
	}
	var repair Table_Struct.Repair_table
	db.Where("Rnumber = ?", rnumber).Delete(&repair)
}


//更新表单
func Update_Table(c *gin.Context) {
	db := Database.GetDB()
	//rnumber := c.PostForm("number")
	//rtime := c.PostForm("time")
	//rphenomenon := c.PostForm("phenomenon")
	//rfault := c.PostForm("fault")
	//rproduct := c.PostForm("product")
	//rremarks := c.PostForm("remarks")
	//rid := c.PostForm("rid")
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	rnumber := m["rnumber"]
	rtime := m["rtime"]
	rphenomenon := m["rphenomenon"]
	rfault := m["rfault "]
	rproduct := m["rproduct"]
	rremarks := m["rremarks"]
	rid := m["rid"]
	var repair Table_Struct.Repair_table
	db.Where("Ris = ?",rid).First(&repair)
	Predictprice := repair.Predictprice
	Predicttime := repair.Predicttime
	var user Table_Struct.Repair_table
	db.Where("Rnumber = ?", rnumber).First(&user)
	if len(user.Rnumber)==0 {
		c.JSON(200,gin.H{"code":400, "msg":"该订单不存在"})
		return
	}
	Ctime := util.DeleteTailBlank(user.Rtime)
	if Ctime!=rtime {
		c.JSON(200,gin.H{"code":400,"msg":"订单创建时间不能修改哦"})
		return
	}
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
	c.JSON(200, gin.H{
		"msg": "修改成功",
	})
}
//查看结算单
func Settlecustomer_sheet(c *gin.Context) {
	db := Database.GetDB()
	//cid := c.PostForm("Cid")
	log.Println("111111111111111111111111111111111")
	cid, ok := c.Params.Get("Cid")
	if !ok {
		c.JSON(200, gin.H{"error": "无效身份证号"})
		return
	}
	log.Println("111111111111111111111111111111111")
	log.Println(cid)
	var user Table_Struct.Repair_table
	db.Where("Rid = ?", cid).First(&user)
	var cost Table_Struct.Table_Cost
	db.Where("Tnumber = ?", user.Rnumber).First(&cost)
	type settable struct {
		Rnumber         string
		Rtime           string
		Rphenomenon     string
		Rfault          string
		Rproduct        string
		Rremarks        string
		Tproblem_cost   string
		Tengineer_cost  string
		Tsparepart_cost string
		Tcost           string
	}
	settlement_sheet1 := settable{
		Rnumber:         user.Rnumber,
		Rtime:           user.Rtime,
		Rphenomenon:     user.Rphenomenon,
		Rfault:          user.Rfault,
		Rproduct:        user.Rproduct,
		Rremarks:        user.Rremarks,
		Tproblem_cost:   cost.Tproblem_cost,
		Tengineer_cost:  cost.Tengineer_cost,
		Tsparepart_cost: cost.Tsparepart_cost,
		Tcost:           cost.Tcost,
	}
	log.Println(settlement_sheet1.Tengineer_cost)
	//var settlement_sheet = map[string] string{"Rnumber":user.Rnumber,"Rtime":user.Rtime,"Rphenomenon":user.Rphenomenon,"Rproduct":user.Rproduct,"Rfault":user.Rfault,"Rremarks":user.Rremarks,"Tproblem_cost":cost.Tproblem_cost,"Tengineer_cost":cost.Tengineer_cost,"Tsparepart_cost":cost.Tsparepart_cost,"Tcost":cost.Tcost}
	c.JSON(200, settlement_sheet1)
	//response.SUCCESS(c, gin.H{"data":settlement_sheet}, "修改成功")
}
//删除完成订单后的用户订单
func Shanchu_repairtable(c *gin.Context) {
	db :=Database.GetDB()
	cid, ok := c.Params.Get("Cid")
	if !ok {
		c.JSON(200, gin.H{"error": "无效身份证号"})
		return
	}
	var repair Table_Struct.Repair_table
	var table Table_Struct.Distribution
	var cost Table_Struct.Table_Cost
	db.Where("Rid = ?",cid).First(&repair)
	log.Println(repair.Rnumber)
	db.Where("Dtable_number = ?",repair.Rnumber).Delete(&table)
	db.Where("Tnumber = ?",repair.Rnumber).Delete(&cost)
	db.Where("Rid = ?",cid).Delete(&repair)
	response.SUCCESS(c, gin.H{}, "删除成功")

}