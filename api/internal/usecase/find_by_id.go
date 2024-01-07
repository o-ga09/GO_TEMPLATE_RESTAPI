package usecase

import (
	"context"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
)

type FindUserByIdUsecase struct {
	uds user.UserDomainService
	ads administrator.AdminDomainService
}

func NewFindUserByIdUsecase(uds user.UserDomainService, ads administrator.AdminDomainService) *FindUserByIdUsecase {
	return &FindUserByIdUsecase{uds: uds, ads: ads}
}

func (us *FindUserByIdUsecase) Run(ctx context.Context, id string) (*user.User, error) {
	// context.Contextの値を取り出す
	value, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, INVALID_USER_ID
	}

	adminUser, _ := us.ads.FindUser(ctx, value)
	if adminUser.GetAdmin() == 1 || id == value {
		user, err := us.uds.FindUserById(ctx, id)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	if adminUser.GetAdmin() != 1 {
		return nil, INVALID_ADMIN
	}

	return nil, INVALID_USER_ID
}
