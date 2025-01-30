package controller

import (
	"testing"

	"github.com/gomdori-foo/bearden/internal/core/common"
	"github.com/stretchr/testify/assert"
)	

type TestController struct {}

func NewTestController() *TestController {
	return &TestController{}
}

func TestControllers(t *testing.T) {
	t.Run("should create controller factories from constructors", func(t *testing.T) {
		// given

		// when
		controllers := Controllers(NewTestController)

		// then
		assert.Equal(t, len(controllers), 1, "expected 1 controller")

		for _, controller := range controllers {
			assert.Equal(t, controller.options.scope, common.ScopeDefault, "expected scope to be default")
		}
	})

	t.Run("should create controller factories from constructor with options", func(t *testing.T) {
		// given

		// when
		controllers := Controllers(
			Use(NewTestController, With(ControllerFactoryOptions{ scope: common.ScopeRequest })),
		)

		// then
		assert.Equal(t, len(controllers), 1, "expected 1 controller")

		for _, controller := range controllers {
			assert.Equal(t, controller.options.scope, common.ScopeRequest, "expected scope to be request")
		}
	})
}