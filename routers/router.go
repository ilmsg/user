package routers

import (
	"github.com/ilmsg/user/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:Index" )
	
	beego.Router("/register", &controllers.UserController{}, "get:Register;post:DoRegister" )
	beego.Router("/login", &controllers.UserController{}, "get:Login;post:DoLogin" )	
	beego.Router("/profile", &controllers.UserController{}, "get:Profile" )
	beego.Router("/logout", &controllers.UserController{}, "get:Logout" )
}
