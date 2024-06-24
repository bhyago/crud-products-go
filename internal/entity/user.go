package entity

import (
	"github.com/bhyago/crud-products-go/internal/entity"
	"github.com/bhyago/crud-products-go/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
}

func NewUser(id int, name, password, email string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(id),
		Name:     name,
		Password: string(hash),
		Email:    email,
	}, nil
}

func (u *User) ValidadePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
