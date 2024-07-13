package main

import (
	"net/http"

	"github.com/bhyago/crud-products-go/configs"
	"github.com/bhyago/crud-products-go/internal/entity"
	"github.com/bhyago/crud-products-go/internal/infra/database"
	"github.com/bhyago/crud-products-go/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	ProductHandle := handlers.NewProductHandle(productDB)

	userDB := database.NewUser(db)
	userHandle := handlers.NewUserHandle(userDB, configs.TokenAuthKey, configs.JWTExpiresIn)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/products", ProductHandle.CreateProduct)
	router.Get("/products/{id}", ProductHandle.GetProduct)
	router.Get("/products", ProductHandle.GetProducts)
	router.Put("/products/{id}", ProductHandle.UpdateProduct)
	router.Delete("/products/{id}", ProductHandle.DeleteProduct)

	router.Post("/users", userHandle.CreateUser)
	router.Post("/users/generate_token", userHandle.GetJWT)

	http.ListenAndServe(":3333", router)
}
