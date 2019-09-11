package main

import (
	"comadmin/job/parse"
	"comadmin/tools/utils"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

/**
调度任务
*/

var (
	Id       = ""
	Interval = 0
	JobCount = 0
)

func main() {
	if register() == nil {
		loadDetailJob()
		unregister()
	}

}

const (
	//queue
	url        = "http://api.pipenv.com/v1/wx/getTasks"
	onlineJob  = "http://api.pipenv.com/v1/wx/onlineJob"
	offlineJob = "http://api.pipenv.com/v1/wx/offlineJob"
	//detail

)

//注册函数 onlineJob
func register() error {
	type params struct {
		Ip string `json:"ip"`
		Id string `json:"id"`
	}
	ip := utils.GetClientIp()
	p := params{Ip: ip, Id: ""}
	bytes, err := json.Marshal(p)

	if err != nil {
		return err
	}

	payload := strings.NewReader(string(bytes))
	header := map[string]string{"Content-Type": "application/json"}
	r := parse.Request{Url: onlineJob, Interval: 10, Timeout: 100, Body: payload, Header: header, Method: "POST"}
	data, err := r.Fetch()
	//{"code":200,"data":"2dda16efea31ea25df7c5fd57dfd242b","msg":"成功"}
	type Ret struct {
		Code int `json:"code"`
		Data struct {
			Id       string `json:"id"`
			Interval int    `json:"interval"`
			JobCount int    `json:"job_count"`
		} `json:"data"`
		Msg string `json:"msg"`
	}
	var ret Ret
	json.Unmarshal(data, &ret)

	if ret.Code == 200 {
		Id = ret.Data.Id
		Interval = ret.Data.Interval
		JobCount = ret.Data.JobCount
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func unregister() error {
	type params struct {
		Ip string `json:"ip"`
		Id string `json:"id"`
	}
	ip := utils.GetClientIp()
	p := params{Ip: ip, Id: Id}
	bytes, err := json.Marshal(p)

	if err != nil {
		return err
	}

	payload := strings.NewReader(string(bytes))
	header := map[string]string{"Content-Type": "application/json"}
	r := parse.Request{Url: offlineJob, Interval: 10, Timeout: 100, Body: payload, Header: header, Method: "POST"}
	_, err = r.Fetch()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func loadDetailJob() {

	type RetData struct {
		Code int `json:"code"`
		Data struct {
			Count int          `json:"count"`
			List  []parse.Info `json:"list"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	var ret RetData

	payload := strings.NewReader("{\"num\": " + strconv.Itoa(JobCount) + " }")
	header := map[string]string{"Content-Type": "application/json"}
	r := parse.Request{Url: url, Interval: Interval, Timeout: 10, Body: payload, Header: header}
	bytes, err := r.Fetch()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(bytes, &ret)

	if ret.Code == 200 {
		if ret.Data.Count == 0 {
			log.Printf("暂时没有任务")
		}
		for _, v := range ret.Data.List {

			parse.Detail(v)
		}
	}

}
