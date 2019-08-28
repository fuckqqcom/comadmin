package main

import (
	"comadmin/job/parse"
	"comadmin/tools/utils"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"strings"
)

/**
调度任务
*/

func main() {
	loadJob()
}

const (
	url = "http://127.0.0.1:1234/v1/wx/pop"
	//detail
	detail = "http://127.0.0.1:1234/v1/wx/detail"
)

func loadJob() {

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

			parseDetail(v)
		}
	}

}

func parseDetail(i parse.Info) {
	if i.Url == "" {
		return
	}
	r := parse.Request{
		Id:          i.Id,
		Url:         i.Url,
		Body:        nil,
		Retry:       3,
		Timeout:     6,
		Interval:    10,
		Method:      "GET",
		Header:      nil,
		VerifyProxy: false,
		VerifyTLS:   false,
	}

	bytes, err := r.Fetch()
	if err != nil {
		log.Println(err)
	}

	newReader := strings.NewReader(string(bytes))
	reader := io.Reader(newReader)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	nickName := doc.Find("#js_name").Text()
	contentStyle, _ := doc.Find("#js_content").Html()
	content := doc.Find("#js_content").Text()
	fmt.Println(nickName, contentStyle, content)

	params := parse.Params{
		Id:        i.Id,
		Title:     i.Url,
		Text:      content,
		TextStyle: contentStyle,
		Biz:       i.Biz,
		Ptime:     utils.Time2Str(i.Ptime, "2006-01-02 15:04:05"),
		Author:    nickName,
		From:      "wx",
	}
	uploadData(params)

}

//传递数据
func uploadData(params parse.Params) {
	bytes, err := json.Marshal(params)

	if err != nil {
		return
	}

	payload := strings.NewReader(string(bytes))
	header := map[string]string{"Content-Type": "application/json"}
	r := parse.Request{
		Id:          "",
		Url:         detail,
		Body:        payload,
		Retry:       3,
		Timeout:     10,
		Interval:    1,
		Method:      "POST",
		Header:      header,
		VerifyProxy: false,
		VerifyTLS:   false,
	}
	r.Fetch()
}
