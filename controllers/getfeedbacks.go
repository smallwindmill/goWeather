package controllers

import (
	"time"
	"strconv"
	"mindplus_statistic/models"
	"github.com/astaxie/beego"
)

type GetFeedBacksController struct {
	beego.Controller
}

func (c *GetFeedBacksController) Post() {
	var createdStartTime time.Time 
	createdStartTime,_ = time.ParseInLocation("2006-01-02 15:04:05", "2018-01-01 00:00:00", time.Local) 
	var createdEndTime time.Time = time.Now()
	var offset int = 0
	var limit int = -1
	
	beego.Debug("input=",c.Input())

	if t := c.Input().Get("offset");t != ""{
		offset, _ = strconv.Atoi(t)
	}

	if t := c.Input().Get("limit"); t != ""{
		limit, _ = strconv.Atoi(t)
	}

	if t := c.Input().Get("createdStartTime"); t != ""{
		createdStartTime,_ = time.ParseInLocation("2006-01-02 15:04:05", t, time.Local) 
	}
	if t := c.Input().Get("createdEndTime"); t != ""{
		createdEndTime,_ = time.ParseInLocation("2006-01-02 15:04:05", t, time.Local) 
	}

	beego.Debug(offset)
	beego.Debug(limit)
	beego.Debug(createdStartTime)
	beego.Debug(createdEndTime)
	data := models.GetFeedBacks(createdStartTime, createdEndTime, offset, limit)
	c.Ctx.WriteString(data)
}