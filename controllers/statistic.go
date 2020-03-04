package controllers

import (
	"time"
	"strconv"
	"mindplus_weather/models"
	"github.com/astaxie/beego"
)

type StatisticController struct {
	beego.Controller
}

func (c *StatisticController) Post() {
	beego.Debug("input=",c.Input())
	var result string
	var value string = "user"
	var visitsMin int = 0
	var visitsMax int = 1000000
	var updatedStartTime time.Time
	var updatedEndTime time.Time
	var createdStartTime time.Time
	var createdEndTime time.Time

	updatedStartTime,_ = time.ParseInLocation("2006-01-02 15:04:05", c.Input().Get("updatedStartTime"),time.Local)
	updatedEndTime = time.Now()
	createdStartTime,_ = time.ParseInLocation("2006-01-02 15:04:05", c.Input().Get("createdStartTime"),time.Local)
	createdEndTime = time.Now()

	if v := c.Input().Get("visitsMin"); v!=""{
		visitsMin, _ = strconv.Atoi(v)
	}
	if v := c.Input().Get("visitsMax"); v!=""{
		visitsMax, _ = strconv.Atoi(v)
	}
	if v := c.Input().Get("updatedStartTime"); v!=""{
		updatedStartTime, _ =  time.ParseInLocation("2006-01-02 15:04:05", v,time.Local)
	}
	if v := c.Input().Get("updatedEndTime"); v!=""{
		updatedEndTime, _ = time.ParseInLocation("2006-01-02 15:04:05", v,time.Local)
	}
	if v := c.Input().Get("createdStartTime"); v!=""{
		createdStartTime, _ = time.ParseInLocation("2006-01-02 15:04:05", v,time.Local)
	}
	if v := c.Input().Get("createdEndTime"); v!=""{
		createdEndTime, _ = time.ParseInLocation("2006-01-02 15:04:05", v,time.Local)
	}

	if v := c.Input().Get("value"); v != ""{
		value = v
	}
	beego.Debug(visitsMin)
	beego.Debug(visitsMax)
	beego.Debug(updatedStartTime)
	beego.Debug(updatedEndTime)
	beego.Debug(createdStartTime)
	beego.Debug(createdEndTime)

	//t,_ := time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06",time.Local)
	//log.Println(t)
	//log.Println(time.Now())
	//m_96h,_ := time.ParseDuration("-97h");
	//m_1h,_ := time.ParseDuration("1h");
	//data := models.GetStatistic(visitsMin, visitsMax, time.Now().Add(m_96h), time.Now().Add(m_1h), time.Now().Add(m_96h), time.Now().Add(m_1h))
	if(value == "user"){
		result = models.GetUserStatistic(visitsMin, visitsMax, updatedStartTime, updatedEndTime, createdStartTime, createdEndTime)
	}else{
		result = models.GetFeedBackStatistic(createdStartTime, createdEndTime)
	}
	c.Ctx.WriteString(result)
}
