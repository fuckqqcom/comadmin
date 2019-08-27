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
