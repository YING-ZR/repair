package Database

import (
	"github.com/jinzhu/gorm"
	"log"
	"repair/Table_Struct"
)

var DB *gorm.DB

//连接数据库
func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=repair password=010923 sslmode=disable")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("连接")
	db.AutoMigrate(&Table_Struct.Customer{})
	db.SingularTable(true)
	DB = db
	return db
}


func GetDB() *gorm.DB{
	return DB
}
