package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["back/beego_api/services/user_service:UserController"] = append(beego.GlobalControllerRouter["back/beego_api/services/user_service:UserController"],
		beego.ControllerComments{
			Method: "Auth",
			Router: `/auth`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/beego_api/services/user_service:UserController"] = append(beego.GlobalControllerRouter["back/beego_api/services/user_service:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/beego_api/services/user_service:UserController"] = append(beego.GlobalControllerRouter["back/beego_api/services/user_service:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
