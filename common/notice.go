package common

import (
	"fmt"

	"github.com/braior/brtool"
	"github.com/spf13/viper"
	"github.com/braior/devops-api/utils"
)

// SendByDingTalkRobot 通过钉钉发送消息通知
func SendByDingTalkRobot(messageType, message, title, robotURL string) (bool, error) {
	var url string
	if robotURL == "" {
		url = viper.GetString("notice.dingTalkRobot")
	} else {
		url = robotURL
	}

	dingTalk := &brtool.DingTalkClient{
		RobotURL: url,
		MsgInfo: &brtool.DingTalkMessage{
			Type:    messageType,
			Message: message,
			Title:   title,
		},
	}
	ok, err := dingTalk.SendMessage()
	if err != nil {
		dingFields := map[string]interface{}{
			"entryType":     "DingTalkRobot",
			"dingTalkRobot": url,
		}
		utils.Logger.Fatal(dingFields,fmt.Sprintf("发送钉钉通知失败了: %s",err))
		//fmt.Printf("noticd %v 发送钉钉通知失败了: %s", dingFields, err)
		return false, err
	}
	return ok, nil
}
