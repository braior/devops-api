package utils

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/braior/brtool"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	LogPathFromCli string
	Debug          bool
	Logger         *brtool.BRFileLogger
)

// NewLogger return a log instance for *logrus.Logger
func LogInit() {

	var logPath string

	if LogPathFromCli == "" {
		logPath = viper.GetString("log.path")
	} else {
		logPath = LogPathFromCli
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
	blog, err := brtool.NewBRFileLog(logPath)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if viper.GetString("log.rotateTime") == "" {
		blog.IsSeparateLevelLog = false
	} else {
		blog.IsSeparateLevelLog = true
	}

	if Debug {
		blog.LogLevel = logrus.DebugLevel
	}

	Logger, err = blog.GetLogger()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
