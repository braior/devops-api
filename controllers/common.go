package controllers

import (
	"devops-api/common"

	"github.com/astaxie/beego"
)

func getUniqueIDName() string {
	// 从配置文件中获取 RequestID或者TraceID,如果配置文件中没有配置默认就是 RequestId
	uniqueIDName := beego.AppConfig.String("uniqueIDName")
	if uniqueIDName == "" {
		uniqueIDName = "RequestID"
	}
	return uniqueIDName
}

var (
	UniQueIDName   = getUniqueIDName()
	NeedTokenError = "need DEVOPS-API-TOKEN header"
	TokenAuthError = "DEVOPS-API-TOKEN auth fail"
)

type StringMap map[string]interface{}

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
}

func (b *BaseController) log(msg StringMap) StringMap {
	if _, ok := msg["requestId"]; !ok {
		msg["requestId"] = b.Data[UniQueIDName]
	}

	if _, ok := msg["clientIP"]; !ok {
		msg["clientIP"] = b.Data["RemoteIP"]
	}

	if _, ok := msg["token"]; !ok {
		msg["token"] = b.Data["token"]
	}
	return msg
}

func (b *BaseController) LogInfo(entryType string, msg StringMap) {
	message := b.log(msg)
	if _, ok := msg["statuscode"]; !ok {
		message["statuscode"] = 0
	}
	common.GetLogger().Info(message, entryType)
}

func (b *BaseController) LogError(entryType string, msg StringMap) {
	message := b.log(msg)
	if _, ok := msg["statuscode"]; !ok {
		message["statuscode"] = 1
	}
	common.GetLogger().Error(message, entryType)
}

// json 将请求的信息
func (b *BaseController) json(entryType, errmsg string, statuscode int, data interface{}, isLog bool) {
	msg := map[string]interface{}{
		"entryType":  entryType,
		"requestID":  b.Data[UniQueIDName],
		"errmsg":     errmsg,
		"statuscode": statuscode,
		"data":       data,
	}

	b.Data["json"] = msg
	b.ServeJSON()

	msg["clientIP"] = b.Data["RemoteIP"]
	msg["token"] = b.Data["token"]

	if isLog {
		go func() {
			if statuscode == 1 {
				b.LogError(entryType, msg)
			} else {
				b.LogInfo(entryType, msg)
			}
		}()
	}
}

func (b *BaseController) JsonError(entryType, errmsg string, data interface{}, isLog bool) {
	b.json(entryType, errmsg, 1, data, isLog)
}

func (b *BaseController) JsonOK(entryType string, data interface{}, isLog bool) {
	b.json(entryType, "", 0, data, isLog)
}

// PhoneController 手机归属地查询
type PhoneController struct {
	BaseController
}
