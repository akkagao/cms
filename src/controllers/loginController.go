package controllers

import (
	"cms/src/common"
	"cms/src/service"
	"strconv"

	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

/**
进入登录页面
*/
func (this *LoginController) Tologin() {
	this.show("common/loginPage.html")
}

/**
登陆
*/
func (this *LoginController) Login() {
	accout := this.GetString("accout")
	password := this.GetString("password")
	encodePwd := common.EncodeMessageMd5(password)

	if admusers, err := service.AdmUserService.Authentication(accout, encodePwd); err != nil {
		this.jsonResult(err.Error())
	} else {
		token := strconv.FormatInt(admusers.Id, 10) + "|" + accout + "|" + this.getClientIp()
		token = common.EncryptAes(token)
		this.Ctx.SetCookie("token", token, 0)
		this.jsonResult(SUCCESS)
	}
}

/**
退出登陆
*/
func (this *LoginController) Loginout() {
	this.Ctx.SetCookie("token", "", 0)
	this.redirect(beego.URLFor("LoginController.Tologin"))
}
