package bearden

import "github.com/gomdori-foo/bearden/internal/core/module"

type Module = module.ModuleFactory

func Builder() *module.ModuleFactoryBuilder {
	return module.Builder()
}
