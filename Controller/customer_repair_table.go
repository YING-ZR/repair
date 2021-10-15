package Controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"repair/Database"
	"repair/Table_Struct"
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
	//获取参数
	phenomenon := c.PostForm("phenomenon")
	fault := c.PostForm("fault")
	product := c.PostForm("product")
	remarks := c.PostForm("remarks")
	//数据验证
	if len(phenomenon) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请输入机器故障现象"})
		return
	}
	if len(fault) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请选择故障类型"})
		return
	}
	if len(product) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请选择产品类型"})
		return
	}
    //自动生成订单号
    var number int
	var table []Table_Struct.Repair_table
	db.Find(&table)
	if len(table) == 0 {
		number = 1
	}else{
		number__ := util.DeleteTailBlank(table[len(table)-1].Rnumber)
		number_,_ := strconv.Atoi(number__)
		log.Println(strconv.Atoi(number__))
		number = number_ + 1
	}
	log.Println(number)

	//自动生成订单时间
	//year := time.Now().Year()
	//month := time.Now().Month()
	//day := time.Now().Day()
	//hour := time.Now().Hour()
	//minute := time.Now().Minute()
	time_ := time.Now().Unix()
	time := time.Unix(time_,0).Format("2006-01-02 15:04:05")
	log.Println(time)

	ID := 150203200109233123

	//创建维修单
	var newtable = Table_Struct.Repair_table{
		Rnumber: strconv.Itoa(number),
		Rtime : time,
		Rphenomenon : phenomenon,
		Rfault : fault,
		Rproduct : product,
		Rremarks : remarks,
		Rid : strconv.Itoa(ID),
		Rdistribution : "否",
	}
	db.Create(&newtable)
	//返回结果
	c.JSON(200, gin.H{
		"msg": "填写成功",
	})
}

//获取信息
func GetTable(c *gin.Context){
	DB := Database.GetDB()
	var table []Table_Struct.Repair_table
	log.Println(table)
	if err := DB.Find(&table).Error; err != nil{
		c.JSON(200,gin.H{"error": err.Error()})
	}else{
		c.JSON(200, table)
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
	rnumber := c.PostForm("number")
	rtime := c.PostForm("time")
	rphenomenon := c.PostForm("phenomenon")
	rfault := c.PostForm("fault")
	rproduct := c.PostForm("product")
	rremarks := c.PostForm("remarks")
	rid := c.PostForm("rid")
	//b, _ := c.GetRawData()
	//var m map[string]string
	//_ = json.Unmarshal(b, &m)
	//rnumber := m["rnumber"]
	//rtime := m["rtime"]
	//rphenomenon := m["rphenomenon"]
	//rfault := m["rfault "]
	//rproduct := m["rproduct"]
	//rremarks := m["rremarks"]
	//rid := m["rid"]
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
	}
	var order Table_Struct.Repair_table
	db.Where("Rnumber = ?", rnumber).Delete(&order)
	db.Create(&rorder)
	c.JSON(200, gin.H{
		"msg": "修改成功",
	})
}
