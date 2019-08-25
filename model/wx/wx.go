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

//微信详情数据(es)
//全文搜索准备 加入 title  biz
type WeiXinDetail struct {
	Id        string    `json:"id"`         //主键id  ArticleId
	Title     string    `json:"title"`      //标题
	Text      string    `json:"text"`       //正文
	TextStyle string    `json:"text_style"` //带样式的正文
	Biz       string    `json:"biz"`        //biz
	Ctime     time.Time `json:"ctime"`      //入库时间
	Mtime     time.Time `json:"mtime"`      //修改时间
	Ptime     time.Time `json:"ptime"`      //发布时间
	Author    time.Time `json:"author"`     //作者
	Original  int       `json:"original"`   //原创
	WordCloud string    `json:"word_cloud"` //词云数据
	Summary   string    `json:"summary"`    //摘要
}

type WeiXinParams struct {
	Biz      string `json:"biz"`      //查询biz
	Keywords string `json:"keywords"` //关键字
	From     string `json:"from"`     //发布时间起始
	To       string `json:"to"`       // 发布时间截止
	Title    string `json:"title"`    //标题模糊查询
	Pn       int    `json:"pn"`
	Ps       int    `json:"ps"`
}
