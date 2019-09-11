package wx

import "time"

/**
第一次运行的时候注册，客户端分配Id
*/

type Job struct {
	Id       string    `json:"id"`                    //客户端id
	Ip       string    `json:"ip"`                    //客户端ip
	Ctime    time.Time `json:"ctime"  xorm:"created"` //注册时间
	Count    int       `json:"count"`                 //运行次数
	JobCount int       `json:"job_count"`             //任务数量
	Status   int       `json:"status"`                //是否在线运行
	Etime    time.Time `json:"etime" xorm:"updated"`  //运行结束时间
}
