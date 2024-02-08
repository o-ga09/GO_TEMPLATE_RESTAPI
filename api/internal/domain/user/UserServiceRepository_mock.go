// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package user

import (
	"context"
	"sync"
)

// Ensure, that UserServiceRepositoryMock does implement UserServiceRepository.
// If this is not the case, regenerate this file with moq.
var _ UserServiceRepository = &UserServiceRepositoryMock{}

// UserServiceRepositoryMock is a mock implementation of UserServiceRepository.
//
//	func TestSomethingThatUsesUserServiceRepository(t *testing.T) {
//
//		// make and configure a mocked UserServiceRepository
//		mockedUserServiceRepository := &UserServiceRepositoryMock{
//			DeleteFunc: func(ctx context.Context, id string) error {
//				panic("mock out the Delete method")
//			},
//			FindUserFunc: func(ctx context.Context) ([]*User, int64, error) {
//				panic("mock out the FindUser method")
//			},
//			FindUserByIdFunc: func(ctx context.Context, id string) (*User, error) {
//				panic("mock out the FindUserById method")
//			},
//			SaveFunc: func(ctx context.Context, param *User) error {
//				panic("mock out the Save method")
//			},
//		}
//
//		// use mockedUserServiceRepository in code that requires UserServiceRepository
//		// and then make assertions.
//
//	}
type UserServiceRepositoryMock struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, id string) error

	// FindUserFunc mocks the FindUser method.
	FindUserFunc func(ctx context.Context) ([]*User, int64, error)

	// FindUserByIdFunc mocks the FindUserById method.
	FindUserByIdFunc func(ctx context.Context, id string) (*User, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(ctx context.Context, param *User) error

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// FindUser holds details about calls to the FindUser method.
		FindUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// FindUserById holds details about calls to the FindUserById method.
		FindUserById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Param is the param argument value.
			Param *User
		}
	}
	lockDelete       sync.RWMutex
	lockFindUser     sync.RWMutex
	lockFindUserById sync.RWMutex
	lockSave         sync.RWMutex
}

// Delete calls DeleteFunc.
func (mock *UserServiceRepositoryMock) Delete(ctx context.Context, id string) error {
	if mock.DeleteFunc == nil {
		panic("UserServiceRepositoryMock.DeleteFunc: method is nil but UserServiceRepository.Delete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedUserServiceRepository.DeleteCalls())
func (mock *UserServiceRepositoryMock) DeleteCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// FindUser calls FindUserFunc.
func (mock *UserServiceRepositoryMock) FindUser(ctx context.Context) ([]*User, int64, error) {
	if mock.FindUserFunc == nil {
		panic("UserServiceRepositoryMock.FindUserFunc: method is nil but UserServiceRepository.FindUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockFindUser.Lock()
	mock.calls.FindUser = append(mock.calls.FindUser, callInfo)
	mock.lockFindUser.Unlock()
	return mock.FindUserFunc(ctx)
}

// FindUserCalls gets all the calls that were made to FindUser.
// Check the length with:
//
//	len(mockedUserServiceRepository.FindUserCalls())
func (mock *UserServiceRepositoryMock) FindUserCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockFindUser.RLock()
	calls = mock.calls.FindUser
	mock.lockFindUser.RUnlock()
	return calls
}

// FindUserById calls FindUserByIdFunc.
func (mock *UserServiceRepositoryMock) FindUserById(ctx context.Context, id string) (*User, error) {
	if mock.FindUserByIdFunc == nil {
		panic("UserServiceRepositoryMock.FindUserByIdFunc: method is nil but UserServiceRepository.FindUserById was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFindUserById.Lock()
	mock.calls.FindUserById = append(mock.calls.FindUserById, callInfo)
	mock.lockFindUserById.Unlock()
	return mock.FindUserByIdFunc(ctx, id)
}

// FindUserByIdCalls gets all the calls that were made to FindUserById.
// Check the length with:
//
//	len(mockedUserServiceRepository.FindUserByIdCalls())
func (mock *UserServiceRepositoryMock) FindUserByIdCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockFindUserById.RLock()
	calls = mock.calls.FindUserById
	mock.lockFindUserById.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *UserServiceRepositoryMock) Save(ctx context.Context, param *User) error {
	if mock.SaveFunc == nil {
		panic("UserServiceRepositoryMock.SaveFunc: method is nil but UserServiceRepository.Save was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Param *User
	}{
		Ctx:   ctx,
		Param: param,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(ctx, param)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//
//	len(mockedUserServiceRepository.SaveCalls())
func (mock *UserServiceRepositoryMock) SaveCalls() []struct {
	Ctx   context.Context
	Param *User
} {
	var calls []struct {
		Ctx   context.Context
		Param *User
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}