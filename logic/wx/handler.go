package wx

import (
	wxd "comadmin/dao/wx"
)

type LogicHandler interface {
	jobHandler
	logicHandler
}

type jobHandler interface {
	FindCountByIdAndIp(id, ip string) int
	Register(interface{}) (interface{}, int)
	FindList(interface{}, int, int) (interface{}, int)
}

type logicHandler interface {
	Create(interface{}) int //添加数据接口
	Update(interface{}, []string) int
	FindBiz(string) (interface{}, int) //获取公号信息
	FindApi() (interface{}, int)       //获取接口
	PostData(interface{}) int          //提交数据接口
	Find(interface{}, int, int) (interface{}, interface{})
}
type Logic struct {
	Db wxd.DbHandler
}

func (l Logic) FindCountByIdAndIp(id, ip string) int {
	return l.Db.FindCountByIdAndIp(id, ip)
}

func (l Logic) Register(i interface{}) (interface{}, int) {
	return l.Db.Register(i)
}

func (l Logic) Create(i interface{}) int {
	return l.Db.Create(i)
}

func (l Logic) FindBiz(id string) (interface{}, int) {
	return l.Db.FindBiz(id)
}

func (l Logic) FindApi() (interface{}, int) {
	return l.Db.FindApi()
}

func (l Logic) PostData(i interface{}) int {
	return l.Db.PostData(i)
}

func (l Logic) Find(i interface{}, pn, ps int) (interface{}, interface{}) {
	return l.Db.Find(i, pn, ps)
}

func (l Logic) FindList(i interface{}, pn, ps int) (interface{}, int) {
	return l.Db.FindList(i, pn, ps)
}
func (l Logic) Update(i interface{}, cols []string) int {
	return l.Db.Update(i, cols)
}

var _ LogicHandler = Logic{}

func NewLogic(path string) LogicHandler {
	return &Logic{Db: wxd.NewDb(path)}
}
