package user

import "context"

//go:generate moq -out UserServiceRepository_mock.go . UserServiceRepository
type UserServiceRepository interface {
	FindUser(ctx context.Context) ([]*User, error)
	FindUserById(ctx context.Context, id string) (*User, error)
	Save(ctx context.Context, param *User) error
	Delete(ctx context.Context, id string) error
}
