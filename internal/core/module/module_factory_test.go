package module

import (
	"testing"

	"github.com/gomdori-foo/bearden/internal/core/provider"
	"github.com/stretchr/testify/assert"
)

type TestController struct {
	testUseCase TestUseCase
}

func NewTestController(testUseCase TestUseCase) *TestController {
	return &TestController{
		testUseCase: testUseCase,
	}
}	

type TestUseCase interface {}

type TestService struct {}

func NewTestService() *TestService {
	return &TestService{}
}

type AppModule struct {}

func NewAppModule() *ModuleFactory {
	return Builder().
		Controllers(
			NewTestController,
		).
		Providers(
			provider.Use(new(TestUseCase), provider.As(NewTestService)),
		).
		Build()
}

func TestModuleFactoryBuilder(t *testing.T) {
	t.Run("should create module factory from builder", func(t *testing.T) {
		// given

		// when
		moduleFactory := NewAppModule()

		// then
		assert.Equal(t, len(moduleFactory.controllers), 1, "expected 1 controller")
		assert.Equal(t, len(moduleFactory.providers), 1, "expected 1 provider")
	})
}