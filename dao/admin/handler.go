package admin

import (
	"comadmin/model/admin"
	"comadmin/pkg/config"
	"fmt"
	"github.com/xormplus/xorm"
)

type DbHandler interface {
	Login(interface{}) int
	Create(interface{}) int           //创建
	Delete(interface{}) int           //删除
	Update(interface{}, []string) int //修改
	Find(i interface{}, pn, ps int) (interface{}, int)
	FindById(id interface{}) int
} //查询

type Dao struct {
	c      config.Config
	Engine *xorm.Engine
}

func NewDb(path string) *Dao {
	return &Dao{Engine: config.EngDb, c: config.NewConfig(path)}
}

//接口   结构体
var _ DbHandler = Dao{}

func (d Dao) Login(i interface{}) int {
	return d.login(i.(admin.User))
}

func (d Dao) Create(i interface{}) int {
	switch t := i.(type) {
	case admin.Domain, admin.DomainApp:
		return d.create(t)
	default:
		fmt.Println("create other ...")
	}
	return 0
}

func (d Dao) Delete(i interface{}) int {
	switch t := i.(type) {
	case admin.Domain:
	case admin.DomainApp:
		return d.delete(t.Id, t)
	default:
		fmt.Println("delete other ...")
	}
	return 0
}

func (d Dao) Update(i interface{}, cols []string) int {
	//todo 怎么优化
	switch t := i.(type) {
	case admin.Domain:
		return d.update(t.Id, t, cols...)
	case admin.DomainApp:
		return d.update(t.Id, t, cols...)
	case admin.User:
		return d.update(t.Id, t, cols...)
	default:
		fmt.Println("update other ...")
		return 0
	}
}

func (d Dao) Find(i interface{}, pn, ps int) (interface{}, int) {

	switch t := i.(type) {
	case admin.Domain:
		return d.findDomain(t, pn, ps)
	case admin.DomainApp:
		return d.findApp(t, pn, ps)
	case admin.User:
		return d.findUser(t, pn, ps)
	default:
		fmt.Println("update other ...")
	}
	return nil, 0
}

func (d Dao) FindById(i interface{}) int {
	switch t := i.(type) {
	case admin.Domain:
	case admin.DomainApp:
		return d.findById(t)
	default:
		fmt.Println("update other ...")
	}
	return 0
}
