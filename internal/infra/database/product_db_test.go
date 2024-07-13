package database

import (
	"testing"

	"github.com/bhyago/crud-products-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)

	err = productDB.Save(product)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	err = productDB.Save(product)
	assert.Nil(t, err)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(products))
}

func TestFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	err = productDB.Save(product)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	err = productDB.Save(product)
	assert.Nil(t, err)

	product.Name = "Product 2"
	product.Price = 20
	err = productDB.Update(product)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	err = productDB.Save(product)
	assert.Nil(t, err)

	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NotNil(t, err)
	assert.Nil(t, productFound)
}

func TestFindAllProductsEmpty(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindByIDNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	productFound, err := productDB.FindByID("1")
	assert.NotNil(t, err)
	assert.Nil(t, productFound)
}

func TestUpdateProductNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	product, err := entity.NewProduct("Product 1", 10)
	if err != nil {
		t.Error(err)
	}
	err = productDB.Update(product)
	assert.NotNil(t, err)
}

func TestDeleteProductNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	err = productDB.Delete("1")
	assert.NotNil(t, err)
}

func TestFindAllProductsPageZeroLimitZeroSortEmpty(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(0, 0, "")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindAllProductsPageZeroLimitZeroSortAsc(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(0, 0, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindAllProductsPageZeroLimitZeroSortDesc(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(0, 0, "desc")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindAllProductsPageZeroLimitZeroSortInvalid(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(0, 0, "invalid")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindAllProductsPageZeroLimitSortAsc(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(0, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindAllProductsPageZeroLimitSortDesc(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(0, 10, "desc")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindAllProductsPageZeroLimitSortInvalid(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(0, 10, "invalid")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindAllProductsPageLimitZeroSortAsc(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 0, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindAllProductsPageLimitZeroSortDesc(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 0, "desc")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestFindAllProductsPageLimitZeroSortInvalid(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 0, "invalid")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}
