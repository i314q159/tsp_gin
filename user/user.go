package user

type User struct {
	ID       int    `gorm:"AUTO_INCREMENT"`
	Email    string `gorm:"type:varchar(25)"`
	PassWord string `gorm:"type:varchar(50);unique_index"`
	NickName string `gorm:"size:50"`
}
