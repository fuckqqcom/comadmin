package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
	"time"
)

const (
	str = `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
)

/**
pageSize:每页展示多少
pageNum:页码
defaultPageSize:默认每页展示多少
*/

func Pagination(pageSize, pageNum, defaultPageSize int) (ps, pn int) {

	if pageNum <= 1 {
		pn = 1
	} else {
		pn = pageNum
	}

	if pageSize <= 0 || pageSize >= defaultPageSize {
		ps = defaultPageSize
	} else {
		ps = pageSize
	}
	return
}

func CheckError(err error, v interface{}) bool {
	if err != nil {
		log.Printf("err is %s,%s", err, v)
		return false
	}
	return true
}

func FindBizStr(url string) (arr []string) {
	fmt.Println(url)
	//a := "http://mp.weixin.qq.com/s?__biz=MzU3ODE2NTMxNQ==&MID=2247485961&idx=1&sn=431af867d04efd973fd16df359365dd6&chksm=fd78c525ca0f4c334da2c677c1622f32058b7d3b89d255d5bb6e21a11a7f32407b67b13245bd&scene=27#wechat_redirect"

	//index := strings.Index(a, "__biz=")
	//fmt.Println(index)
	//print(a[index:])
	bizIndex := strings.Index(url, "__biz=")
	if bizIndex == -1 {
		return nil
	}
	bizEnd := strings.Index(url[bizIndex:], "&")
	biz := url[bizIndex+6 : bizEnd+bizIndex]
	arr = append(arr, biz)

	//mid
	midIndex := strings.Index(url, "mid=")
	if midIndex == -1 {
		midIndex = strings.Index(url, "MID=")
	}
	if midIndex == -1 {
		return nil
	}
	midEnd := strings.Index(url[midIndex:], "&")
	mid := url[midIndex+4 : midIndex+midEnd]
	arr = append(arr, mid)

	idxIndex := strings.Index(url, "&idx=")
	if midIndex == -1 {
		idxIndex = strings.Index(url, "&idx=")
	}
	if midIndex == -1 {
		return nil
	}
	idxEnd := strings.Index(url[idxIndex+5:], "&")
	idx := url[idxIndex+5 : idxEnd+idxIndex+5]
	arr = append(arr, idx)

	return
}

func QueryCols(m map[string]interface{}) (query string, value []interface{}) {
	//待测试 这里是 value是int 和字符串时间等
	count := 0
	if m != nil {
		for k, v := range m {
			if count == 0 {
				query += k + "  ? "
			} else {
				query += " and " + k + "  ? "
			}
			value = append(value, v)
			count += 1
		}
	}
	return
}

func EncodeMd5(value string) string {

	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func SqlRegex(param string) bool {

	re, err := regexp.Compile(str)
	if err != nil {
		log.Printf("error param error is %s", err.Error())
		return true
	}

	return re.MatchString(param)
}

func StringJoin(a ...string) string {
	var buf bytes.Buffer
	for _, k := range a {
		buf.WriteString(k)
	}
	return buf.String()
}

func Time2Str(t time.Time, format string) string {
	//now := time.Now()
	//formatNow := now.Format("2006-01-02 15:04:05")
	formatNow := t.Format(format)
	fmt.Println(formatNow)
	return formatNow
}

func NowTime() time.Time {
	timeUnix := time.Now().Format("2006-01-02")
	location, _ := time.ParseInLocation("2006-01-02", timeUnix, time.Local)
	return location
}

func Str2Time(format, value string) time.Time {
	local, _ := time.LoadLocation("Local")
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-06-20 18:16:15", local)
	t, _ := time.ParseInLocation(format, value, local)

	fmt.Println(t)
	return t
}

/**
打印es sql
*/
func PrintQuery(src interface{}) {
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	log.Printf("es sql--->%s", string(data))
}

/**
获取本地ip地址
*/

func GetClientIp() string {

	netInterfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			adders, _ := netInterfaces[i].Addrs()

			for _, address := range adders {
				if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
					if inet.IP.To4() != nil {
						return inet.IP.String()
					}
				}
			}
		}

	}
	return ""
}
