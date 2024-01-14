package usecase

import (
	"context"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
)

type SaveUserUsecase struct {
	uds user.IUserDomainService
	ads administrator.IAdminDomainService
}

func NewSaveUserUsecase(uds user.IUserDomainService, ads administrator.IAdminDomainService) *SaveUserUsecase {
	return &SaveUserUsecase{uds: uds, ads: ads}
}

func (us *SaveUserUsecase) Run(ctx context.Context, param *user.User) error {
	// context.Contextの値を取り出す
	value, ok := ctx.Value("user_id").(string)
	if !ok {
		return INVALID_USER_ID
	}

	adminUser, err := us.ads.FindUser(ctx, value)
	if err != nil {
		return INVALID_ADMIN
	}

	if adminUser.GetAdmin() == 1 || param.GetUUID() == value {
		if err := us.uds.EditUser(ctx, param); err != nil {
			return err
		}
		return nil
	}

	if adminUser.GetAdmin() != 1 {
		return INVALID_ADMIN
	}

	return INVALID_USER_ID
}
