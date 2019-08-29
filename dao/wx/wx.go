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

func (d Dao) findBizList(mobileId string) (interface{}, int) {
	type WeiXin struct {
		Id   string `json:"id"`
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
	if err == nil {
		data = string(marshal)
	}
	type A struct {
		Name string `json:"name"`
	}
	do, err := d.es.Index().Index(*index).Id(id).BodyString(data).Do(context.Background())
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

func (d Dao) wxList(list wx.WeiXinList, pn, ps int) (interface{}, int) {
	sql := "select biz , url from wei_xin_list where 1=1  "
	if list.Biz != "" {
		sql += fmt.Sprintf(" and biz = '%s'", list.Biz)
	}
	sql += fmt.Sprintf(" and ptime >= '%s' ", utils.Time2Str(time.Now().AddDate(0, 0, -7), "2006-01-02 15:04:05"))
	type ret struct {
		Biz string
		Url string
	}
	w := make([]ret, 0)
	count, err := d.engine.SQL(sql).Limit(ps, (pn-1)*ps).FindAndCount(&w)
	if utils.CheckError(err, count) {
		return w, int(count)
	}
	return nil, 0
}

//查询数据
func (d Dao) findArticle(detail wx.WeiXinParams, pn, ps int) (interface{}, interface{}) {
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

/**
查询公号是否存在
*/

func (d Dao) existWx(u wx.UserWx) int {
	w := &wx.WeiXin{Name: u.Name}
	affect, err := d.engine.Exist(w)
	//存在就返回，不存在就创建
	if utils.CheckError(err, affect) && affect {
		return e.ExistError
	} else {
		//
		return d.create(u)
	}
}

/**
审核数据接口(讲用户提交的数据同步到抓取数据) 审核通过的时候手动把biz补到页面中
*/

func (d Dao) verify(u wx.UserWx, id interface{}, cols ...string) int {
	//在http层就限制死  status == 1
	w := wx.WeiXin{Biz: u.Biz, Name: u.Name, Forbid: 1}
	if d.create(w) == e.Success && d.update(id, u, cols...) == e.Success {
		return e.Success
	} else {
		return e.Errors
	}

}

//更新

func (d Dao) updateWx(id interface{}, bean interface{}, cols ...string) int {

	affect, err := d.engine.Where(" id = ? ", id).Cols(cols...).Update(bean)
	if utils.CheckError(err, affect) {
		return e.Success
	}
	return e.Errors
}
