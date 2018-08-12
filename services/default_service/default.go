package default_service

import (
	"back/beego_api/services/base_service"
)
// Operations about object
type DefaultController struct {
	base_service.BaseController
}

// @Title 欢迎信息
// @Description API 欢迎信息
// @Success 200 {string}
// @router / [any]
func (o *DefaultController) GetAll() {
	o.Data["json"] = base_service.Response{0, "success.", "API 1.0"}
	o.ServeJSON()
}
