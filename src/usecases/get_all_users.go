package usecases

import (
	"github.com/jojoarianto/learn-1-go-api/src/entities"
	"github.com/jojoarianto/learn-1-go-api/src/repositories"
)

type getAllUsers struct {
	usr repositories.UsersRepository
}

func NewGetAllUsers(usr repositories.UsersRepository) *getAllUsers {
	return &getAllUsers{usr: usr}
}

func (gau getAllUsers) GetAllUsers() ([]entities.Users, error) {
	data, _ := gau.usr.GetAll()

	return data, nil
}
