package user

import (
	"context"
	"reflect"
	"testing"
)

var (
	user_1 = NewUser("999999999","hoge@email.com","ewerf2","testuser1","田中","太郎","male","1999/01/01","999-9999-9999","999-9999","東京都","渋谷区","道玄坂")
	user_2 = NewUser("000000000","hoge@email.com","@42312","testuser2","佐藤","二郎","famale","2003/01/01","999-9999-9999","999-9999","東京都","渋谷区","道玄坂")
)

func TestUserDomainService_FindUser(t *testing.T) {
	type field struct {
		userRepo []*User
		userRepoErr error
	}
	tests := []struct {
		name    string
		mockValue  field
		want    []*User
		wantErr bool
	}{
		{name: "正常系",mockValue: field{userRepo: []*User{user_1}, userRepoErr: nil},want: []*User{user_1},wantErr: false},
		{name: "異常系",mockValue: field{userRepo: []*User{}, userRepoErr: INVALID_BIRTH_DAY},want: []*User{},wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := UserServiceRepositoryMock{
				FindUserFunc: func(ctx context.Context) ([]*User, error) {
					return tt.mockValue.userRepo, tt.mockValue.userRepoErr
				},
			}
			mockUDS := NewUserDomainService(&mockUserRepo)
			
			got, err := mockUDS.FindUser(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("UserDomainService.FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("UserDomainService.FindUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDomainService_FindUserById(t *testing.T) {
	type field struct {
		userRepo *User
		userRepoErr error
	}
	type args struct {
		id  string
	}
	tests := []struct {
		name    string
		mockValue  field
		args    args
		want    *User
		wantErr bool
	}{
		{name: "正常系",mockValue: field{userRepo: user_1, userRepoErr: nil},args: args{id: "999999999"},want: user_1,wantErr: false},
		{name: "異常系",mockValue: field{userRepo: nil, userRepoErr: INVALID_BIRTH_DAY}, args: args{id: "000000000"},want: nil,wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := UserServiceRepositoryMock{
				FindUserByIdFunc: func(ctx context.Context, id string) (*User, error) {
					return tt.mockValue.userRepo, tt.mockValue.userRepoErr
				},
			}
			mockUDS := NewUserDomainService(&mockUserRepo)
			got, err := mockUDS.FindUserById(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserDomainService.FindUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("UserDomainService.FindUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDomainService_EditUser(t *testing.T) {
	type field struct {
		userRepoErr error
	}
	type args struct {
		param *User
	}
	tests := []struct {
		name    string
		mockValue  field
		args    args
		wantErr bool
	}{
		{name: "正常系",mockValue: field{ userRepoErr: nil},args: args{param: user_1},wantErr: false},
		{name: "異常系",mockValue: field{ userRepoErr: INVALID_BIRTH_DAY},args: args{param: user_1},wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := UserServiceRepositoryMock{
				SaveFunc: func(ctx context.Context, param *User) error {
					return tt.mockValue.userRepoErr
				},
			}
			mockUDS := NewUserDomainService(&mockUserRepo)
			if err := mockUDS.EditUser(context.Background(), tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("UserDomainService.EditUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDomainService_DeleteUser(t *testing.T) {
	type field struct {
		userRepoErr error
	}
	type args struct {
		id  string
	}
	tests := []struct {
		name    string
		mockValue  field
		args    args
		wantErr bool
	}{
		{name: "正常系",mockValue: field{ userRepoErr: nil},args: args{id: "999999999"},wantErr: false},
		{name: "異常系",mockValue: field{ userRepoErr: INVALID_BIRTH_DAY},args: args{id: "000000000"},wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := UserServiceRepositoryMock{
				DeleteFunc: func(ctx context.Context, id string) error {
					return tt.mockValue.userRepoErr
				},
			}
			mockUDS := NewUserDomainService(&mockUserRepo)
			if err := mockUDS.DeleteUser(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserDomainService.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
