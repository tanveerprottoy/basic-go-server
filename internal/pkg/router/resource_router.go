package router

import (
	"github.com/go-chi/chi"
	"github.com/tanveerprottoy/basic-go-server/internal/app/basicserver/module/resource"
	"github.com/tanveerprottoy/basic-go-server/internal/pkg/constant"
)

func RegisterUserRoutes(router *Router, version string, module *resource.Module) {
	router.Mux.Route(
		constant.ApiPattern+version+constant.ResourcesPattern,
		func(r chi.Router) {
			r.Get(constant.RootPattern, module.Handler.ReadMany)
			r.Get(constant.RootPattern+"get-basic", module.Handler.GetBasicData)
			r.Get(constant.RootPattern+"{id}", module.Handler.ReadOne)
			r.Post(constant.RootPattern, module.Handler.Create)
			r.Patch(constant.RootPattern+"{id}", module.Handler.Update)
			r.Delete(constant.RootPattern+"{id}", module.Handler.Delete)
		},
	)
}
