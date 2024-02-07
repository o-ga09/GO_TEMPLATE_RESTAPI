package usecase

import (
	"context"
	"errors"
	"log/slog"
	"strconv"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
	"github.com/o-ga09/api/pkg"
)

var (
	INVALID_USER_ID    = errors.New("invalid user id")
	INVALID_ADMIN      = errors.New("invalid admin")
	INVALID_REQUEST_ID = errors.New("invalid request id")
)

type FindUserUsecase struct {
	uds user.IUserDomainService
	ads administrator.IAdminDomainService
}

type FindUserUsecaseDto struct {
	TotalCount int64                     `json:"total_count,omitempty"`
	Offset     int                       `json:"offset,omitempty"`
	User       []FindUserUsecaseDtoModel `json:"user,omitempty"`
}

type FindUserUsecaseDtoModel struct {
	ID               string `json:"id,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	User_ID          string `json:"user_id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Gender           string `json:"gender,omitempty"`
	BirthDay         string `json:"birth_day,omitempty"`
	PhoneNumber      string `json:"phone_number,omitempty"`
	PostOfficeNumber string `json:"post_office_number,omitempty"`
	Pref             string `json:"pref,omitempty"`
	City             string `json:"city,omitempty"`
	Extra            string `json:"extra,omitempty"`
}

func NewFindUserUsecase(uds user.IUserDomainService, ads administrator.IAdminDomainService) *FindUserUsecase {
	return &FindUserUsecase{uds: uds, ads: ads}
}

func (us *FindUserUsecase) Run(ctx context.Context) (*FindUserUsecaseDto, error) {
	// context.Contextの値を取り出す
	value, ok := ctx.Value("ctxInfo").(pkg.CtxInfo)
	if !ok {
		return nil, INVALID_USER_ID
	}

	adminUser, err := us.ads.FindUser(ctx, value.UserId)
	if err != nil {
		slog.Error("can not get admin user", "error msg", err, "request id", value.RequestId)
		return nil, err
	}

	if adminUser.GetAdmin() != 1 {
		slog.Error("invalid user (not admin)", "request id", value.RequestId)
		return nil, INVALID_ADMIN
	}

	users, count, err := us.uds.FindUser(ctx)
	if err != nil {
		slog.Error("can not process FindUser Usecase", "error msg", err, "request id", value.RequestId)
		return nil, err
	}

	dtouser := []FindUserUsecaseDtoModel{}
	dto := FindUserUsecaseDto{}
	for _, u := range users {
		r := FindUserUsecaseDtoModel{
			ID:               u.GetUUID(),
			Email:            u.GetEmail(),
			Password:         u.GetPassWord(),
			User_ID:          u.GetID(),
			FirstName:        u.GetFirstName(),
			LastName:         u.GetLastName(),
			Gender:           u.GetGender(),
			BirthDay:         u.GetBirthDay(),
			PhoneNumber:      u.GetPhoneNumber(),
			PostOfficeNumber: u.GetPostOfficeNumber(),
			Pref:             u.GetPref(),
			City:             u.GetCity(),
			Extra:            u.GetExtra(),
		}
		dtouser = append(dtouser, r)
	}
	limit, _ := strconv.Atoi(value.Limit)
	offset, _ := strconv.Atoi(value.Offset)

	dto.TotalCount = count
	dto.Offset = limit + offset

	slog.Info("FindUserUsecase processing done ", "request_id", value.RequestId)
	return &dto, nil
}
