package main

import (
	_ "github.com/astaxie/beego/session/redis"

	"back/beego_api/models"
	"back/beego_api/utils/define"

	_ "back/beego_api/routers"
	"back/beego_api/services/base_service"
	_ "back/beego_api/utils/util"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"

	"time"
	"github.com/satori/go.uuid"
	"github.com/astaxie/beego/logs"
	"back/beego_api/utils/sentry"
)

func init() {
	sentry.Init() // sentry
	models.Init() // 模型
	corsHandler := func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Domain())
		ctx.Output.Header("Access-Control-Allow-Methods", "*")
		ctx.Request.Header.Add(define.TraceId, uuid.NewV4().String()) // trace id
	}
	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)
}

func main() {
	
	logger := logs.GetLogger()
	t1 := time.Now()
	for ii := 0; ii < 1000000; ii++ {
		logger.Println("xxxxx")
	}
	t2 := time.Since(t1).Nanoseconds()

	logger.Println(float64(t2) / float64(time.Second))

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

func destroy()  {
	
}