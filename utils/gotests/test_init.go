package gotests

import (
	"runtime"
	"path/filepath"
	"github.com/astaxie/beego"
	"back/beego_api/utils"
	_ "github.com/astaxie/beego/session/redis"
)

func init() {
	_, file, _, _ := runtime.Caller(1)

	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../../")))
	exist, _ := utils.PathExists(filepath.Join(apppath, "/conf/"))
	if !exist {
		apppath, _ = filepath.Abs(filepath.Dir(filepath.Join(file, "../../../")))
		exist, _ := utils.PathExists(filepath.Join(apppath, "/conf/"))
		if !exist {
			apppath, _ = filepath.Abs(filepath.Dir(filepath.Join(file, "../../../../")))

		}
	}
	beego.TestBeegoInit(apppath)
}

