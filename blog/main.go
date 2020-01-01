package main

// spaces for import priority order
import (
	_ "examples/blog/Config"

	_ "examples/blog/App/Providers"

	_ "examples/blog/App/Models"

	_ "examples/blog/App/Middleware"

	_ "examples/blog/Route"

	"github.com/goKLC/goKLC"
)

func main() {
	app := goKLC.GetApp()

	app.Run()
}
