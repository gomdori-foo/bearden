package provider

import (
	"reflect"
	"runtime"
	"strings"
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
	options *ProviderFactoryOptions
}

// Object used when creating a provider.
type ProviderConstructor struct {
	constructor interface{}
}

// Object used when creating a provider.
type ProviderFactoryOptions struct {
	scope ProviderScope
}

type ProviderScope string

const (
	ScopeDefault ProviderScope = "DEFAULT"
	ScopeTransient ProviderScope = "TRANSIENT"
	ScopeRequest ProviderScope = "REQUEST"
)

func Providers(constructors ...interface{}) []*ProviderFactory {
	providerFactories := make([]*ProviderFactory, len(constructors))

	for i, constructor := range constructors {
		if providerFactory, ok := constructor.(*ProviderFactory); ok {
			providerFactories[i] = providerFactory
		} else if isConstructor(constructor) {
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


func Use(providerType interface{}, providerConstructor *ProviderConstructor, providerFactoryOptions ...*ProviderFactoryOptions) *ProviderFactory {
	if providerConstructor == nil {
		panic("provider constructor is required")
	}

	constructor := providerConstructor.constructor

	options := NewDefaultOptions()

	if len(providerFactoryOptions) > 0 {
		customOptions := providerFactoryOptions[0]
		options.scope = customOptions.scope
	}

	return &ProviderFactory{
		providerType: providerType,
		constructor: constructor,
		options: options,
	}
}

// Used when receiving an implementation for an interface
func As(constructor interface{}) *ProviderConstructor {
	if !isConstructor(constructor) {
		panic("invalid provider constructor")	
	} 

	return &ProviderConstructor{
		constructor: constructor,
	}
}

// Will be used in the future
func With(options *ProviderFactoryOptions) *ProviderFactoryOptions {
	return options
}

func isConstructor(constructor interface{}) bool {
	constructorType := reflect.TypeOf(constructor)
	constructorValue := reflect.ValueOf(constructor)
	constructorName := runtime.FuncForPC(constructorValue.Pointer()).Name()

	if lastDot := strings.LastIndex(constructorName, "."); lastDot >= 0 {
		constructorName = constructorName[lastDot+1:]
	}

	return constructorType.Kind() == reflect.Func && strings.HasPrefix(constructorName, "New")
}

func NewDefaultOptions() *ProviderFactoryOptions {
	return &ProviderFactoryOptions{
		scope: ScopeDefault,
	}
}
