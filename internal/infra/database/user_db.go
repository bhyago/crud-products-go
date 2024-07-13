package database

import (
	"log"

	"github.com/bhyago/crud-products-go/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Save(user *entity.User) error {
	log.Println(user)
	err := u.DB.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}
