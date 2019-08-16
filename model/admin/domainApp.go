package admin

import "time"

/**
Id 用Did和Aid计算出来一个值，然后代表当前唯一id,这样方便查询在插入数据的时候查询是否存在

*/
type DomainApp struct {
	Id     string    `json:"id" xorm:"varchar(64) notnull unique 'id'"`
	Name   string    `json:"name" xorm:"varchar(256) notnull unique 'name'"`
	Did    string    `json:"did" xorm:"varchar(64) notnull unique 'did'"`
	Status int       `json:"status" xorm:"default 1"`
	Ctime  time.Time `json:"ctime" xorm:"created 'ctime'" `
	Mtime  time.Time `json:"mtime" xorm:"updated 'mtime'" `
}
