package main

import (
	"github.com/braior/devops-api/cmd"
	_ "github.com/braior/devops-api/routers"
)

func main() {

	// 初始化获取命令行参数
	// cmd.LogInit()
	cmd.Execute()

	// args := cmd.RootCmd.Args
	// fmt.Printf("%v:", args)
	// switch args {
	// case "server":
	// 	// 获取app运行模式
	// 	if cmd.RunMode != "" {
	// 		if _, ok := brtool.InstringSlice([]string{"dev", "test", "prod"}, *runMode); !ok {
	// 			log.Fatalln("get run mode input error, mode: dev|test|prod")
	// 		}
	// 	}
	// 	beego.BConfig.RunMode = cmd.RunMode
	// }
	// beego.SetStaticPath("/api/static/download/qr", "static/download/qr")
	// beego.Run()
}
