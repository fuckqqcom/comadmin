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
	Create(interface{}) int           //添加数据接口
	Update(interface{}, []string) int //修改

	FindBiz() (interface{}, int) //获取公号信息
	FindApi() (interface{}, int) //获取接口
	PostData(interface{}) int    //提交数据接口
	Find(i interface{}, pn, ps int) (interface{}, interface{})
} //查询

type Dao struct {
	c      config.Config
	engine *xorm.Engine
	es     *elastic.Client
	rs     *redis.Client
}

var (
	index = config.EsIndex
	tp    = config.EsType
)

func NewDb(path string) *Dao {
	return &Dao{engine: config.EngDb, es: config.EsClient, rs: config.RedisClient, c: config.NewConfig(path)}
}

//接口   结构体
var _ DbHandler = Dao{}

//后台创建数据
func (d Dao) Create(i interface{}) int {
	switch t := i.(type) {
	case wx.WeiXin:
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

func (d Dao) FindBiz() (interface{}, int) {
	return d.findBizList()
}

func (d Dao) FindApi() (interface{}, int) {
	return d.findApi()
}

func (d Dao) PostData(i interface{}) int {
	switch t := i.(type) {
	case wx.WeiXinCount:
		return d.create(t)
	default:
		fmt.Println("create other ...")
		return e.Errors
	}
}

func (d Dao) Find(i interface{}, pn, ps int) (interface{}, interface{}) {

	switch t := i.(type) {
	case wx.WeiXinParams:
		return d.findArticle(t, pn, ps)
	case wx.WeiXinList:
		return d.popQueue(t, ps)
	default:
		fmt.Println("update other ...")
		return nil, e.Errors
	}
}

func (d Dao) Update(i interface{}, cols []string) int {
	//todo 怎么优化
	switch t := i.(type) {
	case wx.UserWx:
		fmt.Println(t)
		return 0
	default:
		fmt.Println("update other ...")
		return e.Errors
	}
}
