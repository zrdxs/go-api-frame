package service

import (
	"github.com/MarceloZardoBR/go-api-frame/domain/interfaces"
)

type Service struct {
	db interfaces.Repos
}

func NewServices(db interfaces.Repos) *Service {
	return &Service{
		db: db,
	}
}
