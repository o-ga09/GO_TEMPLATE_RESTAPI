package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
	"github.com/o-ga09/api/pkg"
)

func TestFindUserByIdUsecase_Run(t *testing.T) {
	type field struct {
		user         user.User
		userRepoErr  error
		admin        *administrator.Administrator
		adminRepoErr error
	}
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		fields    field
		args      args
		mockValue field
		user_id   pkg.CtxInfo
		want      *user.User
		wantErr   bool
	}{
		{name: "正常系 - 管理者はユーザ一詳細情報を取得できる", args: args{id: "9999999999"}, want: user_1, mockValue: field{user: *user_1, userRepoErr: nil, admin: admin_1, adminRepoErr: nil}, user_id: pkg.CtxInfo{RequestId: "9999999999"}, wantErr: false},
		{name: "正常系 - 一般ユーザは自身の詳細情報を取得できる", args: args{id: "0000000000"}, want: user_1, mockValue: field{user: *user_1, userRepoErr: nil, admin: admin_2, adminRepoErr: nil}, user_id: pkg.CtxInfo{RequestId: "0000000000"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUDS := user.IUserDomainServiceMock{
				FindUserByIdFunc: func(ctx context.Context, id string) (*user.User, error) {
					return &tt.mockValue.user, tt.fields.userRepoErr
				},
			}
			mockADS := administrator.IAdminDomainServiceMock{
				FindUserFunc: func(ctx context.Context, id string) (*administrator.Administrator, error) {
					return tt.mockValue.admin, tt.mockValue.adminRepoErr
				},
			}

			uc := NewFindUserByIdUsecase(&mockUDS, &mockADS)
			ctx := context.Background()
			ctx = context.WithValue(ctx, "ctxInfo", tt.user_id)
			got, err := uc.Run(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByIdUsecase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("FindUserByIdUsecase.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
