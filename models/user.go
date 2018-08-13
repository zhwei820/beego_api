package models

import (
	"github.com/astaxie/beego/orm"
	// "log"
	// "fmt"
	"time"
	"back/beego_api/utils"
)

type StaffUser struct {
	Id       int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Phone    string    `json:"phone" orm:"column(phone);unique;size(11)"`
	Email    string    `json:"email" orm:"column(email);unique;size(50)"`
	Nickname string    `json:"nickname" orm:"column(nickname);unique;size(100);"`
	Password string    `json:"-" orm:"column(password);size(300)"`
	City     string    `json:"city" orm:"column(city);size(80);"`
	Created  time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated  time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
}

func (u *StaffUser) TableName() string {
	return TableName("user")
}
func init() {
	orm.RegisterModel(new(StaffUser))
}
func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(StaffUser))
}

// 检测手机号是否注册
func CheckUserPhone(phone string) bool {
	exist := Users().Filter("phone", phone).Exist()
	return exist
}

// 检测用户昵称是否存在
func CheckUserNickname(nickname string) bool {
	exist := Users().Filter("nickname", nickname).Exist()
	return exist
}

//创建用户
func CreateUser(user StaffUser) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		return -1
	}
	return id
}

//检测手机和昵称是否注册
func CheckUserPhoneOrNickname(phone string, nickname string) bool {
	cond := orm.NewCondition()
	count, _ := Users().SetCond(cond.And("phone", phone).Or("nickname", nickname)).Count()
	if count <= int64(0) {
		return false
	}
	return true
}
func CheckUserAuth(nickname string, password string) (StaffUser, bool) {
	var staffuser StaffUser
	err := Users().Filter("nickname", nickname).One(&staffuser)

	if err != nil || staffuser.Password != utils.TransPassword(password) {
		return staffuser, false
	}
	return staffuser, true
}

// StaffUser database CRUD methods include Insert, Read, Update and Delete
func (usr *StaffUser) Insert() error {
	if _, err := orm.NewOrm().Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *StaffUser) Read(fields ...string) error {
	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *StaffUser) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *StaffUser) Delete() error {
	if _, err := orm.NewOrm().Delete(usr); err != nil {
		return err
	}
	return nil
}
