package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bhyago/crud-products-go/internal/dto"
	"github.com/bhyago/crud-products-go/internal/entity"
	"github.com/bhyago/crud-products-go/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandle struct {
	UserDB        database.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

type Error struct {
	Message string `json:"message"`
}

func NewUserHandle(db database.UserInterface, JwtExpiriesIn int) *UserHandle {
	return &UserHandle{
		UserDB: db,
	}
}

// GetJWT godoc
// @Summary Get JWT
// @Description Get JWT
// @Tags users
// @Accept  json
// @Produce  json
// @Param request body dto.GetJWTInput true "User credentials"
// @Success 200 {object} dto.GetJWrOutput
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /users/generate_token [post]
func (h *UserHandle) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiriesIn := r.Context().Value("jwtExpiresIn").(int)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !u.ValidadePassword(user.Password) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Duration(jwtExpiriesIn) * time.Second).Unix(),
	})

	acessToken := dto.GetJWrOutput{AccessToken: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acessToken)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param request body dto.CreateUserInput true "User request"
// @Success 201
// @Failure 500
// @Router /users [post]
func (h *UserHandle) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Password, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Save(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
