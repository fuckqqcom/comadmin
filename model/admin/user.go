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

//需要其他的查询参数继续添加
type UserParams struct {
	Id     string `json:"id"`     //根据id查询
	Name   string `json:"name"`   //根据name模糊查询
	Status int    `json:"status"` //根据状态精确查询
	Phone  string `json:"phone"`  //根据手机号码模糊查询
	Did    string `json:"did"`    //根据域id精确查询
	Aid    string `json:"aid"`    //根据应用id精确查询
	Pn     int    `json:"pn"`
	Ps     int    `json:"ps"`
}

//需要其他的返回字段可以继续添加
type UserResult struct {
	Id     string `json:"id"`     //根据id查询
	Name   string `json:"name"`   //根据name模糊查询
	Status int    `json:"status"` //根据状态精确查询
	Phone  string `json:"phone"`  //根据手机号码模糊查询
	Did    string `json:"did"`    //根据域id精确查询
	Aid    string `json:"aid"`    //根据应用id精确查询
}
