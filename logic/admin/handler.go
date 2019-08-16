package admin

import (
	"comadmin/dao/admin"
)

type LogicHandler interface {
	Create(i interface{}) int
	Delete(i interface{}) int
	Update(i interface{}) int
	Find(i interface{}, pn, ps int) (interface{}, int)
}

type Logic struct {
	Db admin.DbHandler
}

func NewLogic(path string) LogicHandler {
	return &Logic{Db: admin.NewDb(path)}
}
func (l Logic) Create(i interface{}) int {
	return l.Db.Create(i)
}
func (l Logic) Delete(i interface{}) int {
	return l.Db.Delete(i)
}

func (l Logic) Update(i interface{}) int {
	return l.Db.Update(i)
}

func (l Logic) Find(i interface{}, pn, ps int) (interface{}, int) {
	return l.Db.Find(i, pn, ps)
}
