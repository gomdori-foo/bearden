package main

import (
	"github.com/gomdori-foo/bearden/factory"
)

func main() {
	app := factory.Create(NewAppModule)
	app.Listen(":8080")
}
