package Router

import (
	"github.com/gin-gonic/gin"
	"repair/Controller"
)

func CollocetRouter(r *gin.Engine) *gin.Engine {
	//用户router
	r.POST("/repair/customer/register", Controller.Register)
	r.POST("/repair/customer/login", Controller.Customer_Login)
	r.PUT("/repair/customer/modify",Controller.Customer_modify)
	r.POST("/repair/customer/set_table", Controller.Set_Table)
	r.GET("/repair/customer/gettable", Controller.GetTable)
	r.GET("/repair/customer/get_table", Controller.Select_Table)//查询维修表
	r.DELETE("/repair/customer/delete_table",Controller.Delete_Table)
	r.PUT("/repair/customer/update_table",Controller.Update_Table)
	//员工router
	r.POST("/repair/staff/login", Controller.Staff_Login)
	r.GET("/repair/staff/task_scheduling/Undistribution", Controller.Get_Undistribution)//查询未分配工程师的维修单
	r.GET("/repair/staff/task_scheduling/distribution", Controller.Distribution_Engineer)
	//库存router
	r.POST("/repairock/ad",Controller.Setsparepart)
	r.POST("/repairock/in",Controller.Stock_in)
	r.POST("/repairock/out",Controller.StockOut)
	r.POST("/repairock/inquire",Controller.Stock_inquire)
	r.POST("/repairock/inquire/single",Controller.Stock_inquire_single)

	return r
}
