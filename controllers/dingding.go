package controllers

import (
	"fmt"

	"github.com/braior/devops-api/common"
	"github.com/sirupsen/logrus"
)

var (
	dingdingEnterType = "SendDingdingMessage"
)

// SendMessage 发送钉钉消息
func (d *DingdingController) SendMessage() {
	msgType := d.GetString("msgType")
	msg := d.GetString("msg")
	title := d.GetString("title")
	robotURL := d.GetString("url")

	ok, err := common.SendByDingTalkRobot(msgType, msg, title, robotURL)
	if err != nil || !ok {
		d.json(dingdingEnterType, fmt.Sprintf("%s", err), 1, logrus.ErrorLevel, LogMap{"result": "send dingding message failed"}, true)
		return
	}

	d.json(dingdingEnterType, "", 0, logrus.InfoLevel, LogMap{"result": "send dingtalk message succeed"}, true)
}
