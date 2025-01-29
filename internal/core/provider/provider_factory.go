package provider

import (
	"github.com/gomdori-foo/bear-den/internal/core/common"
	"github.com/gomdori-foo/bear-den/internal/core/utils"
)

// Object used when using a provider in the application.
type Provider struct {
	providerType interface{}
	constructor func() interface{}
}

// Object used when creating a provider.
type ProviderFactory struct {
	providerType interface{}
	constructor interface{}
	options ProviderFactoryOptions
}

// Object used when creating a provider.
type ProviderFactoryAs struct {
	constructor interface{}
}

// Object used when creating a provider.
type ProviderFactoryOptions struct {
	scope common.Scope
}

type ProviderFactoryWith struct {
	options ProviderFactoryOptions
}

func Providers(constructors ...interface{}) []*ProviderFactory {
	providerFactories := make([]*ProviderFactory, len(constructors))

	for i, constructor := range constructors {
		if providerFactory, ok := constructor.(*ProviderFactory); ok {
			providerFactories[i] = providerFactory
		} else if utils.IsConstructor(constructor) {
			providerFactories[i] = &ProviderFactory{
				providerType: constructor,
				constructor: constructor,
				options: NewDefaultOptions(),
			}
		} else {
			panic("invalid provider constructor")
		}
	}

	return providerFactories
}


func Use(providerType interface{}, as *ProviderFactoryAs, args ...*ProviderFactoryWith) *ProviderFactory {
	if as == nil {
		panic("provider constructor is required")
	}

	constructor := as.constructor

	options := NewDefaultOptions()

	if len(args) > 0 {
		with := args[0]
		options.scope = with.options.scope
	}

	return &ProviderFactory{
		providerType: providerType,
		constructor: constructor,
		options: options,
	}
}

// Used when receiving an implementation for an interface
func As(constructor interface{}) *ProviderFactoryAs {
	if !utils.IsConstructor(constructor) {
		panic("invalid provider constructor")	
	} 

	return &ProviderFactoryAs{
		constructor: constructor,
	}
}

// Will be used in the future
func With(options ProviderFactoryOptions) *ProviderFactoryWith {
	return &ProviderFactoryWith{
		options: options,
	}
}


func NewDefaultOptions() ProviderFactoryOptions {
	return ProviderFactoryOptions{
		scope: common.ScopeDefault,
	}
}
