package admin

import (
	"comadmin/dao/admin"
)

type LogicHandler interface {
	adminHandler
}

type adminHandler interface {
	Login(interface{}) int
	Create(interface{}) int
	Delete(interface{}) int
	Update(interface{}, []string) int
	FindById(interface{}) int
	Find(interface{}, int, int) (interface{}, int)
}

type commonHandler interface {
	Create(interface{}) int                                        //添加数据接口
	Update(interface{}, []string, map[string]interface{}) int      //更新那些字段
	List(interface{}, int, int) (interface{}, interface{})         //第一参数是model对象 ，第二个是ps，第三个是pn
	Exist(interface{}, map[string]interface{}) bool                //判断是否存在
	Delete(interface{}, interface{}) int                           //删除  ids可以只一个 可以是多个
	Get(interface{}, []string, map[string]interface{}) interface{} //查询单个对象,返回对象

	//CreateOrDiscard(interface{}, []string, map[string]interface{}) int //创建 如果存在就丢弃
}

type Logic struct {
	Db admin.DbHandler
}

var _ LogicHandler = Logic{}

func NewLogic(path string) LogicHandler {
	return &Logic{Db: admin.NewDb(path)}
}

func (l Logic) Login(i interface{}) int {
	return l.Db.Login(i)
}
func (l Logic) Create(i interface{}) int {
	return l.Db.Create(i)
}
func (l Logic) Delete(i interface{}) int {
	return l.Db.Delete(i)
}

func (l Logic) Update(i interface{}, cols []string) int {
	return l.Db.Update(i, cols)
}

func (l Logic) Find(i interface{}, pn, ps int) (interface{}, int) {
	return l.Db.Find(i, pn, ps)
}

func (l Logic) FindById(id interface{}) int {
	return l.Db.FindById(id)
}
