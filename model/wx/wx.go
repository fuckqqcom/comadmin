package wx

import "time"

//公号属性
type WeiXin struct {
	Biz   string    `json:"biz"` //公号biz
	Name  string    `json:"name"`
	Desc  string    `json:"detail"` //公号描述
	Count int       `json:"count"`  //公号文章数
	Ctime time.Time `json:"ctime"`  //创建时间
	Mtime time.Time `json:"mtime"`  //最后一次更新时间
}

//阅读量和点赞量
type WeiXinCount struct {
	Biz        string `json:"biz"`
	ArticleId  string `json:"article_id"`  //文章id
	ReadCount  int    `json:"read_count"`  //阅读量
	ThumbCount int    `json:"thumb_count"` //点赞数
}

type Api struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Category int    `json:"category"` // 1是阅读点赞接口  2是详情接口 3是其他接口等
}
