package routers

import (
	"devops-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	apins := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSNamespace("/queryphone",
				beego.NSRouter("", &controllers.PhoneController{}),
			),
		),
	)
	beego.AddNamespace(apins)
}
