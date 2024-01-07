package usecase

import (
	"context"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
)

type DeleteUserUsecase struct {
	uds user.UserDomainService
	ads administrator.AdminDomainService
}

func NewDeleteUserUsecase(uds user.UserDomainService, ads administrator.AdminDomainService) *DeleteUserUsecase {
	return &DeleteUserUsecase{uds: uds, ads: ads}
}

func (us *DeleteUserUsecase) Run(ctx context.Context, id string) error {
	// context.Contextの値を取り出す
	value, ok := ctx.Value("user_id").(string)
	if !ok {
		return INVALID_USER_ID
	}

	adminUser, _ := us.ads.FindUser(ctx, value)
	if adminUser.GetAdmin() == 1 || id == value {
		if err := us.uds.DeleteUser(ctx, id); err != nil {
			return err
		}
		return nil
	}

	if adminUser.GetAdmin() != 1 {
		return INVALID_ADMIN
	}

	return INVALID_USER_ID
}