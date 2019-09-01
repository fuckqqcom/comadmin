package wxd

import (
	"comadmin/model/wx"
	"comadmin/pkg/config"
	"comadmin/pkg/e"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/xormplus/xorm"
)

type DbHandler interface {
	wxHandler
	jobHandler
	daoHandler
}

type wxHandler interface {
}

type daoHandler interface {
	Create(interface{}) int                                   //添加数据接口
	Update(interface{}, []string, map[string]interface{}) int //修改
	Delete(interface{}, interface{}) int
	Exist(interface{}, map[string]interface{}) bool
	List(interface{}, int, int) (interface{}, interface{})         //查询,返回列表
	Get(interface{}, []string, map[string]interface{}) interface{} //查询单个对象,返回对象
	//CreateOrUpdate(interface{}, []string, map[string]interface{}) int  //创建或者更新 存在就更新,不存在就创建
	//CreateOrDiscard(interface{}, []string, map[string]interface{}) int //创建 如果存在就丢弃
}

type jobHandler interface {
	Register(interface{}) (interface{}, int)
}

type Dao struct {
	c      config.Config
	engine *xorm.Engine
	es     *elastic.Client
	rs     *redis.Client
}

var (
	index = &config.EsIndex
)

func NewDb(path string) *Dao {
	return &Dao{engine: config.EngDb, es: config.EsClient, rs: config.RedisClient, c: config.NewConfig(path)}
}

//接口   结构体
var _ DbHandler = Dao{}

//后台创建数据
func (d Dao) Create(i interface{}) int {
	switch t := i.(type) {
	case wx.WeiXin, wx.WeiXinCount:
		return d.create(t)
	case wx.WeiXinDetail:
		return d.insertArticleDetail(t.Id, t)
	case wx.WeiXinList:
		return d.addQueue(t)

	default:
		fmt.Println("create other ...")
		return e.Errors
	}
}

//func (d Dao) CreateOrUpdate(bean interface{}, cols []string, colsValue map[string]interface{}) int {
//	switch t := bean.(type) {
//	case wx.WeiXinCount:
//		//存在就更新
//		if d.existRecord(t, &wx.WeiXinCount{Biz: t.Biz, ArticleId: t.ArticleId}) {
//			return d.updateRecord(t, cols, colsValue)
//		}
//		return d.create(t)
//	}
//	return 0
//}

func (d Dao) List(i interface{}, ps, pn int) (interface{}, interface{}) {

	switch t := i.(type) {
	case wx.WeiXinParams:
		return d.findArticle(t, ps, pn)
	case wx.WeiXinList:
		return d.popQueue(t, ps)
	case wx.WeiXin:
		return d.findBizList(t.MobileId, ps, pn)
	case wx.Api:
		return d.findApi(ps, pn)
		//查询近7天数据
	case wx.Nearly7Day:
		return d.wxNearly7Day(t, ps, pn)

	default:
		fmt.Println("update other ...")
		return nil, e.Errors
	}
}

func (d Dao) Delete(interface{}, interface{}) int {
	return 0

}

func (d Dao) Exist(bean interface{}, m map[string]interface{}) bool {
	switch t := bean.(type) {
	case *wx.WeiXinCount:
		return d.existRecord(t)
	default:
		return false
	}
}

func (d Dao) Update(bean interface{}, cols []string, m map[string]interface{}) int {
	//todo 怎么优化
	switch t := bean.(type) {
	case wx.UserWx:
		fmt.Println(t)
		return 0
	case wx.Job:
		return d.update(t.Id, t, cols...)
	case wx.WeiXinCount:
		return d.updateRecord(bean, cols, m)
	default:
		fmt.Println("update other ...")
		return e.Errors
	}
}

//func (d Dao) CreateOrDiscard(bean interface{}, cols []string, colsValue map[string]interface{}) int {
//	switch t := bean.(type) {
//	case wx.AddWxParams:
//		if !d.existRecord(t, &wx.WeiXinList{HashId: t.HashId}) {
//			w := wx.WeiXinList{
//				HashId:    t.HashId,
//				SourceUrl: t.SourceUrl,
//				Url:       t.Url,
//				Title:     t.Title,
//				Ptime:     t.Ptime,
//				Biz:       t.Biz,
//				Digest:    t.Digest,
//				Original:  t.Original,
//				Type:      t.Type,
//				DelFlag:   t.DelFlag,
//				Cover:     t.Cover,
//			}
//			return d.create(w)
//		} else {
//			return e.ExistError
//		}
//	default:
//		return e.Errors
//	}
//}

func (d Dao) Get(bean interface{}, cols []string, colsValue map[string]interface{}) interface{} {
	switch t := bean.(type) {
	case *wx.Job:
		return d.get(t, cols, colsValue)
	default:
		return nil
	}
}

func (d Dao) Register(bean interface{}) (interface{}, int) {
	switch t := bean.(type) {
	case wx.Job:
		return d.register(t)
	default:
		return nil, e.Errors
	}
}
