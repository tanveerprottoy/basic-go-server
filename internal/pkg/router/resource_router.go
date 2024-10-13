package router

import (
	"github.com/go-chi/chi"
	"github.com/tanveerprottoy/basic-go-server/internal/pkg/constant"
	"github.com/tanveerprottoy/basic-go-server/internal/server/resource"
)

func RegisterUserRoutes(router *Router, version string, handler *resource.Handler) {
	router.Mux.Route(
		constant.ApiPattern+version+constant.ResourcesPattern,
		func(r chi.Router) {
			r.Get(constant.RootPattern, handler.ReadMany)
			r.Get(constant.RootPattern+"get-basic", handler.GetBasicData)
			r.Get(constant.RootPattern+"{id}", handler.ReadOne)
			r.Post(constant.RootPattern, handler.Create)
			r.Patch(constant.RootPattern+"{id}", handler.Update)
			r.Delete(constant.RootPattern+"{id}", handler.Delete)
		},
	)
}
