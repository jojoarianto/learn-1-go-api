package usecases

import (
	"github.com/jojoarianto/learn-1-go-api/src/dto"
	"github.com/jojoarianto/learn-1-go-api/src/entities"
	"github.com/jojoarianto/learn-1-go-api/src/repositories"
)

type createUser struct {
	usr repositories.UsersRepository
}

func NewCreateUser(usr repositories.UsersRepository) *createUser {
	return &createUser{usr: usr}
}

func (cu createUser) CreateUser(userDto dto.CreateUserRequest) error {
	user := entities.Users{
		Email: userDto.Email,
	}

	err := cu.usr.CreateUser(user)
	if err != nil {
		return err
	}

	return err
}
