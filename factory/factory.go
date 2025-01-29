package factory

import "github.com/gomdori-foo/bear-den/application"

func Create() *application.Application {
	return &application.Application{}
}