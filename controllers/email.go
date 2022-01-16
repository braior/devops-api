package controllers

import (
	"fmt"
	"path"
	"strings"

	"github.com/braior/devops-api/cmd"
	"github.com/braior/devops-api/common"
	"github.com/sirupsen/logrus"
)

var (
	mailEntryType = "SendMail"
)

// SendMessage 发送邮件
func (e *EmailController) SendMessage() {
	subject := e.GetString("subject")
	content := e.GetString("content")
	contentType := e.GetString("type")
	to := e.GetString("to")
	cc := e.GetString("cc")

	isAttach, err := e.GetBool("isAttach")
	if err != nil {
		isAttach = false
	}

	var attachFilename string
	if isAttach {
		f, h, err := e.GetFile("attach")
		if err != nil {
			errs := fmt.Sprintf("获取文件失败：%s", err)
			e.LogError(mailEntryType, LogMap{"errMsg": errs})
			e.Json(mailEntryType, errs, -1, logrus.ErrorLevel, LogMap{}, true)
		}
		defer f.Close()
		attachFilename = path.Join(cmd.UploadPath, h.Filename)
		err = e.SaveToFile("attach", attachFilename)
		if err != nil {
			e.LogError(mailEntryType, LogMap{"errMsg": err})
		}
	}

	if subject == "" || content == "" {
		e.Json(mailEntryType, "发送邮件失败：主题或者内容不能为空", 1, logrus.ErrorLevel, LogMap{"result": "send mail failed"}, true)
		return
	}

	if to == "" {
		e.Json(mailEntryType, "发送邮件失败：收件人不能为空", 1, logrus.ErrorLevel, LogMap{"result": "send mail failed"}, true)
		return
	}

	if contentType == "" {
		contentType = "text/plain"
	}

	toMail := strings.Split(to, ",")

	var ccMail []string
	if cc == "" {
		ccMail = []string{}
	} else {
		ccMail = strings.Split(cc, ",")
	}

	_, err = common.SendByEmail(subject, content, contentType, attachFilename, toMail, ccMail)
	if err == nil {
		e.Json(mailEntryType, "", 0, logrus.InfoLevel, LogMap{"result": "send ok"}, true)
		return
	}
	e.Json(mailEntryType, fmt.Sprintf("error: %s", err), 1, logrus.ErrorLevel, LogMap{"result": "send fail"}, true)
}
