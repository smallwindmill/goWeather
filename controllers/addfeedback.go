package controllers

import (
	"time"
	"mindplus_weather/models"
	"github.com/astaxie/beego"
)

type AddFeedBackController struct {
	beego.Controller
}

func (c *AddFeedBackController) Post() {
	beego.Debug("input=",c.Input())
	mac := c.Input().Get("mac")
	feedBackMsg := c.Input().Get("feedBackMsg")
	email := c.Input().Get("email")
	path := c.Input().Get("path")

	createdTime := time.Now()

	if v := c.Input().Get("Created"); v != ""{
		createdTime,_ = time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
	}

	beego.Debug(mac)
	beego.Debug(createdTime)
	beego.Debug(feedBackMsg)
	beego.Debug(email)
	beego.Debug(path)
	models.AddFeedBack(mac, createdTime, feedBackMsg, email, path)
	c.Ctx.WriteString("ok")
}
