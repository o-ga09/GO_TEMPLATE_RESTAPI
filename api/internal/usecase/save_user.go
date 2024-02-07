package usecase

import (
	"context"
	"log/slog"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
	"github.com/o-ga09/api/pkg"
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
	value, ok := ctx.Value("ctxInfo").(pkg.CtxInfo)
	if !ok {
		return INVALID_REQUEST_ID
	}

	if err := us.uds.EditUser(ctx, param); err != nil {
		slog.Error("can not complete SaveUser Usecase", "error msg", err, "request id", value.RequestId)
		return err
	}

	slog.Info("process done SaveUser Usecase", "request id", value.RequestId)
	return nil
}
