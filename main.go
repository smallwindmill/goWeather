package main

import (
	"fmt"
	"mindplus_statistic/routers"
	"mindplus_statistic/models"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println(routers.Sub(9, 3))
	fmt.Println(models.Add(9, 3))
	beego.Run()
}

