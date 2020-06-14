package common

import (
	"log"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New(AppName, AppDescription)

	inits            = app.Command("init", "Init action")
	refreshRootToken = inits.Flag("refresh-root-token", "refresh root token").Bool()
	server           = app.Command("server", "Server mode")
	logPath          = server.Flag("log", "Log Path, In Configgure File, Default: logs/devops-api.log").String()
	runMode          = server.Flag("mode", "Run Mode: dev|prod|test, In Configure File, Default: dev").String()

	token       = app.Command("token", "Token Manage")
	tokenRoot   = token.Flag("root-token", "Specify Root Token").Required().String()
	tokenCreate = token.Flag("create", "Create a Token, Special a Name").String()
	tokenDelete = token.Flag("delete", "Delete a Token, Special a Name").String()

	backup   = app.Command("backup", "Backup BoltDB DB File")
	filepath = backup.Flag("filepath", "Special Backup FilePath").String()
)

// InitCli 初始化命令行参数
func InitCli() {
	app.Author(Author).Version(AppVersion)

	c, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("parse cli args error: %s\n", err)
	}

	switch c{
	case "init":
		var token *Token
		var err error
		token,err= NewToken()
		if *tokenCreate != ""{
			err = token.
		}

	}
}
