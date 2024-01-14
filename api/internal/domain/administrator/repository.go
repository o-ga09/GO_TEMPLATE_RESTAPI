package administrator

import "context"

//go:generate moq -out AdminServiceRepository_mock.go . AdminServiceRepository
type AdminServiceRepository interface {
	FindUser(ctx context.Context, id string) (*Administrator, error)
	Save(ctx context.Context, param *Administrator) error
	Delete(ctx context.Context, id string) error
}
