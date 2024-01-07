package scheme

import "gorm.io/gorm"

type Administrator struct {
	gorm.Model
	User_id string `gorm:"primaryKey;column:user_id;type:varchar(255)"`
	Admin   int `gorm:"column:admin;type:int"`
}
