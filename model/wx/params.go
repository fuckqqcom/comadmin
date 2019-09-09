package wx

//查询近7天数据 weixinList
type Nearly7Day struct {
	Biz string `json:"biz"`
	Pn  int
	Ps  int
}

//提交列表数据
type AddWxParams struct {
	Url   string `json:"url"  binding:"required" `   //文章url
	Title string `json:"title"  binding:"required" ` //文章标题
	Ptime int64  `json:"ptime"`                      //发布时间
}

type BizParams struct {
	MobileId string `json:"mobile_id"  binding:"required"` //手机id
	Ps       int    `json:"ps"`
	Pn       int    `json:"pn"`
}

type ApiParams struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Category int    `json:"category"` // 1是阅读点赞接口  2是详情接口 3是其他接口等
}

type UpdateKeyParams struct {
	Url string `json:"url" binding:"required"`
	Key string `json:"key" binding:"required"`
}

//到处biz的pdf
type BizPdf struct {
	Biz string `json:"biz"`
	Ids string `json:"ids"`
}
