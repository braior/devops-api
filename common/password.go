package main

import (
	"github.com/astaxie/beego"
	"github.com/robfig/cron"
)

var (
	// WillAuthPassword 定时生成的密码
	WillAuthPassword string
	passwordFields   = map[string]interface{}{
		"entryType": "GenPassword",
	}
)

// GenPassword 生成验证密码
func GenPassword() map[string]bool {

}

// GenPassword 生成验证密码
func CronGenAuthPassword() {
	c := cron.New()
	c.AddFunc(beego.AppConfig.String("genAuthPasswordCrontab"), func() {
		GenPassword()
	})
	c.Start()
}
