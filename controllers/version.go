package controllers

import (
	"github.com/braior/devops-api/common"
)

// GetVersion 获取程序版本信息
func (v *VersionController) Get() {
	v.JsonOK("Get APP Version", "", common.GetVersion(), true)
}
