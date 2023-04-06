package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func orm() {
	dsn := "root:i314q159@tcp(127.0.0.1:3306)/tsp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{
		ID: 0,
		//Email: "i314q159@outlook.com",
		//PassWord: "314159",
		NickName: "i314q159",
	})
}
