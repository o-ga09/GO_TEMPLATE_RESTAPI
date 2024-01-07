package administrator

import "context"

type AdminDomainService struct {
	repo AdminServiceRepository
}

func NewAdminDomainService(repo AdminServiceRepository) *AdminDomainService {
	return &AdminDomainService{repo: repo}
}

func (uds *AdminDomainService) FindUser(ctx context.Context, id string) (*Administrator, error) {
	user, err := uds.repo.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uds *AdminDomainService) EditUser(ctx context.Context, param *Administrator) error {
	err := uds.repo.Save(ctx, param)
	if err != nil {
		return err
	}
	return nil
}

func (uds *AdminDomainService) DeleteUser(ctx context.Context, id string) error {
	err := uds.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
