package provider

import (
	"reflect"

	"github.com/gomdori-foo/bearden/internal/core/common"
	"github.com/gomdori-foo/bearden/internal/core/utils"
)

// Object used when creating a provider.
type ProviderFactory struct {
	providerType reflect.Type
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
			providerType := reflect.TypeOf(constructor).Out(0)
			if providerType.Kind() == reflect.Ptr {
				providerType = providerType.Elem()
			}
			
			providerFactories[i] = &ProviderFactory{
				providerType: providerType,
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
		providerType: reflect.TypeOf(providerType),
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

func (p *ProviderFactory) Create(providers []*Provider) *Provider {
	constructorType := reflect.TypeOf(p.constructor)
	if constructorType.Kind() != reflect.Func {
		panic("provider constructor is not a function")
	}

	provider := FindProvider(providers, constructorType)
	if provider != nil {
		return provider
	}

	params := make([]interface{}, constructorType.NumIn())

	for i := 0; i < len(params); i++ {
		paramType := constructorType.In(i)
		provider := FindProvider(providers, paramType)
		if provider == nil {
			return nil
		}

		params[i] = provider.Instance()
	}


	constructor := func() interface{} {
		constructorValue := reflect.ValueOf(p.constructor)
		paramValues := make([]reflect.Value, len(params))
		for i, param := range params {
			paramValues[i] = reflect.ValueOf(param)
		}
		result := constructorValue.Call(paramValues)[0].Interface()
		return result
	}

	options := NewProviderOptions(p.options.scope)

	provider = NewProvider(p.providerType, constructor, options)

	return provider
}

func (p *ProviderFactory) FindProvider(providers []*Provider) *Provider {
	constructorType := reflect.TypeOf(p.constructor)
	return FindProvider(providers, constructorType)
}

func FindProvider(providers []*Provider, reflectType reflect.Type) *Provider {
	for _, provider := range providers {
		providerType := provider.providerType

		if providerType.Kind() == reflect.Ptr {
			providerType = providerType.Elem()
		}
	
		compareType := reflectType
		if reflectType.Kind() == reflect.Ptr {
			compareType = reflectType.Elem()
		}

		if providerType == compareType {
			return provider
		}
	}

	return nil
}