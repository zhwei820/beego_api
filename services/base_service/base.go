package base_service

import (
	"github.com/astaxie/beego"
	"back/beego_api/utils/define"
	"encoding/json"
	"back/beego_api/utils"
	"back/beego_api/utils/util"
	"github.com/rs/zerolog"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) GetLogger() zerolog.Logger {
	// zero log with trace id
	return util.Logger.With().Str("trace_id", this.Ctx.Request.Header.Get(define.TraceId)).Logger()
}

func (this *BaseController) GetJson(ob interface{}) (error) {
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, ob); err == nil {
		err := utils.Validate(ob) // 根据struct tag验证
		if err != nil {
			return err
		}

		return nil
	}
	return err
}

func (this *BaseController) WriteJson(jsonData interface{}) {
	this.Data["json"] = jsonData
	this.ServeJSON()
}

func (this *BaseController) WriteJsonWithCode(code int, jsonData interface{}) {
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
