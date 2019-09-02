package admin

import (
	"comadmin/dao/admin"
)

type LogicHandler interface {
	adminHandler
}

type adminHandler interface {
	Login(interface{}) int
	Add(interface{}) int                                                      //添加数据接口
	Update(interface{}, []string, map[string]interface{}) int                 //更新那些字段
	FindOne(interface{}, map[string]interface{}, int, int) (interface{}, int) //第一参数是model对象 ，第二个是ps，第三个是pn
	Exist(interface{}, map[string]interface{}) bool                           //判断是否存在
	Delete(interface{}, map[string]interface{}) int                           //删除  ids可以只一个 可以是多个
	Get(interface{}, []string, map[string]interface{}) interface{}            //查询单个对象,返回对象
	FindMany(string) (interface{}, int)                                       //复杂sql，直接写sql吧
	FindById(interface{}) int
}

type Logic struct {
	Db admin.DbHandler
}

func (l Logic) FindMany(string) (interface{}, int) {
	panic("implement me")
}

func (l Logic) Get(interface{}, []string, map[string]interface{}) interface{} {
	panic("implement me")
}

var _ LogicHandler = Logic{}

func NewLogic(path string) LogicHandler {
	return &Logic{Db: admin.NewDb(path)}
}

func (l Logic) Login(bean interface{}) int {
	return l.Db.Login(bean)
}
func (l Logic) Add(bean interface{}) int {
	return l.Db.Add(bean)
}
func (l Logic) Delete(bean interface{}, queryValue map[string]interface{}) int {
	return l.Db.Delete(bean, queryValue)
}

func (l Logic) Update(bean interface{}, cols []string, queryValue map[string]interface{}) int {
	return l.Db.Update(bean, cols, queryValue)
}

func (l Logic) FindOne(bean interface{}, m map[string]interface{}, ps, pn int) (interface{}, int) {
	return l.Db.FindOne(bean, m, pn, ps)
}

func (l Logic) FindById(id interface{}) int {
	return l.Db.FindById(id)
}

func (l Logic) Exist(interface{}, map[string]interface{}) bool {
	panic("implement me")
}
