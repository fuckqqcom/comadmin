package admin

import (
	"comadmin/model/admin"
	"comadmin/pkg/config"
	"comadmin/pkg/e"
	"fmt"
	"github.com/xormplus/xorm"
)

type DbHandler interface {
	Create(interface{}) int //创建
	Delete(interface{}) int //删除
	Update(interface{}) int //修改
	Find(i interface{}, pn, ps int) (interface{}, int)
} //查询

type Dao struct {
	c      config.Config
	Engine *xorm.Engine
}

func NewDb(path string) *Dao {
	return &Dao{Engine: config.EngDb, c: config.NewConfig(path)}
}

func (d Dao) Create(i interface{}) int {
	switch i.(type) {
	case admin.Domain:
		insert, err := d.Engine.Insert(i)
		fmt.Println("create domain me", insert, err)
		return e.Success
	default:
		fmt.Println("create other ...")
	}
	return 0
}

func (d Dao) Delete(i interface{}) int {
	switch i.(type) {
	case admin.Domain:
		fmt.Println("delete domain me")
	default:
		fmt.Println("delete other ...")
	}
	return 0
}

func (d Dao) Update(i interface{}) int {
	switch i.(type) {
	case admin.Domain:
		fmt.Println("update domain me")
	default:
		fmt.Println("update other ...")
	}
	return 0
}

func (d Dao) Find(i interface{}, pn, ps int) (interface{}, int) {

	switch i.(type) {
	case admin.Domain:
		d.FindDomain(i.(admin.Domain), pn, ps)
	default:
		fmt.Println("update other ...")
	}
	return nil, 0
}
