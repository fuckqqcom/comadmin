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
	//在所有未被管放拉黑的数据中查找
	query.Must(elastic.NewMatchQuery("forbid", 1))

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
