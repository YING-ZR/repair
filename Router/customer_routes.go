package Router

import (
	"github.com/gin-gonic/gin"
	"repair/Controller"
)

func CollocetRouter(r *gin.Engine) *gin.Engine {
	r.POST("/repair/customer/register", Controller.Register)
	r.POST("/repair/customer/login", Controller.Customer_Login)
	return r
}
