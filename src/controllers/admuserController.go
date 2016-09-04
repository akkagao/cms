package controllers

import (
	"cms/src/common"
	"cms/src/model"
	"cms/src/service"
	"time"
)

type AdmUserController struct {
	BaseController
}

/**
进入管理员列表页面
*/
func (this *AdmUserController) List() {
	this.show("admUser/admUserList.html")
}

/**
获取分页展示数据
*/
func (this *AdmUserController) Gridlist() {
	pageNum, _ := this.GetInt("page")
	rowsNum, _ := this.GetInt("rows")
	admusermail := this.GetString("admusermail")
	admuserphone := this.GetString("admuserphone")
	admusername := this.GetString("admusername")
	admuserid := this.GetString("admuserid")
	account := this.GetString("admaccout")
	p := common.NewPager(pageNum, rowsNum)
	count, admuser := service.AdmUserService.Gridlist(p, admuserid, admusermail, admusername, admuserphone, account)
	this.jsonResultPager(count, admuser)
}

/**
进入添加页面
*/
func (this *AdmUserController) Toaddadmuser() {
	this.show("admUser/addAdmUser.html")
}

/**
添加管理员
*/
func (this *AdmUserController) Addadmuser() {
	account := this.GetString("account")
	mail := this.GetString("mail")
	name := this.GetString("name")
	phone := this.GetString("phone")
	department := this.GetString("department")
	password := this.GetString("password")
	groupIds := this.GetString("ids")
	password = common.EncodeMessageMd5(password)

	admuser := &model.Admuser{
		Accout:     account,
		Name:       name,
		Mail:       mail,
		Phone:      phone,
		Department: department,
		Password:   password,
		Createtime: time.Now(),
		Updatetime: time.Now(),
		Isdel:      1}
	if err := service.AdmUserService.AddAdmUser(admuser, groupIds); err != nil {
		this.jsonResult(err.Error())
	}
	this.jsonResult(SUCCESS)
}

/**
进入修改页面
*/
func (this *AdmUserController) Tomodifyadmuser() {
	admUserId, _ := this.GetInt64("admUserId")
	admUser, _ := service.AdmUserService.GetUserById(admUserId)
	this.Data["admuser"] = admUser
	this.show("admUser/modifyAdmUser.html")
}

/**
修改管理员
*/
func (this *AdmUserController) Modifyyadmuser() {
	userId, _ := this.GetInt64("userId")
	account := this.GetString("account")
	mail := this.GetString("mail")
	name := this.GetString("name")
	phone := this.GetString("phone")
	department := this.GetString("department")
	password := this.GetString("password")
	groupIds := this.GetString("groupids")

	if len(password) != 0 {
		password = common.EncodeMessageMd5(password)
	}

	admuser := &model.Admuser{
		Id:         userId,
		Accout:     account,
		Name:       name,
		Mail:       mail,
		Phone:      phone,
		Department: department,
		Password:   password,
		Createtime: time.Now(),
		Updatetime: time.Now(),
		Isdel:      1}

	if err := service.AdmUserService.ModifyAdmUser(admuser, groupIds); err != nil {
		this.jsonResult(err.Error())
	}
	this.jsonResult(SUCCESS)
}

/**
删除
*/
func (this *AdmUserController) Delete() {
	userids := this.GetString("userids")
	if err := service.AdmUserService.Delete(userids); err != nil {
		this.jsonResult(err.Error())
	}
	this.jsonResult(SUCCESS)
}

/**
获取管理员组列表数据
修改管理员的时候需要加载管理员组列表，并且设置已经选择的权限为选中状态
*/
func (this *AdmUserController) Gridgrouplist() {
	admUserId, _ := this.GetInt64("admUserId")
	groupName := this.GetString("groupName")
	pageNum, _ := this.GetInt("page")
	rowsNum, _ := this.GetInt("rows")
	p := common.NewPager(pageNum, rowsNum)

	count, admuserGroup := service.AdmUserGroupService.Gridlist(groupName, p)
	checkedGroupId := service.AdmUserService.GetAllCheckGroup(admUserId)

	admUserCheckGroup := make([]model.Admusergroupcheck, len(admuserGroup))

	for index, admuser := range admuserGroup {
		admUserCheck := model.Admusergroupcheck{
			Id:         admuser.Id,
			Groupname:  admuser.Groupname,
			Des:        admuser.Des,
			Createtime: admuser.Createtime,
			Updatetime: admuser.Updatetime,
			Isdel:      admuser.Isdel,
			Check:      checkedGroupId[admuser.Id]}
		admUserCheckGroup[index] = admUserCheck
	}

	this.jsonResultPager(count, admUserCheckGroup)
}
