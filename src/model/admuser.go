package model

import "time"

type Admuser struct {
	Id         int64     `json:"id"`
	Accout     string    `json:"accout",orm:"unique"`
	Mail       string    `json:"mail"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Department string    `json:"department"`
	Password   string    `json:"password"`
	Createtime time.Time `json:"createtime",orm:"auto_now_add;type(datetime)"`
	Updatetime time.Time `json:"updatetime",orm:"auto_now;type(datetime)"` // 更新时间
	Isdel      int8      `json:"isdel",orm:"default(1)"`
}
