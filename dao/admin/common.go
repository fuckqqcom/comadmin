package admin

import (
	"comadmin/pkg/e"
	"comadmin/tools/utils"
)

/**
一些公用的操作,比如删除
*/

//删除 通过id删除
func (d Dao) delete(id, bean interface{}) int {
	affect, err := d.Engine.Where(" id = ? ", id).Delete(&bean)
	if utils.CheckError(err, affect) {
		return e.Success
	}
	return e.Errors
}

//创建
func (d Dao) create(bean interface{}) int {
	affect, err := d.Engine.Insert(bean)
	if utils.CheckError(err, affect) {
		return e.Success
	}
	return e.Errors
}

//通过id查找

func (d Dao) findById(bean interface{}) int {
	affect, err := d.Engine.Exist(bean)
	if utils.CheckError(err, affect) && affect {
		return e.Success
	}
	return e.ExistError
}

/**
用过id更新  name 和status状态
*/

func (d Dao) update(id interface{}, bean interface{}, cols ...string) int {

	affect, err := d.Engine.Where(" id = ? ", id).Cols(cols...).Update(bean)
	if utils.CheckError(err, affect) {
		return e.Success
	}
	return e.Errors
}
