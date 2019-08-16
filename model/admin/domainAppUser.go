package admin

import "time"

/**
domain app user关联关系表(admin用)
*/

type DomainAppUser struct {
	Id     string    `json:"id" xorm:"varchar(64) notnull unique 'id'"`
	Did    string    `json:"did" xorm:"varchar(64)   'did'"`
	Aid    string    `json:"aid" xorm:"varchar(64)   'aid'"`
	Uid    string    `json:"uid" xorm:"varchar(64)   'uid'"`
	Status int       `json:"status" xorm:"default 1"`
	CTime  time.Time `json:"ctime" xorm:"created"`
	MTime  time.Time `json:"mtime" xorm:"updated" `
}
