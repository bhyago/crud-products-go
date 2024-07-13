package database

import "github.com/bhyago/crud-products-go/internal/entity"

type UserInterface interface {
	FindByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
}

type ProductInterface interface {
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Save(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id string) error
}
