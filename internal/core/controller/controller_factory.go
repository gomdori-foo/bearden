package controller

import (
	"github.com/gomdori-foo/bear-den/internal/core/common"
	"github.com/gomdori-foo/bear-den/internal/core/utils"
)

type Controller struct {
	constructor interface{}
}

type ControllerFactory struct {
	constructor interface{}
	options ControllerFactoryOptions
}

type ControllerFactoryOptions struct {
	scope common.Scope
}

type ControllerFactoryWith struct {
	options ControllerFactoryOptions
}

func Controllers(constructors ...interface{}) []*ControllerFactory {
	controllerFactories := make([]*ControllerFactory, len(constructors))

	for i, constructor := range constructors {
		if controllerFactory, ok := constructor.(*ControllerFactory); ok {
			controllerFactories[i] = controllerFactory
		} else if utils.IsConstructor(constructor) {
			controllerFactories[i] = &ControllerFactory{
				constructor: constructor,
				options: NewDefaultOptions(),
			}
		} else {
			panic("invalid controller constructor")
		}
	}

	return controllerFactories
}

func Use(constructor interface{}, args ...*ControllerFactoryWith) *ControllerFactory {
	if !utils.IsConstructor(constructor) {
		panic("invalid controller constructor")
	}

	options := NewDefaultOptions()

	if len(args) > 0 {
		with := args[0]
		options.scope = with.options.scope
	}

	return &ControllerFactory{
		constructor: constructor,
		options: options,
	}
}

func With(options ControllerFactoryOptions) *ControllerFactoryWith {
	return &ControllerFactoryWith{
		options: options,
	}
}

func NewDefaultOptions() ControllerFactoryOptions {
	return ControllerFactoryOptions{
		scope: common.ScopeDefault,
	}
}
