package wxd

import (
	"comadmin/model/wx"
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

func (d Dao) exist(bean interface{}) bool {
	affect, err := d.engine.Exist(bean)
	if utils.CheckError(err, affect) {
		return affect
	}
	return false
}

/**
查询记录是否存在
*/
func (d Dao) update(bean interface{}, cols []string, m map[string]interface{}) int {
	query, value := utils.QueryCols(m)
	affect, err := d.engine.Where(query, value...).Cols(cols...).Update(bean)
	if utils.CheckError(err, affect) && affect >= 1 {
		return e.Success
	}
	return e.UpdateError
}

//删除
func (d Dao) delete(bean interface{}, m map[string]interface{}) int {
	query, value := utils.QueryCols(m)
	affect, err := d.engine.Where(query, value...).Delete(&bean)
	if utils.CheckError(err, affect) {
		return e.Success
	}
	return e.Errors
}

//创建
func (d Dao) add(bean interface{}) int {
	affect, err := d.engine.Insert(bean)
	if utils.CheckError(err, affect) {
		return e.Success
	}
	return e.Errors
}

//
////通过id查找
//
//func (d Dao) findById(bean interface{}) int {
//	affect, err := d.engine.Exist(bean)
//	if utils.CheckError(err, affect) && affect {
//		return e.Success
//	}
//	return e.ExistError
//}

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

func (d Dao) get(bean interface{}, cols []string, m map[string]interface{}) interface{} {

	//query, value := utils.QueryCols(m)
	join := strings.Join(cols, ",")
	affect, err := d.engine.Select(join).Get(bean)
	if utils.CheckError(err, affect) && affect {
		return bean
	}
	return nil
}

//单表查询 暂时没有排序
func (d Dao) find(w interface{}, queryValue map[string]interface{}, ps, pn int, orderQuery string) (interface{}, int) {
	query, value := utils.QueryCols(queryValue)
	count, err := d.engine.Where(query, value...).OrderBy(orderQuery).Limit(ps, (pn-1)*ps).FindAndCount(w)
	if utils.CheckError(err, count) {
		return w, int(count)
	}
	return nil, 0
}

func (d Dao) findBizById(id int) string {
	w := new(wx.WeiXin)
	affect, err := d.engine.Where("id = ? ", id).Get(w)
	if utils.CheckError(err, affect) {
		return w.Biz
	}
	return ""
}
