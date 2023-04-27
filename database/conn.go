package database

import (
	"fmt"
	"tsp_gin/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conn() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DB_USER, conf.DB_PASSWORD, conf.DB_IP, conf.DB_PORT, conf.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	createTable(db)

	return db
}

func createTable(db *gorm.DB) {
	// db.AutoMigrate(&User{}, &Product{}, &Order{})
	db.AutoMigrate(&User{})
}
