package routers

import "github.com/astaxie/beego"

func init(){
	apins :=beego.NewNamespace("/api",
	beego.NSNamespace("/queryphone",
	beego.NSRouter("",&con)
}