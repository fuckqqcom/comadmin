package wx

import (
	"comadmin/model/wx"
	"comadmin/pkg/config"
	"comadmin/pkg/e"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/xormplus/xorm"
)

type DbHandler interface {
	Create(interface{}) int      //添加数据接口
	FindBiz() (interface{}, int) //获取公号信息
	FindApi() (interface{}, int) //获取接口
	PostData(interface{}) int    //提交数据接口
} //查询

type Dao struct {
	c      config.Config
	engine *xorm.Engine
	es     *elastic.Client
}

func NewDb(path string) *Dao {
	return &Dao{engine: config.EngDb, c: config.NewConfig(path)}
}

//接口   结构体
var _ DbHandler = Dao{}

//后台创建数据
func (d Dao) Create(i interface{}) int {
	switch t := i.(type) {
	case wx.WeiXin:
		return d.create(t)
	case wx.WeiXinDetail:
		return d.insertDetail(t.ArticleId, t)
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
