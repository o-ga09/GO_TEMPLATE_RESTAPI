// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package administrator

import (
	"context"
	"sync"
)

// Ensure, that IAdminDomainServiceMock does implement IAdminDomainService.
// If this is not the case, regenerate this file with moq.
var _ IAdminDomainService = &IAdminDomainServiceMock{}

// IAdminDomainServiceMock is a mock implementation of IAdminDomainService.
//
//	func TestSomethingThatUsesIAdminDomainService(t *testing.T) {
//
//		// make and configure a mocked IAdminDomainService
//		mockedIAdminDomainService := &IAdminDomainServiceMock{
//			DeleteUserFunc: func(ctx context.Context, id string) error {
//				panic("mock out the DeleteUser method")
//			},
//			EditUserFunc: func(ctx context.Context, param *Administrator) error {
//				panic("mock out the EditUser method")
//			},
//			FindUserFunc: func(ctx context.Context, id string) (*Administrator, error) {
//				panic("mock out the FindUser method")
//			},
//		}
//
//		// use mockedIAdminDomainService in code that requires IAdminDomainService
//		// and then make assertions.
//
//	}
type IAdminDomainServiceMock struct {
	// DeleteUserFunc mocks the DeleteUser method.
	DeleteUserFunc func(ctx context.Context, id string) error

	// EditUserFunc mocks the EditUser method.
	EditUserFunc func(ctx context.Context, param *Administrator) error

	// FindUserFunc mocks the FindUser method.
	FindUserFunc func(ctx context.Context, id string) (*Administrator, error)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteUser holds details about calls to the DeleteUser method.
		DeleteUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// EditUser holds details about calls to the EditUser method.
		EditUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Param is the param argument value.
			Param *Administrator
		}
		// FindUser holds details about calls to the FindUser method.
		FindUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
	}
	lockDeleteUser sync.RWMutex
	lockEditUser   sync.RWMutex
	lockFindUser   sync.RWMutex
}

// DeleteUser calls DeleteUserFunc.
func (mock *IAdminDomainServiceMock) DeleteUser(ctx context.Context, id string) error {
	if mock.DeleteUserFunc == nil {
		panic("IAdminDomainServiceMock.DeleteUserFunc: method is nil but IAdminDomainService.DeleteUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteUser.Lock()
	mock.calls.DeleteUser = append(mock.calls.DeleteUser, callInfo)
	mock.lockDeleteUser.Unlock()
	return mock.DeleteUserFunc(ctx, id)
}

// DeleteUserCalls gets all the calls that were made to DeleteUser.
// Check the length with:
//
//	len(mockedIAdminDomainService.DeleteUserCalls())
func (mock *IAdminDomainServiceMock) DeleteUserCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDeleteUser.RLock()
	calls = mock.calls.DeleteUser
	mock.lockDeleteUser.RUnlock()
	return calls
}

// EditUser calls EditUserFunc.
func (mock *IAdminDomainServiceMock) EditUser(ctx context.Context, param *Administrator) error {
	if mock.EditUserFunc == nil {
		panic("IAdminDomainServiceMock.EditUserFunc: method is nil but IAdminDomainService.EditUser was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Param *Administrator
	}{
		Ctx:   ctx,
		Param: param,
	}
	mock.lockEditUser.Lock()
	mock.calls.EditUser = append(mock.calls.EditUser, callInfo)
	mock.lockEditUser.Unlock()
	return mock.EditUserFunc(ctx, param)
}

// EditUserCalls gets all the calls that were made to EditUser.
// Check the length with:
//
//	len(mockedIAdminDomainService.EditUserCalls())
func (mock *IAdminDomainServiceMock) EditUserCalls() []struct {
	Ctx   context.Context
	Param *Administrator
} {
	var calls []struct {
		Ctx   context.Context
		Param *Administrator
	}
	mock.lockEditUser.RLock()
	calls = mock.calls.EditUser
	mock.lockEditUser.RUnlock()
	return calls
}

// FindUser calls FindUserFunc.
func (mock *IAdminDomainServiceMock) FindUser(ctx context.Context, id string) (*Administrator, error) {
	if mock.FindUserFunc == nil {
		panic("IAdminDomainServiceMock.FindUserFunc: method is nil but IAdminDomainService.FindUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFindUser.Lock()
	mock.calls.FindUser = append(mock.calls.FindUser, callInfo)
	mock.lockFindUser.Unlock()
	return mock.FindUserFunc(ctx, id)
}

// FindUserCalls gets all the calls that were made to FindUser.
// Check the length with:
//
//	len(mockedIAdminDomainService.FindUserCalls())
func (mock *IAdminDomainServiceMock) FindUserCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockFindUser.RLock()
	calls = mock.calls.FindUser
	mock.lockFindUser.RUnlock()
	return calls
}