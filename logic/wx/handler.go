package wx

import (
	wxd "comadmin/dao/wx"
)

type LogicHandler interface {
	jobHandler
	logicHandler
}

/**
map[string]interface
UpdateIterface
*/
type jobHandler interface {
	Register(interface{}) (interface{}, int)
}

type logicHandler interface {
	Create(interface{}) int                                        //添加数据接口
	Update(interface{}, []string, map[string]interface{}) int      //更新那些字段
	List(interface{}, int, int) (interface{}, interface{})         //第一参数是model对象 ，第二个是ps，第三个是pn
	Exist(interface{}, map[string]interface{}) bool                //判断是否存在
	Delete(interface{}, interface{}) int                           //删除  ids可以只一个 可以是多个
	Get(interface{}, []string, map[string]interface{}) interface{} //查询单个对象,返回对象

	//CreateOrDiscard(interface{}, []string, map[string]interface{}) int //创建 如果存在就丢弃
}
type Logic struct {
	Db wxd.DbHandler
}

func (l Logic) Get(bean interface{}, cols []string, colsValue map[string]interface{}) interface{} {
	return l.Db.Get(bean, cols, colsValue)
}

func (l Logic) Exist(bean interface{}, m map[string]interface{}) bool {
	return l.Db.Exist(bean, m)
}

//func (l Logic) CreateOrDiscard(bean interface{}, cols []string, colsValue map[string]interface{}) int {
//	return l.Db.CreateOrDiscard(bean, cols, colsValue)
//}

func (l Logic) Create(bean interface{}) int {
	return l.Db.Create(bean)
}

func (l Logic) List(bean interface{}, pn, ps int) (interface{}, interface{}) {
	return l.Db.List(bean, pn, ps)
}

/**
bean对应model，ids对应id列表
*/
func (l Logic) Delete(bean interface{}, ids interface{}) int {
	return l.Db.Delete(bean, ids)
}

func (l Logic) Update(bean interface{}, cols []string, m map[string]interface{}) int {
	return l.Db.Update(bean, cols, m)
}

func (l Logic) Register(bean interface{}) (interface{}, int) {
	return l.Db.Register(bean)
}

var _ LogicHandler = Logic{}

func NewLogic(path string) LogicHandler {
	return &Logic{Db: wxd.NewDb(path)}
}
