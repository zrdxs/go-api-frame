package service

import (
	"github.com/MarceloZardoBR/go-api-frame/domain/interfaces"
	"github.com/MarceloZardoBR/go-api-frame/infra/config"
)

type Service struct {
	db     interfaces.Repos
	config config.Config
}

func NewServices(db interfaces.Repos, config config.Config) *Service {
	return &Service{
		db:     db,
		config: config,
	}
}
