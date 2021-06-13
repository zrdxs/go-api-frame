package interfaces

import (
	"github.com/MarceloZardoBR/go-api-frame/domain/entity"
)

type Repos interface {
	UserRepo() UserRepo
}

type UserRepo interface {
	GetAll() ([]entity.User, error)
}
