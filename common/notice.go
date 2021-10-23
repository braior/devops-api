package common

import (
	"fmt"

	"github.com/braior/brtool"
	"github.com/braior/devops-api/utils"
	"github.com/spf13/viper"
)

// SendByDingTalkRobot 通过钉钉发送消息通知
func SendByDingTalkRobot(messageType, message, title, robotURL string) (bool, error) {
	var url string
	if robotURL == "" {
		url = viper.GetString("notice.dingding.dingTalkRobot")
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
		utils.Logger.Error(dingFields, fmt.Sprintf("发送钉钉通知失败了: %s", err))
		//fmt.Printf("noticd %v 发送钉钉通知失败了: %s", dingFields, err)
		return false, err
	}
	return ok, nil
}

func SendByEmail(subject, content, contentType, attach string, to, cc []string) (bool, error) {
	host := viper.GetString("notice.email.host")
	port := viper.GetInt("notice.email.port")
	username := viper.GetString("notice.email.username")
	password := viper.GetString("notice.email.password")

	message := brtool.NewEmailMessage(username, subject, contentType, content, attach, to, cc)
	emailClient := brtool.NewEmailClient(host, username, password, port, message)
	ok, err := emailClient.SendMessage()

	if err != nil {
		emailFields := map[string]interface{}{
			"entryType": "SendMail",
			"mail": map[string]interface{}{
				"Username":    username,
				"Host":        host,
				"Port":        port,
				"ContentType": contentType,
				"Attach":      attach,
				"To":          to,
				"Cc":          cc,
			},
		}
		utils.Logger.Error(emailFields, fmt.Sprintf("发送Email通知失败了: %s", err))
		return false, err
	}
	return ok, nil
}
