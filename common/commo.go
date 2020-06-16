package common

import "github.com/astaxie/beego"

var (
	// DBPath 数据库文件路径
	DBPath = beego.AppConfig.String("database::dbPath")

	// LogPathFromCli 从命令行传入日志路径
	LogPathFromCli string
	
	// EnableToken 读取是否启用 Token 认证配置
	EnableToken = getEnableToken()
)

// getEnableToken 读取是否启用 Token 认证配置
func getEnableToken(){
	enableToken,err:=beego.AppConfig.Bool("security::enableToken"){
		if err!=nil{
			enableToken =true
		}

		return enableToken
	}
}