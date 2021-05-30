package service

import (
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

func (s *UserService) GetAll() []entity.User {

	return nil
}
