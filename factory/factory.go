package factory

import (
	application "github.com/gomdori-foo/bear-den"
	"github.com/gomdori-foo/bear-den/internal/core/module"
)

func Create(constructor func() *module.ModuleFactory) *application.BearDenApplication {
	return &application.BearDenApplication{}
}