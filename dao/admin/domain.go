package admin

import (
	"comadmin/model/admin"
	"comadmin/tools/utils"
	"fmt"
)

func (d Dao) FindDomain(domain admin.Domain, pn, ps int) (interface{}, int) {
	type Result struct {
		Name   string `json:"name"`
		Id     string `json:"id"`
		Status int    `json:"status"`
		CTime  string `json:"ctime" xorm:"created"`
		MTime  string `json:"mtime" xorm:"updated" `
	}
	ret := make([]Result, 0)
	sql := "select * from domain where "
	if domain.Name != "" {
		sql += " name like %" + domain.Name + " % "
	}
	if domain.Status != 0 {
		sql += fmt.Sprintf(" and status == %d ", domain.Status)
	}

	count, err := d.Engine.SQL(sql).Limit(ps, (pn-1)*ps).FindAndCount(&ret)
	if !utils.CheckError(err, count) {
		return nil, 0
	}

	return ret, int(count)
}
