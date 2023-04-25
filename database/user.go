package database

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint   `gorm:"AUTO_INCREMENT;primaryKey"`
	Email    string `gorm:"type:varchar(20)"`
	PassWord string `gorm:"type:varchar(20);unique_index"`
	NickName string `gorm:"size:50"`
}

// func AddUserByQuery(context *gin.Context) {
// 	db := Conn()

// 	db.FirstOrCreate(&User{
// 		Email:    context.Query("email"),
// 		PassWord: context.Query("password"),
// 		NickName: context.Query("nickname"),
// 	})
// }

func AddUserByJsonBody(context *gin.Context) {
	var user User
	err := json.NewDecoder(context.Request.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	db := Conn()

	db.Where(User{Email: user.Email}).FirstOrCreate(&User{
		Email:    user.Email,
		PassWord: user.PassWord,
		NickName: user.NickName,
	})

	context.JSON(http.StatusOK, gin.H{
		"nickname": user.NickName,
		"email":    user.Email,
	})
}
