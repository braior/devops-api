package controllers

import (
	"github.com/astaxie/beego"
	"github.com/braior/devops-api/cmd"
	"github.com/braior/devops-api/utils"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// UniQueIDName 唯一id名称
	UniQueIDName = getUniqueIDName()
	// NeedTokenError 需要token错误提示信息
	NeedTokenError = "need DEVOPS-API-TOKEN header"
	// TokenAuthError 错误信息提示
	TokenAuthError = "DEVOPS-API-TOKEN auth failed"
	// ProhibitUseRootToken 不可以使用root token
	ProhibitUseRootToken = "Prohibit the use of root token"
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
// func (b *BaseController) log(msg LogMap) LogMap {

// 	if _, ok := msg["requestID"]; !ok {
// 		msg["requestID"] = b.Data[UniQueIDName]
// 	}

// 	if _, ok := msg["clientIP"]; !ok {
// 		msg["clientIP"] = b.Data["RemoteIP"]
// 	}

// 	if _, ok := msg["token"]; !ok {
// 		msg["token"] = b.Data["token"]
// 	}
// 	return msg
// }

// LogDebug ...
func (b *BaseController) LogDebug(message string, logMap LogMap) {
	// messageMap := b.log(msg)
	// if _, ok := messageMap["statusCode"]; !ok {
	// 	messageMap["statusCode"] = 0
	// }

	utils.Logger.Debug(logMap, message)
}

// LogInfo ...
func (b *BaseController) LogInfo(message string, logMap LogMap) {
	// messageMap := b.log(msg)
	// if _, ok := messageMap["statusCode"]; !ok {
	// 	messageMap["statusCode"] = 0
	// }
	utils.Logger.Info(logMap, message)
}

func (b *BaseController) LogWarn(message string, logMap LogMap) {
	// messageMap := b.log(msg)
	// if _, ok := messageMap["statusCode"]; !ok {
	// 	messageMap["statusCode"] = -1
	// }
	utils.Logger.Warn(logMap, message)
}

func (b *BaseController) LogError(message string, logMap LogMap) {
	// messageMap := b.log(msg)
	// if _, ok := messageMap["statusCode"]; !ok {
	// 	messageMap["statusCode"] = 1
	// }
	utils.Logger.Error(logMap, message)
}

func (b *BaseController) LogFatal(message string, logMap LogMap) {
	// messageMap := b.log(msg)
	// if _, ok := messageMap["statusCode"]; !ok {
	// 	messageMap["statusCode"] = -1
	// }
	utils.Logger.Fatal(logMap, message)
}

func (b *BaseController) Json(entryType string, errMsg string, statusCode int, logLevel logrus.Level, data interface{}, isLog bool) {
	responseData := map[string]interface{}{
		"entryType":  entryType,
		"requestID":  b.Data[UniQueIDName],
		"errMsg":     errMsg,
		"statusCode": statusCode,
		"data":       data,
	}

	b.Data["json"] = responseData
	b.ServeJSON()

	logMsg := make(map[string]interface{})

	logMsg["clientIP"] = b.Data["RemoteIP"]
	logMsg["token"] = b.Data["token"]
	logMsg["requestID"] = b.Data[UniQueIDName]
	logMsg["responseMsg"] = responseData

	if isLog {
		go func() {
			switch logLevel {
			case logrus.DebugLevel:
				b.LogDebug(entryType, logMsg)
			case logrus.InfoLevel:
				b.LogInfo(entryType, logMsg)
			case logrus.WarnLevel:
				b.LogWarn(entryType, logMsg)
			case logrus.ErrorLevel:
				b.LogError(entryType, logMsg)
			case logrus.FatalLevel:
				b.LogFatal(entryType, logMsg)
			}
			// if statusCode == 1 {
			// 	b.LogError(message, logMsg)
			// } else if statusCode == 0 {
			// 	b.LogInfo(message, logMsg)
			// } else if statusCode == -1 {
			// 	b.LogFatal(message, logMsg)
			// }
		}()
	}
}

// func (b *BaseController) Json(entryType, Msg string, statusCode int, logLevel logrus.Level, data interface{}, isLog bool) {
// 	b.json(entryType, Msg, statusCode, logLevel, data, isLog)
// }

// func (b *BaseController) JsonInfo(entryType, Msg string, data interface{}, isLog bool) {
// 	b.json(entryType, Msg, 0, data, isLog)
// }

// func (b *BaseController) JsonWarning(entryType, Msg string, data interface{}, isLog bool) {
// 	b.json(entryType, Msg, 0, data, isLog)
// }

// func (b *BaseController) JsonError(entryType, Msg string, data interface{}, isLog bool) {
// 	b.json(entryType, Msg, -1, data, isLog)
// }

// func (b *BaseController) JsonFatal(entryType, Msg string, data interface{}, isLog bool) {
// 	b.json(entryType, Msg, -1, data, isLog)
// }

// Prepare 覆盖beego.Controller的方法
func (b *BaseController) Prepare() {
	// 获取客户端IP
	b.Data["RemoteIP"] = b.Ctx.Input.IP()

	uniqueID := b.Ctx.Input.Header("UniQueIDName")
	if uniqueID == "" {
		uniqueID = uuid.NewV4().String()
	}

	b.Data[UniQueIDName] = uniqueID

	// 配置文件文件中启用了token功能,才验证token
	if viper.GetBool("security.enableToken") {
		// 获取 DEVOPS-API-TOKEN 头信息
		token := b.Ctx.Input.Header("DEVOPS-API-TOKEN")
		if token == "" {
			b.Json("JWToken Auth", NeedTokenError, 1, logrus.ErrorLevel, LogMap{}, true)
			b.StopRun()
		}
		b.Data["token"] = token

		// 验证 DEVOPS-API-TOKEN 是否有效
		jwtoken, err := cmd.NewToken()
		if err != nil {
			b.Json("JWToken Auth", TokenAuthError, 1, logrus.ErrorLevel, LogMap{}, true)
			b.StopRun()
		}

		// 验证是否是root token 不能使用root token
		isroot, err := jwtoken.IsRootToken(token)
		if err != nil {
			b.Json("JWToken Auth", TokenAuthError, 1, logrus.ErrorLevel, LogMap{}, true)
			b.StopRun()
		}
		if isroot {
			b.Json("JWToken Auth", ProhibitUseRootToken, 1, logrus.ErrorLevel, LogMap{}, true)
			b.StopRun()
		}

		_, err = jwtoken.IsTokenValid(token)
		if err != nil {
			b.Json("JWToken Auth", TokenAuthError, 1, logrus.ErrorLevel, LogMap{}, true)
			b.StopRun()
		}
	}
}

// VersionController 程序自身版本管理控制器
type VersionController struct {
	BaseController
}

// MD5Controller 程序自身版本管理控制器
type MD5Controller struct {
	BaseController
}

// PasswordController
type PasswordController struct {
	BaseController
}

// DingdingController
type DingdingController struct {
	BaseController
}

// EmailController
type EmailController struct {
	BaseController
}

// TwoStepAuthController
type TwoStepAuthController struct {
	BaseController
}
