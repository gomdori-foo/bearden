package main

import bearden "github.com/gomdori-foo/bear-den"

func NewAppModule() *bearden.ModuleFactory {
	return bearden.Builder().
		Controllers(NewAppController).
		Providers(NewAppService).
		Build()
}