package controllers

import "cms/src/service"

type MainController struct {
	BaseController
}

/**
进入首页
*/
func (this *MainController) Index() {
	this.show("index.html")
}

/**
欢迎页面
*/
func (this *MainController) Welcome() {
	this.show("common/welcome.html")
}

/**
左侧菜单
*/
func (this *MainController) LeftMenu() {
	this.show("common/leftMenu.html")
}

/**
头页面
*/
func (this *MainController) Header() {
	this.Data["user"] = this.admUser
	this.show("common/header.html")
}

/**
进入没有权限页面
*/
func (this *MainController) Norole() {
	this.show("common/noRole.html")
}

/**
加载主页面权限tree
*/
func (this *MainController) LoadMenu() {
	id := this.admUser.Id
	roles := service.RoleService.LoadMenu(id)
	this.jsonResult(roles)
}
