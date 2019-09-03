package parse

import (
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Request struct {
	Id          string
	Url         string
	Body        io.Reader
	Retry       int               //重试次数
	Timeout     time.Duration     //超时时间
	Interval    int               //间隔时间
	Method      string            //请求方式
	Header      map[string]string //每个请求自带的header
	VerifyProxy bool              //是否设置代理
	VerifyTLS   bool              //http false  or https true
}

type Info struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Url      string    `json:"url"`
	Original int       `json:"origin"`
	Biz      string    `json:"biz"`
	Ptime    time.Time `json:"ptime"`
}

type Params struct {
	Id        string    `json:"id"  binding:"required" `         //主键id  ArticleId
	Title     string    `json:"title"  binding:"required" `      //标题
	Text      string    `json:"text"  binding:"required" `       //正文
	TextStyle string    `json:"text_style"  binding:"required" ` //带样式的正文
	Biz       string    `json:"biz"  binding:"required" `        //biz
	Url       string    `json:"url"`
	Ctime     time.Time `json:"ctime"`                      //入库时间
	Mtime     time.Time `json:"mtime"`                      //修改时间
	Ptime     time.Time `json:"ptime" binding:"required" `  //发布时间
	Author    string    `json:"author" binding:"required" ` //作者
	From      string    `json:"from"  binding:"required"`
	//Original int       `json:"original" binding:"required"` //原创
	//WordCloud string    `json:"word_cloud"` //词云数据
	//Summary   string    `json:"summary"`    //摘要
	//Forbid    int       `json:"forbid"`     //公号是否被微信官方搞事了
}

func (r *Request) Fetch() ([]byte, error) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(r.Interval) //生成0-99随机整数
	time.Sleep(time.Duration(x) * time.Second)
	log.Printf("now is run url %s", r.Url)
	tr := &http.Transport{}
	if r.VerifyTLS == true {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	//if r.VerifyProxy == true {
	//	tr.Proxy = proxy
	//}

	client := &http.Client{Timeout: r.Timeout * time.Second, Transport: tr}
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		log.Printf("http request url error:(%v)", err)
		return nil, err
	}
	if r.Header != nil {
		for k, v := range r.Header {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http response url error:(%v)", err)
		return nil, errors.New("http response error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("wrong status code")
	}

	return ioutil.ReadAll(resp.Body)
}
