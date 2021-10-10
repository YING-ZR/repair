package Router

import (
	"github.com/gin-gonic/gin"
	"repair/Controller"
)

func CollocetRouter(r *gin.Engine) *gin.Engine {
	r.POST("/repair/customer/register", Controller.Register)
	r.POST("/repair/customer/login", Controller.Customer_Login)
	r.PUT("/repair/customer/modify",Controller.Customer_modify)
	r.POST("/repair/customer/set_table", Controller.Set_Table)
	r.GET("/repair/customer/gettable", Controller.GetTable)
	r.GET("/repair/customer/get_table", Controller.Get_Table)
	r.DELETE("/repair/customer/delete_table",Controller.Delete_Table)
	r.PUT("/repair/customer/update_table",Controller.Update_Table)
	return r
}
