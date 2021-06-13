package service

import (
	"encoding/json"
	"fmt"

	"github.com/MarceloZardoBR/go-api-frame/domain/entity"
	"github.com/MarceloZardoBR/go-api-frame/domain/interfaces"
)

type UserService struct {
	svc *Service
}

func NewUserService(svc *Service) interfaces.UserService {
	return &UserService{
		svc: svc,
	}
}

func (s *UserService) GetAll() ([]entity.User, error) {

	user, err := s.svc.db.UserRepo().GetAll()
	if err != nil {
		return user, err
	}

	json2, _ := json.MarshalIndent(user, "", " ")
	fmt.Println(string(json2))

	return user, nil
}
