package routers

import (
	"mindplus_weather/controllers"
	// "mindplus_weather/plugin"

	"github.com/astaxie/beego"
)


func init() {

	// 跨域
	/*beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
      AllowAllOrigins:  true,
      //AllowOrigins:      []string{"https://192.168.0.102"},
      AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
      AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
      ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
      AllowCredentials: true,
  }))*/



	// 在线http请求转发
	beego.Router("/api/weather", &controllers.WeatherController{}, "get:Weather")
}


