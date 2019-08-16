package admin

import "time"

/**
role

角色分可读 可写 读写
*/

type Role struct {
	Id     string    `json:"id" xorm:"varchar(64) notnull unique 'id'"`
	Name   string    `json:"name" xorm:"varchar(256) notnull unique 'name'"`
	Status int       `json:"status" xorm:"default -1"`
	Rw     int       `json:"rw" xorm:"default 1"` //读写标志  1为读 2为写 3可读可写
	CTime  time.Time `json:"ctime" xorm:"created"`
	MTime  time.Time `json:"mtime" xorm:"updated" `
}
