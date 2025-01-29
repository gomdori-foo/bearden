package main

import (
	"github.com/gomdori-foo/bear-den/factory"
)

func main() {
	app := factory.Create()
	app.Listen(":8080")
}
