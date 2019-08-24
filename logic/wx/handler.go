package wx

import (
	wxd "comadmin/dao/wx"
)

type LogicHandler interface {
	Create(interface{}) int      //添加数据接口
	FindBiz() (interface{}, int) //获取公号信息
	FindApi() (interface{}, int) //获取接口
	PostData(interface{}) int    //提交数据接口
	Find(interface{}, int, int) (interface{}, interface{})
}

type Logic struct {
	Db wxd.DbHandler
}

func (l Logic) Create(interface{}) int {
	panic("implement me")
}

func (l Logic) FindBiz() (interface{}, int) {
	return l.Db.FindBiz()
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

var _ LogicHandler = Logic{}

func NewLogic(path string) LogicHandler {
	return &Logic{Db: wxd.NewDb(path)}
}
