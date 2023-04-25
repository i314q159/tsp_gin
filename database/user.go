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

// add user by json
func AddUser(context *gin.Context) {
	var user User
	err := json.NewDecoder(context.Request.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	db := Conn()

	// TODO: 邮箱被占用，返回一个gin.H{"err": "邮箱被占用"}
	// TODO: 其他字段不存在，会报错，值非空检测
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

// update user by json
func UpdateUser(context *gin.Context) {

	var user User
	err := json.NewDecoder(context.Request.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	db := Conn()

	// TODO: 邮箱不存在，返回一个值
	// TODO: 其他字段不存在，会报错；值为""，不会更新原有数据
	db.Where(User{Email: user.Email}).Model(&user).Updates(&User{
		Email:    user.Email,
		PassWord: user.PassWord,
		NickName: user.NickName,
	})

	context.JSON(http.StatusOK, gin.H{
		"nickname": user.NickName,
		"email":    user.Email,
	})
}
