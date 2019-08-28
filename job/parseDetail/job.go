package parseDetail

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

/**
获取任务
*/

//http请求

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
