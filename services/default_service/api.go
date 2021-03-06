package default_service

import (
	. "back/beego_api/services/base_service"
)

// Operations about object
type DefaultController struct {
	BaseController
}

// @Title 欢迎信息
// @Description API 欢迎信息
// @Success 200 {string}
// @router / [get]
func (this *DefaultController) ApiGetAll() {
	this.WriteJson(Response{0, "success.", "API 1.0"})
}

func (this *DefaultController) Html() {
	dir := this.Ctx.Input.Param(":dir")
	html := this.Ctx.Input.Param(":html")

	this.Layout = "templates/base/userinfo.html"
	this.TplName = "templates/" + dir + "/" + html + ".html"
}
