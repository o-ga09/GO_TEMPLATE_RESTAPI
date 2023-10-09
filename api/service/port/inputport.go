package port

import "github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/domain"

//go:generate moq -out inputport_mock.go . InputPort
type InputPort interface {
	GetAll() (domain.Entities, error)
	GetById(id domain.UserID) (domain.Entity, error)
	Create(domain.CreateJson) error
	Update(id domain.UserID,param domain.CreateJson) error
	Delete(id domain.UserID) error
}