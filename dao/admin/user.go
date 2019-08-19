package admin

import (
	"comadmin/model/admin"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"fmt"
)

//TODO 这里需要加入 域显示中文名称  app显示中文名称
func (d Dao) findUser(u admin.User, pn, ps int) (interface{}, int) {

	/**
	Id     string `json:"id"`     //根据id查询
	Name   string `json:"name"`   //根据name模糊查询
	Status int    `json:"status"` //根据状态精确查询
	Phone  string `json:"phone"`  //根据手机号码模糊查询
	Did    string `json:"did"`    //根据域id精确查询
	Aid    string `json:"aid"`    //根据应用id精确查询
	*/

	ret := make([]admin.UserResult, 0)
	sql := "select id,`name`,status,ctime,mtime from domain where 1= 1  "
	if u.Name != "" {
		sql += "  and `name` like " + `"%` + u.Name + `%" `
	}
	if u.Status != 0 {
		sql += fmt.Sprintf(" and status == %d ", u.Status)
	}

	if u.Phone != "" {
		sql += "  and `phone` like " + `"%` + u.Phone + `%" `
	}

	if u.Did != "" {
		sql += fmt.Sprintf(" and did == '%s' ", u.Did)
	}

	if u.Aid != "" {
		sql += fmt.Sprintf(" and aid == '%s' ", u.Aid)
	}

	count, err := d.Engine.SQL(sql).OrderBy(" ctime desc ").Limit(ps, (pn-1)*ps).FindAndCount(&ret)
	if !utils.CheckError(err, count) {
		return nil, 0
	}
	return ret, int(count)
}

//登录接口
func (d Dao) login(u admin.User) int {
	affect, err := d.Engine.Get(u)
	if utils.CheckError(err, affect) && affect {
		return e.Success
	}
	return e.NotExistError
}
