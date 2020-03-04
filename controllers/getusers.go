package controllers

import (
	"time"
	"strconv"
	"mindplus_weather/models"
	"github.com/astaxie/beego"
)

type GetUsersController struct {
	beego.Controller
}

func (c *GetUsersController) Post() {
	var visitsMin int = 0
	var visitsMax int = 999999
	var updatedStartTime time.Time
	updatedStartTime,_ = time.ParseInLocation("2006-01-02 15:04:05", "2018-01-01 00:00:00", time.Local)
	var updatedEndTime time.Time = time.Now()
	var createdStartTime time.Time
	createdStartTime,_ = time.ParseInLocation("2006-01-02 15:04:05", "2018-01-01 00:00:00", time.Local)
	var createdEndTime time.Time = time.Now()
	var city string
	var orderby string = "id desc"
	var offset int = 0
	var limit int = -1

	beego.Debug("input=",c.Input())

	if t := c.Input().Get("orderby");t != ""{
		orderby = t
	}

	if t := c.Input().Get("city");t != ""{
		city = t
	}

	if t := c.Input().Get("offset");t != ""{
		offset, _ = strconv.Atoi(t)
	}

	if t := c.Input().Get("limit"); t != ""{
		limit, _ = strconv.Atoi(t)
	}

	if t := c.Input().Get("visitsMin");t != ""{
		visitsMin, _ = strconv.Atoi(t)
	}

	if t := c.Input().Get("visitsMax"); t != ""{
		visitsMax, _ = strconv.Atoi(t)
	}

	if t := c.Input().Get("updatedStartTime"); t != ""{
		updatedStartTime,_ = time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	}
	if t := c.Input().Get("updatedEndTime"); t != ""{
		updatedEndTime,_ = time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	}
	if t := c.Input().Get("createdStartTime"); t != ""{
		createdStartTime,_ = time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	}
	if t := c.Input().Get("createdEndTime"); t != ""{
		createdEndTime,_ = time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	}

	beego.Debug(offset)
	beego.Debug(limit)
	beego.Debug(city)
	beego.Debug(orderby)
	beego.Debug(visitsMin)
	beego.Debug(visitsMax)
	beego.Debug(updatedStartTime)
	beego.Debug(updatedEndTime)
	beego.Debug(createdStartTime)
	beego.Debug(createdEndTime)
	data := models.GetUsers(visitsMin, visitsMax, updatedStartTime, updatedEndTime, createdStartTime, createdEndTime, city , offset, limit, orderby);
	c.Ctx.WriteString(data)
}
