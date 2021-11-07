package service_test

import (
	"errors"
	"go-unit-testing/core/domain"
	"go-unit-testing/core/service"
	"go-unit-testing/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	method = struct {
		userRepo struct {
			CreateUser string
		}
	}{
		userRepo: struct {
			CreateUser string
		}{
			"CreateUser",
		},
	}
	mocker    *mockFactory
	mockError = errors.New("error from mocker")
)

type mockFactory struct {
	userRepo *mocks.UserRepository
}

func TestMain(t *testing.M) {
	mocker = &mockFactory{
		userRepo: new(mocks.UserRepository),
	}
	exitCode := t.Run()
	os.Exit(exitCode)
}

func TestCreateUser(t *testing.T) {
	svc := service.New(mocker.userRepo)
	defaultRequest := domain.CreateUserRequest{
		FirstName: "hello",
		LastName:  "world",
	}
	// The naming Pattern is [MethodName_StateUnderTest_ExpectedBehavior]
	t.Run("CreateUser() with valid input should get succeed", func(t *testing.T) {
		createUserInput := createUserStruct(defaultRequest)
		mocker.userRepo.On(method.userRepo.CreateUser, createUserInput).Return(nil).Once()
		err := svc.CreateUser(defaultRequest)
		assert.NoError(t, err)
		mocker.userRepo.AssertExpectations(t)
	})

	t.Run("CreateUser() but got error from userRepo.CreateUser() should get an error", func(t *testing.T) {
		createUserInput := createUserStruct(defaultRequest)
		mocker.userRepo.On(method.userRepo.CreateUser, createUserInput).Return(mockError).Once()
		err := svc.CreateUser(defaultRequest)
		assert.Error(t, err)
		mocker.userRepo.AssertExpectations(t)
	})
}

func createUserStruct(request domain.CreateUserRequest) domain.CreateUser {
	return domain.CreateUser{
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}
}
