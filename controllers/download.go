package controllers

import (
	"os"
	"github.com/astaxie/beego"
)

type DownloadController struct {
	beego.Controller
}

func (c *DownloadController) Post() {
	var url string
	beego.Debug("input=",c.Input())

	if v := c.Input().Get("url");v != ""{
		url = v
	}

	beego.Debug(url)
	if(url != ""){
		url = "/home/dfrobot/attachment/"+url
		_, err := os.Stat(url)
		if err == nil {
			c.Ctx.Output.Download(url)
		}else{
			c.Ctx.WriteString("file do not exist")
		}
	}
	c.Ctx.WriteString("url empty")
}