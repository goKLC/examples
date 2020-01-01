package Controller

import (
	"examples/blog/App/Models"
	"github.com/goKLC/goKLC"
)

type indexController goKLC.Controller

var IndexController = indexController{}

func (ic indexController) Index(r *goKLC.Request) *goKLC.Response {
	var blogs []Models.Blog

	goKLC.GetApp().DB().Find(&blogs)

	return goKLC.NewResponse().Ok(goKLC.NewView("index", goKLC.Context{"blogs": blogs}).Render())
}

func (ic indexController) Post(r *goKLC.Request) *goKLC.Response {

	return goKLC.NewResponse().Ok(goKLC.NewView("post", goKLC.Context{}).Render())
}
