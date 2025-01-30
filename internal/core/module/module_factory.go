package module

import (
	"github.com/gomdori-foo/bearden/internal/core/controller"
	"github.com/gomdori-foo/bearden/internal/core/provider"
)

type Module struct {
	imports []*Module
	controllers []*controller.Controller
	providers []*provider.Provider
	exports []*provider.Provider
}

type ModuleFactoryConstructor func() *ModuleFactory

type ModuleFactory struct {
	imports []ModuleFactoryConstructor
	providers []*provider.ProviderFactory
	controllers []*controller.ControllerFactory
	exports []interface{}
}

type ModuleFactoryBuilder struct {
	imports []ModuleFactoryConstructor
	providers []*provider.ProviderFactory
	controllers []*controller.ControllerFactory
	exports []interface{}
}

func Builder() *ModuleFactoryBuilder {
	return &ModuleFactoryBuilder{}
}

func (m *ModuleFactoryBuilder) Imports(constructors ...ModuleFactoryConstructor) *ModuleFactoryBuilder {
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

// Module Factory
func (m *ModuleFactory) Imports() []ModuleFactoryConstructor {
	return m.imports
}

func (m *ModuleFactory) Providers() []*provider.ProviderFactory {
	return m.providers
}

func (m *ModuleFactory) Controllers() []*controller.ControllerFactory {
	return m.controllers
}

func (m *ModuleFactory) Exports() []interface{} {
	return m.exports
}

// Module
func (m *Module) AppendImports(imports ...*Module) {
	m.imports = append(m.imports, imports...)
}

func (m *Module) AppendProviders(providers ...*provider.Provider) {
	m.providers = append(m.providers, providers...)
}

func (m *Module) AppendControllers(controllers ...*controller.Controller) {
	m.controllers = append(m.controllers, controllers...)
}

func (m *Module) AppendExports(exports ...*provider.Provider) {
	m.exports = append(m.exports, exports...)
}

func (m *Module) Imports() []*Module {
	return m.imports
}

func (m *Module) Providers() []*provider.Provider {
	return m.providers
}

func (m *Module) Controllers() []*controller.Controller {
	return m.controllers
}

func (m *Module) Exports() []*provider.Provider {
	return m.exports
}
