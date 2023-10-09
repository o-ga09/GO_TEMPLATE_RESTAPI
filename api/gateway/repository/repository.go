package repository

import (
	"github.com/google/uuid"
)

//go:generate moq -out repository_mock.go . RepositoryInterface
type RepositoryInterface interface {
	GetAll() (RepositoryJsons, error)
	GetById(id uuid.UUID) (RepositoryJson, error)
	Create(RepositoryParamJson) error
	Update(id uuid.UUID, param RepositoryParamJson) error
	Delete(id uuid.UUID) error
}

type RepositoryJsons []RepositoryJson
type RepositoryJson struct {
	Id       int64  `db:"id" gorm:"primaryKey,autoincrement" json:"id,omitempty"`
	Userid   string `db:"userid" gorm:"size:255;not null" json:"userid,omitempty" validate:"required"`
	Username string `db:"username" gorm:"size:255;not null" json:"username,omitempty" validate:"required"`
}

type RepositoryParamJson struct {
	Userid string `db:"userid" gorm:"size:255;not null"`
	Username string `db:"username" gorm:"size:255;not null"`
}

func (RepositoryJson) TableName() string {
	return "user"
}

func (RepositoryParamJson) TableName() string {
	return "user"
}

