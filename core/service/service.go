package service

import (
	"go-unit-testing/core/domain"
	"go-unit-testing/core/port"
)

type Service struct {
	userRepo port.UserRepository
}

func New(userRepo port.UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (svc *Service) CreateUser(request domain.CreateUserRequest) error {
	err := svc.userRepo.CreateUser(domain.CreateUser{
		FirstName: request.FirstName,
		LastName:  request.LastName,
	})
	if err != nil {
		return err
	}
	return nil
}
