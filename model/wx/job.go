package wx

import "time"

/**
第一次运行的时候注册，客户端分配Id
*/
type Job struct {
	Id       string    `json:"id"`                                          //客户端id
	Ip       string    `json:"ip"`                                          //客户端ip
	Ctime    time.Time `json:"ctime"  xorm:"created"`                       //注册时间
	Count    int       `json:"count"`                                       //运行次数
	JobCount int       `json:"job_count" xorm:"int(11) not null default 2"` //任务数量
	Interval int       `json:"interval" xorm:"int(11) not null default 20"` //时间间隔 ，默认值是 1分钟 配置的
	Status   int       `json:"status"`                                      //是否在线运行
	Etime    time.Time `json:"etime"`                                       //运行结束时间
}
