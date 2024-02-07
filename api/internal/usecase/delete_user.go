package usecase

import (
	"context"
	"log/slog"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
	"github.com/o-ga09/api/pkg"
)

type DeleteUserUsecase struct {
	uds user.IUserDomainService
	ads administrator.IAdminDomainService
}

func NewDeleteUserUsecase(uds user.IUserDomainService, ads administrator.IAdminDomainService) *DeleteUserUsecase {
	return &DeleteUserUsecase{uds: uds, ads: ads}
}

func (us *DeleteUserUsecase) Run(ctx context.Context, id string) error {
	// context.Contextの値を取り出す
	value, ok := ctx.Value("ctxInfo").(pkg.CtxInfo)
	if !ok {
		return INVALID_REQUEST_ID
	}

	if err := us.uds.DeleteUser(ctx, id); err != nil {
		slog.Error("process done DeleteUser Usecase", "error msg", err, "request id", value.RequestId)
		return err
	}
	slog.Info("process done DeleteUser Usecase", "request id", value.RequestId)
	return nil
}
