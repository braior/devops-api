package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/braior/brtool"
	"github.com/braior/devops-api/common"
	"github.com/sirupsen/logrus"
)

var (
	genAuthPasswordEntryType   = "GenAuthPassword"
	checkAuthPasswordEntryType = "CheckAuthPassword"
)

// GenPassword 生成指定长度的密码
func (p *PasswordController) GenPassword() {

	length, err := strconv.Atoi(p.GetString("length"))

	if err != nil {
		p.Json(genAuthPasswordEntryType,
			fmt.Sprintf("%s", err), -1, logrus.ErrorLevel,
			LogMap{"errMsg": "can't convert parameter length to int, must be a int number"}, true)
	}

	name := p.GetString("name")
	charSet := strings.ToLower(p.GetString("charSet"))

	if name == "" {
		p.Json(genAuthPasswordEntryType, "", 0, logrus.InfoLevel,
			LogMap{"passowrd": brtool.GenRandomString(length, charSet)}, true)
	} else {
		m := make(map[string]string)
		names := strings.Split(name, ",")
		for _, name := range names {
			m[name] = brtool.GenRandomString(length, charSet)
		}
		p.Json(genAuthPasswordEntryType, "", 0, logrus.InfoLevel, LogMap{"password": m}, true)
	}
}

// GenAuthPassword 生成验证密码
func (p *PasswordController) GenAuthPassword() {

	username := p.GetString("username")
	email := strings.Split(p.GetString("email"), ",")

	if username == "" || len(email) == 0 {
		p.Json(genAuthPasswordEntryType, "username and email must be required", 1, logrus.InfoLevel, LogMap{"genAuthPassword": false}, true)
		return
	}
	if ok := common.GenPassword(username, email); ok {
		p.Json(genAuthPasswordEntryType, "", 0, logrus.InfoLevel, LogMap{"genAuthPassword": true}, true)
	} else {
		p.Json(genAuthPasswordEntryType, "generate auth password error", 1, logrus.ErrorLevel, LogMap{"genAuthPassword": false}, true)
	}
}

// CheckAuthPassword 密码验证
func (p *PasswordController) CheckAuthPassword() {
	username := p.GetString("username")
	unAuthPassword := p.GetString("password")

	if username == "" || unAuthPassword == "" {
		p.Json(checkAuthPasswordEntryType, "username and password must be provide", 1, logrus.ErrorLevel, LogMap{"auth": false}, true)
		return

	}
	if ok, _ := common.CheckPassword(username, unAuthPassword); ok {
		p.Json(checkAuthPasswordEntryType, "", 0, logrus.ErrorLevel, LogMap{"auth": true}, true)
	} else {
		p.Json(checkAuthPasswordEntryType, "password is not exist or expired", 1, logrus.ErrorLevel, LogMap{"auth": false}, true)
	}
}
