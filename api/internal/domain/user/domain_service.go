package user

import "context"

type UserDomainService struct {
	repo UserServiceRepository
}

func NewUserDomainService(repo UserServiceRepository) *UserDomainService {
	return &UserDomainService{repo: repo}
}

func (uds *UserDomainService) FindUser(ctx context.Context) ([]*User, error) {
	users, err := uds.repo.FindUser(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uds *UserDomainService) FindUserById(ctx context.Context, id string) (*User, error) {
	user, err := uds.repo.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uds *UserDomainService) EditUser(ctx context.Context, param *User) error {
	err := uds.repo.Save(ctx, param)
	if err != nil {
		return err
	}
	return nil
}

func (uds *UserDomainService) DeleteUser(ctx context.Context, id string) error {
	err := uds.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
