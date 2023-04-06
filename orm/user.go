package orm

type User struct {
	ID uint `gorm:"AUTO_INCREMENT;primaryKey"`
	//Email string `gorm:"type:varchar()"`
	//PassWord string `gorm:"type:varchar();unique_index"`
	NickName string `gorm:"size:50"`
}
