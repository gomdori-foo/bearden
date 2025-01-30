package module

import (
	"github.com/gomdori-foo/bearden/internal/core/controller"
	"github.com/gomdori-foo/bearden/internal/core/provider"
)

type Module struct {
	controllers []*controller.Controller
	providers []*provider.Provider
}

type ModuleFactory struct {
	imports []func() *ModuleFactory
	providers []*provider.ProviderFactory
	controllers []*controller.ControllerFactory
	exports []interface{}
}

type ModuleFactoryBuilder struct {
	imports []func() *ModuleFactory
	providers []*provider.ProviderFactory
	controllers []*controller.ControllerFactory
	exports []interface{}
}

func Builder() *ModuleFactoryBuilder {
	return &ModuleFactoryBuilder{}
}

func (m *ModuleFactoryBuilder) Imports(constructors ...func() *ModuleFactory) *ModuleFactoryBuilder {
	m.imports = append(m.imports, constructors...)
	return m
}

func (m *ModuleFactoryBuilder) Controllers(constructors ...interface{}) *ModuleFactoryBuilder {
	controllerFactories := controller.Controllers(constructors...)
	m.controllers = append(m.controllers, controllerFactories...)
	return m
}

func (m *ModuleFactoryBuilder) Providers(constructors ...interface{}) *ModuleFactoryBuilder {
	providerFactories := provider.Providers(constructors...)
	m.providers = append(m.providers, providerFactories...)
	return m
}

func (m *ModuleFactoryBuilder) Exports(exports ...interface{}) *ModuleFactoryBuilder {
	m.exports = append(m.exports, exports...)
	return m
}

func (m *ModuleFactoryBuilder) Build() *ModuleFactory {
	return &ModuleFactory{
		imports: m.imports,
		providers: m.providers,
		controllers: m.controllers,
		exports: m.exports,
	}
}
	