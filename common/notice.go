package common

import (
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"

	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/hltool"
)

// SendByDingTalkRobot 通过钉钉发送消息通知
func SendByDingTalkRobot(messageType, message, title, robotURL string) (bool, error) {
	var url string
	if robotURL == "" {
		url = beego.AppConfig.String("DingTalkRobot")
	} else {
		url = robotURL
	}

	dingtalk := &hltool.DingTalkClient{
		RobotURL: url,
		Message: &hltool.DingTalkMessage{
			Type:    messageType,
			Message: message,
			Title:   title,
		},
	}
	ok, err := hltool.SendMessage(dingtalk)
	if err != nil {
		dingField := map[string]interface{}{
			"entryType":     "DingTalkRobot",
			"dingTalkRobot": url,
		}

		Logger.Error(dingField, fmt.Sprintf("发送钉钉通知失败了: %s", err))
		return false, err
	}

	return ok, nil
}

// SendByEmail 通过Email发送消息通知
func SendByEmail(subject, content, contentType, attach string, to, cc []string) (bool, error) {
	username := beego.AppConfig.String("email::username")
	host := beego.AppConfig.String("email::host")
	password := beego.AppConfig.String("email::password")
	port, err := beego.AppConfig.Int("email::port")
	if err != nil {
		confFields := map[string]interface{}{
			"entryType": "Parse email Configure File",
		}
		Logger.Error(confFields, fmt.Sprintf("从配置文件中解析邮件端口失败: %s", err))
		return false, err
	}

	message := hltool.NewEmailMessage(username, subject, contentType, content, attach, to, cc)
	email := hltool.NewEmailClient(host, username, password, port, message)
	ok, err := hltool.SendMessage(email)
	if err != nil {
		emailFields := map[string]interface{}{
			"entryType": "SendMial",
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
		Logger.Error(emailFields, fmt.Sprintf("发送邮件失败了: %s", err))
		return false, err
	}
	return ok, nil
}

// SendByAWSEmail 通过aws Email发送消息通知
func SendByAWSEmail(subject, content, contentType string, to []string) (bool, error) {
	host := beego.AppConfig.String("aws_ses_email::host")
	sendername := beego.AppConfig.String("aws_ses_email::sendername")
	smtpuser := beego.AppConfig.String("aws_ses_email::smtpuser")
	smtpPass := beego.AppConfig.String("aws_ses_email::smtpPass")

	sender := beego.AppConfig.String("aws_ses_email::sender")
	// configurationSet := beego.AppConfig.String("aws_ses_email::configurationSet")
	// awsRegion := beego.AppConfig.String("aws_ses_email::awsRegion")

	HtmlBody := "<h1>Amazon SES Test Email (适用于 Go 的 AWS 开发工具包)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>适用于 Go 的 AWS 开发工具包</a>.</p>" + content

	// The character encoding for the email.
	// CharSet := "UTF-8"

	//The email body for recipients with non-HTML email clients.
	TextBody := "This email was sent with Amazon SES using the 适用于 Go 的 AWS 开发工具包."

	// The tags to apply to this message. Separate multiple key-value pairs
	// with commas.
	// If you comment out or remove this variable, you will also need to
	// comment out or remove the header on line 80.
	// Tags := "genre=test,genre2=test2"

	port, err := beego.AppConfig.Int("aws_ses_email::port")
	if err != nil {
		confFields := map[string]interface{}{
			"entryType": "Parse aws_ses Configure File",
		}
		Logger.Error(confFields, fmt.Sprintf("从配置文件中解析邮件端口失败: %s", err))
		return false, err
	}

	// 设置系统环境变量，为aws_ses发送邮件做准备
	/*
		awsAccessKey := beego.AppConfig.String("aws_ses_email::smtpUser")
		awsSecretKey := beego.AppConfig.String("aws_ses_email::smtpPass")
		if awsAccessKey == "" || awsSecretKey == "" {
			confFields := map[string]interface{}{
				"entryType": "Parse aws_ses Configure File",
			}
			Logger.Error(confFields, fmt.Sprintf("从配置文件中解析awsAccessKey|awsSecretKey失败"))
			return false, errors.New("从配置文件中解析awsAccessKey|awsSecretKey失败")
		}

		err := os.Setenv("AWS_ACCESS_KEY_ID", awsAccessKey)
		if err != nil {
			confFields := map[string]interface{}{
				"entryType": "Set AWS_SES ENV",
			}
			Logger.Error(confFields, fmt.Sprintf("设置AWS_ACCESS_KEY_ID环境变量失败：%s", err))
		}

		err = os.Setenv("AWS_SECRET_ACCESS_KEY", awsSecretKey)
		if err != nil {
			confFields := map[string]interface{}{
				"entryType": "Set AWS_SES ENV",
			}
			Logger.Error(confFields, fmt.Sprintf("设置AWS_SECRET_ACCESS_KEY环境变量失败：%s", err))
		}
	*/

	// Create a new message.
	m := gomail.NewMessage()

	// Set the main email part to use HTML.
	m.SetBody("text/html", HtmlBody)

	// Set the alternative part to plain text.
	m.AddAlternative("text/plain", TextBody)

	str := strings.Replace(strings.Trim(fmt.Sprint(to), "[]"), " ", ",", -1)

	// Construct the message headers, including a Configuration Set and a Tag.
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(sender, sendername)},
		"To":      {str},
		"Subject": {subject},
		// Comment or remove the next line if you are not using a configuration set
		// "X-SES-CONFIGURATION-SET": {ConfigSet},
		// Comment or remove the next line if you are not using custom tag
		//"X-SES-MESSAGE-TAGS": {Tags},

	})

	// Send the email.
	d := gomail.NewPlainDialer(host, port, smtpuser, smtpPass)

	// Display an error message if something goes wrong; otherwise,
	// display a message confirming that the message was sent.
	// Assemble the email.
	if err := d.DialAndSend(m); err != nil {
		emailFields := map[string]interface{}{
			"entryType": "SendMial",
			"mail": map[string]interface{}{
				"Username":    "",
				"Host":        sender,
				"Port":        "",
				"ContentType": contentType,
				"Attach":      "",
				"To":          to,
				// "Cc":          cc,
			},
		}
		Logger.Error(emailFields, fmt.Sprintf("发送邮件失败了: %s", err))
		return false, err
	}
	return true, nil
}
