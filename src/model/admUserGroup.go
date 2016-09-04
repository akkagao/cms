package model

import "time"

type Admusergroup struct {
	Id         int64     `json:"id"`
	Groupname  string    `json:"groupname"`
	Des        string    `json:"describe"`
	Createtime time.Time `json:"createtime",orm:"auto_now_add;type(datetime)"`
	Updatetime time.Time `json:"updatetime",orm:"auto_now;type(datetime)"`
	Isdel      int8      `json:"isdel",orm:"default(1)"`
}

type Admusergroupcheck struct {
	Id         int64     `json:"id"`
	Groupname  string    `json:"groupname"`
	Des        string    `json:"describe"`
	Createtime time.Time `json:"createtime",orm:"auto_now_add;type(datetime)"`
	Updatetime time.Time `json:"updatetime",orm:"auto_now;type(datetime)"`
	Isdel      int8      `json:"isdel",orm:"default(1)"`
	Check      bool      `json:"check"`
}
