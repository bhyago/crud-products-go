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

func NewUserHandle(db database.UserInterface, jwt *jwtauth.JWTAuth, JwtExpiriesIn int) *UserHandle {
	return &UserHandle{
		UserDB:        db,
		Jwt:           jwt,
		JwtExpiriesIn: JwtExpiriesIn,
	}
}

func (h *UserHandle) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !u.ValidadePassword(user.Password) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	_, token, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Duration(h.JwtExpiriesIn) * time.Second).Unix(),
	})

	acessToken :=
		struct {
			AcessToken string `json:"access_token"`
		}{
			AcessToken: token,
		}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acessToken)
}

func (h *UserHandle) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Password, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UserDB.Save(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
