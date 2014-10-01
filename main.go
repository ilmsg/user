package main

import (
	_ "github.com/ilmsg/user/routers"
	"github.com/ilmsg/user/models"
	"github.com/astaxie/beego"
)

func main() {
	models.Connect();
	beego.Run()
}

