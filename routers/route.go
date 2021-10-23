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
			),
			beego.NSNamespace("/version",
				beego.NSRouter("", &controllers.VersionController{}),
			),
			beego.NSNamespace("/md5",
				beego.NSRouter("", &controllers.MD5Controller{}),
			),
			beego.NSNamespace("/sendmsg",
				beego.NSNamespace("/dingding",
					beego.NSRouter("", &controllers.DingdingController{}, "post:SendMessage"),
				),
			),
		),
	)

	beego.AddNamespace(apiNS)
}
