package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(100)" json:"Name"`
	Email    string `gorm:"type:varchar(200)" json:"Email"`
	Password string `gorm:"type:varchar(200)" json:"Password"`
}
