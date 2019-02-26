package usecases

import (
	"github.com/jojoarianto/learn-1-go-api/src/dto"
	"github.com/jojoarianto/learn-1-go-api/src/repositories"
)

type getUserById struct {
	usr repositories.UsersRepository
}

func NewGetUserById(usr repositories.UsersRepository) *getUserById {
	return &getUserById{usr: usr}
}

func (gubi getUserById) GetUserById(id int) (dto.UsersResponse, error) {
	data, _ := gubi.usr.Get(id)

	return dto.UsersResponse{
		User: data,
	}, nil
}
