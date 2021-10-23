package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/braior/brtool"
	"github.com/sirupsen/logrus"
)

var genPasswordEntryType = "GenPassword"

// GenPassword 生成指定长度的密码
func (p *PasswordController) GenPassword() {

	length, err := strconv.Atoi(p.GetString("length"))

	if err != nil {
		p.Json(genPasswordEntryType,
			fmt.Sprintf("%s", err), -1, logrus.ErrorLevel,
			LogMap{"errMsg": "can't convert parameter length to int, must be a int number"}, true)
	}

	name := p.GetString("name")
	charSet := strings.ToLower(p.GetString("charSet"))

	if name == "" {
		p.Json(genPasswordEntryType, "", 0, logrus.InfoLevel,
			LogMap{"passowrd": brtool.GenRandomString(length, charSet)}, true)
	} else {
		m := make(map[string]string)
		names := strings.Split(name, ",")
		for _, name := range names {
			m[name] = brtool.GenRandomString(length, charSet)
		}
		p.Json(genPasswordEntryType, "", 0, logrus.InfoLevel, LogMap{"password": m}, true)
	}
}
