package routers

import (
	"github.com/astaxie/beego"
	"github.com/braior/devops-api/controllers"
)

func init() {

	apiNS := beego.NewNamespace("/apis",
		beego.NSNamespace("/v1",
			beego.NSNamespace("password",
				beego.NSRouter("/generation", &controllers.PasswordController{}, "get:GenPassword"),
				beego.NSRouter("/genAuthPassword", &controllers.PasswordController{}, "get:GenAuthPassword"),
				beego.NSRouter("/checkAuthPassword", &controllers.PasswordController{}, "post:CheckAuthPassword")),
			beego.NSNamespace("/version",
				beego.NSRouter("", &controllers.VersionController{}),
			),
			beego.NSNamespace("/md5",
				beego.NSRouter("", &controllers.MD5Controller{}),
			),
			beego.NSNamespace("/sendmsg",
				beego.NSNamespace("/dingding",
					beego.NSRouter("", &controllers.DingTalkController{}, "post:SendMessage"),
				),
				beego.NSNamespace("/mail",
					beego.NSRouter("", &controllers.EmailController{}, "post:SendMessage"),
				),
			),
			beego.NSNamespace("/twostepauth",

				beego.NSRouter("/enable", &controllers.TwoStepAuthController{}, "get:Enable"),
				beego.NSRouter("/disable", &controllers.TwoStepAuthController{}, "get:Disable"),
				beego.NSRouter("/auth", &controllers.TwoStepAuthController{}, "post:Auth"),
			),
		),
	)
	beego.AddNamespace(apiNS)
}
