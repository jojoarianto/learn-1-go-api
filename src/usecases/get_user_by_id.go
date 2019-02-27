package usecases

import (
	"github.com/jojoarianto/learn-1-go-api/src/entities"
	"github.com/jojoarianto/learn-1-go-api/src/repositories"
)

type getUserById struct {
	usr repositories.UsersRepository
}

func NewGetUserById(usr repositories.UsersRepository) *getUserById {
	return &getUserById{usr: usr}
}

func (gubi getUserById) GetUserById(id int) (entities.Users, error) {
	data, err := gubi.usr.Get(id)
	if err != nil {
		return entities.Users{}, err
	}

	return data, nil
}
