package repositories

import (
	"github.com/jojoarianto/learn-1-go-api/src/entities"
)

type UsersRepository interface {
	Get(id int) (entities.Users, error)
	GetAll() ([]entities.Users, error)
}
