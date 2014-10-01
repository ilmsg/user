package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ilmsg/user/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Prepare() {
	this.Layout = "layout/default.html"
}

func (this *UserController) Login() {
	user := this.GetSession("user")
	if user != nil {
		this.Ctx.Redirect(302, "/")
	}
	this.TplNames = "user/login.html"
}

func (this *UserController) DoLogin() {
	username := this.GetString("username")
	password := this.GetString("password")
	
	user, err := models.CheckLogin(username, password)
	if err == nil {
		this.SetSession("user", user)
		this.Ctx.Redirect(302, "/")
	}else{
		this.Data["username"] = username
		this.Data["password"] = password
		this.TplNames = "user/login.html"
	}
}

func (this *UserController) Register() {
	this.TplNames = "user/register.html"
}

func (this *UserController) DoRegister() {
	this.TplNames = "user/register.html"
}

func (this *UserController) Logout() {
	this.DelSession("user")
	this.Ctx.Redirect(302, "/")
}

func (this *UserController) Profile() {
	user := this.GetSession("user")
	if user == nil {
		this.Ctx.Redirect(302, "/")
	}
	this.Data["user"] = user
	this.TplNames = "user/profile.html"
}
