package admin

import "time"

/**
加上冗余两个字段
*/
type User struct {
	//Id  int64
	Id     string    `json:"id" xorm:"varchar(64) notnull unique 'id'"`
	Name   string    `json:"name" xorm:"varchar(256) "`
	Nick   string    `json:"name" xorm:"varchar(256) "`
	Pwd    string    `json:"pwd" xorm:"varchar(64)"`
	Phone  string    `json:"phone"`
	Status int       `json:"status" xorm:"default -1"`
	CTime  time.Time `json:"ctime" xorm:"created"`
	MTime  time.Time `json:"mtime" xorm:"updated"`
	Did    string    `json:"did" xorm:"varchar(64) 'did'"` //冗余字段，为了查询方便
	Aid    string    `json:"adi" xorm:"varchar(64) 'aid'"`
}
