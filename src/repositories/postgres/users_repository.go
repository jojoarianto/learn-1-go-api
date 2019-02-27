package postgres

import (
	"github.com/jinzhu/gorm"

	"github.com/jojoarianto/learn-1-go-api/src/entities"
	"github.com/jojoarianto/learn-1-go-api/src/repositories"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UsersRepository {
	return &usersRepository{db: db}
}

func (u *usersRepository) Get(id int) (entities.Users, error) {
	usr := entities.Users{}
	if err := u.db.First(&usr, id).Error; err != nil {
		return entities.Users{}, entities.ErrUserNotfound
	}
	return usr, nil
}

func (u *usersRepository) GetAll() ([]entities.Users, error) {
	usr := []entities.Users{}
	if err := u.db.Find(&usr).Error; err != nil {
		return nil, entities.ErrUserNotfound
	}
	return usr, nil
}

func (u *usersRepository) CreateUser(usr entities.Users) error {
	if err := u.db.Save(&usr).Error; err != nil {
		return err
	}

	return nil
}
