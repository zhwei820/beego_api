package routers

import (
	"github.com/astaxie/beego"
	"back/beego_api/services/user_service"
	"back/beego_api/services/default_service"
)
// @APIVersion 1.0.0
// @Description beego Test API
// 使用注释路由
// @SecurityDefinition jwt apiKey Authorization header header
func init() {
	beego.Router("/", &default_service.DefaultController{}, "*:GetAll")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&user_service.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
