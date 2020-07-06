package common

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/hltool"
)

var (

	// UploadPath 上传目录
	UploadPath string

	// Logger 日志对象
	Logger *hltool.HLogger

	// QrImageDir 二维码图片目录
	QrImageDir string

	// LogPathFromCli 从命令行传入日志路径
	LogPathFromCli string

	// EnableToken 读取是否启用 Token 认证配置
	EnableToken = getEnableToken()

	// DBPath 数据库文件路径
	DBPath = beego.AppConfig.String("database::dbPath")
)

func init() {

	// 2步验证图片目录是否存在
	QrImageDir = beego.AppConfig.String("twostepauth::qrImageDir")
	if !hltool.IsExist(QrImageDir) {
		os.MkdirAll(QrImageDir, os.ModePerm)
	}
}

// getEnableToken 读取是否启用 Token 认证配置
func getEnableToken() bool {
	enableToken, err := beego.AppConfig.Bool("security::enableToken")
	if err != nil {
		enableToken = true
	}

	return enableToken
}
