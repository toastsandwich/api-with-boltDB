package handler

import (
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/toastsandwich/restraunt-api-system/model"
	"github.com/toastsandwich/restraunt-api-system/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(db *bolt.DB) (*UserHandler, error) {
	service, err := service.NewUserService(db)
	if err != nil {
		return nil, err
	}
	return &UserHandler{
		UserService: service,
	}, nil
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	email := r.URL.Query().Get("email")
	dob := r.URL.Query().Get("dob")
	password := r.URL.Query().Get("password")

	u := model.User{
		FirstName: firstName,
		LastName:  lastName,
		DOB:       dob,
		Password:  password,
	}

	err := h.UserService.CreateUserService(email, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	w.Header().Set("Content-Type", "application/json")
	u, err := h.UserService.GetUserService(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

