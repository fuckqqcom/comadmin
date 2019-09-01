package parse

import (
	"comadmin/tools/utils"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"strings"
)

/**
获取任务
*/

//http请求

func ParseDetail(i Info) {
	if i.Url == "" {
		return
	}
	r := Request{
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

	params := Params{
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
func uploadData(params Params) {
	bytes, err := json.Marshal(params)

	if err != nil {
		return
	}

	payload := strings.NewReader(string(bytes))
	header := map[string]string{"Content-Type": "application/json"}
	r := Request{
		Id:          "",
		Url:         "",
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
