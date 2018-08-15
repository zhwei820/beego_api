package base_service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"back/beego_api/utils/define"
	"log"
	"encoding/json"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) GetLogger() *log.Logger {
	return logs.GetLogger(this.Ctx.Request.Header.Get(define.TraceId))
}

func (this *BaseController) GetJson(ob interface{}) (error) {
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {
		valid := validation.Validation{}
		b, err := valid.Valid(&ob)
		if err != nil {
			return err
		}
		if !b {
			// validation does not pass
			for _, err := range valid.Errors {
				return err
			}
		}
		return nil
	}
	return err
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
