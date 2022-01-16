package controllers

import (
	"fmt"

	"github.com/braior/devops-api/common"
	"github.com/sirupsen/logrus"
)

var (
	dingTalkEnterType = "SendDingTalkMessage"
)

// SendMessage 发送钉钉消息
func (d *DingTalkController) SendMessage() {
	msgType := d.GetString("msgType")
	msg := d.GetString("msg")
	title := d.GetString("title")
	robotURL := d.GetString("url")

	ok, err := common.SendByDingTalkRobot(msgType, msg, title, robotURL)
	if err != nil || !ok {
		d.Json(dingTalkEnterType, fmt.Sprintf("%s", err), 1, logrus.ErrorLevel, LogMap{"result": "send ding talk message failed"}, true)
		return
	}

	d.Json(dingTalkEnterType, "", 0, logrus.InfoLevel, LogMap{"result": "send ding talk message succeed"}, true)
}
