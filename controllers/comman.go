package controllers

import (
	"github.com/astaxie/beego"
	"github.com/braior/devops-api/common"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

var (
	// UniQueIDName 唯一id名称
	UniQueIDName = getUniqueIDName()
	// NeedTokenError 需要token错误提示信息
	NeedTokenError = "need DEVOPS-API-TOKEN header"
	// TokenAuthError 错误信息提示
	TokenAuthError = "DEVOPS-API-TOKEN auth failed"
)

// LogMap 记录日志数据
type LogMap map[string]interface{}

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
}

func getUniqueIDName() string {
	// 从配置文件中获取 RequestID或者TraceID,如果配置文件中没有配置默认就是 RequestId
	uniqueIDName := viper.GetString("app.requestID")
	if uniqueIDName == "" {
		uniqueIDName = "RequestID"
	}
	// uniqueIDName := beego.AppConfig.String("uniqueIDName")
	// if uniqueIDName == "" {
	// 	uniqueIDName = "requestID"
	// }
	return uniqueIDName
}

// log 记录body中的header信息到日志中
func (b *BaseController) log(msg LogMap) LogMap {
	if _, ok := msg["requestID"]; !ok {
		msg["requestID"] = b.Data[UniQueIDName]
	}

	if _, ok := msg["clientIP"]; !ok {
		msg["clientIP"] = b.Data["RemoteIP"]
	}

	if _, ok := msg["token"]; !ok {
		msg["token"] = b.Data["token"]
	}
	return msg
}

// LogInfo ...
func (b *BaseController) LogInfo(entryType string, msg LogMap) {
	message := b.log(msg)
	if _, ok := msg["statusCode"]; !ok {
		message["statusCode"] = 0
	}
	common.Logger.Info(message, entryType)
}

func (b *BaseController) LogError(entryType string, msg LogMap) {
	message := b.log(msg)
	if _, ok := message["statusCode"]; !ok {
		message["statusCode"] = 1
	}
	common.Logger.Error(message, entryType)
}

func (b *BaseController) json(entryType, errMsg string, statusCode int, data interface{}, isLog bool) {
	msg := map[string]interface{}{
		"entryType":  entryType,
		"requestID":  b.Data[UniQueIDName],
		"errMsg":     errMsg,
		"statusCode": statusCode,
		"data":       data,
	}
	b.Data["json"] = msg
	b.ServeJSON()

	msg["clientIP"] = b.Data["RemotIP"]
	msg["token"] = b.Data["token"]

	if isLog {
		go func() {
			if statusCode == 1 {
				b.LogError(entryType, msg)
			} else {
				b.LogInfo(entryType, msg)
			}
		}()
	}
}

func (b *BaseController) JsonError(entryType, errMsg string, data interface{}, isLog bool) {
	b.json(entryType, errMsg, 1, data, isLog)
}

func (b *BaseController) JsonOK(entryType, errMsg string, data interface{}, isLog bool) {
	b.json(entryType, "", 0, data, isLog)
}

func (b *BaseController) Prepare() {
	// 获取客户端IP
	b.Data["RemoteIP"] = b.Ctx.Input.IP()

	uniqueID := b.Ctx.Input.Header(UniQueIDName)
	if uniqueID == "" {
		uniqueID = uuid.NewV4().String()
	}

	b.Data[UniQueIDName] = uniqueID
}

// VersionController 程序自身版本管理控制器
type VersionController struct {
	BaseController
}
