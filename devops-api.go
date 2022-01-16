package main

import (
	"github.com/braior/devops-api/cmd"
	_ "github.com/braior/devops-api/routers"
)

func main() {

	// 初始化获取命令行参数
	// cmd.LogInit()
	cmd.Execute()

}
