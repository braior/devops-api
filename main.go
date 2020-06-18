package main

import (
	"devops-api/common"
	_ "devops-api/routers"
)

func main() {

	// 是否启用 定时生成验证密码 功能

	// 初始化获取命令行参数
	common.InitCli()

}
