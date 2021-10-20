package controllers

// import (
// 	"fmt"
// 	"path"
// 	"strings"

// 	"github.com/braior/devops-api/cmd"
// 	"github.com/sirupsen/logrus"
// )

// var (
// 	mailEntryType = SendMail
// )

// // SendMail 发送邮件
// func (e *EmailController) SendMail() {
// 	subject := e.GetString("subject")
// 	content := e.GetString("content")
// 	contentType := e.GetString("type")
// 	to := e.GetString("to")
// 	cc := e.GetString("cc")

// 	isAttach, err := e.GetBool("attach")
// 	if err != nil {
// 		isAttach = false
// 	}

// 	var attachFilename string
// 	if isAttach {
// 		f, h, err := e.GetFile("attach")
// 		if err != nil {
// 			errs := fmt.Sprintf("获取文件失败：%s", err)
// 			e.LogError(mailEntryType, LogMap{"errmsg": errs})
// 			e.json(mailEntryType, errs, -1, logrus.ErrorLevel, LogMap{}, true)
// 		}
// 		defer f.Close()
// 		attachFilename = path.Join(cmd.UploadPath, h.Filename)
// 		e.SaveToFile("attach", attachFilename)
// 	}

// 	if subject == "" || content == "" {
// 		e.json(mailEntryType, "发送邮件失败：主题或者内容不能为空", -1, logrus.ErrorLevel, LogMap{"result": "send mail failed"}, true)
// 		return
// 	}

// 	if to == "" {
// 		e.json(mailEntryType, "发送邮件失败：收件人不能为空", -1, logrus.ErrorLevel, LogMap{"result": "send mail failed"}, true)
// 		return
// 	}

// 	if contentType == "" {
// 		contentType = "text/plain"
// 	}

// 	toMail := strings.Split(to, ",")

// 	var ccMail []string
// 	if cc == "" {
// 		ccMail = []string{}
// 	} else {
// 		ccMail = strings.Split(cc, ",")
// 	}

// 	_,err:=common.
// }
