package controllers

import (
	"fmt"

	"github.com/braior/devops-api/common"
	"github.com/sirupsen/logrus"
)

var (
	twoStepAuthEntryType = "TwoStepAuth"
)

// Enable open two-step auth
func (t *TwoStepAuthController) Enable() {
	username := t.GetString("username")
	issuer := t.GetString("issuer")

	twoStepAuth := common.NewTwoStepAuth(username)
	twoStepAuth.Issuer = issuer
	twoStepAuth.Digits = common.TwoStepAuthDigits

	twoStepAuthInfo, err := twoStepAuth.Enable()
	if err != nil {
		t.Json(twoStepAuthEntryType, fmt.Sprintf("%s", err), 1, logrus.ErrorLevel, LogMap{"enable": "no", "username": username}, true)
		return
	}
	t.Json(twoStepAuthEntryType, "", 0, logrus.WarnLevel, twoStepAuthInfo, true)

}

// Disable ...
func (t *TwoStepAuthController) Disable() {
	username := t.GetString("username")
	twoStepAuth := common.NewTwoStepAuth(username)
	err := twoStepAuth.Disable()
	if err != nil {
		t.Json(twoStepAuthEntryType,
			fmt.Sprintf("%s", err), 1, logrus.ErrorLevel,
			LogMap{"disable": "yes", "username": username}, true)
		return
	}
	t.Json(twoStepAuthEntryType, "", 0, logrus.InfoLevel,
		LogMap{"disable": "no", "username": username}, true)
}

// Auth 验证用户输入的6位数字
func (t *TwoStepAuthController) Auth() {
	username := t.GetString("username")
	issuer := t.GetString("issuer")
	token := t.GetString("token")

	twoStepAuth := common.NewTwoStepAuth(username)
	twoStepAuth.Issuer = issuer
	isok, err := twoStepAuth.Auth(token)

	if err != nil {
		t.Json(twoStepAuthEntryType, "", 0, logrus.InfoLevel,
			LogMap{"username": username, "issuer": issuer, "auth": isok}, true)
		return
	}
	t.Json(twoStepAuthEntryType, fmt.Sprintf("%s", err), 1, logrus.ErrorLevel,
		LogMap{"username": username, "issuer": issuer, "auth": isok}, true)
}
