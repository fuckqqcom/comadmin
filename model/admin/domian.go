package admin

import "time"

/**
域
*/
type Domain struct {
	Id     string    `json:"id" xorm:"varchar(64) notnull unique 'id'"`
	Name   string    `json:"name" xorm:"varchar(256) notnull unique 'name'"`
	Status int       `json:"status" xorm:"default -1"`
	Ctime  time.Time `json:"ctime" xorm:"created 'ctime'"`
	Mtime  time.Time `json:"mtime" xorm:"updated 'mtime'" `
}
