package administrator

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

var (
	admin_1 = NewAdministrator("999999999",1)
	admin_2 = NewAdministrator("000000000",0)
)


func TestAdminDomainService_FindUser(t *testing.T) {
	type field struct {
		adminRepo Administrator
		adminRepoErr error
	}
	type args struct {
		id  string
	}
	tests := []struct {
		name    string
		mockValue  field
		args    args
		want    *Administrator
		wantErr bool
	}{
		{name: "正常系",mockValue: field{adminRepo: *admin_1, adminRepoErr: nil},want: admin_1,wantErr: false},
		{name: "異常系",mockValue: field{adminRepo: Administrator{}, adminRepoErr: errors.New("not found")},want: nil,wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAdminRepo := AdminServiceRepositoryMock{
				FindUserFunc: func(ctx context.Context, id string) (*Administrator, error) {
					return &tt.mockValue.adminRepo, tt.mockValue.adminRepoErr
				},
			}
			mockADS := NewAdminDomainService(&mockAdminRepo)
			got, err := mockADS.FindUser(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminDomainService.FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("AdminDomainService.FindUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminDomainService_EditUser(t *testing.T) {
	type field struct {
		adminRepoErr error
	}
	type args struct {
		param *Administrator
	}
	tests := []struct {
		name    string
		mockValue  field
		args    args
		wantErr bool
	}{
		{name: "正常系",mockValue: field{adminRepoErr: nil},args: args{param: admin_1},wantErr: false},
		{name: "異常系",mockValue: field{adminRepoErr: errors.New("not found")},args: args{param: admin_1},wantErr: true},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAdminRepo := AdminServiceRepositoryMock{
				SaveFunc: func(ctx context.Context, param *Administrator) error {
					return tt.mockValue.adminRepoErr
				},
			}
			mockADS := NewAdminDomainService(&mockAdminRepo)
			if err := mockADS.EditUser(context.Background(), tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("AdminDomainService.EditUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdminDomainService_DeleteUser(t *testing.T) {
	type field struct {
		adminRepo Administrator
		adminRepoErr error
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
		{name: "正常系",mockValue: field{adminRepo: *admin_1, adminRepoErr: nil},args: args{id: "999999999"},wantErr: false},
		{name: "異常系",mockValue: field{adminRepo: Administrator{}, adminRepoErr: errors.New("not found")},args: args{id: "000000000"},wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAdminRepo := AdminServiceRepositoryMock{
				DeleteFunc: func(ctx context.Context, id string) error {
					return tt.mockValue.adminRepoErr
				},
			}
			mockADS := NewAdminDomainService(&mockAdminRepo)
			if err := mockADS.DeleteUser(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AdminDomainService.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
