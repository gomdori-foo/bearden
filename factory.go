package bearden

import "github.com/gomdori-foo/bearden/internal/core/module"

func Create(constructor func() *module.ModuleFactory) *BearDenApplication {
	return &BearDenApplication{}
}
