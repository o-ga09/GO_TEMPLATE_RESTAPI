package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/o-ga09/api/internal/domain/administrator"
	"github.com/o-ga09/api/internal/domain/user"
)

var (
	user_1  = user.NewUser("999999999", "hoge@email.com", "ewerf2", "testuser1", "田中", "太郎", "male", "1999/01/01", "999-9999-9999", "999-9999", "東京都", "渋谷区", "道玄坂")
	user_2  = user.NewUser("000000000", "hoge@email.com", "@42312", "testuser2", "佐藤", "二郎", "famale", "2003/01/01", "999-9999-9999", "999-9999", "東京都", "渋谷区", "道玄坂")
	admin_1 = administrator.NewAdministrator("999999999", 1)
	admin_2 = administrator.NewAdministrator("000000000", 0)
)

func TestFindUserUsecase_Run(t *testing.T) {
	type field struct {
		user         []*user.User
		userRepoErr  error
		admin        *administrator.Administrator
		adminRepoErr error
	}
	tests := []struct {
		name      string
		want      []*user.User
		mockValue field
		user_id   string
		wantErr   bool
	}{
		{name: "正常系 - 管理者はユーザ一覧を取得できる", want: []*user.User{user_1}, mockValue: field{user: []*user.User{user_1}, userRepoErr: nil, admin: admin_1, adminRepoErr: nil}, user_id: "999999999", wantErr: false},
		{name: "異常系 - 一般ユーザ一覧を取得できない", want: []*user.User{}, mockValue: field{user: []*user.User{user_2}, userRepoErr: nil, admin: nil, adminRepoErr: INVALID_ADMIN}, user_id: "000000000", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUDS := user.IUserDomainServiceMock{
				FindUserFunc: func(ctx context.Context) ([]*user.User, int64, error) {
					return tt.mockValue.user, int64(len(tt.mockValue.user)), tt.mockValue.userRepoErr
				},
			}
			mockADS := administrator.IAdminDomainServiceMock{
				FindUserFunc: func(ctx context.Context, id string) (*administrator.Administrator, error) {
					return tt.mockValue.admin, tt.mockValue.adminRepoErr
				},
			}

			uc := NewFindUserUsecase(&mockUDS, &mockADS)
			ctx := context.Background()
			ctx = context.WithValue(ctx, "user_id", tt.user_id)

			got, err := uc.Run(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserUsecase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("FindUserUsecase.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
