package util

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var Logger *log.Logger

func init() {
	path := filepath.Join("data", "log")
	os.MkdirAll(path, 0777)

	logName := beego.AppConfig.DefaultString("log_name", "example.log")
	logs.EnableFuncCallDepth(true)
	logs.Async(1e4)
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"./data/log/%s","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`, logName))
	if beego.AppConfig.String("runmode") == "dev" {
		logs.SetLogger("console")
	}
	Logger = logs.GetLogger()
}
