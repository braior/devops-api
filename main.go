package main

import (
	_ "devops-api/routers"

	"github.com/astaxie/beego"
)

func main() {

	// 是否启用 定时生成验证密码 功能
	if ok, _ := beego.AppConfig.Bool("authpassword::enableCrontabAuthPassword"); ok {
		common.
	}

}
