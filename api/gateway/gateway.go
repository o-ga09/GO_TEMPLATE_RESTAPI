package gateway

import (
	"errors"

	"github.com/google/uuid"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/domain"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/gateway/repository"
)

type UserGateway struct {
	driver repository.RepositoryInterface
}

func NewUserGateWay(r repository.RepositoryInterface) *UserGateway {
	return &UserGateway{r}
}

func (g *UserGateway) GetAll() (domain.Entities, error) {
	res, err := g.driver.GetAll()
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("0件のレコードを取得しました")
	}

	users := domain.Entities{}

	for _, r := range res {
		suuid, _ := uuid.Parse(r.Userid)
		u := domain.Entity{
			Id: r.Id,
			User: domain.UserID{V:suuid},
			Name: domain.UserName{V: r.Username},
		}

		users = append(users,u)
	}

	return users, nil
}

func (g *UserGateway) GetById(id domain.UserID) (domain.Entity, error) {
	userid := id.V
	res, err := g.driver.GetById(userid)

	if err != nil {
		return domain.Entity{}, err
	}

	suuid, err := uuid.Parse(res.Userid)
	user := domain.Entity{
		Id: res.Id,
		User: domain.UserID{V: suuid},
		Name: domain.UserName{V: res.Username},
	}

	return user, nil
}

func (g *UserGateway) Create(p domain.CreateJson) error {
	param := repository.RepositoryParamJson{
		Userid: p.User.V.String(),
		Username: p.Name.V,
	}
	err := g.driver.Create(param)

	return err
}

func (g *UserGateway) Update(id domain.UserID,p domain.CreateJson) error {
	userid := id.V
	param := repository.RepositoryParamJson{
		Userid: p.User.V.String(),
		Username: p.Name.V,
	}
	err := g.driver.Update(userid,param)

	return err
}

func (g *UserGateway) Delete(id domain.UserID) error {
	userid := id.V
	err := g.driver.Delete(userid)

	return err
}