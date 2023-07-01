package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	middlewarepkg "github.com/tanveerprottoy/basic-go-server/pkg/middleware"
)

// Router struct
type Router struct {
	Mux *chi.Mux
}

func NewRouter() *Router {
	r := &Router{}
	r.Mux = chi.NewRouter()
	r.registerGlobalMiddlewares()
	return r
}

func (r *Router) registerGlobalMiddlewares() {
	r.Mux.Use(
		// built in provided by chi
		// base middleware stack
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,

		// custom global middlewares
		middlewarepkg.JSONContentTypeMiddleWare,
		middlewarepkg.CORSEnableMiddleWare,
	)
}
