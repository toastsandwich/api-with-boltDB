package service

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/toastsandwich/restraunt-api-system/model"
	"github.com/toastsandwich/restraunt-api-system/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(db *bolt.DB) (*UserService, error) {
	repo, err := repository.NewUserRepository(db, "food")
	if err != nil {
		return nil, err
	}

	return &UserService{
		UserRepository: repo,
	}, nil
}

func (s *UserService) CreateUserService(email string, u model.User) error {

	buf, err := json.Marshal(u)
	if err != nil {
		return err
	}

	err = s.UserRepository.Set([]byte(email), buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserService(email string) (*model.User, error) {
	var u *model.User = &model.User{}
	data, err := s.UserRepository.Get([]byte(email))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, u)
	if err != nil {
		return nil, fmt.Errorf("error in marshalling data: %s", err.Error())
	}
	return u, nil
}

func (s *UserService) GetAllUserService() (map[string]model.User, error) {
	m := make(map[string]model.User)
	data, err := s.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	for k, v := range data {
		u := model.User{}
		err := json.Unmarshal(v, &u)
		if err != nil {
			return nil, err
		}
		m[k] = u
	}
	return m, nil
}
