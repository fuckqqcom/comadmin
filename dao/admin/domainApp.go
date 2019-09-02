package admin

import (
	"comadmin/model/admin"
	"comadmin/tools/utils"
	"fmt"
)

//func (d Dao) updateApp(app admin.DomainApp) int {
//	cols := make([]string, 0)
//	if app.Name != "" {
//		cols = append(cols, "name")
//	}
//	if app.Status != 0 {
//		cols = append(cols, "status")
//	}
//	affect, err := d.Engine.Where(" id = ? ", app.Id).Cols(cols...).Update(app)
//	if utils.CheckError(err, affect) {
//		return e.Success
//	}
//	return e.Errors
//}

/**
join联查
*/
func (d Dao) findApp(app admin.DomainApp, pn, ps int) (interface{}, int) {
	type Result struct {
		Name   string `json:"name"`
		Id     string `json:"id"`
		Status int    `json:"status"`
		Ctime  string `json:"ctime" `
		Mtime  string `json:"mtime" `
	}
	ret := make([]Result, 0)
	sql := fmt.Sprintf("select id,`name`,status,ctime,mtime from domain where 1= 1 and did = '%s' ", app.Did)
	if app.Name != "" {
		sql += "  and `name` like " + `"%` + app.Name + `%" `
	}
	if app.Status != 0 {
		sql += fmt.Sprintf(" and status == %d ", app.Status)
	}
	count, err := d.engine.SQL(sql).Limit(ps, (pn-1)*ps).FindAndCount(&ret)
	if !utils.CheckError(err, count) {
		return nil, 0
	}
	return ret, int(count)
}
