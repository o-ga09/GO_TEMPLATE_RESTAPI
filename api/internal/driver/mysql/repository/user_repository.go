package repository

import (
	"context"
	"errors"
	"log/slog"
	"strconv"

	userDomain "github.com/o-ga09/api/internal/domain/user"
	"github.com/o-ga09/api/internal/driver/mysql/scheme"
	"github.com/o-ga09/api/pkg"

	"gorm.io/gorm"
)

type UserDriver struct {
	conn *gorm.DB
}

// Delete implements user.UserServiceRepository.
func (ud *UserDriver) Delete(ctx context.Context, id string) error {
	ctxValue := ctx.Value("ctxInfo").(pkg.CtxInfo)
	user := scheme.User{}
	err := ud.conn.Where("uid = ?", id).Delete(&user).Error
	if err != nil {
		slog.Info("can not complete DeleteUser Repository", "request id", ctxValue.RequestId)
		return err
	}

	slog.Info("process done DeleteUser Repository", "request id", ctxValue.RequestId)
	return nil
}

// FindUser implements user.UserServiceRepository.
func (ud *UserDriver) FindUser(ctx context.Context) ([]*userDomain.User, int64, error) {
	ctxValue := ctx.Value("ctxInfo").(pkg.CtxInfo)
	res := []*userDomain.User{}
	users := []*scheme.User{}
	var totalCount int64

	limit, err := strconv.Atoi(ctxValue.PageLimit)
	if err != nil {
		limit = 100
	}
	offset, err := strconv.Atoi(ctxValue.PageOffset)
	if err != nil {
		offset = 0
	}

	err = ud.conn.Limit(limit).Offset(offset).Order("id ASC").Find(&users).Count(&totalCount).Error
	if err != nil {
		slog.Error("can not complate FindByID Repository", "error msg", err, "request id", ctxValue.RequestId)
		return nil, 0, err
	}

	for _, user := range users {
		u := userDomain.NewUser(user.UID, user.Email, user.Password, user.UserID, user.FirstName, user.LastName, user.Gender, user.BirthDay, user.PhoneNumber, user.PostOfficeNumber, user.Pref, user.City, user.Extra)
		if err != nil {
			slog.Error("can not complate FindByID Repository", "error msg", err, "request id", ctxValue.RequestId)
			return nil, 0, err
		}

		res = append(res, u)
	}

	slog.Info("process done FindByID Repository", "request id", ctxValue.RequestId, "total count", totalCount)
	return res, totalCount, nil
}

// FindUserById implements user.UserServiceRepository.
func (ud *UserDriver) FindUserById(ctx context.Context, id string) (*userDomain.User, error) {
	ctxValue := ctx.Value("ctxInfo").(pkg.CtxInfo)
	user := scheme.User{}

	err := ud.conn.Where("user_id = ?", id).Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("can not complate FindByID Repository", "request id", ctxValue.RequestId)
		return nil, gorm.ErrRecordNotFound
	}

	res := userDomain.NewUser(user.UID, user.Email, user.Password, user.UserID, user.FirstName, user.LastName, user.Gender, user.BirthDay, user.PhoneNumber, user.PostOfficeNumber, user.Pref, user.City, user.Extra)
	if err != nil {
		slog.Error("can not complete FindByID Repository", "request id", ctxValue.RequestId)
		return nil, err
	}

	slog.Info("process done FindByID Repository", "request id", ctxValue.RequestId)
	return res, nil
}

// Save implements user.UserServiceRepository.
func (ud *UserDriver) Save(ctx context.Context, param *userDomain.User) error {
	ctxValue := ctx.Value("ctxInfo").(pkg.CtxInfo)

	repoUser := scheme.User{
		UID:              param.GetUUID(),
		UserID:           param.GetID(),
		Email:            param.GetEmail(),
		Password:         param.GetPassWord(),
		BirthDay:         param.GetBirthDay(),
		Gender:           param.GetGender(),
		PhoneNumber:      param.GetPhoneNumber(),
		PostOfficeNumber: param.GetPostOfficeNumber(),
		LastName:         param.GetLastName(),
		FirstName:        param.GetFirstName(),
		Pref:             param.GetPref(),
		City:             param.GetCity(),
		Extra:            param.GetExtra(),
	}

	err := ud.conn.Save(&repoUser).Error
	if err != nil {
		slog.Error("can not complete SaveUser Repository", "request id", ctxValue.RequestId)
		return err
	}

	slog.Info("process done SaveUser Repository", "request id", ctxValue.RequestId)
	return err
}

func NewUserDriver(conn *gorm.DB) userDomain.UserServiceRepository {
	return &UserDriver{conn: conn}
}
