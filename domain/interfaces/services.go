package interfaces

import (
	"github.com/MarceloZardoBR/go-api-frame/domain/entity"
)

type UserService interface {
	GetAll() ([]entity.User, error)
}
