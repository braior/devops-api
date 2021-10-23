package controllers

import (
	"github.com/braior/devops-api/common"
	"github.com/sirupsen/logrus"
)

// GetVersion 获取程序版本信息
func (v *VersionController) Get() {
	v.Json("Get APP Version", "", 0, logrus.InfoLevel, common.GetVersion(), true)
	// v.Json("Get APP Version", "", -1, logrus.ErrorLevel, common.GetVersion(), true)
	//v.JsonError("Get APP Version", "error", common.GetVersion(), true)
	//v.JsonFatal("Get APP Version", "fatal", common.GetVersion(), true)
}
