package default_service

import (
	. "back/beego_api/services/base_service"
	"encoding/json"
)

// Operations about object
type DefaultController struct {
	BaseController
}

// @Title 欢迎信息
// @Description API 欢迎信息
// @Success 200 {string}
// @router / [get]
func (this *DefaultController) GetAll() {

	this.Data["json"] = Response{0, "success.", "API 1.0"}
	this.ServeJSON()
}

type TestController struct {
	BaseController
}

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// @Title PostTest
// @Description PostTest
// @Param   body  body default_service.User 用户信息
// @Success 200 {string}
// @router / [post]
func (this *TestController) PostTest() {
	var user User
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		this.Data["json"] = user
	} else {
		this.Data["json"] = err.Error()
	}

	this.ServeJSON()
}
