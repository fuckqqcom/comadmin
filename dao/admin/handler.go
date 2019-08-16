package admin

import (
	"comadmin/model/admin"
	"fmt"
)

type DBHandler interface {
	Create(interface{}) int //创建
	Delete(interface{}) int //删除
	Update(interface{}) int //修改
}

type Dao struct {
	Engine string
}

func NewDB() *Dao {
	return &Dao{Engine: "1"}
}

func (d Dao) Create(i interface{}) int {
	switch i.(type) {
	case admin.Domain:
		fmt.Println("create domain me")
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
