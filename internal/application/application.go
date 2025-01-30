package application

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gomdori-foo/bearden/internal/core/module"
)

type BearDenApplication struct {
	router *gin.Engine
	module *module.Module
}

func NewBearDenApplication(router *gin.Engine, module *module.Module) *BearDenApplication {
	app := &BearDenApplication{
		router: router,
		module: module,
	}

	app.resolveRoutes(module)

	return app;
}

func (a *BearDenApplication) Listen(addr ...string) error {
	return a.router.Run(addr...)
}

func (a *BearDenApplication) resolveRoutes(module *module.Module) {
	for _, controller := range module.Controllers() {
		prefix := controller.Prefix()
		instance := controller.Instance()
		group := a.router.Group(prefix)

		for _, route := range controller.Routes() {
			method := strings.ToUpper(route.Method())
			path := route.Path()
			handler := reflect.ValueOf(instance).MethodByName(route.HandlerName())
			handlerFunc := func(c *gin.Context) {
				handler.Call([]reflect.Value{reflect.ValueOf(c)})
			}
			group.Handle(method, path, handlerFunc)
		}
	}

	for _, module := range module.Imports() {
		a.resolveRoutes(module)
	}
}