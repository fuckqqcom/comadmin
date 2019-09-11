package wx

import "time"

//公号属性
/**
查询的时候拿自增主键id当做传过来的id转下 找到对应的biz
*/
type WeiXin struct {
	Id       int64     //自增中间 id
	MobileId string    `json:"mobile_id"`             //手机id
	Wid      string    `json:"wid" xorm:"wid"`        //微信id
	Biz      string    `json:"biz"`                   //公号biz
	Name     string    `json:"name"`                  //名称
	Key      string    `json:"key"`                   //微信的key
	Url      string    `json:"url"`                   //img url
	Desc     string    `json:"detail"`                //公号描述
	Count    int       `json:"count"`                 //公号文章数
	Uin      string    `json:"uin"`                   //uin
	Ctime    time.Time `json:"ctime" xorm:"created" ` //创建时间
	Mtime    time.Time `json:"mtime" xorm:"updated" ` //最后一次更新时间
	Forbid   int       `json:"forbid"`                //公号是否被微信官方搞事了
	Note     string    `json:"note"`                  //被管放禁用后的提示 (生于xxx,卒于xxx)
	Default  int       `json:"default"`               //默认匿名用户展示
}

//阅读量和点赞量
type WeiXinCount struct {
	Id         int64     //自增主键id
	Biz        string    `json:"biz"`
	ArticleId  string    `json:"article_id"`            //文章id
	ReadCount  int       `json:"read_count"`            //阅读量
	ThumbCount int       `json:"thumb_count"`           //点赞数
	Ctime      time.Time `json:"ctime" xorm:"created" ` //创建时间
	Mtime      time.Time `json:"mtime" xorm:"updated" ` //最后一次更新时间
}

type Api struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Url      string    `json:"url"`
	Category int       `json:"category"`              // 1是阅读点赞接口  2是详情接口 3是其他接口等
	Ctime    time.Time `json:"ctime" xorm:"created" ` //创建时间
	Mtime    time.Time `json:"mtime" xorm:"updated" ` //最后一次更新时间
}

//微信详情数据(es)
//全文搜索准备 加入 title  biz
type WeiXinDetail struct {
	Id        string    `json:"id"`         //主键id  ArticleId
	Title     string    `json:"title"`      //标题
	Text      string    `json:"text"`       //正文
	TextStyle string    `json:"text_style"` //带样式的正文
	Biz       string    `json:"biz"`        //biz
	Url       string    `json:"url"`        //url
	Ctime     time.Time `json:"ctime"`      //入库时间
	Mtime     time.Time `json:"mtime"`      //修改时间
	Ptime     time.Time `json:"ptime"`      //发布时间
	Author    string    `json:"author"`     //作者
	Original  int       `json:"original"`   //原创
	WordCloud string    `json:"word_cloud"` //词云数据
	Summary   string    `json:"summary"`    //摘要
	Forbid    int       `json:"forbid"`     //公号是否被微信官方搞事了
	From      string    `json:"from"`       //来源微信
}

//文章列表
type WeiXinList struct {
	Id        int64     //主键id
	HashId    string    `json:"hash_id"`                     //文章id
	SourceUrl string    `json:"source_url" `                 //原始url
	Url       string    `json:"url"  binding:"required" `    //文章url
	Title     string    `json:"title" binding:"required"   ` //文章标题
	Ptime     time.Time `json:"ptime"`                       //发布时间
	Biz       string    `json:"biz"`                         //biz信息
	Digest    string    `json:"digest"`                      //摘要
	Original  int       `json:"original"`                    //原型信息
	Type      int       `json:"type"`                        //api接口中的字段
	DelFlag   int       `json:"del_flag"`                    //是否删除
	Cover     string    `json:"cover"`                       //图链接
	Ctime     time.Time `json:"ctime" xorm:"created" `       //创建时间
	Mtime     time.Time `json:"mtime" xorm:"updated" `       //最后一次更新时间
}

//用户提交信息的微信号
/**
审核通过做数据迁移，只把name迁移进去(同时精确查询name是否存在)
用户提交的时候也进行精确查询，如果存在提示已经存在,不在用户能看到的公号范围，提示用户xx
*/
type UserWx struct {
	Uid    string    `json:"uid"`
	Biz    string    `json:"biz"`                   //biz信息
	Name   string    `json:"name"`                  //微信名称
	Status int       `json:"status"`                //同步到抓取数据库 status是1 否则是-1 提交初始状态是0
	Note   string    `json:"note"`                  //未审核通过原因
	Ctime  time.Time `json:"ctime" xorm:"created" ` //数据提交时间
	Mtime  time.Time `json:"mtime" xorm:"updated"`  //审核时间
}

//查询参数
type WeiXinParams struct {
	Biz      string `json:"biz"`      //查询biz
	Keywords string `json:"keywords"` //关键字
	From     string `json:"from"`     //发布时间起始
	To       string `json:"to"`       // 发布时间截止
	Title    string `json:"title"`    //标题模糊查询
	Pn       int    `json:"pn"`
	Ps       int    `json:"ps"`
	Type     int    `json:"type"` //导出当前页还是导出全部数据  1是导出全部  -1导出当前页
}
