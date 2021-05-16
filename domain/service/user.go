package service

import (
	"github.com/MarceloZardoBR/go-api-frame/domain/entity"
	"github.com/MarceloZardoBR/go-api-frame/domain/interfaces"
)

type UserService struct {
}

func NewService() interfaces.UserService {
	return &UserService{}
}

func (s *UserService) GetAll() []entity.User {

	return nil
}
