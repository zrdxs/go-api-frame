package interfaces

import (
	"github.com/MarceloZardoBR/go-api-frame/domain/entity"
)

type UserRepo interface {
	GetAll() []entity.User
}
