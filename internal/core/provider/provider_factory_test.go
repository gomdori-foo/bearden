package provider

import (
	"testing"

	"github.com/gomdori-foo/bearden/internal/core/common"
	"github.com/stretchr/testify/assert"
)

type TestInterface interface {}

type TestService struct {}

func NewTestService() *TestService {
	return &TestService{}
}

func TestProviders(t *testing.T) {
	t.Run("should create provider factories from constructors", func(t *testing.T) {
		// given

		// when
		providers := Providers(NewTestService)
		
		// then
		assert.Equal(t, len(providers), 1, "expected 1 provider")
		
		for _, provider := range providers {
			assert.Equal(t, provider.options.scope, common.ScopeDefault, "expected scope to be default")
		}
	})

	t.Run("should create provider factories from interface and constructor", func(t *testing.T) {
		// given

		// when 
		providers := Providers(
			Use(new(TestInterface), As(NewTestService)),
		)

		// then
		assert.Equal(t, len(providers), 1, "expected 1 provider")

		for _, provider := range providers {
			assert.Equal(t, provider.options.scope, common.ScopeDefault, "expected scope to be default")
		}
	})

	t.Run("should create provider factories from interface, constructor and options", func(t *testing.T) {
		// given

		// when 
		providers := Providers(
			Use(new(TestInterface), As(NewTestService), With(ProviderFactoryOptions{ scope: common.ScopeRequest })),
		)

		// then
		assert.Equal(t, len(providers), 1, "expected 1 provider")

		for _, provider := range providers {
			assert.Equal(t, provider.options.scope, common.ScopeRequest, "expected scope to be request")
		}
	})
}
