package wx

import (
	wx2 "comadmin/model/wx"
	"comadmin/tools/utils"
)

func (d Dao) findBizList() (interface{}, int) {
	type Wx struct {
		Biz  string `json:"biz"`
		Name string `json:"name"`
	}

	wx := make([]Wx, 0)

	count, err := d.Engine.FindAndCount(&wx)
	if utils.CheckError(err, count) {
		return wx, int(count)
	}
	return nil, 0
}

func (d Dao) findApi() (interface{}, int) {

	wx := make([]wx2.API, 0)

	count, err := d.Engine.FindAndCount(&wx)
	if utils.CheckError(err, count) {
		return wx, int(count)
	}
	return nil, 0
}
