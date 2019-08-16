package admin

import "time"

/**
user role
*/

type UserRole struct {
	Id     string    `json:"id" xorm:"varchar(64) notnull unique 'id'"`
	Rid    string    `json:"rid" xorm:"varchar(64) "`
	Uid    string    `json:"uid" xorm:"varchar(64) "`
	Status int       `json:"status" xorm:"default 1"`
	CTime  time.Time `json:"ctime" xorm:"created"`
	MTime  time.Time `json:"mtime" xorm:"updated" `
}
