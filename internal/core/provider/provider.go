package provider

import (
	"reflect"

	"github.com/gomdori-foo/bearden/internal/core/common"
)

// Object used when using a provider in the application.
type Provider struct {
	providerType reflect.Type
	constructor func() interface{}
	instance interface{}
	options ProviderOptions
}

type ProviderOptions struct {
	scope common.Scope
}

func NewProvider(providerType reflect.Type, constructor func() interface{}, options ProviderOptions) *Provider {
	return &Provider{
		providerType: providerType,
		constructor: constructor,
		options: options,
		instance: nil,
	}
}

func NewProviderOptions(scope common.Scope) ProviderOptions {
	return ProviderOptions{
		scope: common.ScopeDefault,
	}
}

func (p *Provider) Instance() interface{} {
	if p.options.scope == common.ScopeDefault {
		if p.instance == nil {
			p.instance = p.constructor()
		}
		return p.instance
	}

	return p.constructor()
}
