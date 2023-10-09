package service

import (
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/domain"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/service/port"
)

type UserService struct {
	inputport port.InputPort
}

func NewUserService(inputport port.InputPort) *UserService {
	return &UserService{inputport}
}

func (u *UserService) FindUsers() (domain.Entities, error) {
	res, err := u.inputport.GetAll()
	if err != nil {
		return domain.Entities{}, err
	}

	return res, nil
}

func (u *UserService) FindUserById(id domain.UserID) (domain.Entity, error) {
	res, err := u.inputport.GetById(id)
	if err != nil {
		return domain.Entity{}, err
	}

	return res, nil
}

func (u * UserService) CreateUser(param domain.CreateJson) error {
	err := u.inputport.Create(param)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) UpdateUser(id domain.UserID, param domain.CreateJson) error {
	err := u.inputport.Update(id,param)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) DeleteUser(id domain.UserID) error {
	err := u.inputport.Delete(id)
	if err != nil {
		return err
	}

	return nil
}