package admin

import (
	"comadmin/pkg/e"
	"comadmin/tools/utils"
)

/**
一些公用的操作,比如删除
*/

//删除 通过id删除
func (d Dao) delete(bean interface{}, m map[string]interface{}) int {
	query, value := utils.QueryCols(m)
	affect, err := d.engine.Where(query, value...).Delete(bean)
	if utils.CheckError(err, affect) && affect >= 1 {
		return e.Success
	}
	return e.Errors
}

//创建
//func (d Dao) create(bean interface{}) int {
//	affect, err := d.Engine.Insert(bean)
//	if utils.CheckError(err, affect) {
//		return e.Success
//	}
//	return e.Errors
//}

func (d Dao) add(bean interface{}) int {
	affect, err := d.engine.Insert(bean)
	if utils.CheckError(err, affect) {
		return e.Success
	}
	return e.Errors
}

//通过id查找

func (d Dao) findById(bean interface{}) int {
	affect, err := d.engine.Exist(bean)
	if utils.CheckError(err, affect) && affect {
		return e.Success
	}
	return e.ExistError
}

/**
用过id更新  name 和status状态
*/

//func (d Dao) update(id interface{}, bean interface{}, cols ...string) int {
//
//	affect, err := d.engine.Where(" id = ? ", id).Cols(cols...).Update(bean)
//	if utils.CheckError(err, affect) {
//		return e.Success
//	}
//	return e.Errors
//}
func (d Dao) update(bean interface{}, cols []string, m map[string]interface{}) int {
	query, value := utils.QueryCols(m)
	affect, err := d.engine.Where(query, value...).Cols(cols...).Update(bean)
	if utils.CheckError(err, affect) && affect >= 1 {
		return e.Success
	}
	return e.UpdateError
}

//单表查询
func (d Dao) findOneTable(w interface{}, queryValue map[string]interface{}, ps, pn int) (interface{}, int) {
	query, value := utils.QueryCols(queryValue)
	count, err := d.engine.Where(query, value...).Limit(ps, (pn-1)*ps).FindAndCount(w)
	if utils.CheckError(err, count) {
		return w, int(count)
	}
	return nil, 0
}

func (d Dao) exist(bean interface{}) bool {
	affect, err := d.engine.Exist(bean)
	if utils.CheckError(err, affect) {
		return affect
	}
	return false
}
