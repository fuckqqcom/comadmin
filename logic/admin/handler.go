package admin

import (
	"comadmin/dao/admin"
)

type LogicHandler interface {
	Create(i interface{}) int
	Delete(i interface{}) int
	Update(i interface{}) int
}

type Logic struct {
	DB admin.DBHandler
}

func NewLogic(path string) LogicHandler {
	return &Logic{DB: admin.NewDB(path)}
}
func (l Logic) Create(i interface{}) int {
	return l.DB.Create(i)
}
func (l Logic) Delete(i interface{}) int {
	return l.DB.Delete(i)
}

func (l Logic) Update(i interface{}) int {
	return l.DB.Update(i)
}
