package controllers

import (
	"github.com/braior/devops-api/common"
	"github.com/sirupsen/logrus"
)

// Get 获取程序版本信息
func (v *VersionController) Get() {
	v.Json("Get APP Version", "", 0, logrus.InfoLevel, common.GetVersion(), true)
}
