package service

import (
	"cms/src/model"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	o           orm.Ormer
	tablePrefix string // 表前缀

	RoleService         *roleService
	AdmUserGroupService *admUserGroupService
	AdmUserService      *admUserService
)

func init() {
	beego.Info("init orm start...")
	tablePrefix = beego.AppConfig.String("db.prefix")

	dbType := beego.AppConfig.String("db_type")
	dsn := generateDSN()
	orm.RegisterDataBase("default", dbType, dsn)

	orm.RegisterModelWithPrefix(tablePrefix,
		new(model.Role),
		new(model.Admusergroup),
		new(model.GroupRoleRel),
		new(model.Admuser),
		new(model.UserGroupRel),
	)
	orm.RunSyncdb("default", false, true)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	o = orm.NewOrm()
	orm.RunCommand()

	beego.Info("init orm end.")
	//初始化service
	initService()
}

func generateDSN() string {
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbUser := beego.AppConfig.String("db_user")
	dbPassword := beego.AppConfig.String("db_pass")
	dbName := beego.AppConfig.String("db_name")

	//beego.Debug(dbHost, dbPort, dbUser, dbPassword, dbName, dbType)
	// root:@tcp(127.0.0.1:3306)/test?charset=utf8
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	return dsn
}

func initService() {
	RoleService = &roleService{}
	AdmUserGroupService = &admUserGroupService{}
	AdmUserService = &admUserService{}
}
