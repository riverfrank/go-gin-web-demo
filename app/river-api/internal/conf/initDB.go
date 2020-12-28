package conf

import (
	"go-gin-web-demo/app/river-api/internal/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func SetupDB() {

	dsn := "root:root@tcp(127.0.0.1:3306)/ginhello?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(1)
	dao.DB = db
}
