package resourcecfg

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/basic-go-server/internal/server/resource"
)

// Config holds the components of the current package
type Config struct {
	Handler    *resource.Handler
	Service    *resource.Service
	Repository *resource.Repository
}

// NewConfig initializes a new NewConfig
func NewConfig(db *sqlx.DB, validate *validator.Validate) *Config {
	c := new(Config)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	c.Repository = resource.NewRepository(db)
	c.Service = resource.NewService(c.Repository)
	c.Handler = resource.NewHandler(c.Service, validate)
	return c
}
