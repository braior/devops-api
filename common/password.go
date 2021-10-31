package common

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/braior/brtool"
	"github.com/braior/devops-api/model"
	"github.com/braior/devops-api/utils"
	"github.com/spf13/viper"
)

var (
	genPasswordFields = map[string]interface{}{
		"entryType": "GenPassword",
	}

	expiredTime string
)

// // GenPassword 生成验证密码
func GenPassword(username string, email []string) bool {
	var result bool

	genAuthPassword := brtool.GenRandomString(32, "mix")
	info := fmt.Sprintf("请注意,本次验证密码为: %s 生成时间: %s", genAuthPassword, time.Now().Format("20060102150405"))
	utils.Logger.Info(genPasswordFields, info)

	rc := model.RedisPool().Get()
	defer rc.Close()

	expiredTime = viper.GetString("authPassword.authPasswordExpiration")
	if expiredTime == "" {
		expiredTime = "60"
	}

	_, err := rc.Do("set", username, genAuthPassword, "EX", expiredTime)
	if err != nil {
		errLog := fmt.Sprintf("set redis key err: %s", err)
		beego.BeeLogger.Error(errLog)
	}

	sendPasswordResult := make(chan bool)
	if ok := viper.GetBool("authPassword.enableDingtalkReciveGenPassword"); ok {
		go func(ch chan bool, messageType, message string) {
			ok, _ := SendByDingTalkRobot(messageType, message, "", "")
			ch <- ok
		}(sendPasswordResult, "text", info)
	}

	if ok := viper.GetBool("authPassword.enableEmailReciveGenAuthPassword"); ok {

		go func(ch chan bool, subject, content, contentType, attach string, to, cc []string) {
			ok, _ := SendByEmail(subject, content, contentType, attach, to, cc)
			ch <- ok
		}(sendPasswordResult, "验证码", info, "text/plain", "", email, []string{})
	}

	for {
		if <-sendPasswordResult {
			result = true
			break
		}
	}
	return result

}

func CheckPassword(username, password string) (bool, error) {
	rc := model.RedisPool().Get()
	defer rc.Close()

	authPassword, err := rc.Do("get", username)
	if err != nil {
		errLog := fmt.Sprintf("set redis key err: %s", err)
		beego.BeeLogger.Error(errLog)
	}
	fmt.Println(utils.Strval(authPassword))

	if password == utils.Strval(authPassword) {
		return true, nil
	}
	return false, nil
}
