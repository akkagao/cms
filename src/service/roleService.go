package service

import (
	"bytes"
	"cms/src/common"
	"cms/src/model"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type roleService struct{}

/**
添加权限
*/
func (this *roleService) AddRole(role *model.Role) error {
	if _, err := o.Insert(role); err != nil {
		return &common.BizError{"添加失败"}
	}
	return nil
}

/**
查询列表的分页数据
*/
func (this *roleService) Gridlist(pager *common.Pager, roleid int, roleName, roleUrl string) (int, []model.Role) {
	//查询总数
	contsql := "SELECT count(1) from t_role t where t.pid = ?"
	condition := genCondition(roleName, roleUrl)
	var count int
	err := o.Raw(contsql+condition, roleid).QueryRow(&count)
	if err != nil {
		beego.Error("查询Pid为", roleid, "的role总数异常，error message：", err.Error())
	}
	beego.Debug("pid 为", roleid, "的role有", count, "个")

	if count < 1 {
		beego.Info("没有pid 为", roleid, "的role")
		return 0, nil
	}

	// 从数据库查询数据
	var roles []model.Role
	listsql := "SELECT id, pid, name, roleurl, module, action, ismenu, des from t_role t where t.pid = ?  "
	_, err = o.Raw(listsql+condition+common.LIMIT, roleid, pager.GetBegin(), pager.GetLen()).QueryRows(&roles)
	if err != nil {
		beego.Error("查询Pid为", roleid, "的role列表异常，error message：", err.Error())
	}

	return count, roles
}

func genCondition(roleName, roleUrl string) (condition string) {
	if !strings.EqualFold(roleName, "") {
		condition += " and t.name = '" + roleName + "'"
	}
	if !strings.EqualFold(roleUrl, "") {
		condition += " and t.roleurl = '" + roleUrl + "'"
	}
	return
}

/**
查询树
@param needRoot:查询的数据集中是否需要包含root节点
*/
func (this *roleService) Listtree(needRoot bool) []model.RoleTree {
	var buf bytes.Buffer
	buf.WriteString("SELECT id, pid, name, roleurl, ismenu, des from t_role t ")
	if !needRoot {
		buf.WriteString(" where t.id != 0")
	}
	var roles []model.RoleTree
	beego.Debug("查询权限树sql：", buf.String())
	_, err := o.Raw(buf.String()).QueryRows(&roles)
	if err != nil {
		beego.Error("查询权限树的role列表异常，error message：", err.Error())
	}
	beego.Debug("生成权限树的数据：", roles)
	return roles
}

/**
根据ID查询role
*/
func (this *roleService) GetRoleById(id int64) (model.Role, error) {
	role := model.Role{Id: id}
	if err := o.Read(&role); err != nil {
		return model.Role{}, err
	}
	return role, nil
}

/**
修改权限
*/
func (this *roleService) ModifyRole(r *model.Role) error {
	role := model.Role{Id: r.Id}
	//根据ID读取
	if err := o.Read(&role); err != nil {
		return err
	}
	//修改
	if num, err := o.Update(r); num <= 0 && err != nil {
		return err
	}
	return nil
}

/**
删除权限
*/
func (this *roleService) DeleteRole(ids []string) error {
	idstr := strings.Join(ids, ",")

	var count int
	countSubRoleSql := "select count(1) from t_role where pid in (" + idstr + ")"
	o.Raw(countSubRoleSql).QueryRow(&count)
	if count > 0 {
		return &common.BizError{"不能删除有子节点的权限，请先删除所有子节点！"}
	}

	sql := "DELETE from t_role  where id in (" + idstr + ")"
	if _, err := o.Raw(sql).Exec(); err != nil {
		return &common.BizError{"删除失败！"}
	}
	return nil
}

/**
权限校验
*/
func (this *roleService) ValidateRole(controllerName, actionName string, id int64) error {
	if this.isAdministrator(id) {
		beego.Debug("用户属于超级管理员，不用校验权限")
		return nil
	}
	selectSql := "SELECT COUNT(1) FROM t_user_group_rel ur,t_role r ,t_group_role_rel gr where r.module = ? and r.action = ? and ur.userid = ? and ur.groupid = gr.groupid and r.id = gr.roleid and ur.isdel = 1 and gr.isdel = 1"
	var count int
	o.Raw(selectSql, controllerName, actionName, id).QueryRow(&count)
	if count > 0 {
		return nil
	}
	return &common.BizError{"您没有权限执行此操作，请联系系统管理员。"}
}

/**
加载权限树
*/
func (this *roleService) LoadMenu(id int64) []model.RoleTree {

	var roles []model.RoleTree
	if this.isAdministrator(id) {
		selectSql := "SELECT t.id, pid, name, roleurl , ismenu, des from t_role t where t.id != 0 and t.ismenu = 0"
		if _, err := o.Raw(selectSql).QueryRows(&roles); err != nil {
			beego.Error("查询权限树的role列表异常，error message：", err.Error())
			return roles
		}
	} else {
		selectSql := "SELECT DISTINCT t.id, pid, name, roleurl , ismenu, des from t_role t,t_user_group_rel ug,t_group_role_rel gr where t.id != 0 and t.ismenu = 0 and t.id = gr.roleid and ug.userid=? and ug.groupid = gr.groupid and ug.isdel=1 and gr.isdel =1"
		if _, err := o.Raw(selectSql, id).QueryRows(&roles); err != nil {
			beego.Error("查询权限树的role列表异常，error message：", err.Error())
			return roles
		}
	}

	pidMap := make(map[int64]bool, 10)
	for _, role := range roles {
		pidMap[role.Pid] = true
	}

	for i, role := range roles {
		//展开所有父节点
		if pidMap[role.Id] {
			roles[i].Open = true
			continue
		}
		if !strings.EqualFold(role.Roleurl, "") {
			click := "click: addTab('" + roles[i].Name + "','" + roles[i].Roleurl + "')"
			roles[i].Click = click
		}
	}

	return roles
}

/*
判断当前用户是否属于 超级管理员
*/
func (this *roleService) isAdministrator(id int64) bool {

	flag := false
	var list orm.ParamsList
	num, err := o.Raw("SELECT groupid from t_user_group_rel t where t.userid = ? and t.isdel =1", id).ValuesFlat(&list)
	if err != nil || num < 1 {
		return flag
	}
	for i := 0; i < len(list); i++ {
		groupId := list[i].(string)
		if id, err := strconv.ParseInt(groupId, 10, 64); err == nil {
			if id == 1 {
				return true
			}
		}
	}
	return flag
}
