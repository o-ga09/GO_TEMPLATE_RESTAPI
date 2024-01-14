package usecase

import (
	"context"
	"testing"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
)

func TestSaveUserUsecase_Run(t *testing.T) {
	type field struct {
		userRepoErr error
		admin *administrator.Administrator
		adminRepoErr error
	}
	type args struct {
		param *user.User
	}
	tests := []struct {
		name    string
		args    args
		mockValue field
		user_id string
		wantErr bool
	}{
		{name: "正常系 - 管理者はユーザを登録できる",args: args{param: user_1},mockValue: field{userRepoErr: nil,admin: admin_1,adminRepoErr: nil},user_id: "999999999",wantErr: false},
		{name: "正常系 - 一般ユーザは、自身の情報を登録できる",args: args{param: user_2},mockValue: field{userRepoErr: nil,admin: admin_2,adminRepoErr: nil},user_id: "000000000",wantErr: false},
		{name: "異常系 - 一般ユーザは、他の情報を登録できない",args: args{param: user_1},mockValue: field{userRepoErr: nil,admin: admin_2,adminRepoErr: INVALID_ADMIN},user_id: "000000000",wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUDS := user.IUserDomainServiceMock{
				EditUserFunc: func(ctx context.Context, param *user.User) error {
					return tt.mockValue.userRepoErr
				},
			}
			mockADS := administrator.IAdminDomainServiceMock{
				FindUserFunc: func(ctx context.Context, id string) (*administrator.Administrator, error) {
					return tt.mockValue.admin, tt.mockValue.adminRepoErr
				},
			}
		
			uc := NewSaveUserUsecase(&mockUDS,&mockADS)
			ctx := context.Background()
			ctx = context.WithValue(ctx,"user_id",tt.user_id)
			if err := uc.Run(ctx, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("SaveUserUsecase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
