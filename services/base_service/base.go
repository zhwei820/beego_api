package base_service

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) WriteJson(jsonData interface{})  {
	this.Data["json"] = jsonData
	this.ServeJSON()
}

func (this *BaseController) WriteJsonWithCode(code int, jsonData interface{})  {
	this.Ctx.ResponseWriter.WriteHeader(code)
	this.Data["json"] = jsonData
	this.ServeJSON()
}

//Response 结构体
type Response struct {
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

//Response 结构体
type ErrResponse struct {
	Errcode int         `json:"errcode"`
	Errmsg  interface{} `json:"errmsg"`
}
