package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"repair/Database"
	"repair/Router"
	"repair/middleware"
)


func main() {

	db := Database.ConnectDB()
	defer db.Close()
	r := gin.Default()
	r.Use(middleware.Cors())
	r = Router.CollocetRouter(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务

}
