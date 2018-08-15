package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["back/beego_api/services/default_service:DefaultController"] = append(beego.GlobalControllerRouter["back/beego_api/services/default_service:DefaultController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/beego_api/services/default_service:TestController"] = append(beego.GlobalControllerRouter["back/beego_api/services/default_service:TestController"],
		beego.ControllerComments{
			Method: "PostTest",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
