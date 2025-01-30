package main

import (
	"github.com/gomdori-foo/bearden"
)

// @Module()
func NewAppModule() *bearden.Module {
	return bearden.Builder().
		Controllers(NewAppController).
		Providers(NewAppService).
		Build()
}