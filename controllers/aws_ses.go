package controllers

import (
	"devops-api/common"
	"fmt"
	"path"
	"strings"
)

var (
	mailAWSEntryType = "AWS SES SendMail"
)

// SendMail aws邮件送
func (a *AWSEMailControler) SendMail() {
	subject := a.GetString("subject")
	content := a.GetString("content")
	contentType := a.GetString("type")
	to := a.GetString("to")
	// cc := a.GetString("cc")

	isattach, err := a.GetBool("isattach")
	if err != nil {
		isattach = false
	}

	var attachFileName string
	if isattach {
		f, h, err := a.GetFile("attach")
		if err != nil {
			errs := fmt.Sprintf("获取文件失败：%s", err)
			a.LogError(mailAWSEntryType, StringMap{"errmsg": errs})
			a.JsonError(mailAWSEntryType, errs, StringMap{}, false)
			return
		}

		defer f.Close()
		attachFileName = path.Join(common.UploadPath, h.Filename)
		a.SaveToFile("attach", attachFileName)
	}

	if to == "" {
		a.JsonError(mailAWSEntryType, "发送失败: 收件人不能为空", StringMap{"result": "send fail"}, false)
		return
	}

	if contentType == "" {
		contentType = "text/plain"
	}

	var toMail []string
	toMail = strings.Split(to, ",")

	// var ccMail []string
	// if cc == "" {
	// 	ccMail = []string{}
	// } else {
	// 	ccMail = strings.Split(cc, ",")
	// }

	// 带附件,抄送发送
	// _, err = common.SendByAWSEmail(subject, content, contentType, attachFileName, toMail, ccMail)
	// 不带附件和抄送发送
	_, err = common.SendByAWSEmail(subject, content, contentType, toMail)

	if err == nil {
		a.JsonOK(mailAWSEntryType, StringMap{"result": "send ok"}, true)
	}
	a.JsonError(mailAWSEntryType, fmt.Sprintf("errors: %s", err), StringMap{"result": "send fail"}, false)
}

// SendMail2 aws邮件送
func (a *AWSEMailControler) SendMail2() {
	subject := a.GetString("subject")
	content := a.GetString("content")
	contentType := a.GetString("type")
	to := a.GetString("to")
	cc := a.GetString("cc")

	isattach, err := a.GetBool("isattach")
	if err != nil {
		isattach = false
	}

	var attachFileName string
	if isattach {
		f, h, err := a.GetFile("attach")
		if err != nil {
			errs := fmt.Sprintf("获取文件失败：%s", err)
			a.LogError(mailAWSEntryType, StringMap{"errmsg": errs})
			a.JsonError(mailAWSEntryType, errs, StringMap{}, false)
			return
		}

		defer f.Close()
		attachFileName = path.Join(common.UploadPath, h.Filename)
		a.SaveToFile("attach", attachFileName)
	}

	if to == "" {
		a.JsonError(mailAWSEntryType, "发送失败: 收件人不能为空", StringMap{"result": "send fail"}, false)
		return
	}

	if contentType == "" {
		contentType = "text/plain"
	}

	var toMail []*string
	toMailSplit := strings.Split(to, ",")
	for _,v:=range toMailSplit{
		toMail=append(toMail, &v)
	}

	var ccMail []*string
	if cc == "" {
		ccMail = []*string{}
	} else {
		ccMailSplit := strings.Split(cc, ",")
		for _,v :=range ccMailSplit{
			ccMail=append(ccMail, &v)
		}
	}

	// 带附件,抄送发送
	_, err = common.SendByAWSEmail2(subject, content, contentType, attachFileName, toMail, ccMail)
	// 不带附件和抄送发送
	// _, err = common.SendByAWSEmail(subject, content, contentType, toMail)

	if err == nil {
		a.JsonOK(mailAWSEntryType, StringMap{"result": "send ok"}, true)
	}
	a.JsonError(mailAWSEntryType, fmt.Sprintf("errors: %s", err), StringMap{"result": "send fail"}, false)
}