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
	r.GET("repair/customer/get_customer/:Id",Controller.Get_personal)
	r.POST("/repair/customer/set_table/:Hno", Controller.Set_Table)
	r.GET("/repair/customer/gettable/:Hno", Controller.GetTable)
	r.GET("/repair/customer/get_table", Controller.Select_Table)//查询维修表
	r.DELETE("/repair/customer/delete_table",Controller.Delete_Table)
	r.PUT("/repair/customer/update_table",Controller.Update_Table)
	r.GET("/repair/customer/Maintenance_status/:Hno",Controller.Maintenance_status)
	r.GET("/repair/customer/Engineer_imag/:Hno",Controller.Engineer_imag)
	r.GET("/repair/customer/Settlecustomer_sheet/:Cid",Controller.Settlecustomer_sheet)//客服返回结算单
	r.DELETE("/repair/customer/Shanchu_repairtable/:Cid",Controller.Shanchu_repairtable)
	//员工router
	r.POST("/repair/staff/login", Controller.Staff_Login)//员工登陆
	//任务调度
	r.GET("/repair/staff/task_scheduling/Undistribution", Controller.Get_Undistribution)//查询未分配工程师的维修单
	r.GET("/repair/staff/task_scheduling/Distribution", Controller.Distribution_Engineer)//分配工程师
	r.GET("/repair/staff/task_scheduling/GetDistribution", Controller.Get_Distribution)//返回分配表单
	r.GET("/repair/staff/task_scheduling/Distribution_Table/:Hno", Controller.Get_Distribution_table)//根据分配表订单号查询维修表
	//库存router
	r.POST("/repairock/ad",Controller.Setsparepart)//新增备件
	r.POST("/repairock/in",Controller.Stock_in)//添加备件
	r.POST("/repairock/out",Controller.StockOut)//出库
	r.GET("/repairock/inquire",Controller.Stock_inquire)//查询整个备件库存
	r.GET("/repairock/inquire/single/:Hno",Controller.Stock_inquire_single)//查询单个备件库存
	r.DELETE("/repair/stock/delete",Controller.Stock_delete)//删除备件

	//技术工程师
	r.GET("/repair/staff/engineer/GetTable/:Number", Controller.Engineer_Get_Repair_Table)//返回分配表单
	r.GET("/repair/staff/engineer/GetDistribution/:Hno", Controller.Engineer_Get_Distribution)//查询distribution表
	r.POST("/repair/staff/engineer/Get_Sparepart/:Hno",Controller.Engineer_Get_Sparepart)//备件
	r.PUT("/repair/staff/engineer/Update_Distribution_ddelay_state/:Hno", Controller.Engineer_Update_Distribution_ddelay_state)//修改维修延迟程度状态
	r.PUT("/repair/staff/engineer/Update_Distribution_status/:Hno", Controller.Engineer_Update_Distribution_status)//维修完成
	//财务
	r.GET("/repair/staff/finance/Finance_Get_Distribution", Controller.Finance_Get_Distribution)//查询distribution表
	r.GET("/repair/staff/finance/Finance_Get_table/:Hno", Controller.Finance_Get_table)//查询维修单表
	r.GET("/repair/staff/finance/Finance_Get_Engineer/:Hno", Controller.Finance_Get_Engineer)//查询工程师表
	r.GET("/repair/staff/finance/Finance_Settlement/:Hno", Controller.Finance_Settlement)//结算单计算
	//客服
	r.POST("/repair/service/enter",Controller.Customerservice_enter)//客服录入客户信息
	r.PUT("/repair/service/modify",Controller.Customerservice_modify)//客服修改客户信息
	r.DELETE("/repair/service/delete",Controller.Customer_delete)//客服删除客户信息
	r.GET("/repair/service/inquire",Controller.Customer_inquire)//客户查看整个客户表的信息
	r.GET("/repair/service/inquire/single/:Cid",Controller.Customer_inquire_single)//客服查询某个客户的个人信息
	r.PUT("/repair/service/modify/repair",Controller.Repair_information)//客户进行报修设备的维护
	r.PUT("/repair/service/modify/record",Controller.Update_Table1)//客服进行报修记录的维护
	r.POST("/repair/service/Confirmrepair",Controller.Confirm_repair)//客服返回一个客户确认单
	r.GET("/repair/service/GetcustomerTable",Controller.GetcustomerTable)//返回整个维修单的信息
	r.POST("repair/service/setrepair",Controller.Customer_Settable)//客服填写维修单
	r.GET("/repair/service/settablement_sheet/:Cid",Controller.Settlement_sheet)//客服返回一个结算单
	//r.DELETE("/repair/service/shanchuweixiudan",Controller.Shanchuweixiudan)
	return r
}
