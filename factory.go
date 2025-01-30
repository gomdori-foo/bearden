package bearden

import (
	"github.com/gomdori-foo/bearden/internal/application"
	"github.com/gomdori-foo/bearden/internal/core/module"
	"github.com/gomdori-foo/bearden/internal/factory"
)


func Create(constructor func() *module.ModuleFactory) *application.BearDenApplication {
	return factory.Create(constructor)
}
