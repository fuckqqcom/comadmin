package wxd

import (
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"strings"
)

/**
一些公用的操作,比如删除
*/

/**
查询记录是否存在
*/

func (d Dao) existRecord(bean interface{}) bool {
	//query, value := utils.QueryCols(m)
	affect, err := d.engine.Exist(bean)
	//affect, err := d.engine.Exist(query)
	if utils.CheckError(err, affect) {
		return affect
	}
	return false
}

/**
查询记录是否存在
*/
func (d Dao) updateRecord(bean interface{}, cols []string, m map[string]interface{}) int {
	//query := ""
	//value := make([]interface{}, 0)
	//count := 0
	//if m != nil {
	//	for k, v := range m {
	//		if count == 0 {
	//			query += k + " = ? "
	//		} else {
	//			query += " and " + k + " = ? "
	//		}
	//		value = append(value, v)
	//		count += 1
	//	}
	//}
	query, value := utils.QueryCols(m)
	affect, err := d.engine.Where(query, value...).Cols(cols...).Update(bean)
	if utils.CheckError(err, affect) && affect >= 1 {
		return e.Success
	}
	return e.UpdateError
}

//删除 通过id删除
func (d Dao) delete(id, bean interface{}) int {
	affect, err := d.engine.Where(" id = ? ", id).Delete(&bean)
	if utils.CheckError(err, affect) {
		return e.Success
	}
	return e.Errors
}

//创建
func (d Dao) create(bean interface{}) int {
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

func (d Dao) update(id interface{}, bean interface{}, cols ...string) int {

	affect, err := d.engine.Where(" id = ? ", id).Cols(cols...).Update(bean)
	if utils.CheckError(err, affect) {
		return e.Success
	}
	return e.Errors
}

func (d Dao) get(bean interface{}, cols []string, m map[string]interface{}) interface{} {

	//query, value := utils.QueryCols(m)
	join := strings.Join(cols, ",")
	affect, err := d.engine.Select(join).Get(bean)
	if utils.CheckError(err, affect) && affect {
		return bean
	}
	return nil

}
