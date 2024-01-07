package repository

import (
	"context"
	"errors"

	userDomain "github.com/o-ga09/api/internal/domain/user"
	"github.com/o-ga09/api/internal/driver/mysql/scheme"

	"gorm.io/gorm"
)

type UserDriver struct {
	conn *gorm.DB
}

// Delete implements user.UserServiceRepository.
func (ud *UserDriver) Delete(ctx context.Context, id string) error {
	user := scheme.User{}
	err := ud.conn.Where("uid = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// FindUser implements user.UserServiceRepository.
func (ud *UserDriver) FindUser(ctx context.Context) ([]*userDomain.User, error) {
	res := []*userDomain.User{}
	users := []*scheme.User{}
	err := ud.conn.Find(&users).Error
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		u := userDomain.NewUser(user.UID, user.Email, user.Password, user.UserID, user.FirstName, user.LastName, user.Gender, user.BirthDay, user.PhoneNumber, user.PostOfficeNumber, user.Pref, user.City, user.Extra)
		if err != nil {
			return nil, err
		}

		res = append(res, u)
	}

	return res, nil
}

// FindUserById implements user.UserServiceRepository.
func (ud *UserDriver) FindUserById(ctx context.Context, id string) (*userDomain.User, error) {
	user := scheme.User{}

	err := ud.conn.Where("user_id = ?", id).Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	res := userDomain.NewUser(user.UID, user.Email, user.Password, user.UserID, user.FirstName, user.LastName, user.Gender, user.BirthDay, user.PhoneNumber, user.PostOfficeNumber, user.Pref, user.City, user.Extra)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Save implements user.UserServiceRepository.
func (ud *UserDriver) Save(ctx context.Context, param *userDomain.User) error {
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
	return err
}

func NewUserDriver(conn *gorm.DB) userDomain.UserServiceRepository {
	return &UserDriver{conn: conn}
}
