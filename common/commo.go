package common

import "github.com/astaxie/beego"

var (
	// DBPath 数据库文件路径
	DBPath = beego.AppConfig.String("database::dbPath")
)
