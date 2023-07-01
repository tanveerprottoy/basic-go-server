package resource

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/basic-go-server/internal/app/basicserver/module/resource/entity"
	"github.com/tanveerprottoy/basic-go-server/pkg/data/sqlxpkg"
)

type Module struct {
	Handler    *Handler
	Service    *Service
	Repository sqlxpkg.Repository[entity.Resource]
}

func NewModule(db *sqlx.DB, validate *validator.Validate) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = NewRepository(db)
	m.Service = NewService(m.Repository)
	m.Handler = NewHandler(m.Service, validate)
	return m
}
