package main

import (
	"comadmin/job/parse"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/**
调度任务
*/

func main() {
	loadDetailJob()
}

const (
	//queue
	url = "http://api.pipenv.com/v1/wx/getTasks"
	//detail

)

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

	payload := strings.NewReader("{\"num\":10}")
	header := map[string]string{"Content-Type": "application/json"}
	r := parse.Request{Url: url, Interval: 10, Timeout: 10, Body: payload, Header: header}
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
		for _, v := range ret.Data.List {

			parse.Detail(v)
		}
	}

}
