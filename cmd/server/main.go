package main

import (
	"log"
	"net/http"

	"github.com/bhyago/crud-products-go/configs"
	_ "github.com/bhyago/crud-products-go/docs"
	"github.com/bhyago/crud-products-go/internal/entity"
	"github.com/bhyago/crud-products-go/internal/infra/database"
	"github.com/bhyago/crud-products-go/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title CRUD Products API
// @version 1
// @description This is a CRUD Products API
// @host localhost:3333
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io
// contact.email test@test

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host localhost:3333
// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

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
	userHandle := handlers.NewUserHandle(userDB, configs.JWTExpiresIn)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// router.Use(LogRequest)
	router.Use(middleware.WithValue("jwt", configs.TokenAuthKey))
	router.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))

	router.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuthKey))
		r.Use(jwtauth.Authenticator)

		r.Post("/", ProductHandle.CreateProduct)
		r.Get("/{id}", ProductHandle.GetProduct)
		r.Get("/", ProductHandle.GetProducts)
		r.Put("/{id}", ProductHandle.UpdateProduct)
		r.Delete("/{id}", ProductHandle.DeleteProduct)
	})

	router.Post("/users", userHandle.CreateUser)
	router.Post("/users/generate_token", userHandle.GetJWT)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:3333/docs/doc.json")))
	http.ListenAndServe(":3333", router)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
