package factory

import (
	application "github.com/gomdori-foo/bearden"
	"github.com/gomdori-foo/bearden/internal/core/module"
)

func Create(constructor func() *module.ModuleFactory) *application.BearDenApplication {
	return &application.BearDenApplication{}
}