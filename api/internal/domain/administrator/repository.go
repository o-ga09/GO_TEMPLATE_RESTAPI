package administrator

import "context"

type AdminServiceRepository interface {
	FindUser(ctx context.Context, id string) (*Administrator, error)
	Save(ctx context.Context, param *Administrator) error
	Delete(ctx context.Context, id string) error
}
