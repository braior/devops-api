package main

import (
	"github.com/astaxie/beego"
)

func main(){
	if ok,_ := beego.AppConfig("authpassword::enableCrontabAuthPassword");ok{
		common
	}
}