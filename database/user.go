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

func UserRegister(context *gin.Context) {
	var user User
	err := json.NewDecoder(context.Request.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	db := Conn()

	var user2 User
	//判断邮箱是否存在
	db.Where("email = ?", user.Email).First(&user2)
	if user2.ID != 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户已存在",
		})
		return
	}

	//创建用户
	db.Create(&user)

	context.JSON(http.StatusOK, gin.H{
		"nickname": user.NickName,
		"email":    user.Email,
		"msg":      "用户注册成功",
	})
}

func UserLogin(context *gin.Context) {
	var user User
	err := json.NewDecoder(context.Request.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	db := Conn()

	var user2 User
	//判断邮箱是否存在
	db.Where("email = ?", user.Email).First(&user2)
	if user2.ID == 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}
	if user.PassWord != user2.PassWord {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"nickname": user.NickName,
		"email":    user.Email,
		"msg":      "用户登录成功",
	})
}
