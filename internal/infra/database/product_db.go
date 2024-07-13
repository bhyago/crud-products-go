package database

import (
	"github.com/bhyago/crud-products-go/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort == "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order(sort).Find(&products).Error
	} else {
		err = p.DB.Find(&products).Order("created_at " + sort).Error
	}
	return products, err
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) Save(product *entity.Product) error {
	if err := p.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	if err := p.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *Product) Delete(id string) error {
	_, err := p.FindByID(id)
	if err != nil {
		return err
	}

	if err := p.DB.Where("id = ?", id).Delete(&entity.Product{}).Error; err != nil {
		return err
	}
	return nil
}
