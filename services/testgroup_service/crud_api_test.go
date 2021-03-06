package testgroup_service_test

import (
	_ "back/beego_api/utils/gotests"

	"net/http"
	"net/http/httptest"
	"testing"
	_ "back/beego_api/routers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/bitly/go-simplejson"
	"fmt"
	"bytes"
	"back/beego_api/utils/random"
	"io"
)

func getRandomInput() io.Reader {
	str := random.RandNumString(6)

	return bytes.NewBuffer([]byte(fmt.Sprintf(`{
	                            "groupname": "%v","group": "%v","permission": "%v"
								}`,
								 str,  str,  str, 
								)))
}

func TestCreate(t *testing.T) {
	r, _ := http.NewRequest("POST", fmt.Sprintf("/api/testgroup/testgroup/?"),
		getRandomInput())

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("TestDetail", "code", w.Code, w.Body.String())

	var resId int

	Convey("detail\n", t, func() {
		Convey("status Code ", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("result Code ", func() {
			j, _ := simplejson.NewJson(w.Body.Bytes())
			code := j.Get("code").MustInt()
			So(code, ShouldEqual, 0)

			resId = j.Get("data").MustInt() // model_data.obj id

		})
	})

	testList(t)
	testDetail(t, resId)
	testUpdate(t, resId)

}
func testList(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/testgroup/testgroup?page=1&page_size=10", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("TestList", "code", w.Code, w.Body.String())

	Convey("list\n", t, func() {
		Convey("status Code ", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("result Code ", func() {
			j, _ := simplejson.NewJson(w.Body.Bytes())

			code := j.Get("code").MustInt()
			So(code, ShouldEqual, 0)
		})
	})

}

func testDetail(t *testing.T, resId int) {
	r, _ := http.NewRequest("GET", fmt.Sprintf("/api/testgroup/testgroup/%v?", resId), nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("TestDetail", "code", w.Code, w.Body.String())

	Convey("detail\n", t, func() {
		Convey("status Code ", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("result Code ", func() {
			j, _ := simplejson.NewJson(w.Body.Bytes())
			code := j.Get("code").MustInt()
			So(code, ShouldEqual, 0)
		})
	})
}

func testUpdate(t *testing.T, resId int) {

	r, _ := http.NewRequest("PUT", fmt.Sprintf("/api/testgroup/testgroup/%v?", resId),
		getRandomInput())

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("TestDetail", "code", w.Code, w.Body.String())

	Convey("detail\n", t, func() {
		Convey("status Code ", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("result Code ", func() {
			j, _ := simplejson.NewJson(w.Body.Bytes())
			code := j.Get("code").MustInt()
			So(code, ShouldEqual, 0)
		})
	})
}