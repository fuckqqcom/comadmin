package wx

import (
	wx2 "comadmin/model/wx"
	"comadmin/tools/utils"
)

func (d Dao) findBizList() (interface{}, int) {
	type WeiXin struct {
		Biz  string `json:"biz"`
		Name string `json:"name"`
	}

	wx := make([]WeiXin, 0)

	count, err := d.Engine.FindAndCount(&wx)
	if utils.CheckError(err, count) {
		return wx, int(count)
	}
	return nil, 0
}

func (d Dao) findApi() (interface{}, int) {

	wx := make([]wx2.Api, 0)

	count, err := d.Engine.FindAndCount(&wx)
	if utils.CheckError(err, count) {
		return wx, int(count)
	}
	return nil, 0
}

func (d Dao) insertDetail() {
	d.es.Create()
}
