package admin

import (
	"comadmin/model/admin"
	"comadmin/tools/utils"
	"fmt"
)

func (d Dao) findRole(r admin.Role, pn, ps int) (interface{}, int) {
	type Result struct {
		Name   string `json:"name"`
		Id     string `json:"id"`
		Status int    `json:"status"`
		Ctime  string `json:"ctime" `
		Mtime  string `json:"mtime" `
	}
	ret := make([]Result, 0)
	sql := fmt.Sprintf("select id,`name`,status,ctime,mtime from domain where 1= 1 and id = '%s' ", r.Id)
	if r.Name != "" {
		sql += "  and `name` like " + `"%` + r.Name + `%" `
	}
	if r.Status != 0 {
		sql += fmt.Sprintf(" and status == %d ", r.Status)
	}
	count, err := d.Engine.SQL(sql).Limit(ps, (pn-1)*ps).FindAndCount(&ret)
	if !utils.CheckError(err, count) {
		return nil, 0
	}
	return ret, int(count)
}
