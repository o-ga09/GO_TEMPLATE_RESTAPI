package usecase

import (
	"context"
	"testing"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
)

func TestDeleteUserUsecase_Run(t *testing.T) {
	type field struct {
		userRepoErr  error
		admin        *administrator.Administrator
		adminRepoErr error
	}
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		field     field
		args      args
		mockValue field
		user_id   string
		wantErr   bool
	}{
		{name: "正常系 - 管理者はユーザを削除できる", args: args{id: "000000000"}, mockValue: field{userRepoErr: nil, admin: admin_1, adminRepoErr: nil}, user_id: "999999999", wantErr: false},
		{name: "正常系 - 一般ユーザは、自身の情報を削除できる", args: args{id: "000000000"}, mockValue: field{userRepoErr: nil, admin: admin_2, adminRepoErr: nil}, user_id: "000000000", wantErr: false},
		{name: "異常系 - 一般ユーザは、他の情報を削除できない", args: args{id: "999999999"}, mockValue: field{userRepoErr: nil, admin: admin_2, adminRepoErr: nil}, user_id: "000000000", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUDS := user.IUserDomainServiceMock{
				DeleteUserFunc: func(ctx context.Context, id string) error {
					return tt.field.userRepoErr
				},
			}
			mockADS := administrator.IAdminDomainServiceMock{
				FindUserFunc: func(ctx context.Context, id string) (*administrator.Administrator, error) {
					return tt.mockValue.admin, tt.field.adminRepoErr
				},
			}

			uc := NewDeleteUserUsecase(&mockUDS, &mockADS)
			ctx := context.Background()
			ctx = context.WithValue(ctx, "user_id", tt.user_id)
			if err := uc.Run(ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserUsecase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
