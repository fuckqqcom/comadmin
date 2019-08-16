package admin

import (
	"comadmin/dao/admin"
)

type LogicHandler interface {
	Create(i interface{}) interface{}
	Delete(i interface{}) interface{}
}

type Logic struct {
	DB admin.DBHandler
}

func NewLogic() LogicHandler {
	return &Logic{DB: admin.NewDB()}
}
func (l Logic) Create(i interface{}) interface{} {
	return l.DB.Create(i)
}
func (l Logic) Delete(i interface{}) interface{} {
	return l.DB.Delete(i)
}
