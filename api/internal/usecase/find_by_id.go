package usecase

import (
	"context"
	"log/slog"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
	"github.com/o-ga09/api/pkg"
)

type FindUserByIdUsecase struct {
	uds user.IUserDomainService
	ads administrator.IAdminDomainService
}

func NewFindUserByIdUsecase(uds user.IUserDomainService, ads administrator.IAdminDomainService) *FindUserByIdUsecase {
	return &FindUserByIdUsecase{uds: uds, ads: ads}
}

func (us *FindUserByIdUsecase) Run(ctx context.Context, id string) (*user.User, error) {
	value, ok := ctx.Value("ctxInfo").(pkg.CtxInfo)
	if !ok {
		return nil, INVALID_REQUEST_ID
	}

	user, err := us.uds.FindUserById(ctx, id)
	if err != nil {
		slog.Info("can not complete FindById usecase", "request id", value.RequestId)
		return nil, err
	}

	slog.Info("process done FindById usecase", "request id", value.RequestId)
	return user, nil
}
