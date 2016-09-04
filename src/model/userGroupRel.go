package model

type UserGroupRel struct {
	Id      int64
	Userid  int64
	Groupid int64
	Isdel   int8 `orm:"default(1)"`
}
