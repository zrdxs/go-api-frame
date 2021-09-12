package services

import (
	"github.com/MarceloZardoBR/go-api-frame/domain/interfaces"
	"github.com/MarceloZardoBR/go-api-frame/infra/config"
)

type Service struct {
	db  interfaces.Repos
	cfg *config.Config
}

func NewServices(db interfaces.Repos, cfg *config.Config) *Service {
	return &Service{
		db:  db,
		cfg: cfg,
	}
}
