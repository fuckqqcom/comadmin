package wxd

import (
	"comadmin/model/wx"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

//func (d Dao) bizs(m map[string]interface{}, ps, pn int) (interface{}, int) {
//	type WeiXin struct {
//		Id   string `json:"id"`
//		Biz  string `json:"biz"`
//		Name string `json:"name"`
//	}
//
//	w := make([]WeiXin, 0)
//	query, value := utils.QueryCols(m)
//	count, err := d.engine.Where(query, value...).Limit(ps, (pn-1)*ps).OrderBy("mtime desc ").FindAndCount(&w)
//	if utils.CheckError(err, count) {
//		return w, int(count)
//	}
//	return nil, 0
//}

//func (d Dao) apis(m map[string]interface{}, ps, pn int) (interface{}, int) {
//
//	type Api struct {
//		Id       string `json:"id"`
//		Name     string `json:"name"`
//		Url      string `json:"url"`
//		Category int    `json:"category"` // 1是阅读点赞接口  2是详情接口 3是其他接口等
//	}
//
//	w := make([]Api, 0)
//	query, value := utils.QueryCols(m)
//	count, err := d.engine.Where(query, value).Limit(ps, (pn-1)*ps).FindAndCount(&w)
//	if utils.CheckError(err, count) {
//		return w, int(count)
//	}
//	return nil, 0
//}

//插入es数据
func (d Dao) addArticleDetail(id string, bean interface{}) int {
	data := ""
	marshal, err := json.Marshal(bean)
	if err == nil {
		data = string(marshal)
	}
	fmt.Println("data--->", data)
	type A struct {
		Name string `json:"name"`
	}
	do, err := d.es.Index().Index(*index).Id(id).BodyString(data).Do(context.Background())
	if utils.CheckError(err, do) {
		return e.Success
	}
	return e.Errors
}

//func (d Dao) wxNearly7Days(m map[string]interface{}, ps, pn int) (interface{}, int) {
//
//	type WeiXinList struct {
//		Biz string
//		Url string
//	}
//	w := make([]WeiXinList, 0)
//	query, value := utils.QueryCols(m)
//	count, err := d.engine.Where(query, value...).Limit(ps, (pn-1)*ps).FindAndCount(&w)
//	if utils.CheckError(err, count) {
//		return w, int(count)
//	}
//	return nil, 0
//}

//查询数据
func (d Dao) articles(detail wx.WeiXinParams, ps, pn int) (interface{}, interface{}) {
	/**
	这里拼接es sql


	select a from table where forbid = 1 and ( title like "%aaa%" or text like "%bbb%")

	https://www.tuicool.com/articles/NFVzeqy
	*/
	//zuiwaiceng  query bool
	query := elastic.NewBoolQuery()
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewTermQuery("forbid", 1))
	query.Filter(boolQuery)
	filter := elastic.NewBoolQuery()
	newBoolQuery := elastic.NewBoolQuery()
	queryFilter := filter.Must(newBoolQuery)
	if detail.Title != "" {
		newBoolQuery.Should(elastic.NewWildcardQuery("title", detail.Title))
	}

	if detail.Keywords != "" {
		split := strings.Split(detail.Keywords, ",")
		for _, v := range split {
			newBoolQuery.Should(elastic.NewWildcardQuery("title", v))
			newBoolQuery.Should(elastic.NewWildcardQuery("text", v))
		}
	}
	if detail.Biz != "" {
		query.Must(elastic.NewMatchQuery("biz", detail.Biz))
	}
	const t = "2006-01-02 15:04:05"
	if detail.From != "" && detail.To != "" {
		query.Filter(elastic.NewRangeQuery("ptime").Gte(utils.Str2Time(t, detail.From)).Lte(utils.Str2Time(t, detail.To)))
	} else if detail.From != "" {
		query.Filter(elastic.NewRangeQuery("ptime").Gte(utils.Str2Time(t, detail.From)).Lte(time.Now()))
	}

	query.Filter(queryFilter)

	field := elastic.NewFetchSourceContext(true)
	field.Include("id", "text", "text_style", "biz", "author", "original", "word_cloud", "summary", "title", "forbid")
	source, _ := query.Source()
	utils.PrintQuery(source)
	result, err := d.es.Search().FetchSourceContext(field).Index(*index).Query(query).Do(context.Background())
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
