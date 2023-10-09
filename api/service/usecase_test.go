package service

import (
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/domain"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/service/port"
)

func TestFindUsers(t *testing.T) {
	userServiceMock := &port.InputPortMock{
		GetAllFunc: func() (domain.Entities, error) {
			return domain.Entities{}, nil
		},
	}

	userservice := NewUserService(userServiceMock)
	res, err := userservice.FindUsers()

	expected := domain.Entities{}
	assert.Equal(t,res,expected)
	assert.Equal(t,err,nil)
}

func TestFindUserById(t *testing.T) {
	userServiceMock := &port.InputPortMock{
		GetByIdFunc: func(id domain.UserID) (domain.Entity, error) {
			return domain.Entity{}, nil
		},
	}

	userservice := NewUserService(userServiceMock)

	id := domain.UserID{V: uuid.New()}
	res, err := userservice.FindUserById(id)

	expected := domain.Entity{}
	assert.Equal(t,res,expected)
	assert.Equal(t,err,nil)
}

func TestCreateUser(t *testing.T) {
	userServiceMock := &port.InputPortMock{
		CreateFunc: func(param domain.CreateJson) error {
			return nil
		},
	}

	userservice := NewUserService(userServiceMock)
	
	param := domain.CreateJson{User: domain.UserID{V: uuid.New()}, Name: domain.UserName{V: "testUser1"}}
	err := userservice.CreateUser(param)

	assert.Equal(t,err,nil)
}

func TestUpdateUser(t *testing.T) {
	userServiceMock := &port.InputPortMock{
		UpdateFunc: func(id domain.UserID ,param domain.CreateJson) error {
			return nil
		},
	}

	userservice := NewUserService(userServiceMock)
	
	id := domain.UserID{V: uuid.New()}
	param := domain.CreateJson{User: domain.UserID{V: uuid.New()}, Name: domain.UserName{V: "testUser1"}}
	err := userservice.UpdateUser(id,param)

	assert.Equal(t,err,nil)
}

func TestDeleteUser(t *testing.T) {
	userServiceMock := &port.InputPortMock{
		DeleteFunc: func(id domain.UserID) error {
			return nil
		},
	}

	userservice := NewUserService(userServiceMock)
	
	id := domain.UserID{V: uuid.New()}
	err := userservice.DeleteUser(id)

	assert.Equal(t,err,nil)
}