package main

import (
	"back/beego_api/models"
	"back/beego_api/utils/define"

	_ "back/beego_api/routers"
	"back/beego_api/services/base_service"
	_ "back/beego_api/utils/util"
	_ "github.com/gomodule/redigo/redis"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"

	"time"
	"github.com/satori/go.uuid"
)

func init() {
	models.Init()
	corsHandler := func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Domain())
		ctx.Output.Header("Access-Control-Allow-Methods", "*")
		ctx.Request.Header.Add(define.TraceId, uuid.NewV4().String()) // trace id
	}
	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)
}
func main() {
	debug, _ := beego.AppConfig.Bool("debug")
	if debug {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "./swagger"
		orm.Debug = true
	}
	orm.DefaultTimeLoc = time.UTC
	beego.ErrorController(&base_service.ErrorController{})
	beego.BConfig.ServerName = "snail server 1.0"

	beego.Run()
}
