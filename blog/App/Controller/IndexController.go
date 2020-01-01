package Controller

import (
	"github.com/goKLC/goKLC"
)

type indexController goKLC.Controller

var IndexController = indexController{}

func (ic indexController) Index(r *goKLC.Request) *goKLC.Response {

	return goKLC.NewResponse().Ok(goKLC.NewView("index", goKLC.Context{}).Render())
}
