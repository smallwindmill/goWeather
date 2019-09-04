package controllers

import (
	"mindplus_statistic/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {

	beego.Debug("input=",c.Input())
	username := c.Input().Get("username")
	password := c.Input().Get("password")

	beego.Debug(username)
	beego.Debug(password)

	err := models.Login(username,password)
	if(err != nil){
		c.Ctx.WriteString("failure")
	}else{
		c.Ctx.WriteString("ok")
	}
}