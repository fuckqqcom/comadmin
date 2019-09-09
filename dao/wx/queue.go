package wxd

import (
	"comadmin/model/wx"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"encoding/json"
	"log"
	"time"
)

/**
微信的一些redis队列数据
*/

/**
详情页抓取队列:
	入队和出队
*/

const (
	detailKey = "detailKey"
)

func (d Dao) addQueue(w wx.WeiXinList) int {
	code := d.add(w)
	//先判断是否创建成功
	if code != e.Success {
		return code
	}
	type Params struct {
		Id       string    `json:"id"`
		Title    string    `json:"title"`
		Url      string    `json:"url"`
		Original int       `json:"origin"`
		Biz      string    `json:"biz"`
		Ptime    time.Time `json:"ptime"`
	}

	p := Params{
		Id:       w.HashId,
		Title:    w.Title,
		Url:      w.Url,
		Original: w.Original,
		Biz:      w.Biz,
		Ptime:    w.Ptime,
	}
	bytes, err := json.Marshal(p)
	if utils.CheckError(err, string(bytes)) {
		add, err := d.rs.SAdd(detailKey, string(bytes)).Result()
		if utils.CheckError(err, add) {
			log.Printf("addQueue ret is %v", add)
			return e.Success
		}
		return e.Errors
	}
	return e.Errors
}

/**
从队列中取数据
*/
func (d Dao) queues(l wx.WeiXinList, num, pn int) (interface{}, int) {
	if num <= 0 || num >= 10 {
		num = 2
	}
	s := make([]interface{}, 0)
	/**
	先根据出来的数据考虑是否序列化下
	*/
	type Params struct {
		Id       string    `json:"id"`
		Title    string    `json:"title"`
		Url      string    `json:"url"`
		Original int       `json:"origin"`
		Biz      string    `json:"biz"`
		Ptime    time.Time `json:"ptime"`
	}
	for i := 0; i <= num; i++ {
		var p Params
		pop, err := d.rs.SPop(detailKey).Result()
		if pop == "" {
			utils.CheckError(err, pop)
			break
		}
		if utils.CheckError(err, pop) {
			err := json.Unmarshal([]byte(pop), &p)
			if utils.CheckError(err, "json to struct") {
				s = append(s, p)
			}
			//s[i] = p
		}
	}
	return s, len(s)
}
