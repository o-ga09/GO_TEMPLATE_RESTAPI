package scheme

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UID              string `gorm:"primaryKey;column:uid;type:varchar(255)"`
	Email            string `gorm:"column:email;type:varchar(255)"`
	Password         string `gorm:"column:password;type:varchar(255)"`
	UserID           string `gorm:"column:user_id;type:varchar(50)"`
	FirstName        string `gorm:"column:first_name;type:varchar(50)"`
	LastName         string `gorm:"column:last_name;type:varchar(50)"`
	Gender           string `gorm:"column:gender;type:varchar(255)"`
	BirthDay         string `gorm:"column:birth_day;type:date"`
	PhoneNumber      string `gorm:"column:phone_number;type:varchar(20)"`
	PostOfficeNumber string `gorm:"column:post_office_number;type:varchar(20)"`
	Pref             string `gorm:"column:pref;type:varchar(20)"`
	City             string `gorm:"column:city;type:varchar(50)"`
	Extra            string `gorm:"column:extra;type:varchar(255)"`
}
