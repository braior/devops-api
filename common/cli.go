package common

import (
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/braior/brtool"
	"github.com/spf13/viper"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New(AppName, AppDescription)

	inits            = app.Command("init", "Init Action")
	refreshRootToken = inits.Flag("refresh-root-token", "Refresh Root Token").Bool()

	server  = app.Command("server", "Server Mode")
	LogPathFromCli = server.Flag("log", "Log Path, In Configure File, Default:./logs/devops-api.log").String()
	runMode = server.Flag("mode", "Run Mode: dev|test|prod, In Configure File, Default: dev").String()

	token       = app.Command("token", "Token Manage")
	tokenRoot   = token.Flag("root-token", "Specify Root Token").Required().String()
	tokenCreate = token.Flag("create", "Create a Token, Special a Name").String()
	tokenDelete = token.Flag("delete", "Delete a Token, Special a Name").String()

	backup   = app.Command("backup", "Backup BoltDB DB File")
	filepath = backup.Flag("filepath", "Special Backup FilePath").String()
)

// InitCli 初始化命令行参数
func InitCli() {

	//获取项目的执行路径
	appPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 指定配置文件路径，配置文件名称及类型
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(appPath + "/conf/")

	// 读取配置文件内容
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 绑定环境变量，优先从环境变量读取
	// viper.BindEnv("kafka.peers", "KAFKA_PEERS")

	// 从命令行获取配置信息，覆盖配置文件提供的配置
	// flag.Int("app.httpPort", 8080, "set app http port")
	// flag.String("log.logPath", "logs/devops-api.log", "set app log path")

	// 解析命令行传入的参数
	// flag.Parse()

	// viper 绑定flag传入的参数
	// viper.BindPFlags(flag.CommandLine)

	// 加载系统环境变量
	// viper.AutomaticEnv()

	// 初始化日志对象
	initLog()

	app.Author(Author).Version(AppVersion)

	c, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("parse cli args error: %s\n", err)
	}

	switch c {
	case "server":
		// 获取app运行模式
		if *runMode != "" {
			if _, ok := brtool.InstringSlice([]string{"dev", "test", "prod"}, *runMode); !ok {
				log.Fatalln("get run mode input error, mode: dev|test|prod")
			}
		}
		beego.BConfig.RunMode = *runMode
	}
	beego.SetStaticPath("/api/static/download/qr", "static/download/qr")
	beego.Run()
}
