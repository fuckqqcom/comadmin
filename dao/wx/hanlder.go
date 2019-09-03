package wxd

import (
	"comadmin/model/wx"
	"comadmin/pkg/config"
	"comadmin/pkg/e"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/xormplus/xorm"
	"log"
)

type DbHandler interface {
	wxHandler
	jobHandler
	daoHandler
}

type wxHandler interface {
}

type daoHandler interface {
	Add(interface{}) int                                      //添加数据接口
	Update(interface{}, []string, map[string]interface{}) int //修改
	Delete(interface{}, map[string]interface{}) int
	Exist(interface{}, map[string]interface{}) bool
	FindOne(interface{}, map[string]interface{}, int, int, string) (interface{}, interface{}) //查询,返回列表 map装载查询条件  单表查询
	Get(interface{}, []string, map[string]interface{}) interface{}                            //查询单个对象,返回对象
	//AddOrUpdate(interface{}, []string, map[string]interface{}) int  //创建或者更新 存在就更新,不存在就创建
	//AddOrDiscard(interface{}, []string, map[string]interface{}) int //创建 如果存在就丢弃
}

type jobHandler interface {
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
func (d Dao) Add(i interface{}) int {
	switch t := i.(type) {
	case wx.WeiXin, wx.WeiXinCount, wx.Job:
		//mysql
		return d.add(t)
	case wx.WeiXinDetail:
		//入库es
		return d.addArticleDetail(t.Id, t)
	case wx.WeiXinList:
		//入库队列
		return d.addQueue(t)
	default:
		fmt.Println("add other ...")
		return e.Errors
	}
}

func (d Dao) Delete(bean interface{}, m map[string]interface{}) int {
	switch t := bean.(type) {
	case *wx.WeiXinCount, *wx.WeiXinList, *wx.Job:
		return d.delete(t, m)
	default:
		log.Println("delete other error")
		return e.Errors
	}

}

func (d Dao) Exist(bean interface{}, m map[string]interface{}) bool {
	switch t := bean.(type) {
	case *wx.WeiXinCount, *wx.WeiXinList, *wx.Job:
		return d.exist(t)
	default:
		log.Println("exist other error")
		return true
	}
}

func (d Dao) Update(bean interface{}, cols []string, m map[string]interface{}) int {
	//todo 怎么优化
	switch t := bean.(type) {
	case wx.Job, wx.WeiXinCount:
		return d.update(t, cols, m)
	default:
		fmt.Println("update other ...")
		return e.Errors
	}
}

func (d Dao) Get(bean interface{}, cols []string, colsValue map[string]interface{}) interface{} {
	switch t := bean.(type) {
	case *wx.Job:
		return d.get(t, cols, colsValue)
	default:
		return nil
	}
}

func (d Dao) FindOne(i interface{}, m map[string]interface{}, ps, pn int, orderQuery string) (interface{}, interface{}) {

	switch t := i.(type) {
	case wx.WeiXinParams:
		return d.articles(t, ps, pn)
	case wx.WeiXinList:
		return d.queues(t, ps, pn)
	case wx.ApiParams:
		w := make([]wx.Api, 0)
		return d.find(&w, m, ps, pn, orderQuery)
	case wx.Nearly7Day:
		type WeiXinList struct {
			Biz    string `json:"biz"`
			Url    string `json:"url"`
			HashId string `json:"hash_id"`
		}
		w := make([]WeiXinList, 0)
		return d.find(&w, m, ps, pn, orderQuery)
	case wx.BizParams:
		type WeiXin struct {
			Id   string `json:"id"`
			Biz  string `json:"biz"`
			Name string `json:"name"`
		}
		w := make([]WeiXin, 0)
		return d.find(&w, m, ps, pn, orderQuery)
	default:
		fmt.Println("update other ...")
		return nil, e.Errors
	}
}
