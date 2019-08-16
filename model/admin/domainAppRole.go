package admin

import "time"

/**
domain app role关联关系(admin用)
*/

type DomainAppRole struct {
	Id     string    `json:"id" xorm:"varchar(64) notnull unique 'id'"`
	Did    string    `json:"did" xorm:"varchar(64)   'did'"`
	Aid    string    `json:"aid" xorm:"varchar(64)   'aid'"`
	Rid    string    `json:"rid" xorm:"varchar(64)   'rid'"`
	Status int       `json:"status" xorm:"default 1"`
	CTime  time.Time `json:"ctime" xorm:"created"`
	MTime  time.Time `json:"mtime" xorm:"updated" `
}
