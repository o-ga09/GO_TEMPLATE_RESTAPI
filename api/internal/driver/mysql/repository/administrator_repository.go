package repository

import (
	"context"
	"errors"

	administratorDomain "github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/driver/mysql/scheme"
	"gorm.io/gorm"
)

type AdminDriver struct {
	conn *gorm.DB
}

// Delete implements administrator.AdminServiceRepository.
func (ad *AdminDriver) Delete(ctx context.Context, id string) error {
	user := scheme.Administrator{}
	err := ad.conn.Where("user_id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// FindUser implements administrator.AdminServiceRepository.
func (ad *AdminDriver) FindUser(ctx context.Context, id string) (*administratorDomain.Administrator, error) {
	user := scheme.Administrator{}

	err := ad.conn.Where("user_id = ?", id).Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	res := administratorDomain.NewAdministrator(user.User_id, user.Admin)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Save implements administrator.AdminServiceRepository.
func (ad *AdminDriver) Save(ctx context.Context, param *administratorDomain.Administrator) error {
	repoUser := scheme.Administrator{
		User_id: param.GetUUID(),
		Admin:   param.GetAdmin(),
	}

	err := ad.conn.Save(&repoUser).Error
	return err
}

func NewAdminDriver(conn *gorm.DB) administratorDomain.AdminServiceRepository {
	return &AdminDriver{conn: conn}
}
