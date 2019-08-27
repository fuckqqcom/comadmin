package wx

import "time"

/**
第一次运行的时候注册，客户端分配Id
*/

type Job struct {
	Id     string    `json:"id"`     //客户端id
	IP     string    `json:"ip"`     //客户端ip
	Ctime  time.Time `json:"ctime"`  //注册时间
	Count  int       `json:"count"`  //运行次数
	Status int       `json:"status"` //是否在线运行
	Etime  time.Time `json:"etime"`  //运行结束时间
}
