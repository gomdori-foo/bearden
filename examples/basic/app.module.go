package main

import bearden "github.com/gomdori-foo/bearden"

// @Module()
func NewAppModule() *bearden.ModuleFactory {
	return bearden.Builder().
		Controllers(NewAppController).
		Providers(NewAppService).
		Build()
}