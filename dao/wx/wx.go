package wxd

import (
	"comadmin/model/wx"
	"comadmin/pkg/config"
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
	do, err := d.es.Index().Index(index).OpType(tp).Id(id).BodyString(data).Do(context.Background())
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
	t := false
	if detail.Keywords != "" {
		split := strings.Split(detail.Keywords, ",")
		for _, v := range split {
			query.Should(elastic.NewMatchPhraseQuery("text", v))
		}
		t = true
	}
	if detail.Title != "" {
		query.Should(elastic.NewMatchPhraseQuery("title", detail.Title))
		t = true
	}
	if detail.From != "" {
		query.Filter(elastic.NewRangeQuery("ptime").Gte(detail.From).Lte(detail.To))
	}
	if detail.Biz != "" {
		query.Should(elastic.NewMatchQuery("biz", detail.Biz))
		t = true
	}

	if detail.Pn <= 1 {
		detail.Pn = 1
	}

	if detail.Ps <= 0 || detail.Ps > 50 {
		detail.Ps = 50
	}

	if !t {
		query.Should(elastic.NewMatchAllQuery())
	}
	//查询单个id的文档
	//result, err := d.es.Get().Index(config.EsIndex).Id("vfNTx2wBXVO2c-XIzCHy").Do(context.Background())
	//bytes, err := result.Source.MarshalJSON()
	//fmt.Print(string(bytes),err)
	//查询所有
	//query := elastic.NewMatchAllQuery()
	//
	//result, err := d.es.Search().Index(config.EsIndex).Query(query).Do(context.Background())
	//if utils.CheckError(err, result) {
	//	array := make([]interface{}, 0)
	//	for _, hit := range result.Hits.Hits {
	//		array = append(array, hit.Source)
	//	}
	//	return array, result.Hits.TotalHits.Value
	//}

	field := elastic.NewFetchSourceContext(true)
	field.Include("id", "text", "text_style", "biz", "author", "original", "word_cloud", "summary")

	result, err := d.es.Search().FetchSourceContext(field).Index(config.EsIndex).Query(query).Do(context.Background())
	if utils.CheckError(err, result) {
		array := make([]interface{}, len(result.Hits.Hits))
		for index, hit := range result.Hits.Hits {
			//var r ret
			//json.Unmarshal(hit.Source, &r)
			array[index] = hit.Source
		}
		return array, result.Hits.TotalHits.Value
	}
	return nil, 0
}
