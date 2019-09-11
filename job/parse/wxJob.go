package parse

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"strings"
	"time"
)

/**
获取任务
*/

//http请求

const (
	//detail = "http://api.pipenv.com/v1/wx/addDetail"
	detail = "http://127.0.0.1:1234/v1/wx/addDetail"
)

func Detail(i Info) {
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
		Url:       i.Url,
		Title:     i.Title,
		Text:      strings.TrimSpace(content),
		TextStyle: strings.TrimSpace(strings.ReplaceAll(contentStyle, `data-src="`, `src="`)),
		Biz:       i.Biz,
		Ptime:     i.Ptime,
		Author:    strings.ReplaceAll(strings.TrimSpace(nickName), "\n", ""),
		From:      "wx",
		Ctime:     time.Now(),
		Mtime:     time.Now(),
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
	fetch, err := r.Fetch()
	log.Printf("upload data is %v %s", string(fetch), err)
}
