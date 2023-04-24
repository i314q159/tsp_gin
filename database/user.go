package database

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint   `gorm:"AUTO_INCREMENT;primaryKey"`
	Email    string `gorm:"type:varchar(20)"`
	PassWord string `gorm:"type:varchar(20);unique_index"`
	NickName string `gorm:"size:50"`
}

// TODO: 以邮箱为准，防止重复插入
func AddUserByQuery(context *gin.Context) {
	db := Conn()

	// Create
	db.Create(&User{
		ID:       0,
		Email:    context.Query("email"),
		PassWord: context.Query("password"),
		NickName: context.Query("nickname"),
	})
}

// TODO: 以邮箱为准，防止重复插入
func AddUserByBody(context *gin.Context) {
	var user User
	err := json.NewDecoder(context.Request.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	db := Conn()
	db.Create(&User{
		ID:       0,
		Email:    user.Email,
		PassWord: user.PassWord,
		NickName: user.NickName,
	})
}
