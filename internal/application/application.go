package application

import (
	"github.com/gin-gonic/gin"
	"github.com/gomdori-foo/bearden/internal/core/module"
)

type BearDenApplication struct {
	router *gin.Engine
	module *module.Module
}

func NewBearDenApplication(router *gin.Engine, module *module.Module) *BearDenApplication {
	return &BearDenApplication{
		router: router,
		module: module,
	}
}

func (a *BearDenApplication) Listen(addr ...string) error {
	err := a.router.Run(addr...)
	if err != nil {
		return err
	}

	return nil
}

