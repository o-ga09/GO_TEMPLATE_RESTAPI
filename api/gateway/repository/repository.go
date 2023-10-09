package repository

import "github.com/google/uuid"

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
	Id int64
	UserID string
	UserName string
}

type RepositoryParamJson struct {
	UserID string
	UserName string
}