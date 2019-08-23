package wx

import (
	wx2 "comadmin/model/wx"
	"comadmin/tools/utils"
	"context"
	"encoding/json"
)

func (d Dao) findBizList() (interface{}, int) {
	type WeiXin struct {
		Biz  string `json:"biz"`
		Name string `json:"name"`
	}

	wx := make([]WeiXin, 0)

	count, err := d.engine.FindAndCount(&wx)
	if utils.CheckError(err, count) {
		return wx, int(count)
	}
	return nil, 0
}

func (d Dao) findApi() (interface{}, int) {

	wx := make([]wx2.Api, 0)

	count, err := d.engine.FindAndCount(&wx)
	if utils.CheckError(err, count) {
		return wx, int(count)
	}
	return nil, 0
}

func (d Dao) insertDetail(id string, bean interface{}) int {
	data := ""
	marshal, err := json.Marshal(bean)
	if err != nil {
		data = string(marshal)
	}
	type A struct {
		Name string `json:"name"`
	}
	do, err := d.es.Index().Index("xxx").OpType("xx").Id(id).BodyString(data).Do(context.Background())
	if utils.CheckError(err, do) {

	}
	return 0
}
