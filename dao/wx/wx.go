package wxd

import (
	"comadmin/model/wx"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"strings"
)

func (d Dao) findBizList() (interface{}, int) {
	type WeiXin struct {
		Biz  string `json:"biz"`
		Name string `json:"name"`
	}

	w := make([]WeiXin, 0)

	count, err := d.engine.FindAndCount(&w)
	if utils.CheckError(err, count) {
		return w, int(count)
	}
	return nil, 0
}

func (d Dao) findApi() (interface{}, int) {

	w := make([]wx.Api, 0)

	count, err := d.engine.FindAndCount(&w)
	if utils.CheckError(err, count) {
		return w, int(count)
	}
	return nil, 0
}

//插入es数据
func (d Dao) insertArticleDetail(id string, bean interface{}) int {
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
		return e.Success
	}
	return e.Errors
}

//插入列表数据
func (d Dao) insertArticleList(bean interface{}) int {
	affect, err := d.engine.Insert(bean)
	if utils.CheckError(err, affect) {
		return e.Success
	}
	return e.Errors
}

//查询数据
func (d Dao) findArticle(detail wx.WeiXinParams, pn, ps int) (interface{}, interface{}) {
	/**
	这里拼接es sql
	*/
	query := elastic.NewBoolQuery()
	if detail.Keywords != "" {
		split := strings.Split(detail.Keywords, ",")
		for _, v := range split {
			query.Should(elastic.NewMatchPhraseQuery("text", v))

		}
	}
	if detail.Title != "" {
		query.Should(elastic.NewMatchPhraseQuery("title", detail.Title))
	}
	if detail.From != "" {
		query.Filter(elastic.NewRangeQuery("ptime").Gte(detail.From).Lte(detail.To))
	}
	if detail.Biz != "" {
		query.Should(elastic.NewMatchQuery("biz", detail.Biz))
	}

	if detail.Pn <= 1 {
		detail.Pn = 1
	}

	if detail.Ps < 0 || detail.Ps > 50 {
		detail.Ps = 50
	}
	result, err := d.es.Search("index").SearchType("type").Query(query).From(detail.Pn * detail.Ps).Size(detail.Ps).Pretty(true).Do(context.Background())
	if utils.CheckError(err, result) {
		return result.Hits.Hits, result.Hits.TotalHits
	}

	return nil, 0
}
