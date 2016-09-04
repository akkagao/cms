package model

type GroupRoleRel struct {
	Id      int64 `json:"id"`
	Groupid int64 `json:"groupid"`
	Roleid  int64 `json:"roleid"`
	Isdel   int8  `json:"isdel",orm:"default(1)"`
}
