package port

import "go-unit-testing/core/domain"

type UserRepository interface {
	CreateUser(user domain.CreateUser) error
}
