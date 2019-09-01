package wx

import "time"

//查询近7天数据 weixinList
type Nearly7Day struct {
	Biz string `json:"biz"`
	Pn  int
	Ps  int
}

//提交列表数据
type AddWxParams struct {
	Id        int64
	HashId    string    `json:"id"`                         //文章id
	SourceUrl string    `json:"source_url" `                //原始url
	Url       string    `json:"url"  binding:"required" `   //文章url
	Title     string    `json:"title"  binding:"required" ` //文章标题
	Ptime     time.Time `json:"ptime"`                      //发布时间
	Biz       string    `json:"biz"  binding:"required"`    //biz信息
	Digest    string    `json:"digest"`                     //摘要
	Original  int       `json:"original"`                   //原型信息
	Type      int       `json:"type"`                       //api接口中的字段
	DelFlag   int       `json:"del_flag"`                   //是否删除
	Cover     string    `json:"cover"`                      //图链接
}