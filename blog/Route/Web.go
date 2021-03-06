package Route

import (
	"examples/blog/App/Controller"
	"github.com/goKLC/goKLC"
)

var app *goKLC.App

func init() {
	app = goKLC.GetApp()
	Router()
}

func Router() {
	app.Route().Get("/", Controller.IndexController.Index).Name("index")
	app.Route().Get("/post", Controller.IndexController.Post).Name("post")
}
