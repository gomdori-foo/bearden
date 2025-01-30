package factory

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gomdori-foo/bearden/internal/application"
	"github.com/gomdori-foo/bearden/internal/core/controller"
	"github.com/gomdori-foo/bearden/internal/core/flag"
	"github.com/gomdori-foo/bearden/internal/core/module"
	"github.com/gomdori-foo/bearden/internal/core/provider"
)

func Create(constructor func() *module.ModuleFactory) *application.BearDenApplication {
	appModuleFactory := constructor()

	if flag.IsProduction() {
		return createProductionApplication(appModuleFactory)
	}

	return createDevelopmentApplication(appModuleFactory)
}

func createProductionApplication(appModuleFactory *module.ModuleFactory) *application.BearDenApplication {
	return &application.BearDenApplication{}
}

func createDevelopmentApplication(appModuleFactory *module.ModuleFactory) *application.BearDenApplication {
	router := gin.Default()
	module := createModule(appModuleFactory)

	return application.NewBearDenApplication(router, module)
}

func createModule(appModuleFactory *module.ModuleFactory) *module.Module {
	module := &module.Module{}

	// Create imported modules
	createSubModules(module, appModuleFactory)

	// Create providers
	createProviders(module, appModuleFactory, appModuleFactory.Providers())

	// Create controllers
	createControllers(module, appModuleFactory.Controllers())

	return module
}

func createSubModules(module *module.Module, moduleFactory *module.ModuleFactory) {
	for _, moduleFactoryConstructor := range moduleFactory.Imports() {
		moduleFactory := moduleFactoryConstructor()
		generatedModule := createModule(moduleFactory)
		module.AppendProviders(generatedModule.Providers()...)

		pc := reflect.ValueOf(moduleFactoryConstructor).Pointer()
		for _, export := range moduleFactory.Exports() {
			if reflect.ValueOf(export).Pointer() == pc {
				module.AppendExports(generatedModule.Providers()...)
			}
		}
	}
}

func createProviders(module *module.Module, moduleFactory *module.ModuleFactory, providerFactories []*provider.ProviderFactory) {
	var retryProviderFactories []*provider.ProviderFactory

	for _, providerFactory := range providerFactories {
		provider := providerFactory.Create(module.Providers())
		if provider == nil {
			retryProviderFactories = append(retryProviderFactories, providerFactory)
			continue
		}

		module.AppendProviders(provider)

		pc := reflect.ValueOf(providerFactory).Pointer()
		for _, export := range moduleFactory.Exports() {
			if reflect.ValueOf(export).Pointer() == pc {
				module.AppendExports(provider)
			}
		}
	}

	if len(retryProviderFactories) > 0 {
		if len(retryProviderFactories) == len(providerFactories) {
			panic("Failed to generate providers")
		}

		createProviders(module, moduleFactory, retryProviderFactories)
	}
}

func createControllers(module *module.Module, controllerFactories []*controller.ControllerFactory) {
	for _, controllerFactory := range controllerFactories {
		controller := controllerFactory.Create(module.Providers())
		if controller == nil {
			panic("Failed to generate controller")
		}
		module.AppendControllers(controller)
	}
}