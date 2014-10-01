package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {
	this.Layout = "layout/default.html"
}

func (this *MainController) Index() {
	this.TplNames = "main/index.html"
}
