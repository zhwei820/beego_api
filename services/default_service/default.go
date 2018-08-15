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
// @router / [any]
func (o *DefaultController) GetAll() {

	o.Data["json"] = Response{0, "success.", "API 1.0"}
	o.ServeJSON()
}
