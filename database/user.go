package database

type User struct {
	ID       uint   `gorm:"AUTO_INCREMENT;primaryKey"`
	Email    string `gorm:"type:varchar(20)"`
	PassWord string `gorm:"type:varchar(20);unique_index"`
	NickName string `gorm:"size:50"`
}

func AddUser() {
	db := Conn()

	// Create
	db.Create(&User{
		ID:       0,
		Email:    "i314q159@outlook.com",
		PassWord: "314159",
		NickName: "he wen bao",
	})
}
