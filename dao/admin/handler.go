package admin

import (
	"comadmin/model/admin"
	"comadmin/pkg/config"
	"comadmin/pkg/e"
	"fmt"
	"github.com/xormplus/xorm"
)

type DbHandler interface {
	Login(interface{}) int
	Add(interface{}) int //创建
	Delete(interface{}, map[string]interface{}) int
	Update(interface{}, []string, map[string]interface{}) int                 //修改
	FindOne(interface{}, map[string]interface{}, int, int) (interface{}, int) //查询,返回列表 map装载查询条件  单表查询
	FindById(id interface{}) int
} //查询

type daoHandler interface {
	Add(interface{}) int                                      //添加数据接口
	Update(interface{}, []string, map[string]interface{}) int //修改
	Delete(interface{}, map[string]interface{}) int
	Exist(interface{}, map[string]interface{}) bool
	FindOne(interface{}, map[string]interface{}, int, int) (interface{}, interface{}) //查询,返回列表 map装载查询条件  单表查询
	Get(interface{}, []string, map[string]interface{}) interface{}                    //查询单个对象,返回对象
	//AddOrUpdate(interface{}, []string, map[string]interface{}) int  //创建或者更新 存在就更新,不存在就创建
	//AddOrDiscard(interface{}, []string, map[string]interface{}) int //创建 如果存在就丢弃
}

type Dao struct {
	c      config.Config
	engine *xorm.Engine
}

func NewDb(path string) *Dao {
	return &Dao{engine: config.EngDb, c: config.NewConfig(path)}
}

//接口   结构体
var _ DbHandler = Dao{}

func (d Dao) Login(i interface{}) int {

	switch t := i.(type) {
	case *admin.User:
		return d.login(t)
	default:
		return e.Errors
	}
}

func (d Dao) Add(i interface{}) int {
	switch t := i.(type) {
	case admin.Domain, admin.DomainApp, admin.User, admin.DomainAppUser:
		return d.add(t)
	default:
		fmt.Println("create other ...")
		return e.Errors
	}
}

func (d Dao) Delete(bean interface{}, queryValue map[string]interface{}) int {
	switch t := bean.(type) {
	case admin.Domain:
		return d.delete(t, queryValue)
	//case admin.DomainApp:
	//	return d.delete(t.Id, t)
	default:
		fmt.Println("delete other ...")
		return e.Errors
	}
}

func (d Dao) Update(bean interface{}, cols []string, queryValue map[string]interface{}) int {
	//todo 怎么优化
	switch t := bean.(type) {
	case admin.Domain:
		return d.update(t, cols, queryValue)
	//case admin.DomainApp:
	//	return d.update(t.Id, t, cols...)
	//case admin.User:
	//	return d.update(t.Id, t, cols...)
	//case admin.Role:
	//	return d.update(t.Id, t, cols...)
	//case admin.DomainAppRole:
	//	return d.update(t.Id, t, cols...)

	default:
		fmt.Println("update other ...")
		return e.Errors
	}
}

func (d Dao) FindOne(bean interface{}, m map[string]interface{}, ps, pn int) (interface{}, int) {

	switch t := bean.(type) {
	case admin.Domain:
		type domain struct {
			Name   string `json:"name"`
			Id     string `json:"id"`
			Status int    `json:"status"`
			Ctime  string `json:"ctime" `
			Mtime  string `json:"mtime" `
		}
		dm := make([]domain, 0)
		return d.findOneTable(&dm, m, pn, ps)
		//return d.findDomain(t, pn, ps)
	case admin.DomainApp:
		return d.findApp(t, pn, ps)
	case admin.User:
		return d.findUser(t, pn, ps)
	case admin.Role:
		return d.findRole(t, pn, ps)

	default:
		fmt.Println("update other ...")
		return nil, e.Errors
	}
}

func (d Dao) FindById(i interface{}) int {
	switch t := i.(type) {
	case *admin.Domain:
		return d.findById(t)
	case *admin.DomainApp:
		return d.findById(t)
	default:
		fmt.Println("update other ...")
		return e.Errors
	}
}
