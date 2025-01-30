package controller

import "github.com/gomdori-foo/bearden/internal/core/common"

type Controller struct {
	prefix string
	routes []Route
	constructor func() interface{}
	instance interface{}
	options ControllerOptions
}

type ControllerOptions struct {
	scope common.Scope
}

func NewController(prefix string, routes []Route, constructor func() interface{}) *Controller {
	return &Controller{
		prefix: prefix,
		routes: routes,
		constructor: constructor,
	}
}

func (c *Controller) Instance() interface{} {
	if c.options.scope == common.ScopeDefault {
		if c.instance == nil {
			c.instance = c.constructor()
		}
		return c.instance
	}

	return c.constructor()
}

func (c *Controller) Prefix() string {
	return c.prefix
}

func (c *Controller) Routes() []Route {
	return c.routes
}

func (c *Controller) Options() ControllerOptions {
	return c.options
}