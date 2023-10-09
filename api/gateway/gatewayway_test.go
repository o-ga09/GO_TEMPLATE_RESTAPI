package gateway

import (
	"errors"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/domain"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/gateway/repository"
)

func TestGetAll(t *testing.T) {
	userGatewayMock := &repository.RepositoryInterfaceMock{
		GetAllFunc: func() (repository.RepositoryJsons, error) {
			return repository.RepositoryJsons{}, nil
		},
	}

	usergateway := NewUserGateWay(userGatewayMock)
	res, err := usergateway.GetAll()

	expected := domain.Entities{}
	assert.Equal(t, len(res), len(expected))
	assert.Equal(t, err, errors.New("0件のレコードを取得しました"))
}

func TestGetById(t *testing.T) {
	userGatewayMock := &repository.RepositoryInterfaceMock{
		GetByIdFunc: func(id uuid.UUID) (repository.RepositoryJson, error) {
			return repository.RepositoryJson{}, nil
		},
	}

	id := domain.UserID{V: uuid.New()}
	usergateway := NewUserGateWay(userGatewayMock)
	res, err := usergateway.GetById(id)

	expected := domain.Entity{}
	assert.Equal(t, res, expected)
	assert.Equal(t, err, nil)
}

func TestCreate(t *testing.T) {
	userGatewayMock := &repository.RepositoryInterfaceMock{
		CreateFunc: func(param repository.RepositoryParamJson) error {
			return nil
		},
	}

	usergateway := NewUserGateWay(userGatewayMock)
	
	param := domain.CreateJson{}
	err := usergateway.Create(param)

	assert.Equal(t, err, nil)
}

func TestUpdate(t *testing.T) {
	userGatewayMock := &repository.RepositoryInterfaceMock{
		UpdateFunc: func(id uuid.UUID, param repository.RepositoryParamJson) error {
			return nil
		},
	}

	usergateway := NewUserGateWay(userGatewayMock)
	
	id := domain.UserID{V: uuid.New()}
	param := domain.CreateJson{}
	err := usergateway.Update(id,param)

	assert.Equal(t, err, nil)
}

func TestDelete(t *testing.T) {
	userGatewayMock := &repository.RepositoryInterfaceMock{
		DeleteFunc: func(id uuid.UUID) error {
			return nil
		},
	}

	usergateway := NewUserGateWay(userGatewayMock)
	
	id := domain.UserID{V: uuid.New()}
	err := usergateway.Delete(id)

	assert.Equal(t, err, nil)
}