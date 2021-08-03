package common

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/braior/brtool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Logger *brtool.BRLogger

// NewLogger return a log instance for *logrus.Logger
func initLog() {

	var logPath string

	if *LogPathFromCli == "" {
		logPath = viper.GetString("log.logPath")
	} else {
		logPath = *LogPathFromCli
	}
	// 日志中添加文件名和方法信息
	logrus.SetReportCaller(true)

	var currentDir string
	if !strings.HasPrefix(logPath, "/") {
		var err error
		currentDir, err = os.Getwd()
		if err != nil {
			log.Fatalf("error: %s", err)
		}
	}
	logPath = path.Join(currentDir, logPath)
	blog, err := brtool.NewBRLog(logPath)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	Logger, err = blog.GetLogger()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
