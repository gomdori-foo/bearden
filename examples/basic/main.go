package main

import (
	"github.com/gomdori-foo/bearden"
)

func main() {
	app := bearden.Create(NewAppModule)
	app.Listen(":8080")
}
