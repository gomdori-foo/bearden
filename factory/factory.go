package factory

import "github.com/gomdori-foo/bear-den/internal/core/application"

func Create() *application.Application {
	return &application.Application{}
}