package main

import (
	"comadmin/job/parseDetail"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

/**
调度任务
*/

func main() {
	loadJob()
}

const url = "http://127.0.0.1:1234/v1/wx/pop"

func loadJob() {

	type RetData struct {
		Code int `json:"code"`
		Data struct {
			Count int `json:"count"`
			List  []struct {
				Id       string    `json:"id"`
				Title    string    `json:"title"`
				Url      string    `json:"url"`
				Original int       `json:"origin"`
				Biz      string    `json:"biz"`
				Ptime    time.Time `json:"ptime"`
			} `json:"list"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	var ret RetData

	payload := strings.NewReader("{\"num\":10}")
	header := map[string]string{"Content-Type": "application/json"}
	r := parseDetail.Request{Url: url, Interval: 10, Timeout: 10, Body: payload, Header: header}
	bytes, err := r.Fetch()
	fmt.Println(string(bytes))
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(bytes, &ret)

	if ret.Code == 200 {
		if ret.Data.Count == 0 {
			log.Printf("暂时没有任务")
		}
		for k, v := range ret.Data.List {
			fmt.Println(k, v)
		}
	}

}
