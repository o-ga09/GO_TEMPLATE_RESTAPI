package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}
	
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		return gormDB, mock, err
	}
	return gormDB, mock, err
}