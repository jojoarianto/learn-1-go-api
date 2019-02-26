package dto

import (
	"github.com/jojoarianto/learn-1-go-api/src/entities"
)

type ListUsersResponse struct {
	Users []entities.Users `json:list_users`
}

type UsersResponse struct {
	User entities.Users `json:user`
}
