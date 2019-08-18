package admin

import (
	"comadmin/model/admin"
	"comadmin/tools/utils"
	"fmt"
)

func (d Dao) findDomain(domain admin.Domain, pn, ps int) (interface{}, int) {
	type Result struct {
		Name   string `json:"name"`
		Id     string `json:"id"`
		Status int    `json:"status"`
		Ctime  string `json:"ctime" `
		Mtime  string `json:"mtime" `
	}
	ret := make([]Result, 0)
	sql := "select id,`name`,status,ctime,mtime from domain where 1= 1  "
	if domain.Name != "" {
		sql += "  and `name` like " + `"%` + domain.Name + `%" `
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

//func (d Dao) deleteDomain(domain admin.Domain) int {
//	affect, err := d.Engine.Where(" id = ? ", domain.Id).Delete(&domain)
//	if utils.CheckError(err, affect) {
//		return e.Success
//	}
//	return e.Errors
//}

//func (d Dao) updateDomain(domain admin.Domain) int {
//	cols := make([]string, 0)
//	if domain.Name != "" {
//		cols = append(cols, "name")
//	}
//	if domain.Status != 0 {
//		cols = append(cols, "status")
//	}
//	affect, err := d.Engine.Where(" id = ? ", domain.Id).Cols(cols...).Update(domain)
//	if utils.CheckError(err, affect) {
//		return e.Success
//	}
//	return e.Errors
//}

//func (d Dao) findById(domain admin.Domain) int {
//	affect, err := d.Engine.Exist(domain)
//	if utils.CheckError(err, affect) && affect {
//		return e.Success
//	}
//	return e.ExistError
//}
