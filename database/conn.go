package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conn() *gorm.DB {
	dsn := "root:i314q159@tcp(127.0.0.1:3306)/tsp?charset=utf8mb4&parseTime=True&loc=Local"
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
