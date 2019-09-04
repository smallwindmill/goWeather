package routers

import (
	"mindplus_statistic/controllers"
	"github.com/astaxie/beego"
)

func Sub(x,y int) int {
	return x-y
}

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/statistic", &controllers.StatisticController{})
	beego.Router("/adduser", &controllers.AddUserController{})
	beego.Router("/getusers", &controllers.GetUsersController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/addfeedback", &controllers.AddFeedBackController{})
	beego.Router("/getfeedbacks", &controllers.GetFeedBacksController{})
	beego.Router("/download", &controllers.DownloadController{})
	beego.Router("/upload", &controllers.UploadController{})
	// beego.Router("/collect", &controllers.CollectController{}, "get:getCollect;post:postCollect")
	beego.Router("/collect", &controllers.CollectController{}, "get:GetCollect;post:PostCollect")
}


