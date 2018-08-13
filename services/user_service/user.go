package user_service

import (
	"fmt"
	"back/beego_api/models"
	"back/beego_api/utils"
	"strings"
	. "back/beego_api/services/base_service"
)

var (
	ErrPhoneIsRegis     = ErrResponse{422001, "手机用户已经注册"}
	ErrNicknameIsRegis  = ErrResponse{422002, "用户名已经被注册"}
	ErrNicknameOrPasswd = ErrResponse{422003, "账号或密码错误。"}
)

type UserController struct {
	BaseController
}
type LoginToken struct {
	User  models.StaffUser `json:"user"`
	Token string           `json:"token"`
}

// @Title 注册新用户
// @Description 用户注册
// @Param	phone		formData 	string	true 		"用户手机号"
// @Param	nickname	formData 	string	true		"用户昵称"
// @Param	password	formData 	string	true		"密码"
// @Success 200 {object} models.StaffUser
// @Failure 403 参数错误：缺失或格式错误
// @Faulure 422 已被注册
// @router /reg [post]
func (this *UserController) Register() {
	phone := this.GetString("phone")
	email := this.GetString("email")
	nickname := this.GetString("nickname")
	password := this.GetString("password")
	//code := this.GetString("code")

	if this.validate_register(phone, nickname, password, email) != nil {
		return
	}

	if models.CheckUserPhone(phone) {
		this.WriteJsonWithCode(403, ErrPhoneIsRegis)
		return
	}
	if models.CheckUserNickname(nickname) {
		this.WriteJsonWithCode(403, ErrNicknameIsRegis)
		return
	}

	password = utils.TransPassword(password) // 存储密码hash值
	user := models.StaffUser{
		Phone:    phone,
		Email:    email,
		Nickname: nickname,
		Password: password,
	}
	this.WriteJson(Response{0, "success.", models.CreateUser(user)})

}

// @Title 登录
// @Description 账号登录接口
// @Param	nickname	formData 	string	true		"用户昵称"
// @Param	password	formData 	string	true		"密码"
// @Success 200 {string}
// @Failure 401 No Admin
// @router /login [post]
func (this *UserController) Login() {
	nickname := this.GetString("nickname")
	password := this.GetString("password")

	user, ok := models.CheckUserAuth(nickname, password)
	if !ok {
		this.WriteJsonWithCode(403, ErrNicknameOrPasswd)
		return
	}

	et := utils.EasyToken{
		Username: user.Nickname,
		Uid:      user.Id,
		Expires:  utils.GetExpireTime(),
	}

	token, err := et.GetToken()
	if token == "" || err != nil {
		this.WriteJson(ErrResponse{0, err})
	} else {
		this.WriteJson(Response{0, "success.", LoginToken{user, token}})
	}

}

// @Title 认证测试
// @Description 测试错误码
// @Success 200 {string}
// @Failure 401 unauthorized
// @router /auth [get]
// @Security jwt
func (this *UserController) Auth() {
	et := utils.EasyToken{}
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))
	valid, err := et.ValidateToken(authtoken, 0)
	if !valid {
		this.WriteJsonWithCode(401, ErrResponse{-1, fmt.Sprintf("%s", err)})
		return
	}
	this.WriteJson(Response{0, "success.", "is login"})

}
